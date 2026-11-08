package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yesly-dev/yesly-cli/internal/spec"
	"github.com/yesly-dev/yesly-cli/internal/webhook"
)

func webhooksCommand(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhooks",
		Short: "Sign and verify webhook payloads locally",
		Long: `Signs and verifies Yesly-Signature headers on this machine, with no API
call, to test your webhook handler. To manage the endpoints themselves, use
"yesly webhook-endpoints".`,
		Args: noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error { return cmd.Help() },
	}
	cmd.AddCommand(webhooksSignCommand(app), webhooksVerifyCommand(app), webhooksEventTypesCommand(app))
	return cmd
}

func webhooksSignCommand(app *App) *cobra.Command {
	var secret, payload string
	var timestamp int64
	cmd := &cobra.Command{
		Use:   "sign",
		Short: "Print the Yesly-Signature header for a payload",
		Long:  `Reads the payload from --payload (inline or @file) or from stdin, and prints the Yesly-Signature value the API would send with it.`,
		Example: `  yesly webhooks sign --secret whsec_… --payload @event.json
  curl localhost:3000/webhook -H "Yesly-Signature: $(yesly webhooks sign --secret whsec_… --payload @event.json)" \
    -H 'Content-Type: application/json' --data-binary @event.json`,
		Args: noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			body, err := readPayload(app, payload)
			if err != nil {
				return err
			}
			t := app.Now()
			if timestamp != 0 {
				t = time.Unix(timestamp, 0)
			}
			fmt.Fprintln(app.Out, webhook.Sign(body, secret, t))
			return nil
		},
	}
	cmd.Flags().StringVar(&secret, "secret", "", "(required) the endpoint's signing secret, whsec_…")
	cmd.Flags().StringVar(&payload, "payload", "", "payload, inline or @file (default: stdin)")
	cmd.Flags().Int64Var(&timestamp, "timestamp", 0, "Unix time to sign with (default: now)")
	_ = cmd.MarkFlagRequired("secret")
	return cmd
}

func webhooksVerifyCommand(app *App) *cobra.Command {
	var secret, payload, header string
	var tolerance time.Duration
	cmd := &cobra.Command{
		Use:     "verify",
		Short:   "Check a Yesly-Signature header against a payload",
		Long:    `Reads the payload from --payload (inline or @file) or from stdin. Exits 0 when the header matches and is within the tolerance, 1 otherwise.`,
		Example: `  yesly webhooks verify --secret whsec_… --header 't=1762531200,v1=5f0c…' --payload @body.json`,
		Args:    noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			body, err := readPayload(app, payload)
			if err != nil {
				return err
			}
			if err := webhook.Verify(body, header, secret, tolerance, app.Now()); err != nil {
				return err
			}
			fmt.Fprintln(app.Out, "The signature matches.")
			return nil
		},
	}
	cmd.Flags().StringVar(&secret, "secret", "", "(required) the endpoint's signing secret, whsec_…")
	cmd.Flags().StringVar(&header, "header", "", "(required) the Yesly-Signature value")
	cmd.Flags().StringVar(&payload, "payload", "", "raw payload, inline or @file (default: stdin)")
	cmd.Flags().DurationVar(&tolerance, "tolerance", webhook.DefaultTolerance, "maximum age of the timestamp; 0 turns the check off")
	_ = cmd.MarkFlagRequired("secret")
	_ = cmd.MarkFlagRequired("header")
	return cmd
}

func webhooksEventTypesCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "event-types",
		Short: "List every webhook event type",
		Args:  noArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Fprintln(app.Out, strings.Join(spec.EventTypes, "\n"))
			return nil
		},
	}
}

func readPayload(app *App, payload string) ([]byte, error) {
	if payload == "" {
		payload = "@-"
	}
	return readArg(app, payload)
}
