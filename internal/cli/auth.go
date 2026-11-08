package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yesly-dev/yesly-cli/internal/api"
	"github.com/yesly-dev/yesly-cli/internal/config"
)

func loginCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Save an API key in a profile",
		Long: `Saves an API key, after checking it against the API, in the profile named
by --profile ("default" if not given). Paste the key at the prompt, or pipe
it in: echo "$KEY" | yesly login. --base-url is saved with the profile.

Create keys in the dashboard, under API keys. Keys starting with ysk_test_
run in test mode; ysk_live_ keys contact real customers and spend balance.`,
		Example: "  yesly login\n  yesly login --profile mock --base-url http://localhost:4010",
		Args:    noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			key := app.apiKey
			if key == "" {
				var err error
				if key, err = app.promptKey(); err != nil {
					return err
				}
			}
			key = strings.TrimSpace(key)
			if api.Mode(key) == "" {
				return usagef("that does not look like a Yesly API key (ysk_test_… or ysk_live_…)")
			}
			app.apiKey = key
			c, err := app.credentials()
			if err != nil {
				return err
			}
			if _, err := app.client(c).Do(cmd.Context(), api.Request{Method: "GET", Path: "/balance"}); err != nil {
				return fmt.Errorf("the key did not work: %w", err)
			}
			p := c.file.Profile(c.profileName)
			if p == nil {
				p = &config.Profile{}
				c.file.Profiles[c.profileName] = p
			}
			p.APIKey = key
			if app.baseURL != "" {
				p.BaseURL = c.baseURL
			}
			if err := c.file.Save(); err != nil {
				return err
			}
			path, _ := config.Path()
			fmt.Fprintf(app.Err, "Saved %s (%s mode) as profile %q in %s\n", mask(key), api.Mode(key), c.profileName, path)
			return nil
		},
	}
}

func (app *App) promptKey() (string, error) {
	if app.IsTerminal != nil && app.IsTerminal() && app.ReadSecret != nil {
		fmt.Fprint(app.Err, "API key (ysk_test_… or ysk_live_…): ")
		key, err := app.ReadSecret()
		fmt.Fprintln(app.Err)
		return key, err
	}
	line, err := bufio.NewReader(app.In).ReadString('\n')
	if err != nil && line == "" {
		return "", usagef("no API key on stdin")
	}
	return line, nil
}

func logoutCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Remove a saved profile",
		Long:  `Removes the profile named by --profile ("default" if not given) from the config file. The key itself stays valid until you revoke it in the dashboard.`,
		Args:  noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			f, err := config.Load()
			if err != nil {
				return err
			}
			name := app.profileName()
			if f.Profile(name) == nil {
				return usagef("there is no profile %q", name)
			}
			delete(f.Profiles, name)
			if err := f.Save(); err != nil {
				return err
			}
			fmt.Fprintf(app.Err, "Removed profile %q\n", name)
			return nil
		},
	}
}

func whoamiCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "whoami",
		Short: "Show which key and base URL commands use, and check the key",
		Args:  noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			c, err := app.credentials()
			if err != nil {
				return err
			}
			path, _ := config.Path()
			fmt.Fprintf(app.Out, "Config:   %s\n", path)
			if names := c.file.Names(); len(names) > 0 {
				fmt.Fprintf(app.Out, "Profiles: %s\n", strings.Join(names, ", "))
			}
			if c.apiKey == "" {
				fmt.Fprintln(app.Out, "Key:      none (run yesly login or set YESLY_API_KEY)")
				return nil
			}
			fmt.Fprintf(app.Out, "Key:      %s, %s mode, from %s\n", mask(c.apiKey), api.Mode(c.apiKey), c.keySource)
			fmt.Fprintf(app.Out, "Base URL: %s\n", c.baseURL)
			fmt.Fprintf(app.Out, "Version:  %s\n", app.apiVersion)
			if _, err := app.client(c).Do(cmd.Context(), api.Request{Method: "GET", Path: "/balance"}); err != nil {
				return fmt.Errorf("the key did not work: %w", err)
			}
			fmt.Fprintln(app.Out, "Check:    the API accepted the key")
			return nil
		},
	}
}
