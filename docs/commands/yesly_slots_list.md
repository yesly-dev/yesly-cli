## yesly slots list

(deprecated) List slots

### Synopsis

Deprecated. List slots.

Lists published slots, soonest first.

API: GET /slots

```
yesly slots list [flags]
```

### Examples

```
  yesly slots list --type delivery --from 2026-11-13 --to 2026-11-20 --available
```

### Options

```
      --all                     follow every page and print all items as one JSON array
      --available               true for slots with room left.
      --from string             Slots starting on or after this day or time.
  -h, --help                    help for list
      --limit int               Between 1 and 100. Default 20.
      --location string         Only this location.
      --starting-after string   Cursor: the id of the last item of the previous page.
      --to string               Slots starting before this day or time.
      --type string             Only this appointment type.
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

* [yesly slots](yesly_slots.md)	 - (deprecated) Free times published in Yesly

