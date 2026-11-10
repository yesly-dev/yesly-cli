# yesly-cli

The official command-line interface for the Yesly API.

> Status: 0.2.0. Covers API version `2026-10-01`.

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
yesly appointments list -p acme
yesly whoami -p acme      # which key, mode and base URL a command would use
yesly logout -p acme
```

## Usage

Commands follow the API reference: `yesly <resource> <operation>`.

```bash
yesly appointments create --external-id ORD-48213 --type delivery \
  --customer.name Martha --customer.phone +34612345678 \
  --slot.start 2026-11-13T09:00:00+01:00 --slot.end 2026-11-13T11:00:00+01:00 \
  --metadata order_id=PV-18342 --metadata route_id=BIO-07

yesly appointments list --status no_reply --date 2026-11-13
yesly appointments retrieve apt_…
yesly appointments update apt_… --slot.start 2026-11-14T09:00:00+01:00
yesly appointments cancel apt_… --reason "Shop closed"
yesly webhook-endpoints create --url https://example.com/yesly/webhook --enabled-events '*'
yesly balance retrieve
yesly analytics retrieve-summary --range 30d
```

| Resource | Operations |
| --- | --- |
| `appointments` | `list`, `count`, `create`, `retrieve`, `update`, `cancel`, `import`, `recording` |
| `imports` | `retrieve` |
| `appointment-types` | `list`, `create`, `update` |
| `locations` | `list`, `create`, `update` |
| `webhook-endpoints` | `list`, `create`, `update`, `delete`, `test`, `roll-secret` |
| `events` | `list`, `retrieve` |
| `balance` | `retrieve`, `list-transactions` |
| `analytics` | `retrieve-summary` |
| `slots` (deprecated) | `list`, `create`, `delete` |

`yesly <resource> <operation> --help` lists every flag with its meaning, and
[docs/commands](docs/commands/yesly.md) has the same reference as Markdown.

### Flags and bodies

- Fields inside objects use dots: `--customer.name`, `--slot.start`.
- Lists take commas or repeated flags: `--channels whatsapp,voice`.
- Maps take `key=value`, repeated: `--metadata route_id=BIO-07`. In an update,
  an empty value (`--metadata route_id=`) removes the key.
- Flags with fixed values (`--status`, `--time-format`, `--enabled-events`…)
  are checked before sending, and complete in the shell.
- Lists of objects, such as the appointments of an import, take JSON inline
  or from a file: `--appointments @appointments.json`.
- `--data` (`-d`) sends a whole JSON body, inline, from a file (`@file`) or
  from stdin (`@-`). Flags given with it override its fields:

  ```bash
  yesly appointments create -d @appointment.json --external-id ORD-48214
  ```

### Output

Answers are printed as the API sends them, as indented JSON on stdout;
`--compact` prints one line. Errors go to stderr. Pipe to `jq` to pick fields:

```bash
yesly appointments list --status no_reply --compact | jq -r '.data[].external_id'
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
safe to repeat. Creating an appointment is idempotent by `external_id`.

### Any request

`yesly api` sends any method and path, with your key and the same retries:

```bash
yesly api GET /appointments -f status=no_reply -F limit=5
yesly api POST /appointments/apt_…/cancel -f reason="Shop closed"
yesly api PATCH /locations/bilbao --data @location.json
```

`-f key=value` adds a string and `-F key=value` a typed value (`true`,
`false`, `null`, numbers, JSON, or `@file`). For GET and DELETE they become
query parameters; otherwise they go in the body, where `customer.name` or
`customer[name]` nests and `channels[]` appends.

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
yesly listen --events appointment.confirmed,appointment.no_reply
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
five minutes (`--tolerance`), and 1 otherwise. To send a test delivery from
Yesly itself, use `yesly webhook-endpoints test we_… --type appointment.confirmed`.

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
