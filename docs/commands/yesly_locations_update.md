## yesly locations update

Update a location

### Synopsis

Update a location.

Changes a location. Send only what changes.

API: PATCH /locations/{id}

Arguments:
  <id>  Id of the location, such as bilbao.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly locations update <id> [flags]
```

### Examples

```
  yesly locations update bilbao --transfer-phone +34944000001
```

### Options

```
      --address string          New address.
  -d, --data string             JSON body, inline or @file (@- for stdin)
  -h, --help                    help for update
      --name string             New name.
      --transfer-phone string   New transfer phone, E.164. An empty string removes it.
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

