## yesly webhook-endpoints update

Update an endpoint

### Synopsis

Update an endpoint.

Changes an endpoint. Send only what changes.

API: PATCH /webhook_endpoints/{id}

Arguments:
  <id>  Id of the webhook endpoint, prefixed we_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly webhook-endpoints update <id> [flags]
```

### Examples

```
  yesly webhook-endpoints update we_… --disabled
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
      --description string       New note.
      --disabled                 true stops sending events to it and drops its pending retries. false enables it again and resets its count of failed deliveries.
      --enabled-events strings   Replaces the list of event types. One of: appointment.created, appointment.confirmed, appointment.rescheduled, appointment.reschedule_requested, appointment.cancelled, appointment.no_reply, message.sent, message.delivered, message.failed, call.started, call.completed, call.transferred, balance.low, balance.topped_up, *, appointment.*, message.*, call.*, balance.*.
  -h, --help                     help for update
      --url string               New URL.
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

