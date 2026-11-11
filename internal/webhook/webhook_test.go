package webhook

import (
	"strings"
	"testing"
	"time"
)

const (
	secret = "whsec_test"
	body   = `{"id":"evt_1","object":"event","type":"conversation.ended","data":{"conversation":{"id":"cnv_7Qp4Ns2Xk9Lm3D","outcome":{"code":"rescheduled","data":{"delivery_window":"2026-12-03T09:00/13:00"}}}}}`
)

var now = time.Unix(1794474252, 0)

func TestSignMatchesTheOtherLibraries(t *testing.T) {
	// The same vector as the Java and .NET SDK tests, computed with Python's hmac module.
	want := "t=1794474252,v1=f60bd9a59d63d6f2d01cbdcd9aec4781af68e4d1e915aaff424817ce69921f83"
	if got := Sign([]byte(body), secret, now); got != want {
		t.Fatalf("Sign = %s, want %s", got, want)
	}
}

func TestVerify(t *testing.T) {
	good := Sign([]byte(body), secret, now)
	other := Sign([]byte(body), "whsec_old", now)
	cases := []struct {
		name    string
		header  string
		payload string
		at      time.Time
		tol     time.Duration
		wantErr string
	}{
		{"valid", good, body, now, DefaultTolerance, ""},
		{"either signature during a roll", other + ",v1=" + strings.Split(good, "v1=")[1], body, now, DefaultTolerance, ""},
		{"changed payload", good, body + " ", now, DefaultTolerance, "matches"},
		{"wrong secret", other, body, now, DefaultTolerance, "matches"},
		{"too old", good, body, now.Add(301 * time.Second), DefaultTolerance, "tolerance"},
		{"from the future", good, body, now.Add(-301 * time.Second), DefaultTolerance, "tolerance"},
		{"tolerance off", good, body, now.Add(24 * time.Hour), 0, ""},
		{"missing", "", body, now, DefaultTolerance, "missing"},
		{"no v1", "t=1794474252", body, now, DefaultTolerance, "malformed"},
		{"bad t", "t=x,v1=00", body, now, DefaultTolerance, "malformed"},
		{"not hex", "t=1794474252,v1=zz", body, now, DefaultTolerance, "matches"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Verify([]byte(c.payload), c.header, secret, c.tol, c.at)
			if c.wantErr == "" && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if c.wantErr != "" && (err == nil || !strings.Contains(err.Error(), c.wantErr)) {
				t.Fatalf("error = %v, want one containing %q", err, c.wantErr)
			}
		})
	}
}
