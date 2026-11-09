## yesly locations create

Create a location

### Synopsis

Create a location.

Adds a location.

API: POST /locations

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly locations create [flags]
```

### Examples

```
  yesly locations create --id bilbao --name "Bilbao depot" \
    --address "Calle Mayor 1, 48001 Bilbao" --transfer-phone +34944000000
```

### Options

```
      --address string           (required) Full address.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for create
      --id string                (required) 1 to 40 lowercase letters, numbers, - and _. Cannot change later.
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --name string              (required) Name shown to the customer.
      --transfer-phone string    Phone for call transfers, E.164. Not checked; live calls do not transfer today.
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

* [yesly locations](yesly_locations.md)	 - Your centres, shops or depots, and the phone calls are passed to

