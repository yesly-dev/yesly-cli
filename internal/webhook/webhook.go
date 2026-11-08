// Package webhook signs and verifies Yesly-Signature headers, the same way
// the API and the SDKs do: t=<unix>,v1=<hex HMAC-SHA256 of "t.body">.
package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DefaultTolerance is how far a delivery's timestamp may be from now.
const DefaultTolerance = 300 * time.Second

// Sign returns a Yesly-Signature value for payload at timestamp.
func Sign(payload []byte, secret string, timestamp time.Time) string {
	t := timestamp.Unix()
	return fmt.Sprintf("t=%d,v1=%s", t, hex.EncodeToString(mac(payload, secret, t)))
}

// Verify checks header against payload and secret. A tolerance of 0 turns
// the timestamp check off. Any v1 value may match, as during a secret roll.
func Verify(payload []byte, header, secret string, tolerance time.Duration, now time.Time) error {
	if header == "" {
		return errors.New("the Yesly-Signature header is missing")
	}
	var t int64
	var haveT bool
	var sigs []string
	for _, part := range strings.Split(header, ",") {
		k, v, ok := strings.Cut(part, "=")
		if !ok {
			continue
		}
		switch strings.TrimSpace(k) {
		case "t":
			n, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
			if err != nil {
				return errors.New("the Yesly-Signature header is malformed")
			}
			t, haveT = n, true
		case "v1":
			sigs = append(sigs, strings.TrimSpace(v))
		}
	}
	if !haveT || len(sigs) == 0 {
		return errors.New("the Yesly-Signature header is malformed")
	}
	if tolerance > 0 {
		d := now.Sub(time.Unix(t, 0))
		if d < 0 {
			d = -d
		}
		if d > tolerance {
			return errors.New("the Yesly-Signature timestamp is outside the tolerance")
		}
	}
	want := mac(payload, secret, t)
	for _, s := range sigs {
		got, err := hex.DecodeString(s)
		if err == nil && hmac.Equal(got, want) {
			return nil
		}
	}
	return errors.New("no signature in Yesly-Signature matches the payload")
}

func mac(payload []byte, secret string, t int64) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(strconv.FormatInt(t, 10) + "."))
	h.Write(payload)
	return h.Sum(nil)
}
