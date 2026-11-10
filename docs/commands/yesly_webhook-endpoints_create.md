## yesly webhook-endpoints create

Create an endpoint

### Synopsis

Create an endpoint.

Adds a webhook endpoint in the key's mode. Live and test endpoints are separate.

API: POST /webhook_endpoints

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly webhook-endpoints create [flags]
```

### Examples

```
  yesly webhook-endpoints create --url https://example.com/yesly/webhook \
    --enabled-events appointment.confirmed,appointment.no_reply
  yesly webhook-endpoints create --url https://example.com/yesly/webhook --enabled-events '*'
```

### Options

```
      --api-version string       Stored and returned as sent; default 2026-10-01. It does not change the events.
  -d, --data string              JSON body, inline or @file (@- for stdin)
      --description string       Optional note.
      --enabled-events strings   (required) Event types sent to the endpoint. ["*"] for all; a family such as appointment.* for every event of one object, including types added later. One of: appointment.created, appointment.confirmed, appointment.rescheduled, appointment.reschedule_requested, appointment.cancelled, appointment.no_reply, message.sent, message.delivered, message.failed, call.started, call.completed, call.transferred, balance.low, balance.topped_up, *, appointment.*, message.*, call.*, balance.*.
  -h, --help                     help for create
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --url string               (required) HTTPS URL on a public host, without a user name or password (url_invalid otherwise). Deliveries to a self-signed certificate fail.
```

### Options inherited from parent commands

```
      --api-key string     API key; overrides YESLY_API_KEY and the profile
      --base-url string    API base URL; overrides YESLY_BASE_URL and the profile (default https://api.yesly.dev/v1)
      --compact            print JSON on one line
      --dry-run            print the request instead of sending it
      --live               allow requests that change data with a live key
      --max-retries int    retries after a network error, 408, 429 or 5xx (default 2)
  -p, --profile string     profile saved by "yesly login" (default "default", or YESLY_PROFILE)
      --timeout duration   time limit for each attempt (default 30s)
```

### SEE ALSO

* [yesly webhook-endpoints](yesly_webhook-endpoints.md)	 - URLs that receive events

