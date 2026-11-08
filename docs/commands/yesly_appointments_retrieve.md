## yesly appointments retrieve

Retrieve an appointment

### Synopsis

Retrieve an appointment.

Fetches an appointment with its attempts and answer. To find one by your own id, list with `external_id`.

API: GET /appointments/{id}

Arguments:
  <id>  Id of the appointment, prefixed apt_.

```
yesly appointments retrieve <id> [flags]
```

### Examples

```
  yesly appointments retrieve apt_…
```

### Options

```
  -h, --help   help for retrieve
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

