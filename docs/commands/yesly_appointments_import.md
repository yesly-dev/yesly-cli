## yesly appointments import

Import appointments

### Synopsis

Import appointments.

Creates or updates up to 1,000 appointments in one request, for systems
that plan the whole day at once. Each item takes the same fields as
Create an appointment; items whose `external_id` already exists are
updated. The import runs in the background: poll Retrieve an import.

API: POST /appointments/import

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly appointments import [flags]
```

### Examples

```
  yesly appointments import --appointments @appointments.json
  yesly appointments import --data '{"appointments": [ … ]}'
```

### Options

```
      --appointments string      (required) Between 1 and 1,000 appointments. JSON, inline or @file.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for import
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
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

