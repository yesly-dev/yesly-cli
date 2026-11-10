## yesly events list

List events

### Synopsis

List events.

Lists events from the last 30 days, newest first. Use it to catch up after downtime on your side.

API: GET /events

```
yesly events list [flags]
```

### Examples

```
  yesly events list --type 'appointment.*' --from 2026-11-07T00:00:00Z
```

### Options

```
      --all                     follow every page and print all items as one JSON array
      --appointment string      Only events about this appointment.
      --from string             Events created on or after this time. A day alone means 00:00 UTC.
  -h, --help                    help for list
      --limit int               Between 1 and 100. Default 20.
      --starting-after string   Cursor: the id of the last item of the previous page.
      --type string             Only this event type, matched exactly.
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

* [yesly events](yesly_events.md)	 - Everything that happened, as delivered to webhook endpoints

