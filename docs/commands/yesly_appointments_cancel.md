## yesly appointments cancel

Cancel an appointment

### Synopsis

Cancel an appointment.

Cancels an appointment from your side, for example when the order is
cancelled. Yesly stops contacting the customer at once and does not send
them anything else. Attempts already made are still charged. Fires
`appointment.cancelled`.

API: POST /appointments/{id}/cancel

Arguments:
  <id>  Id of the appointment, prefixed apt_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly appointments cancel <id> [flags]
```

### Examples

```
  yesly appointments cancel apt_… --reason "Shop closed"
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for cancel
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --reason string            Optional. Stored in cancellation.reason and shown in the dashboard.
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

* [yesly appointments](yesly_appointments.md)	 - Visits, deliveries or meetings Yesly confirms with your customers

