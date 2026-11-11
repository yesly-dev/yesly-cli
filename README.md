# yesly-cli

The official command-line interface for the Yesly API.

> Status: 0.4.0. Covers API version `2026-11-11`.

Every operation of the API reference is a command, with a flag for each
field. The CLI also signs and verifies webhook payloads and forwards your
account's events to a local server, so you can test a webhook handler
without exposing it to the internet.

## Install

Download the archive for your system from
[Releases](https://github.com/yesly-dev/yesly-cli/releases), unpack it and put
`yesly` on your `PATH`. Builds exist for macOS, Linux and Windows, on amd64
and arm64; `checksums.txt` has the SHA-256 of each archive.

With Go 1.26 or later:

```bash
go install github.com/yesly-dev/yesly-cli/cmd/yesly@latest
```

The binary has no dependencies.

## Log in

```bash
yesly login
```

Paste an API key from the dashboard (API keys). The CLI checks it against
the API and saves it in `~/.config/yesly/config.json`, readable only by you.
A `ysk_test_` key runs everything in test mode; a `ysk_live_` key contacts
real customers and spends balance.

Keys are read from, in order: `--api-key`, `YESLY_API_KEY`, and the profile
(`--profile`, `YESLY_PROFILE`, or `default`). Keep several accounts or
environments as profiles:

```bash
yesly login --profile acme
yesly conversations list -p acme
yesly whoami -p acme      # which key, mode and base URL a command would use
yesly logout -p acme
```

## Usage

Commands follow the API reference: `yesly <resource> <operation>`. The usual
flow is to ask an agent to contact someone, then read the outcome when the
conversation ends:

```bash
# Create the contact, or update it by your own id.
yesly contacts create --external-id customer-3391 --name Marta --phone +34600000001

# Ask the delivery agent to contact her: WhatsApp first, then a call.
yesly conversations create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B \
  --channels whatsapp,call --start-at 2026-12-01T10:00:00+01:00 \
  --context '{"order": "8841", "delivery_window": "2026-12-02T09:00/13:00"}' \
  --idempotency-key order-8841-delivery

yesly conversations retrieve cnv_7Qp4Ns2Xk9Lm3D        # status, outcome, cost
yesly conversations list-messages cnv_7Qp4Ns2Xk9Lm3D   # what was said
yesly conversations cancel cnv_7Qp4Ns2Xk9Lm3D
yesly agents list
yesly messages create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B --channel sms --text "Hola Marta"
yesly webhook-endpoints create --url https://example.com/yesly/webhook --enabled-events conversation.ended
yesly analytics summary --range 30d
```

`--contact` takes a contact id or the data to find or create one, as JSON
(`--contact '{"external_id": "customer-3391", "phone": "+34600000001"}'`), so
the first command is optional.

| Resource | Operations |
| --- | --- |
| `agents` | `list`, `retrieve`, `update` |
| `senders` | `list` |
| `contacts` | `list`, `create`, `retrieve`, `update`, `delete`, `merge`, `list-memory`, `delete-memory` |
| `conversations` | `list`, `create`, `retrieve`, `delete`, `cancel`, `list-messages`, `recording` |
| `messages` | `list`, `create`, `retrieve` |
| `templates` | `list`, `create` |
| `opt-outs` | `list`, `create`, `delete` |
| `analytics` | `summary` |
| `balance` | `retrieve`, `list-transactions` |
| `webhook-endpoints` | `list`, `create`, `update`, `delete`, `test`, `roll-secret` |
| `events` | `list`, `retrieve` |

`yesly <resource> <operation> --help` lists every flag with its meaning, and
[docs/commands](docs/commands/yesly.md) has the same reference as Markdown.

### Flags and bodies

- Fields inside objects use dots: `--limits.daily-spend`, `--template.name`.
- Lists take commas or repeated flags: `--channels whatsapp,call`.
- Maps take `key=value`, repeated: `--metadata segment=premium`.
- Flags with fixed values (`--status`, `--channels`, `--enabled-events`…)
  are checked before sending, and complete in the shell.
- Free-form objects and lists of objects, such as a conversation's
  `--context` or an agent's `--contact-hours`, take JSON inline or from a
  file: `--context @context.json`. `--contact` also takes a plain id.
- `--data` (`-d`) sends a whole JSON body, inline, from a file (`@file`) or
  from stdin (`@-`). Flags given with it override its fields:

  ```bash
  yesly conversations create -d @conversation.json --channels call
  ```

### Output

Answers are printed as the API sends them, as indented JSON on stdout;
`--compact` prints one line. Errors go to stderr. Pipe to `jq` to pick fields:

```bash
yesly conversations list --status ended --compact | jq -r '.data[] | [.id, .outcome.code] | @tsv'
```

List commands print one page (`--limit` up to 100, `--starting-after` for the
next one). `--all` follows every page and prints all items as one JSON array:

```bash
yesly balance list-transactions --from 2026-11-01 --all | jq 'map(.amount) | add'
```

`--dry-run` prints the request instead of sending it, without needing a key.

### Live keys

With a `ysk_live_` key, commands that change data stop unless you add
`--live`. Reads (`list`, `retrieve`) never need it.

### Retries and idempotency

Network errors, timeouts, 408, 429 and 5xx are retried twice with backoff,
honouring `Retry-After` (`--max-retries`, `--timeout` per attempt).
Operations that take an `Idempotency-Key` get a new one per command, reused
on every retry; pass your own with `--idempotency-key` to make a whole command
safe to repeat. Send one with every `conversations create`, derived from your
own data (`order-8841-delivery`): a repeated command then never contacts
anyone twice.

### Any request

`yesly api` sends any method and path, with your key and the same retries:

```bash
yesly api GET /conversations -f status=no_reply -F limit=5
yesly api POST /conversations/cnv_7Qp4Ns2Xk9Lm3D/cancel
yesly api PATCH /contacts/ctc_3Kf9Lq2Wm7Rt4B --data @contact.json
```

`-f key=value` adds a string and `-F key=value` a typed value (`true`,
`false`, `null`, numbers, JSON, or `@file`). For GET and DELETE they become
query parameters; otherwise they go in the body, where `contact.name` or
`contact[name]` nests and `channels[]` appends.

## Integration checklist

The usual integration is two calls: ask an agent to contact someone (create a
conversation), and receive the outcome on a webhook (`conversation.ended`).
These are the points an integration most often gets wrong; check them before
going live. The whole API reference is also in Markdown, for agents and other
tools: [llms-full.txt](https://yesly.dev/llms-full.txt) (index in
[llms.txt](https://yesly.dev/llms.txt), OpenAPI in
[openapi.yaml](https://yesly.dev/openapi.yaml)).

1. **Read your agent first** (`yesly agents retrieve agt_…`). It declares the
   `outcomes` it can end with and the `context` it expects; a context that
   does not fit is a `422` `context_invalid` with the field in `param`.
2. **Use your own id as the contact's `external_id`.** Contacts are matched
   only by exact identifiers (`external_id`, then a phone, then an email), and
   `contacts create` adds what is new to the one it finds. An identifier that
   belongs to another contact is a `409` `contact_identifier_taken`: merge the
   two first (`yesly contacts merge`).
3. **Send an `--idempotency-key` with every new conversation**, derived from
   your own data, so a repeated command never contacts anyone twice.
4. **Send times with an offset** (`2026-12-01T10:00:00+01:00`) and phones in
   E.164. Live contacts must be in Spain, and not premium or special-rate
   numbers (`phone_not_allowed`).
5. **Handle the outcome on a webhook.** Verify the signature over the raw
   body, answer `2xx` within 10 seconds before doing slow work, and skip event
   ids you have already handled: a delivery can arrive more than once.
   `conversation.ended` carries the conversation with its `outcome`
   (`{ "code", "data" }`); `conversation.no_reply`, `conversation.cancelled`
   and `conversation.failed` close the others.
6. **Build and test with a `ysk_test_` key.** The test phone numbers choose
   what happens (`+34600000001` answers and the agent ends with its first
   outcome, `+34600000009` never answers; the full table is in the reference)
   and waits last seconds.
7. **Retry only what can succeed.** The CLI already retries network errors,
   `408`, `429` and `5xx`, with the same idempotency key. A `4xx` fails the
   same way again: fix the request using the error's `code` and `param`.
8. **Keep keys on the server.** A read-only key can list and retrieve but not
   write (`api_key_read_only`); a key limited to one agent gets `403`
   `api_key_agent_scope` on any other. With a `ysk_live_` key, commands that
   change data also need `--live`.

## Webhooks

### Receive events locally

```bash
yesly listen --forward-to http://localhost:3000/yesly/webhook
```

`listen` polls `GET /events` and posts each new event to your URL with a
`Yesly-Signature` header, the same way a webhook delivery arrives. The
signing secret is created on first use and saved in the profile; it is printed
when `listen` starts, so use it as the secret of your handler in development.
Without `--forward-to`, events are printed as JSON lines. `--events` keeps
only some types:

```bash
yesly listen --events conversation.ended,conversation.no_reply
```

Events arrive with up to `--interval` (2 s) of delay. `listen` creates no
webhook endpoint, and your real endpoints keep receiving events as usual.

### Sign and verify payloads

```bash
yesly webhooks sign --secret whsec_… --payload @event.json
yesly webhooks verify --secret whsec_… --header 't=…,v1=…' --payload @event.json
yesly webhooks event-types
```

`verify` exits with 0 when the signature matches and the timestamp is within
five minutes (`--tolerance`), and 1 otherwise. For example, to check a
`conversation.ended` delivery your handler saved, with the raw body in
`event.json` and the `Yesly-Signature` header it came with:

```bash
yesly webhooks verify --secret "$YESLY_WEBHOOK_SECRET" \
  --header "$(cat signature.txt)" --payload @event.json \
  && jq '.data.conversation | {id, outcome}' event.json
```

To send a test delivery from Yesly itself, use
`yesly webhook-endpoints test we_… --type conversation.ended`.

## Shell completion

```bash
yesly completion zsh > "${fpath[1]}/_yesly"      # zsh
yesly completion bash > /etc/bash_completion.d/yesly
yesly completion fish > ~/.config/fish/completions/yesly.fish
yesly completion powershell | Out-String | Invoke-Expression
```

## Environment

| Variable | Use |
| --- | --- |
| `YESLY_API_KEY` | API key, instead of a profile |
| `YESLY_PROFILE` | Profile to use when `--profile` is not given |
| `YESLY_BASE_URL` | API base URL, such as a mock server on `http://localhost:4010` |
| `YESLY_CONFIG` | Path of the config file |
| `XDG_CONFIG_HOME` | Where `yesly/config.json` lives (default `~/.config`) |

The base URL must use HTTPS, except for `localhost`, because the key is sent
with every request.

## Exit codes

| Code | Meaning |
| --- | --- |
| 0 | Done |
| 1 | The API answered with an error, it could not be reached, or a signature did not match |
| 2 | The command line was wrong: unknown flag, bad value, missing argument, or `--live` missing |

## Development

```bash
go test ./...
go run ./cmd/yesly --help
```

`internal/spec/zz_generated.go` is generated from the OpenAPI specification
and lists every operation and its flags; the commands are built from it at
run time. Do not edit it: change `yesly-openapi`, then, with both repos side
by side,

```bash
(cd ../yesly-openapi && npm install)
node ../yesly-openapi/sdk/generate.mjs cli .
go run ./cmd/yesly gen-docs docs/commands
```

Add an example for any new operation to `internal/cli/examples.go`; a test
fails until every operation has one. CI checks that `docs/commands` matches
the code.

## Releasing

Bump `Version` in `internal/version/version.go` and the status line above,
then push a tag `v<version>`. The release workflow runs the tests and
[GoReleaser](https://goreleaser.com), which builds every platform and
publishes the archives and checksums as a GitHub release. The workflow runs
in the `release` environment, so add reviewers there to approve each release.

## License

MIT
