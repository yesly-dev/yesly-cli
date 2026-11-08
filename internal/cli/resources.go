package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/yesly-dev/yesly-cli/internal/api"
	"github.com/yesly-dev/yesly-cli/internal/spec"
)

// kebab turns API names into command and flag names: webhook_endpoints -> webhook-endpoints.
func kebab(s string) string { return strings.ReplaceAll(s, "_", "-") }

// resourceCommands returns one command per resource, with one subcommand per operation.
func resourceCommands(app *App) []*cobra.Command {
	var out []*cobra.Command
	byName := map[string]*cobra.Command{}
	for i := range spec.Operations {
		op := &spec.Operations[i]
		parent := byName[op.Resource]
		if parent == nil {
			parent = &cobra.Command{
				Use:   kebab(op.Resource),
				Short: resourceShort(op.Resource),
				Long:  spec.Resources[op.Resource],
				Args:  noArgs,
				RunE:  func(cmd *cobra.Command, _ []string) error { return cmd.Help() },
			}
			byName[op.Resource] = parent
			out = append(out, parent)
		}
		parent.AddCommand(operationCommand(app, op))
	}
	return out
}

// resourceShort is the first sentence of the resource's description.
func resourceShort(resource string) string {
	d := spec.Resources[resource]
	if d == "" {
		if fallback, ok := resourceFallback[resource]; ok {
			return fallback
		}
		return strings.ReplaceAll(resource, "_", " ")
	}
	sentences := strings.SplitAfter(d, ". ")
	short := strings.TrimSpace(sentences[0])
	if short == "Deprecated." && len(sentences) > 1 {
		short = "(deprecated) " + strings.TrimSpace(sentences[1])
	}
	return strings.TrimSuffix(short, ".")
}

var resourceFallback = map[string]string{
	"imports": "Results of appointment imports",
}

func operationCommand(app *App, op *spec.Operation) *cobra.Command {
	use := kebab(op.Method)
	for _, p := range op.PathParams {
		use += " <" + p.Name + ">"
	}
	long := op.Summary + "."
	short := op.Summary
	if op.Deprecated {
		long = "Deprecated. " + long
		short = "(deprecated) " + short
	}
	if op.Description != "" {
		long += "\n\n" + op.Description
	}
	long += fmt.Sprintf("\n\nAPI: %s %s", op.HTTP, op.Path)
	if len(op.PathParams) > 0 {
		var ps []string
		for _, p := range op.PathParams {
			ps = append(ps, fmt.Sprintf("  <%s>  %s", p.Name, p.Description))
		}
		long += "\n\nArguments:\n" + strings.Join(ps, "\n")
	}
	if op.HasBody {
		long += "\n\nFlags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields."
	}

	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) != len(op.PathParams) {
				names := make([]string, len(op.PathParams))
				for i, p := range op.PathParams {
					names[i] = "<" + p.Name + ">"
				}
				if len(names) == 0 {
					return usagef("unexpected argument %q", args[0])
				}
				return usagef("expected %s", strings.Join(names, " "))
			}
			return nil
		},
	}
	cmd.Example = examples[op.Resource+"."+op.Method]

	f := cmd.Flags()
	query := bindFields(cmd, op.Query)
	body := bindFields(cmd, op.Body)
	var data, idempotencyKey string
	var all bool
	if op.HasBody {
		f.StringVarP(&data, "data", "d", "", "JSON body, inline or @file (@- for stdin)")
	}
	if op.Idempotent {
		f.StringVar(&idempotencyKey, "idempotency-key", "", "Idempotency-Key header (default: a new UUID, reused on retries)")
	}
	if op.List {
		f.BoolVar(&all, "all", false, "follow every page and print all items as one JSON array")
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		path := op.Path
		for i, p := range op.PathParams {
			path = strings.Replace(path, "{"+p.Name+"}", url.PathEscape(args[i]), 1)
		}
		q := url.Values{}
		if err := query.toQuery(app, cmd.Flags(), q); err != nil {
			return err
		}
		req := api.Request{Method: op.HTTP, Path: path, Query: q}
		if op.HasBody {
			obj, err := readJSONObject(app, data)
			if err != nil {
				return err
			}
			if err := body.toBody(app, cmd.Flags(), obj); err != nil {
				return err
			}
			req.Body = obj
		} else if op.HTTP == "POST" {
			// Like the SDKs, a POST always carries a JSON object.
			req.Body = map[string]any{}
		}
		if op.Idempotent {
			req.IdempotencyKey = idempotencyKey
			if req.IdempotencyKey == "" {
				req.IdempotencyKey = api.NewIdempotencyKey()
			}
		}
		if all {
			return app.listAll(cmd.Context(), req)
		}
		res, err := app.send(cmd.Context(), req)
		if err != nil || res == nil {
			return err
		}
		return app.printJSON(res.Body)
	}
	return cmd
}

// readJSONObject parses --data: inline JSON, @file or @- for stdin. Empty means {}.
func readJSONObject(app *App, data string) (map[string]any, error) {
	obj := map[string]any{}
	if data == "" {
		return obj, nil
	}
	b, err := readArg(app, data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &obj); err != nil {
		return nil, usagef("--data is not a JSON object: %v", err)
	}
	return obj, nil
}

// listAll follows starting_after until has_more is false and prints the items.
func (app *App) listAll(ctx context.Context, req api.Request) error {
	items := []json.RawMessage{}
	for {
		res, err := app.send(ctx, req)
		if err != nil || res == nil {
			return err
		}
		var page struct {
			Data    []json.RawMessage `json:"data"`
			HasMore bool              `json:"has_more"`
		}
		if err := json.Unmarshal(res.Body, &page); err != nil {
			return fmt.Errorf("reading the page: %w", err)
		}
		items = append(items, page.Data...)
		if !page.HasMore || len(page.Data) == 0 {
			break
		}
		var last struct {
			ID string `json:"id"`
		}
		if err := json.Unmarshal(page.Data[len(page.Data)-1], &last); err != nil || last.ID == "" {
			return fmt.Errorf("the last item of the page has no id")
		}
		q := url.Values{}
		for k, v := range req.Query {
			q[k] = v
		}
		q.Set("starting_after", last.ID)
		req.Query = q
	}
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return app.printJSON(b)
}

// fieldSet is the flags bound for a list of fields.
type fieldSet []spec.Field

func flagName(field string) string { return kebab(field) }

func flagUsage(fl spec.Field) string {
	u := fl.Description
	if fl.Kind == spec.Map {
		u = strings.TrimSpace(u + " Repeat as key=value.")
	}
	if fl.Kind == spec.JSON {
		u = strings.TrimSpace(u + " JSON, inline or @file.")
	}
	if len(fl.Enum) > 0 {
		u = strings.TrimSpace(u + " One of: " + strings.Join(fl.Enum, ", ") + ".")
	}
	if fl.Kind == spec.Bool && fl.Default != "" && !strings.Contains(u, "efault") {
		u = strings.TrimSpace(u + " Default " + fl.Default + ".")
	}
	if fl.Required {
		u = "(required) " + u
	}
	return u
}

func bindFields(cmd *cobra.Command, fields []spec.Field) fieldSet {
	f := cmd.Flags()
	for _, fl := range fields {
		name := flagName(fl.Name)
		usage := flagUsage(fl)
		switch fl.Kind {
		case spec.Integer:
			f.Int64(name, 0, usage)
		case spec.Number:
			f.Float64(name, 0, usage)
		case spec.Bool:
			f.Bool(name, false, usage)
		case spec.StringList:
			f.StringSlice(name, nil, usage)
		case spec.Map:
			f.StringArray(name, nil, usage)
		default:
			f.String(name, "", usage)
		}
		if fl.Kind == spec.Bool {
			f.Lookup(name).NoOptDefVal = "true"
		}
		if len(fl.Enum) > 0 {
			enum := fl.Enum
			_ = cmd.RegisterFlagCompletionFunc(name, func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
				return enum, cobra.ShellCompDirectiveNoFileComp
			})
		}
	}
	return fieldSet(fields)
}

// value reads a changed flag as the JSON value to send.
func value(fs *pflag.FlagSet, fl spec.Field, app *App) (any, error) {
	name := flagName(fl.Name)
	switch fl.Kind {
	case spec.Integer:
		return fs.GetInt64(name)
	case spec.Number:
		return fs.GetFloat64(name)
	case spec.Bool:
		return fs.GetBool(name)
	case spec.StringList:
		v, err := fs.GetStringSlice(name)
		if err != nil {
			return nil, err
		}
		if err := checkEnum(fl, v...); err != nil {
			return nil, err
		}
		return v, nil
	case spec.Map:
		pairs, err := fs.GetStringArray(name)
		if err != nil {
			return nil, err
		}
		m := map[string]any{}
		for _, p := range pairs {
			k, v, ok := strings.Cut(p, "=")
			if !ok || k == "" {
				return nil, usagef("--%s takes key=value, got %q", name, p)
			}
			m[k] = v
		}
		return m, nil
	case spec.JSON:
		s, _ := fs.GetString(name)
		b, err := readArg(app, s)
		if err != nil {
			return nil, err
		}
		var v any
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, usagef("--%s is not valid JSON: %v", name, err)
		}
		return v, nil
	default:
		s, err := fs.GetString(name)
		if err != nil {
			return nil, err
		}
		if err := checkEnum(fl, s); err != nil {
			return nil, err
		}
		return s, nil
	}
}

func checkEnum(fl spec.Field, vals ...string) error {
	if len(fl.Enum) == 0 {
		return nil
	}
	for _, v := range vals {
		ok := false
		for _, e := range fl.Enum {
			if v == e {
				ok = true
				break
			}
		}
		if !ok {
			return usagef("--%s must be one of %s, got %q", flagName(fl.Name), strings.Join(fl.Enum, ", "), v)
		}
	}
	return nil
}

func (fsx fieldSet) toQuery(app *App, fs *pflag.FlagSet, q url.Values) error {
	for _, fl := range fsx {
		if !fs.Changed(flagName(fl.Name)) {
			continue
		}
		v, err := value(fs, fl, app)
		if err != nil {
			return err
		}
		switch t := v.(type) {
		case []string:
			for _, s := range t {
				q.Add(fl.Name, s)
			}
		case map[string]any:
			for k, s := range t {
				q.Add(fl.Name+"["+k+"]", fmt.Sprint(s))
			}
		default:
			q.Set(fl.Name, fmt.Sprint(t))
		}
	}
	return nil
}

func (fsx fieldSet) toBody(app *App, fs *pflag.FlagSet, obj map[string]any) error {
	for _, fl := range fsx {
		if !fs.Changed(flagName(fl.Name)) {
			continue
		}
		v, err := value(fs, fl, app)
		if err != nil {
			return err
		}
		setPath(obj, strings.Split(fl.Name, "."), v)
	}
	return nil
}

// setPath sets obj[a][b]... = v, creating objects on the way and replacing
// anything that is not one.
func setPath(obj map[string]any, path []string, v any) {
	for _, k := range path[:len(path)-1] {
		next, ok := obj[k].(map[string]any)
		if !ok {
			next = map[string]any{}
			obj[k] = next
		}
		obj = next
	}
	obj[path[len(path)-1]] = v
}
