package cli

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yesly-dev/yesly-cli/internal/api"
	"github.com/yesly-dev/yesly-cli/internal/spec"
)

func apiCommand(app *App) *cobra.Command {
	var raw, typed []string
	var data, idempotencyKey string
	var all bool
	cmd := &cobra.Command{
		Use:   "api <method> <path>",
		Short: "Send any request to the API",
		Long: `Sends a request to any path of the API, relative to the base URL, with
your key, Yesly-Version, retries and, for POSTs the API makes idempotent, an
Idempotency-Key.

-f key=value adds a string; -F key=value adds a typed value (true, false,
null, numbers and JSON stay as they are, @file reads a file). For GET and
DELETE they become query parameters; otherwise they go in the JSON body,
where "customer.name" or "customer[name]" nests and "channels[]" appends.`,
		Example: `  yesly api GET /appointments -f status=no_reply -F limit=5
  yesly api POST /appointments -f external_id=ORD-1 -f customer.name=Martha \
    -f customer.phone=+34600000001 -f slot.start=2026-11-13T09:00:00+01:00
  yesly api PATCH /appointments/apt_… --data @changes.json`,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) != 2 {
				return usagef("expected <method> <path>")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			method := strings.ToUpper(args[0])
			path := "/" + strings.TrimPrefix(args[1], "/")
			req := api.Request{Method: method, Path: path}
			if u, err := url.Parse(path); err == nil && u.RawQuery != "" {
				req.Path = u.Path
				req.Query = u.Query()
			}
			fields, err := parseFields(app, raw, typed)
			if err != nil {
				return err
			}
			if method == "GET" || method == "DELETE" {
				if data != "" {
					return usagef("--data is only for requests with a body")
				}
				if req.Query == nil {
					req.Query = url.Values{}
				}
				for _, f := range fields {
					switch v := f.value.(type) {
					case nil:
					case string:
						req.Query.Add(f.key, v)
					default:
						b, _ := json.Marshal(v)
						req.Query.Add(f.key, strings.Trim(string(b), `"`))
					}
				}
			} else {
				obj, err := readJSONObject(app, data)
				if err != nil {
					return err
				}
				for _, f := range fields {
					if err := setKey(obj, f.key, f.value); err != nil {
						return err
					}
				}
				req.Body = obj
			}
			if idempotencyKey != "" {
				req.IdempotencyKey = idempotencyKey
			} else if method == "POST" && idempotentPath(req.Path) {
				req.IdempotencyKey = api.NewIdempotencyKey()
			}
			if all {
				return app.listAll(cmd.Context(), req)
			}
			res, err := app.send(cmd.Context(), req)
			if err != nil || res == nil {
				return err
			}
			return app.printJSON(res.Body)
		},
	}
	f := cmd.Flags()
	f.StringArrayVarP(&raw, "field", "f", nil, "add a string field, key=value")
	f.StringArrayVarP(&typed, "typed-field", "F", nil, "add a typed field, key=value")
	f.StringVarP(&data, "data", "d", "", "JSON body, inline or @file (@- for stdin)")
	f.StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency-Key header")
	f.BoolVar(&all, "all", false, "follow every page of a list and print all items as one JSON array")
	return cmd
}

type field struct {
	key   string
	value any
}

func parseFields(app *App, raw, typed []string) ([]field, error) {
	var out []field
	for _, s := range raw {
		k, v, ok := strings.Cut(s, "=")
		if !ok || k == "" {
			return nil, usagef("-f takes key=value, got %q", s)
		}
		out = append(out, field{k, v})
	}
	for _, s := range typed {
		k, v, ok := strings.Cut(s, "=")
		if !ok || k == "" {
			return nil, usagef("-F takes key=value, got %q", s)
		}
		val, err := typedValue(app, v)
		if err != nil {
			return nil, err
		}
		out = append(out, field{k, val})
	}
	return out, nil
}

func typedValue(app *App, v string) (any, error) {
	switch v {
	case "true":
		return true, nil
	case "false":
		return false, nil
	case "null":
		return nil, nil
	}
	if strings.HasPrefix(v, "@") {
		b, err := readArg(app, v)
		if err != nil {
			return nil, err
		}
		return strings.TrimRight(string(b), "\n"), nil
	}
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		return n, nil
	}
	if n, err := strconv.ParseFloat(v, 64); err == nil {
		return n, nil
	}
	if strings.HasPrefix(v, "{") || strings.HasPrefix(v, "[") {
		var j any
		if err := json.Unmarshal([]byte(v), &j); err == nil {
			return j, nil
		}
	}
	return v, nil
}

// setKey sets a key such as customer.name, customer[name] or channels[] in obj.
func setKey(obj map[string]any, key string, v any) error {
	parts := splitKey(key)
	cur := obj
	for i, p := range parts {
		last := i == len(parts)-1
		if last {
			cur[p] = v
			return nil
		}
		if parts[i+1] == "" {
			// p[] appends to an array; nothing may follow.
			if i+1 != len(parts)-1 {
				return usagef("%q: [] must come last", key)
			}
			arr, _ := cur[p].([]any)
			cur[p] = append(arr, v)
			return nil
		}
		next, ok := cur[p].(map[string]any)
		if !ok {
			next = map[string]any{}
			cur[p] = next
		}
		cur = next
	}
	return nil
}

var keyPart = regexp.MustCompile(`\[([^\]]*)\]`)

func splitKey(key string) []string {
	head := key
	var rest []string
	if i := strings.Index(key, "["); i > 0 {
		head = key[:i]
		for _, m := range keyPart.FindAllStringSubmatch(key[i:], -1) {
			rest = append(rest, m[1])
		}
	}
	return append(strings.Split(head, "."), rest...)
}

// idempotentPath reports whether a POST to path is an operation that takes an Idempotency-Key.
func idempotentPath(path string) bool {
	for _, op := range spec.Operations {
		if op.HTTP == "POST" && op.Idempotent && matchPath(op.Path, path) {
			return true
		}
	}
	return false
}

func matchPath(template, path string) bool {
	t := strings.Split(strings.Trim(template, "/"), "/")
	p := strings.Split(strings.Trim(path, "/"), "/")
	if len(t) != len(p) {
		return false
	}
	for i := range t {
		if strings.HasPrefix(t[i], "{") {
			continue
		}
		if t[i] != p[i] {
			return false
		}
	}
	return true
}
