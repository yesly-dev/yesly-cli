package api

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func newClient(url string) (*Client, *[]time.Duration) {
	var slept []time.Duration
	return &Client{
		APIKey:     "ysk_test_abc",
		BaseURL:    url,
		APIVersion: "2026-10-01",
		UserAgent:  "yesly-cli/test",
		MaxRetries: 2,
		Sleep:      func(d time.Duration) { slept = append(slept, d) },
	}, &slept
}

func TestHeadersAndBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer ysk_test_abc" {
			t.Errorf("Authorization = %q", got)
		}
		if got := r.Header.Get("Yesly-Version"); got != "2026-10-01" {
			t.Errorf("Yesly-Version = %q", got)
		}
		if got := r.Header.Get("User-Agent"); got != "yesly-cli/test" {
			t.Errorf("User-Agent = %q", got)
		}
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Errorf("Content-Type = %q", got)
		}
		if r.URL.RawQuery != "a=1" {
			t.Errorf("query = %q", r.URL.RawQuery)
		}
		b, _ := io.ReadAll(r.Body)
		if string(b) != `{"reason":"x"}` {
			t.Errorf("body = %s", b)
		}
		w.Header().Set("Request-Id", "req_1")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()
	c, _ := newClient(srv.URL)
	res, err := c.Do(context.Background(), Request{Method: "POST", Path: "/x", Query: url.Values{"a": {"1"}}, Body: map[string]string{"reason": "x"}})
	if err != nil {
		t.Fatal(err)
	}
	if string(res.Body) != `{"ok":true}` || res.RequestID != "req_1" {
		t.Fatalf("response = %+v", res)
	}
}

func TestRetriesKeepTheIdempotencyKey(t *testing.T) {
	var keys []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys = append(keys, r.Header.Get("Idempotency-Key"))
		if len(keys) == 1 {
			w.WriteHeader(503)
			return
		}
		if len(keys) == 2 {
			w.Header().Set("Retry-After", "3")
			w.WriteHeader(429)
			return
		}
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	c, slept := newClient(srv.URL)
	if _, err := c.Do(context.Background(), Request{Method: "POST", Path: "/x", Body: map[string]any{}, IdempotencyKey: "key-1"}); err != nil {
		t.Fatal(err)
	}
	if strings.Join(keys, ",") != "key-1,key-1,key-1" {
		t.Fatalf("keys = %v", keys)
	}
	if len(*slept) != 2 || (*slept)[1] != 3*time.Second {
		t.Fatalf("slept = %v, want a backoff then Retry-After's 3s", *slept)
	}
	if d := (*slept)[0]; d < 250*time.Millisecond || d > 500*time.Millisecond {
		t.Fatalf("first backoff = %v", d)
	}
}

func TestErrorsAfterTheLastRetry(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Request-Id", "req_9")
		w.WriteHeader(500)
		_, _ = w.Write([]byte(`{"error":{"type":"api_error","message":"Boom."}}`))
	}))
	defer srv.Close()
	c, _ := newClient(srv.URL)
	_, err := c.Do(context.Background(), Request{Method: "GET", Path: "/x"})
	var e *Error
	if !errors.As(err, &e) || e.Status != 500 || e.Message != "Boom." || e.RequestID != "req_9" {
		t.Fatalf("err = %#v", err)
	}
	if calls != 3 {
		t.Fatalf("calls = %d, want 3", calls)
	}
}

func TestClientErrorsAreNotRetried(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(422)
		_, _ = w.Write([]byte(`{"error":{"type":"invalid_request","code":"phone_invalid","param":"customer.phone","message":"Bad phone.","request_id":"req_2"}}`))
	}))
	defer srv.Close()
	c, _ := newClient(srv.URL)
	_, err := c.Do(context.Background(), Request{Method: "POST", Path: "/x", Body: map[string]any{}})
	var e *Error
	if !errors.As(err, &e) || e.Code != "phone_invalid" || e.Param != "customer.phone" || e.RequestID != "req_2" {
		t.Fatalf("err = %#v", err)
	}
	if calls != 1 {
		t.Fatalf("calls = %d", calls)
	}
	want := "Bad phone. (status 422, type invalid_request, code phone_invalid, param customer.phone, request req_2)"
	if err.Error() != want {
		t.Fatalf("message = %q", err.Error())
	}
}

func TestRedirectsAreNotFollowed(t *testing.T) {
	var other bool
	target := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { other = true }))
	defer target.Close()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusFound)
	}))
	defer srv.Close()
	c, _ := newClient(srv.URL)
	_, err := c.Do(context.Background(), Request{Method: "GET", Path: "/x"})
	var e *Error
	if !errors.As(err, &e) || e.Status != 302 || other {
		t.Fatalf("err = %v, followed = %v", err, other)
	}
}

func TestConnectionErrors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	srv.Close()
	c, slept := newClient(srv.URL)
	_, err := c.Do(context.Background(), Request{Method: "GET", Path: "/x"})
	var ce *ConnectionError
	if !errors.As(err, &ce) || len(*slept) != 2 {
		t.Fatalf("err = %v, retries = %d", err, len(*slept))
	}
}

func TestNoKey(t *testing.T) {
	c := &Client{BaseURL: "https://example.com"}
	if _, err := c.Do(context.Background(), Request{Method: "GET", Path: "/"}); err == nil || !strings.Contains(err.Error(), "no API key") {
		t.Fatalf("err = %v", err)
	}
}

func TestCheckBaseURL(t *testing.T) {
	for raw, ok := range map[string]bool{
		"https://api.yesly.dev/v1/": true,
		"http://localhost:4010":     true,
		"http://127.0.0.1:4010":     true,
		"http://example.com":        false,
		"ftp://example.com":         false,
		"not a url":                 false,
	} {
		got, err := CheckBaseURL(raw)
		if ok != (err == nil) {
			t.Errorf("CheckBaseURL(%q) error = %v", raw, err)
		}
		if ok && strings.HasSuffix(got, "/") {
			t.Errorf("CheckBaseURL(%q) = %q keeps the slash", raw, got)
		}
	}
}

func TestModeAndKeys(t *testing.T) {
	if Mode("ysk_live_x") != "live" || Mode("ysk_test_x") != "test" || Mode("sk_x") != "" {
		t.Fatal("Mode")
	}
	k := NewIdempotencyKey()
	if len(k) != 36 || k[14] != '4' || NewIdempotencyKey() == k {
		t.Fatalf("NewIdempotencyKey = %s", k)
	}
}
