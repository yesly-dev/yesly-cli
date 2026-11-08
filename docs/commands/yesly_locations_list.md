## yesly locations list

List locations

### Synopsis

List locations.

Lists the account's locations.

API: GET /locations

```
yesly locations list [flags]
```

### Examples

```
  yesly locations list
```

### Options

```
      --all    follow every page and print all items as one JSON array
  -h, --help   help for list
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

