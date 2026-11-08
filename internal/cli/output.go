package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/yesly-dev/yesly-cli/internal/api"
	"golang.org/x/term"
)

// printJSON writes an API answer indented with two spaces, or on one line
// with --compact, keeping the API's key order. Anything that is not JSON is
// written as it came.
func (app *App) printJSON(b []byte) error {
	if len(bytes.TrimSpace(b)) == 0 {
		return nil
	}
	var buf bytes.Buffer
	var err error
	if app.compact {
		err = json.Compact(&buf, b)
	} else {
		err = json.Indent(&buf, b, "", "  ")
	}
	if err != nil {
		_, err = app.Out.Write(b)
		return err
	}
	buf.WriteByte('\n')
	_, err = app.Out.Write(buf.Bytes())
	return err
}

// printDryRun shows what would be sent, without the API key.
func (app *App) printDryRun(c *credentials, r api.Request) error {
	target := c.baseURL + r.Path
	if len(r.Query) > 0 {
		target += "?" + r.Query.Encode()
	}
	fmt.Fprintf(app.Out, "%s %s\n", r.Method, target)
	fmt.Fprintf(app.Out, "Yesly-Version: %s\n", app.apiVersion)
	if r.IdempotencyKey != "" {
		fmt.Fprintf(app.Out, "Idempotency-Key: %s\n", r.IdempotencyKey)
	}
	if r.Body != nil {
		b, err := json.Marshal(r.Body)
		if err != nil {
			return err
		}
		fmt.Fprintln(app.Out)
		return app.printJSON(b)
	}
	return nil
}

// readArg returns the bytes of a value that may be @file or @- for stdin.
func readArg(app *App, s string) ([]byte, error) {
	if s == "@-" {
		return io.ReadAll(app.In)
	}
	if strings.HasPrefix(s, "@") {
		b, err := os.ReadFile(s[1:])
		if err != nil {
			return nil, usageError{err}
		}
		return b, nil
	}
	return []byte(s), nil
}

func stdinIsTerminal() bool { return term.IsTerminal(int(os.Stdin.Fd())) }

func readSecret() (string, error) {
	b, err := term.ReadPassword(int(os.Stdin.Fd()))
	return string(b), err
}

// mask shows the key's prefix and last four characters.
func mask(key string) string {
	if len(key) <= 12 {
		return strings.Repeat("•", len(key))
	}
	prefix := key[:strings.LastIndex(key[:min(len(key), 9)], "_")+1]
	return prefix + "…" + key[len(key)-4:]
}
