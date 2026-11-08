## yesly webhook-endpoints test

Send a test event

### Synopsis

Send a test event.

Sends a sample event, marked `"mode": "test"`, to the endpoint and
returns what it answered. Useful to check the URL and your signature
check before going live.

API: POST /webhook_endpoints/{id}/test

Arguments:
  <id>  Id of the webhook endpoint, prefixed we_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly webhook-endpoints test <id> [flags]
```

### Examples

```
  yesly webhook-endpoints test we_… --type appointment.confirmed
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for test
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --type string              Event type to simulate. One of: appointment.created, appointment.confirmed, appointment.rescheduled, appointment.reschedule_requested, appointment.cancelled, appointment.no_reply, message.sent, message.delivered, message.failed, call.started, call.completed, call.transferred, balance.low, balance.topped_up.
```

### Options inherited from parent commands

```
      --api-key string       API key; overrides YESLY_API_KEY and the profile
      --api-version string   Yesly-Version header (default "2026-10-01")
      --base-url string      API base URL; overrides YESLY_BASE_URL and the profile (default https://api.yesly.dev/v1)
      --compact              print JSON on one line
      --dry-run              print the request instead of sending it
      --live                 allow requests that change data with a live key
      --max-retries int      retries after a network error, 408, 429 or 5xx (default 2)
  -p, --profile string       profile saved by "yesly login" (default "default", or YESLY_PROFILE)
      --timeout duration     time limit for each attempt (default 30s)
```

### SEE ALSO

* [yesly webhook-endpoints](yesly_webhook-endpoints.md)	 - URLs that receive events

