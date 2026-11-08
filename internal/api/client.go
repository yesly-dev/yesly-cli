// Package api sends requests to the Yesly API with the same behaviour as the
// official SDKs: bearer auth, Yesly-Version, retries with backoff and
// Retry-After, and an Idempotency-Key reused across retries.
package api

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	mrand "math/rand/v2"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Client holds what every request needs.
type Client struct {
	APIKey     string
	BaseURL    string
	APIVersion string
	UserAgent  string
	Timeout    time.Duration
	MaxRetries int
	HTTP       *http.Client
	// Sleep waits between retries; tests replace it.
	Sleep func(time.Duration)
}

// Request is one call.
type Request struct {
	Method string
	Path   string
	Query  url.Values
	// Body is sent as JSON when not nil. Use json.RawMessage for bytes you already have.
	Body any
	// IdempotencyKey is sent as Idempotency-Key when not empty.
	IdempotencyKey string
}

// Response is a successful answer.
type Response struct {
	Status    int
	Body      []byte
	RequestID string
}

// Mode is "live" or "test" from the key prefix, or "" when unknown.
func Mode(key string) string {
	switch {
	case strings.HasPrefix(key, "ysk_live_"):
		return "live"
	case strings.HasPrefix(key, "ysk_test_"):
		return "test"
	}
	return ""
}

// CheckBaseURL refuses anything but HTTPS, except http to this machine,
// because the API key goes with every request.
func CheckBaseURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "", fmt.Errorf("the base URL %q is not a valid URL", raw)
	}
	loopback := u.Hostname() == "localhost" || u.Hostname() == "127.0.0.1" || u.Hostname() == "::1"
	if u.Scheme != "https" && !(u.Scheme == "http" && loopback) {
		return "", errors.New("the base URL must use https:// (http:// is only allowed for localhost)")
	}
	return strings.TrimRight(raw, "/"), nil
}

// NewIdempotencyKey returns a random UUID v4.
func NewIdempotencyKey() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = b[6]&0x0f | 0x40
	b[8] = b[8]&0x3f | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// DefaultUserAgent is yesly-cli/<version> go/<version> (<os>; <arch>).
func DefaultUserAgent(version string) string {
	return fmt.Sprintf("yesly-cli/%s %s (%s; %s)", version, strings.Replace(runtime.Version(), "go", "go/", 1), runtime.GOOS, runtime.GOARCH)
}

func retryable(status int) bool {
	return status == 408 || status == 429 || status >= 500
}

// Do sends the request, retrying network errors, timeouts, 408, 429 and 5xx.
// A non-2xx answer comes back as *Error.
func (c *Client) Do(ctx context.Context, r Request) (*Response, error) {
	if c.APIKey == "" {
		return nil, errors.New("no API key: run `yesly login`, set YESLY_API_KEY or pass --api-key")
	}
	target := c.BaseURL + r.Path
	if len(r.Query) > 0 {
		target += "?" + r.Query.Encode()
	}
	var body []byte
	if r.Body != nil {
		var err error
		if body, err = json.Marshal(r.Body); err != nil {
			return nil, err
		}
	}
	httpClient := c.HTTP
	if httpClient == nil {
		// Redirects are not followed, so the API key is never sent to another URL.
		httpClient = &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	}
	sleep := c.Sleep
	if sleep == nil {
		sleep = time.Sleep
	}
	timeout := c.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	for attempt := 0; ; attempt++ {
		res, err := c.attempt(ctx, httpClient, timeout, r, target, body)
		if err != nil {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			if attempt < c.MaxRetries {
				sleep(backoff(attempt))
				continue
			}
			var ne net.Error
			if errors.As(err, &ne) && ne.Timeout() {
				return nil, &ConnectionError{Message: fmt.Sprintf("the request timed out after %s", timeout), Err: err}
			}
			return nil, &ConnectionError{Message: "could not reach the Yesly API", Err: err}
		}
		if res.Status >= 200 && res.Status < 300 {
			return &Response{Status: res.Status, Body: res.Body, RequestID: res.Header.Get("Request-Id")}, nil
		}
		if retryable(res.Status) && attempt < c.MaxRetries {
			if d, ok := retryAfter(res.Header); ok {
				sleep(d)
			} else {
				sleep(backoff(attempt))
			}
			continue
		}
		return nil, errorFrom(res.Status, res.Body, res.Header)
	}
}

type rawResponse struct {
	Status int
	Header http.Header
	Body   []byte
}

func (c *Client) attempt(ctx context.Context, hc *http.Client, timeout time.Duration, r Request, target string, body []byte) (*rawResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, r.Method, target, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Yesly-Version", c.APIVersion)
	req.Header.Set("User-Agent", c.UserAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if r.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", r.IdempotencyKey)
	}
	res, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &rawResponse{Status: res.StatusCode, Header: res.Header, Body: b}, nil
}

// retryAfter reads Retry-After in seconds, up to a minute.
func retryAfter(h http.Header) (time.Duration, bool) {
	s, err := strconv.ParseFloat(h.Get("Retry-After"), 64)
	if err != nil || s < 0 || s > 60 || math.IsNaN(s) {
		return 0, false
	}
	return time.Duration(s * float64(time.Second)), true
}

// backoff is 0.5 s doubled per attempt up to 8 s, with jitter in its upper half.
func backoff(attempt int) time.Duration {
	base := math.Min(500*math.Pow(2, float64(attempt)), 8000)
	ms := base/2 + mrand.Float64()*(base/2)
	return time.Duration(ms * float64(time.Millisecond))
}
