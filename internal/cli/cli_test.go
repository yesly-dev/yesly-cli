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
	r := run(t, "", "conversations", "create", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"--data", `{"agent":"agt_old","start_at":"2026-12-01T10:00:00+01:00"}`,
		"--agent", "agt_8Hd2kQ4mZ1LpXy", "--contact", `{"external_id":"customer-3391","phone":"+34600000001","name":"Marta"}`,
		"--channels", "whatsapp,call", "--context", `{"order":"8841","delivery_window":"2026-12-02T09:00/13:00"}`)
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	c := api.last(t)
	if c.Method != "POST" || c.Path != "/conversations" {
		t.Fatalf("%s %s", c.Method, c.Path)
	}
	var got map[string]any
	_ = json.Unmarshal([]byte(c.Body), &got)
	want := map[string]any{
		"agent":    "agt_8Hd2kQ4mZ1LpXy",
		"contact":  map[string]any{"external_id": "customer-3391", "name": "Marta", "phone": "+34600000001"},
		"channels": []any{"whatsapp", "call"},
		"start_at": "2026-12-01T10:00:00+01:00",
		"context":  map[string]any{"delivery_window": "2026-12-02T09:00/13:00", "order": "8841"},
	}
	wantJSON, _ := json.Marshal(want)
	gotJSON, _ := json.Marshal(got)
	if string(wantJSON) != string(gotJSON) {
		t.Fatalf("body = %s\nwant %s", gotJSON, wantJSON)
	}
	if len(c.Header.Get("Idempotency-Key")) != 36 {
		t.Fatal("creating a conversation should send an Idempotency-Key")
	}
	if r.out != "{\n  \"object\": \"ok\"\n}\n" {
		t.Fatalf("output = %q", r.out)
	}
}

func TestJSONFlagTakesAPlainID(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "messages", "create", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"--agent", "agt_8Hd2kQ4mZ1LpXy", "--contact", "ctc_3Kf9Lq2Wm7Rt4B", "--channel", "sms", "--text", "Hola")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	if b := api.last(t).Body; b != `{"agent":"agt_8Hd2kQ4mZ1LpXy","channel":"sms","contact":"ctc_3Kf9Lq2Wm7Rt4B","text":"Hola"}` {
		t.Fatalf("body = %s", b)
	}
	if r := run(t, "", "messages", "create", "--api-key", "ysk_test_x", "--dry-run",
		"--agent", "agt_1", "--contact", `{"phone":`, "--channel", "sms"); r.code != 2 {
		t.Fatalf("broken JSON object: exit %d", r.code)
	}
}

func TestTypedFlags(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "agents", "update", "agt_8Hd2kQ4mZ1LpXy", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"--limits.messages-per-conversation", "21", "--recording.enabled=false")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	c := api.last(t)
	if c.Method != "PATCH" || c.Path != "/agents/agt_8Hd2kQ4mZ1LpXy" || c.Body != `{"limits":{"messages_per_conversation":21},"recording":{"enabled":false}}` {
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
	run(t, "", "conversations", "cancel", "cnv_1", "--api-key", "ysk_test_x", "--base-url", api.URL, "--idempotency-key", "mine")
	if got := api.last(t).Header.Get("Idempotency-Key"); got != "mine" {
		t.Fatalf("Idempotency-Key = %q", got)
	}
}

func TestQueryFlags(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "conversations", "list", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"--status", "no_reply", "--agent", "agt_8Hd2kQ4mZ1LpXy", "--limit", "5")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	if q := api.last(t).Query; q != "agent=agt_8Hd2kQ4mZ1LpXy&limit=5&status=no_reply" {
		t.Fatalf("query = %s", q)
	}
}

func TestPathArgumentsAreEscaped(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	run(t, "", "contacts", "update", "a/b", "--api-key", "ysk_test_x", "--base-url", api.URL, "--name", "X")
	if p := api.last(t); !strings.HasSuffix(p.Path, "/contacts/a/b") {
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
	r := run(t, "", "conversations", "cancel", "cnv_1", "--api-key", "ysk_live_x", "--base-url", api.URL)
	if r.code != 2 || !strings.Contains(r.err, "--live") || len(api.calls) != 0 {
		t.Fatalf("exit %d, err %q, calls %d", r.code, r.err, len(api.calls))
	}
	if r := run(t, "", "conversations", "list", "--api-key", "ysk_live_x", "--base-url", api.URL); r.code != 0 {
		t.Fatalf("GET with a live key: %s", r.err)
	}
	if r := run(t, "", "conversations", "cancel", "cnv_1", "--live", "--api-key", "ysk_live_x", "--base-url", api.URL); r.code != 0 {
		t.Fatalf("--live: %s", r.err)
	}
}

func TestAPIErrorsExitWithOne(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	api.handler = func(w http.ResponseWriter, _ *http.Request, _ string) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte(`{"error":{"type":"invalid_request","code":"resource_missing","message":"No such conversation.","request_id":"req_1"}}`))
	}
	r := run(t, "", "conversations", "retrieve", "cnv_x", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 1 || r.err != "Error: No such conversation. (status 404, type invalid_request, code resource_missing, request req_1)\n" {
		t.Fatalf("exit %d, err %q", r.code, r.err)
	}
}

func TestUsageErrorsExitWithTwo(t *testing.T) {
	setup(t)
	for _, args := range [][]string{
		{"conversations", "retrieve"},
		{"conversations", "list", "extra"},
		{"conversations", "list", "--status", "nope"},
		{"conversations", "create", "--nope"},
		{"conversations", "create", "--data", "[1]"},
		{"conversations", "create", "--channels", "fax"},
		{"contacts", "create", "--metadata", "novalue"},
		{"conversations", "list", "--base-url", "http://example.com"},
	} {
		r := run(t, "", append(args, "--api-key", "ysk_test_x", "--dry-run")...)
		if r.code != 2 {
			t.Errorf("%v: exit %d, err %q", args, r.code, r.err)
		}
	}
}

func TestDryRunNeedsNoKey(t *testing.T) {
	setup(t)
	r := run(t, "", "opt-outs", "create", "--channel", "sms", "--address", "+34600000001", "--dry-run", "--idempotency-key", "k")
	want := "POST https://api.yesly.dev/v1/opt_outs\nYesly-Version: 2026-11-11\nIdempotency-Key: k\n\n{\n  \"address\": \"+34600000001\",\n  \"channel\": \"sms\"\n}\n"
	if r.code != 0 || r.out != want {
		t.Fatalf("exit %d, out %q, err %q", r.code, r.out, r.err)
	}
}

func TestJSONFlagFromFile(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	file := filepath.Join(t.TempDir(), "a.json")
	_ = os.WriteFile(file, []byte(`[{"name":"Lunes","days":[1],"from":"09:00","to":"14:00"}]`), 0o600)
	r := run(t, "", "agents", "update", "agt_1", "--contact-hours", "@"+file, "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 0 || api.last(t).Body != `{"contact_hours":[{"days":[1],"from":"09:00","name":"Lunes","to":"14:00"}]}` {
		t.Fatalf("exit %d, err %q, body %s", r.code, r.err, api.last(t).Body)
	}
	r = run(t, `{"name":"Marta"}`, "contacts", "update", "ctc_1", "--data", "@-", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if r.code != 0 || api.last(t).Body != `{"name":"Marta"}` {
		t.Fatalf("stdin: exit %d, body %s", r.code, api.last(t).Body)
	}
}

func TestAPICommand(t *testing.T) {
	setup(t)
	api := newFakeAPI(t)
	r := run(t, "", "api", "post", "conversations", "--api-key", "ysk_test_x", "--base-url", api.URL,
		"-f", "agent=agt_1", "-f", "contact.name=Marta", "-f", "contact[phone]=+34600000001",
		"-f", "channels[]=sms", "-f", "channels[]=call", "-F", "context={\"order\":\"8841\"}")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.err)
	}
	c := api.last(t)
	var got map[string]any
	_ = json.Unmarshal([]byte(c.Body), &got)
	gotJSON, _ := json.Marshal(got)
	if string(gotJSON) != `{"agent":"agt_1","channels":["sms","call"],"contact":{"name":"Marta","phone":"+34600000001"},"context":{"order":"8841"}}` {
		t.Fatalf("body = %s", gotJSON)
	}
	if c.Header.Get("Idempotency-Key") == "" {
		t.Fatal("POST /conversations should get an Idempotency-Key")
	}

	run(t, "", "api", "POST", "/conversations/cnv_1/cancel", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if api.last(t).Header.Get("Idempotency-Key") == "" {
		t.Fatal("cancel should get an Idempotency-Key")
	}
	run(t, "", "api", "POST", "/not_an_operation", "--api-key", "ysk_test_x", "--base-url", api.URL)
	if api.last(t).Header.Get("Idempotency-Key") != "" {
		t.Fatal("a POST that is not an idempotent operation takes no Idempotency-Key")
	}

	run(t, "", "api", "GET", "/conversations?status=no_reply", "-F", "limit=5", "-f", "metadata[route_id]=BIO", "--api-key", "ysk_test_x", "--base-url", api.URL)
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
	body := `{"id":"evt_1","object":"event","type":"conversation.ended","data":{"conversation":{"id":"cnv_7Qp4Ns2Xk9Lm3D","outcome":{"code":"rescheduled","data":{"delivery_window":"2026-12-03T09:00/13:00"}}}}}`
	r := run(t, body, "webhooks", "sign", "--secret", "whsec_test")
	want := "t=1794474252,v1=f60bd9a59d63d6f2d01cbdcd9aec4781af68e4d1e915aaff424817ce69921f83\n"
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
		old := `{"id":"evt_old","type":"conversation.created","created":"2026-11-08T10:00:00Z"}`
		if polls == 1 {
			_, _ = w.Write([]byte(`{"data":[` + old + `],"has_more":false}`))
			return
		}
		_, _ = w.Write([]byte(`{"data":[{"id":"evt_new2","type":"message.sent","created":"2026-11-08T10:00:02Z"},` +
			`{"id":"evt_new1","type":"conversation.ended","created":"2026-11-08T10:00:01Z"},` + old + `],"has_more":false}`))
	}

	var errb bytes.Buffer
	app := &App{In: strings.NewReader(""), Out: io.Discard, Err: &errb, Now: time.Now, Sleep: func(time.Duration) {}}
	root := NewRoot(app)
	root.SetArgs([]string{"listen", "--api-key", "ysk_test_x", "--base-url", api.URL, "--forward-to", local.URL,
		"--secret", "whsec_local", "--interval", "10ms", "--events", "conversation.ended"})
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
