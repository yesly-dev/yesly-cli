// Package cli builds the yesly command tree: one command per API operation,
// from the generated table in internal/spec, plus login, api, listen and
// webhooks.
package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yesly-dev/yesly-cli/internal/api"
	"github.com/yesly-dev/yesly-cli/internal/config"
	"github.com/yesly-dev/yesly-cli/internal/spec"
	"github.com/yesly-dev/yesly-cli/internal/version"
)

// App is the CLI's state: where it writes, and the global flags.
type App struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
	// HTTP is the client for API calls and forwarded deliveries; nil for the default.
	HTTP *http.Client
	// Now and Sleep are replaced in tests.
	Now   func() time.Time
	Sleep func(time.Duration)
	// IsTerminal reports whether In is a terminal, for the login prompt.
	IsTerminal func() bool
	// ReadSecret reads a line from the terminal without echo.
	ReadSecret func() (string, error)

	apiKey     string
	profile    string
	baseURL    string
	apiVersion string
	maxRetries int
	timeout    time.Duration
	compact    bool
	live       bool
	dryRun     bool
}

// usageError is a mistake in the command line; it exits with status 2.
type usageError struct{ err error }

func (e usageError) Error() string { return e.err.Error() }
func (e usageError) Unwrap() error { return e.err }

func usagef(format string, args ...any) error { return usageError{fmt.Errorf(format, args...)} }

// ExitCode maps an error from Execute to the process exit status.
func ExitCode(err error) int {
	if err == nil {
		return 0
	}
	var u usageError
	if errors.As(err, &u) {
		return 2
	}
	return 1
}

// NewRoot returns the yesly command.
func NewRoot(app *App) *cobra.Command {
	if app.Now == nil {
		app.Now = time.Now
	}
	if app.Sleep == nil {
		app.Sleep = time.Sleep
	}
	root := &cobra.Command{
		Use:   "yesly",
		Short: "Work with the Yesly API from your terminal",
		Long: `yesly calls every operation of the Yesly API, signs and checks webhook
payloads, and forwards your account's events to a local server.

Authenticate once with "yesly login", or set YESLY_API_KEY.`,
		Version:       fmt.Sprintf("%s (API %s)", version.Version, spec.APIVersion),
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.SetIn(app.In)
	root.SetOut(app.Out)
	root.SetErr(app.Err)
	root.SetFlagErrorFunc(func(_ *cobra.Command, err error) error { return usageError{err} })

	pf := root.PersistentFlags()
	pf.StringVar(&app.apiKey, "api-key", "", "API key; overrides YESLY_API_KEY and the profile")
	pf.StringVarP(&app.profile, "profile", "p", "", `profile saved by "yesly login" (default "default", or YESLY_PROFILE)`)
	pf.StringVar(&app.baseURL, "base-url", "", "API base URL; overrides YESLY_BASE_URL and the profile (default "+spec.BaseURL+")")
	pf.StringVar(&app.apiVersion, "api-version", spec.APIVersion, "Yesly-Version header")
	pf.IntVar(&app.maxRetries, "max-retries", 2, "retries after a network error, 408, 429 or 5xx")
	pf.DurationVar(&app.timeout, "timeout", 30*time.Second, "time limit for each attempt")
	pf.BoolVar(&app.compact, "compact", false, "print JSON on one line")
	pf.BoolVar(&app.live, "live", false, "allow requests that change data with a live key")
	pf.BoolVar(&app.dryRun, "dry-run", false, "print the request instead of sending it")

	root.AddGroup(&cobra.Group{ID: "resources", Title: "API resources:"}, &cobra.Group{ID: "tools", Title: "Tools:"})
	for _, c := range resourceCommands(app) {
		c.GroupID = "resources"
		root.AddCommand(c)
	}
	for _, c := range []*cobra.Command{loginCommand(app), logoutCommand(app), whoamiCommand(app), apiCommand(app), listenCommand(app), webhooksCommand(app), versionCommand(app)} {
		c.GroupID = "tools"
		root.AddCommand(c)
	}
	root.AddCommand(docsCommand(root))
	root.SetHelpCommandGroupID("tools")
	root.SetCompletionCommandGroupID("tools")
	return root
}

// Execute runs the CLI with the process's arguments and streams, and returns the exit status.
func Execute(ctx context.Context) int {
	app := &App{In: os.Stdin, Out: os.Stdout, Err: os.Stderr, IsTerminal: stdinIsTerminal, ReadSecret: readSecret}
	root := NewRoot(app)
	err := root.ExecuteContext(ctx)
	if err != nil {
		printError(app.Err, err)
	}
	return ExitCode(err)
}

func printError(w io.Writer, err error) {
	var apiErr *api.Error
	if errors.As(err, &apiErr) {
		fmt.Fprintf(w, "Error: %s\n", apiErr.Error())
		return
	}
	msg := err.Error()
	var u usageError
	if errors.As(err, &u) && !strings.Contains(msg, "--help") {
		msg += "\nRun with --help for usage."
	}
	fmt.Fprintf(w, "Error: %s\n", msg)
}

// credentials resolves the key, base URL and profile name for a call.
type credentials struct {
	apiKey      string
	keySource   string
	baseURL     string
	profileName string
	profile     *config.Profile
	file        *config.File
}

func (app *App) credentials() (*credentials, error) {
	f, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("reading the config file: %w", err)
	}
	c := &credentials{file: f, profileName: app.profileName()}
	c.profile = f.Profile(c.profileName)
	switch {
	case app.apiKey != "":
		c.apiKey, c.keySource = app.apiKey, "--api-key"
	case os.Getenv("YESLY_API_KEY") != "":
		c.apiKey, c.keySource = os.Getenv("YESLY_API_KEY"), "YESLY_API_KEY"
	case c.profile != nil && c.profile.APIKey != "":
		c.apiKey, c.keySource = c.profile.APIKey, "profile "+c.profileName
	}
	base := spec.BaseURL
	switch {
	case app.baseURL != "":
		base = app.baseURL
	case os.Getenv("YESLY_BASE_URL") != "":
		base = os.Getenv("YESLY_BASE_URL")
	case c.profile != nil && c.profile.BaseURL != "":
		base = c.profile.BaseURL
	}
	if c.baseURL, err = api.CheckBaseURL(base); err != nil {
		return nil, usageError{err}
	}
	return c, nil
}

func (app *App) profileName() string {
	if app.profile != "" {
		return app.profile
	}
	if p := os.Getenv("YESLY_PROFILE"); p != "" {
		return p
	}
	return "default"
}

func (app *App) client(c *credentials) *api.Client {
	return &api.Client{
		APIKey:     c.apiKey,
		BaseURL:    c.baseURL,
		APIVersion: app.apiVersion,
		UserAgent:  api.DefaultUserAgent(version.Version),
		Timeout:    app.timeout,
		MaxRetries: app.maxRetries,
		HTTP:       app.HTTP,
		Sleep:      app.Sleep,
	}
}

// send runs one request, or prints it with --dry-run. Live keys need --live
// for anything but GET.
func (app *App) send(ctx context.Context, r api.Request) (*api.Response, error) {
	c, err := app.credentials()
	if err != nil {
		return nil, err
	}
	if app.dryRun {
		return nil, app.printDryRun(c, r)
	}
	if r.Method != http.MethodGet && api.Mode(c.apiKey) == "live" && !app.live {
		return nil, usagef("this request changes live data; add --live to send it with a live key")
	}
	return app.client(c).Do(ctx, r)
}

func versionCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the CLI and API versions",
		Args:  noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Fprintf(app.Out, "yesly %s\nAPI version %s\n", version.Version, spec.APIVersion)
			if version.Commit != "" && version.Commit != "none" {
				fmt.Fprintf(app.Out, "commit %s %s\n", version.Commit, version.Date)
			}
			return nil
		},
	}
}

func noArgs(_ *cobra.Command, args []string) error {
	if len(args) > 0 {
		return usagef("unexpected argument %q", args[0])
	}
	return nil
}
