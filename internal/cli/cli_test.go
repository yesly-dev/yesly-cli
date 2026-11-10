package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/yesly-dev/yesly-cli/internal/config"
	"github.com/yesly-dev/yesly-cli/internal/spec"
	"github.com/yesly-dev/yesly-cli/internal/webhook"
)

type call struct {
	Method, Path, Query, Body string
	Header                    http.Header
}

// fakeAPI records requests and answers with handler, or {} by default.
type fakeAPI struct {
	*httptest.Server
	mu      sync.Mutex
	calls   []call
	handler func(w http.ResponseWriter, r *http.Request, body string)
}

func newFakeAPI(t *testing.T) *fakeAPI {
	f := &fakeAPI{}
	f.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		f.mu.Lock()
		f.calls = append(f.calls, call{r.Method, r.URL.Path, r.URL.RawQuery, string(b), r.Header.Clone()})
		h := f.handler
		f.mu.Unlock()
		if h != nil {
			h(w, r, string(b))
			return
		}
		_, _ = w.Write([]byte(`{"object":"ok"}`))
	}))
	t.Cleanup(f.Close)
	return f
}

func (f *fakeAPI) last(t *testing.T) call {
	t.Helper()
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.calls) == 0 {
		t.Fatal("no request was sent")
	}
	return f.calls[len(f.calls)-1]
}

// setup isolates the config file and environment.
func setup(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.json")
	t.Setenv("YESLY_CONFIG", path)
	t.Setenv("YESLY_API_KEY", "")
	t.Setenv("YESLY_BASE_URL", "")
	t.Setenv("YESLY_PROFILE", "")
	return path
}

type result struct {
	out, err string
	code     int
}

func run(t *testing.T, stdin string, args ...string) result {
	t.Helper()
	var out, errb bytes.Buffer
	app := &App{
		In:    strings.NewReader(stdin),
		Out:   &out,
		Err:   &errb,
		Sleep: func(time.Duration) {},
		Now:   func() time.Time { return time.Unix(1794474252, 0) },
	}
	root := NewRoot(app)
	root.SetArgs(args)
	err := root.ExecuteContext(context.Background())
	if err != nil {
		printError(&errb, err)
	}
	return result{out.String(), errb.String(), ExitCode(err)}
}

func TestEveryOperationIsACommandWithAnExample(t *testing.T) {
	root := NewRoot(&App{In: strings.NewReader(""), Out: io.Discard, Err: io.Discard})
	for _, op := range spec.Operations {
		cmd, _, err := root.Find([]string{kebab(op.Resource), kebab(op.Method)})
		if err != nil || cmd.Name() != kebab(op.Method) {
			t.Errorf("%s.%s has no command", op.Resource, op.Method)
			continue
		}
		if examples[op.Resource+"."+op.Method] == "" {
			t.Errorf("%s.%s has no example in examples.go", op.Resource, op.Method)
		}
	}
	if len(examples) != len(spec.Operations) {
		t.Errorf("examples.go has %d entries for %d operations", len(examples), len(spec.Operations))
	}
}

func TestCreateBuildsTheBodyFromFlagsAndData(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "appointments", "create", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"--data", `{"type":"delivery","customer":{"email":"m@example.com","name":"Old"}}`,
		"--external-id", "ORD-1", "--customer.name", "Martha", "--customer.phone", "+34600000001",
		"--slot.start", "2026-11-13T09:00:00+01:00", "--channels", "whatsapp,voice",
		"--metadata", "order_id=PV-1", "--metadata", "route_id=BIO-07")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	c := api.last(t)
	if c.Method != "POST" || c.Path != "/appointments" {
		t.Fatalf("%s %s", c.Method, c.Path)
	}
	var got map[string]any
	_ = json.Unmarshal([]byte(c.Body), &got)
	want := map[string]any{
		"type":        "delivery",
		"external_id": "ORD-1",
		"customer":    map[string]any{"email": "m@example.com", "name": "Martha", "phone": "+34600000001"},
		"slot":        map[string]any{"start": "2026-11-13T09:00:00+01:00"},
		"channels":    []any{"whatsapp", "voice"},
		"metadata":    map[string]any{"order_id": "PV-1", "route_id": "BIO-07"},
	}
	wantJSON, _ := json.Marshal(want)
	gotJSON, _ := json.Marshal(got)
	if string(wantJSON) != string(gotJSON) {
		t.Fatalf("body = %s\nwant %s", gotJSON, wantJSON)
	}
	if c.Header.Get("Idempotency-Key") != "" {
		t.Fatal("create is idempotent by external_id and sends no Idempotency-Key")
	}
	if r.out != "{\n  \"object\": \"ok\"\n}\n" {
		t.Fatalf("output = %q", r.out)
	}
}

func TestTypedFlags(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "appointment-types", "update", "delivery", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"--reschedule.window-days", "21", "--reschedulable=false")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	c := api.last(t)
	if c.Method != "PATCH" || c.Path != "/appointment_types/delivery" || c.Body != `{"reschedulable":false,"reschedule":{"window_days":21}}` {
		t.Fatalf("%s %s %s", c.Method, c.Path, c.Body)
	}
}

func TestPostWithoutBodyFlagsSendsAnEmptyObjectAndAnIdempotencyKey(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "webhook-endpoints", "roll-secret", "we_1", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	c := api.last(t)
	if c.Path != "/webhook_endpoints/we_1/roll_secret" || c.Body != "{}" || len(c.Header.Get("Idempotency-Key")) != 36 {
		t.Fatalf("%+v", c)
	}
	run(t, "", "appointments", "cancel", "apt_1", "--api-key", "ysk_test_x", "--base-url", api.URL, "--idempotency-key", "mine")
	if got := api.last(t).Header.Get("Idempotency-Key"); got != "mine" {
		t.Fatalf("Idempotency-Key = %q", got)
	}
}

func TestQueryFlags(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "appointments", "list", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"--status", "no_reply", "--metadata", "route_id=BIO-07", "--limit", "5")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	if q := api.last(t).Query; q != "limit=5&metadata%5Broute_id%5D=BIO-07&status=no_reply" {
		t.Fatalf("query = %s", q)
	}
}

func TestPathArgumentsAreEscaped(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	run(t, "", "locations", "update", "a/b", "--api-key", "ysk_test_x", "--base-url", api.URL, "--name", "X")
	if p := api.last(t); !strings.HasSuffix(p.Path, "/locations/a/b") {
		// Go decodes %2F in URL.Path; the raw request keeps it escaped.
		t.Fatalf("path = %s", p.Path)
	}
}

func TestListAll(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	api.handler = func(w http.ResponseWriter, r *http.Request, _ string) {
		switch r.URL.Query().Get("starting_after") {
		case "":
			_, _ = w.Write([]byte(`{"object":"list","data":[{"id":"evt_3"},{"id":"evt_2"}],"has_more":true}`))
		case "evt_2":
			_, _ = w.Write([]byte(`{"object":"list","data":[{"id":"evt_1"}],"has_more":false}`))
		default:
			w.WriteHeader(400)
		}
	}
	r := run(t, "", "events", "list", "--all", "--compact", "--limit", "2", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 0 || r.out != `[{"id":"evt_3"},{"id":"evt_2"},{"id":"evt_1"}]`+"\n" {
		t.Fatalf("exit %d, out %q, err %q", r.code, r.out, r.err)
	}
	if q := api.last(t).Query; q != "limit=2&starting_after=evt_2" {
		t.Fatalf("query = %s", q)
	}
}

func TestLiveKeysNeedLiveToChangeData(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "appointments", "cancel", "apt_1", "--api-key", "ysk_live_x", "--base-url", api.URL)
	if r.code != 2 || !strings.Contains(r.err, "--live") || len(api.calls) != 0 {
		t.Fatalf("exit %d, err %q, calls %d", r.code, r.err, len(api.calls))
	}
	if r := run(t, "", "appointments", "list", "--api-key", "ysk_live_x", "--base-url", api.URL); r.code != 0 {
		t.Fatalf("GET with a live key: %s", r.err)
	}
	if r := run(t, "", "appointments", "cancel", "apt_1", "--live", "--api-key", "ysk_live_x", "--base-url", api.URL); r.code != 0 {
		t.Fatalf("--live: %s", r.err)
	}
}

func TestAPIErrorsExitWithOne(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	api.handler = func(w http.ResponseWriter, _ *http.Request, _ string) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte(`{"error":{"type":"invalid_request","code":"resource_missing","message":"No such appointment.","request_id":"req_1"}}`))
	}
	r := run(t, "", "appointments", "retrieve", "apt_x", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 1 || r.err != "Error: No such appointment. (status 404, type invalid_request, code resource_missing, request req_1)\n" {
		t.Fatalf("exit %d, err %q", r.code, r.err)
	}
}

func TestUsageErrorsExitWithTwo(t *testing.T) {
	setup(t)
	for _, args := range [][]string{
		{"appointments", "retrieve"},
		{"appointments", "list", "extra"},
		{"appointments", "list", "--status", "nope"},
		{"appointments", "create", "--nope"},
		{"appointments", "create", "--data", "[1]"},
		{"appointments", "create", "--metadata", "novalue"},
		{"appointments", "list", "--base-url", "http://example.com"},
	} {
		r := run(t, "", append(args, "--api-key", "ysk_test_x", "--dry-run")...)
		if r.code != 2 {
			t.Errorf("%v: exit %d, err %q", args, r.code, r.err)
		}
	}
}

func TestDryRunNeedsNoKey(t *testing.T) {
	setup(t)
	r := run(t, "", "appointments", "cancel", "apt_1", "--reason", "Closed", "--dry-run", "--idempotency-key", "k")
	want := "POST https://api.yesly.dev/v1/appointments/apt_1/cancel\nYesly-Version: 2026-10-01\nIdempotency-Key: k\n\n{\n  \"reason\": \"Closed\"\n}\n"
	if r.code != 0 || r.out != want {
		t.Fatalf("exit %d, out %q, err %q", r.code, r.out, r.err)
	}
}

func TestJSONFlagFromFile(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	file := filepath.Join(t.TempDir(), "a.json")
	_ = os.WriteFile(file, []byte(`[{"external_id":"A"}]`), 0o600)
	r := run(t, "", "appointments", "import", "--appointments", "@"+file, "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 0 || api.last(t).Body != `{"appointments":[{"external_id":"A"}]}` {
		t.Fatalf("exit %d, err %q, body %s", r.code, r.err, api.last(t).Body)
	}
	r = run(t, `{"appointments":[]}`, "appointments", "import", "--data", "@-", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 0 || api.last(t).Body != `{"appointments":[]}` {
		t.Fatalf("stdin: exit %d, body %s", r.code, api.last(t).Body)
	}
}

func TestAPICommand(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "api", "post", "appointments", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"-f", "external_id=ORD-1", "-f", "customer.name=Martha", "-f", "customer[phone]=+34600000001",
		"-f", "channels[]=sms", "-f", "channels[]=voice", "-F", "metadata={\"a\":\"b\"}")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	c := api.last(t)
	var got map[string]any
	_ = json.Unmarshal([]byte(c.Body), &got)
	gotJSON, _ := json.Marshal(got)
	if string(gotJSON) != `{"channels":["sms","voice"],"customer":{"name":"Martha","phone":"+34600000001"},"external_id":"ORD-1","metadata":{"a":"b"}}` {
		t.Fatalf("body = %s", gotJSON)
	}
	if c.Header.Get("Idempotency-Key") != "" {
		t.Fatal("POST /appointments takes no Idempotency-Key")
	}

	run(t, "", "api", "POST", "/appointments/apt_1/cancel", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if api.last(t).Header.Get("Idempotency-Key") == "" {
		t.Fatal("cancel should get an Idempotency-Key")
	}

	run(t, "", "api", "GET", "/appointments?status=no_reply", "-F", "limit=5", "-f", "metadata[route_id]=BIO", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if q := api.last(t).Query; q != "limit=5&metadata%5Broute_id%5D=BIO&status=no_reply" {
		t.Fatalf("query = %s", q)
	}
}

func TestLoginSavesAProfile(t *testing.T) {
	path := setup(t)
	api := newFakeAPI(t)
	r := run(t, "ysk_test_abcdef123456\n", "login", "--profile", "staging", "--base-url", api.URL)
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	if c := api.last(t); c.Path != "/balance" || c.Header.Get("Authorization") != "Bearer ysk_test_abcdef123456" {
		t.Fatalf("check call = %+v", c)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	// Windows has no Unix permission bits; the file sits in the user's profile there.
	if runtime.GOOS != "windows" && info.Mode().Perm() != 0o600 {
		t.Fatalf("config file mode = %v", info.Mode())
	}
	f, _ := config.Load()
	p := f.Profile("staging")
	if p == nil || p.APIKey != "ysk_test_abcdef123456" || p.BaseURL != api.URL {
		t.Fatalf("profile = %+v", p)
	}
	if strings.Contains(r.err, "abcdef123456") {
		t.Fatal("login printed the whole key")
	}

	// The profile is used by later commands.
	r = run(t, "", "balance", "retrieve", "-p", "staging")
	if r.code != 0 || api.last(t).Header.Get("Authorization") != "Bearer ysk_test_abcdef123456" {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	t.Setenv("YESLY_API_KEY", "ysk_test_fromenv")
	run(t, "", "balance", "retrieve", "-p", "staging")
	if got := api.last(t).Header.Get("Authorization"); got != "Bearer ysk_test_fromenv" {
		t.Fatalf("the environment should win over the profile, got %s", got)
	}

	r = run(t, "", "logout", "-p", "staging")
	f, _ = config.Load()
	if r.code != 0 || f.Profile("staging") != nil {
		t.Fatalf("logout: exit %d, %s", r.code, r.err)
	}
}

func TestLoginRejectsBadKeys(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	if r := run(t, "sk_nope\n", "login", "--base-url", api.URL); r.code != 2 {
		t.Fatalf("exit %d", r.code)
	}
	api.handler = func(w http.ResponseWriter, _ *http.Request, _ string) {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(`{"error":{"type":"authentication_error","message":"Invalid API key."}}`))
	}
	r := run(t, "ysk_test_revoked0000\n", "login", "--base-url", api.URL)
	f, _ := config.Load()
	if r.code != 1 || !strings.Contains(r.err, "Invalid API key") || len(f.Profiles) != 0 {
		t.Fatalf("exit %d, err %q, profiles %v", r.code, r.err, f.Profiles)
	}
}

func TestWebhooksSignAndVerify(t *testing.T) {
	setup(t)
	body := `{"id":"evt_1","object":"event","type":"appointment.confirmed","data":{"appointment":{"id":"apt_1","external_id":"ORD-48213"}}}`
	r := run(t, body, "webhooks", "sign", "--secret", "whsec_test")
	want := "t=1794474252,v1=62b8ce65d20d98e1d1d8f14ff76317b0bf9b166a3dc5c2d980022e54c62e0a67\n"
	if r.code != 0 || r.out != want {
		t.Fatalf("sign: exit %d, out %q, err %q", r.code, r.out, r.err)
	}
	r = run(t, body, "webhooks", "verify", "--secret", "whsec_test", "--header", strings.TrimSpace(want))
	if r.code != 0 {
		t.Fatalf("verify: exit %d, %s", r.code, r.err)
	}
	r = run(t, body+" ", "webhooks", "verify", "--secret", "whsec_test", "--header", strings.TrimSpace(want))
	if r.code != 1 {
		t.Fatalf("verify of a changed body: exit %d", r.code)
	}
}

func TestListenForwardsNewEventsSigned(t *testing.T) {
	setup(t)
	var mu sync.Mutex
	var delivered []call
	local := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		mu.Lock()
		delivered = append(delivered, call{Body: string(b), Header: r.Header.Clone()})
		mu.Unlock()
	}))
	defer local.Close()

	api := newFakeAPI(t)
	polls := 0
	api.handler = func(w http.ResponseWriter, r *http.Request, _ string) {
		polls++
		old := `{"id":"evt_old","type":"appointment.created","created":"2026-11-08T10:00:00Z"}`
		if polls == 1 {
			_, _ = w.Write([]byte(`{"data":[` + old + `],"has_more":false}`))
			return
		}
		_, _ = w.Write([]byte(`{"data":[{"id":"evt_new2","type":"message.sent","created":"2026-11-08T10:00:02Z"},` +
			`{"id":"evt_new1","type":"appointment.confirmed","created":"2026-11-08T10:00:01Z"},` + old + `],"has_more":false}`))
	}

	var errb bytes.Buffer
	app := &App{In: strings.NewReader(""), Out: io.Discard, Err: &errb, Now: time.Now, Sleep: func(time.Duration) {}}
	root := NewRoot(app)
	root.SetArgs([]string{"listen", "--api-key", "ysk_test_x", "--base-url", api.URL, "--forward-to", local.URL,
		"--secret", "whsec_local", "--interval", "10ms", "--events", "appointment.confirmed"})
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error)
	go func() { done <- root.ExecuteContext(ctx) }()
	deadline := time.Now().Add(5 * time.Second)
	for {
		mu.Lock()
		n := len(delivered)
		mu.Unlock()
		if n > 0 || time.Now().After(deadline) {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(50 * time.Millisecond)
	cancel()
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(delivered) != 1 {
		t.Fatalf("delivered %d events, want only evt_new1; log:\n%s", len(delivered), errb.String())
	}
	d := delivered[0]
	if !strings.Contains(d.Body, "evt_new1") {
		t.Fatalf("delivered %s", d.Body)
	}
	if err := webhook.Verify([]byte(d.Body), d.Header.Get("Yesly-Signature"), "whsec_local", webhook.DefaultTolerance, time.Now()); err != nil {
		t.Fatalf("signature: %v", err)
	}
}
