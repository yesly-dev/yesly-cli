package cli

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/yesly-dev/yesly-cli/internal/api"
	"github.com/yesly-dev/yesly-cli/internal/spec"
	"github.com/yesly-dev/yesly-cli/internal/version"
	"github.com/yesly-dev/yesly-cli/internal/webhook"
)

// lookback is how far before the last event seen each poll asks from, so an
// event stored a little late is not missed. Ids already seen are skipped.
const lookback = 2 * time.Minute

func listenCommand(app *App) *cobra.Command {
	var forwardTo, secret string
	var types []string
	var interval time.Duration
	cmd := &cobra.Command{
		Use:   "listen",
		Short: "Receive your account's events locally, without a public URL",
		Long: `Polls GET /events for new events and prints each one as a line of JSON, or,
with --forward-to, posts it to a local URL signed like a webhook delivery,
so you can test your handler without exposing it to the internet.

Events carry your customers' data, so --forward-to takes http:// only for
this machine (localhost, 127.0.0.1, ::1) and https:// anywhere, and a
redirect is reported, not followed.

Deliveries are signed with --secret, or with a secret this command creates
and saves in the profile; verify them with that secret in development. The
saved secret is printed when listen starts; one given with --secret is not.
Events arrive with up to --interval of delay, from the moment listen starts.
This does not create a webhook endpoint, and real endpoints keep receiving
events.`,
		Example: `  yesly listen
  yesly listen --forward-to http://localhost:3000/yesly/webhook
  yesly listen --events conversation.ended,conversation.no_reply`,
		Args: noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := checkEnum(spec.Field{Name: "events", Enum: spec.EventTypes}, types...); err != nil {
				return err
			}
			showSecret := false
			if forwardTo != "" {
				if err := checkForwardTo(forwardTo); err != nil {
					return err
				}
				if secret == "" {
					var err error
					if secret, err = app.listenSecret(); err != nil {
						return err
					}
					// The CLI's own development secret; one typed by the user stays out of the output.
					showSecret = true
				}
			}
			return app.listen(cmd.Context(), forwardTo, secret, showSecret, types, interval)
		},
	}
	f := cmd.Flags()
	f.StringVar(&forwardTo, "forward-to", "", "URL to post each event to, signed with Yesly-Signature (https://, or http:// on this machine)")
	f.StringVar(&secret, "secret", "", "signing secret for forwarded events (default: one saved in the profile)")
	f.StringSliceVar(&types, "events", nil, "only these event types (default: all)")
	f.DurationVar(&interval, "interval", 2*time.Second, "time between polls")
	_ = cmd.RegisterFlagCompletionFunc("events", func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
		return spec.EventTypes, cobra.ShellCompDirectiveNoFileComp
	})
	return cmd
}

// checkForwardTo refuses a target that would send events, which carry
// customers' names and phone numbers, in clear over the network.
func checkForwardTo(raw string) error {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		return usagef("--forward-to must be an http:// or https:// URL")
	}
	if u.Scheme == "http" && !api.Loopback(u.Hostname()) {
		return usagef("--forward-to takes http:// only for this machine (localhost, 127.0.0.1 or ::1); use https:// for %s, so events do not travel in clear", u.Hostname())
	}
	return nil
}

// listenSecret returns the profile's forwarding secret, creating it the first time.
func (app *App) listenSecret() (string, error) {
	c, err := app.credentials()
	if err != nil {
		return "", err
	}
	if c.profile != nil && c.profile.ListenSecret != "" {
		return c.profile.ListenSecret, nil
	}
	var b [24]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	secret := "whsec_" + hex.EncodeToString(b[:])
	if c.profile == nil {
		fmt.Fprintf(app.Err, "There is no profile %q to save the signing secret in; pass --secret to keep the same one.\n", c.profileName)
		return secret, nil
	}
	c.profile.ListenSecret = secret
	if err := c.file.Save(); err != nil {
		return "", err
	}
	return secret, nil
}

type listenedEvent struct {
	raw     json.RawMessage
	id      string
	typ     string
	created time.Time
}

func (app *App) listen(ctx context.Context, forwardTo, secret string, showSecret bool, types []string, interval time.Duration) error {
	c, err := app.credentials()
	if err != nil {
		return err
	}
	if c.apiKey == "" {
		return usagef("no API key: run yesly login, set YESLY_API_KEY or pass --api-key")
	}
	client := app.client(c)
	wanted := map[string]bool{}
	for _, t := range types {
		wanted[t] = true
	}
	seen := map[string]bool{}
	since := app.Now().Add(-lookback)

	// Events from before listen started are marked seen, not sent.
	first, err := app.pollEvents(ctx, client, since)
	if err != nil {
		return err
	}
	for _, e := range first {
		seen[e.id] = true
	}

	mode := api.Mode(c.apiKey)
	if forwardTo != "" {
		signedWith := "the secret given with --secret"
		if showSecret {
			signedWith = secret
		}
		fmt.Fprintf(app.Err, "Forwarding %s-mode events to %s, signed with %s. Ctrl-C to stop.\n", mode, forwardTo, signedWith)
	} else {
		fmt.Fprintf(app.Err, "Listening for %s-mode events. Ctrl-C to stop.\n", mode)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(interval):
		}
		events, err := app.pollEvents(ctx, client, since)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			fmt.Fprintf(app.Err, "%s  polling failed: %v\n", app.Now().Format("15:04:05"), err)
			continue
		}
		for _, e := range events {
			if seen[e.id] {
				continue
			}
			seen[e.id] = true
			if e.created.Add(-lookback).After(since) {
				since = e.created.Add(-lookback)
			}
			if len(wanted) > 0 && !wanted[e.typ] {
				continue
			}
			if forwardTo == "" {
				var buf bytes.Buffer
				_ = json.Compact(&buf, e.raw)
				fmt.Fprintln(app.Out, buf.String())
				continue
			}
			status, err := app.forward(ctx, forwardTo, secret, e.raw)
			line := fmt.Sprintf("%s  %-34s %s  ", app.Now().Format("15:04:05"), e.typ, e.id)
			if err != nil {
				fmt.Fprintf(app.Err, "%s-> %v\n", line, err)
			} else {
				fmt.Fprintf(app.Err, "%s-> %d\n", line, status)
			}
		}
	}
}

// pollEvents returns every event created since t, oldest first.
func (app *App) pollEvents(ctx context.Context, client *api.Client, t time.Time) ([]listenedEvent, error) {
	q := url.Values{"from": {t.UTC().Format(time.RFC3339)}, "limit": {"100"}}
	var out []listenedEvent
	for {
		res, err := client.Do(ctx, api.Request{Method: "GET", Path: "/events", Query: q})
		if err != nil {
			return nil, err
		}
		var page struct {
			Data    []json.RawMessage `json:"data"`
			HasMore bool              `json:"has_more"`
		}
		if err := json.Unmarshal(res.Body, &page); err != nil {
			return nil, fmt.Errorf("reading events: %w", err)
		}
		for _, raw := range page.Data {
			var head struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Created string `json:"created"`
			}
			if err := json.Unmarshal(raw, &head); err != nil || head.ID == "" {
				continue
			}
			created, _ := time.Parse(time.RFC3339, head.Created)
			out = append(out, listenedEvent{raw: raw, id: head.ID, typ: head.Type, created: created})
		}
		if !page.HasMore || len(page.Data) == 0 {
			break
		}
		q.Set("starting_after", out[len(out)-1].id)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].created.Before(out[j].created) })
	return out, nil
}

// forward posts one event as a webhook delivery would arrive. A redirect is
// an error: following it would post the event and its signature to an
// address the user did not give.
func (app *App) forward(ctx context.Context, target, secret string, raw json.RawMessage) (int, error) {
	var body bytes.Buffer
	if err := json.Compact(&body, raw); err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body.Bytes()))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "yesly-cli/"+version.Version+" (listen)")
	req.Header.Set("Yesly-Signature", webhook.Sign(body.Bytes(), secret, app.Now()))
	hc := &http.Client{}
	if app.HTTP != nil {
		*hc = *app.HTTP
	}
	hc.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	res, err := hc.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	_, _ = io.Copy(io.Discard, res.Body)
	if res.StatusCode >= 300 && res.StatusCode < 400 {
		if to := res.Header.Get("Location"); to != "" {
			return 0, fmt.Errorf("%d: redirect to %s not followed; pass the final URL to --forward-to", res.StatusCode, to)
		}
	}
	return res.StatusCode, nil
}
