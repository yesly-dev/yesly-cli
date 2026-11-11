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
  yesly events list --type conversation.ended --from 2026-12-01T00:00:00+01:00
  yesly events list --conversation cnv_7Qp4Ns2Xk9Lm3D
```

### Options

```
      --all                     follow every page and print all items as one JSON array
      --contact string          Only events about this contact and its opt-outs.
      --conversation string     Only events about this conversation and its messages.
      --from string             Events created on or after this time. A day alone means 00:00 UTC.
  -h, --help                    help for list
      --limit int               Between 1 and 100. Default 20.
      --starting-after string   Cursor: the id of the last item of the previous page.
      --type string             Only this event type, matched exactly.
```

### Options inherited from parent commands

```
      --api-key string       API key; overrides YESLY_API_KEY and the profile
      --api-version string   Yesly-Version header (default "2026-11-11")
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

