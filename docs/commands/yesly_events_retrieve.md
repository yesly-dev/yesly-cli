## yesly events retrieve

Retrieve an event

### Synopsis

Retrieve an event.

Fetches one event from the last 30 days.

API: GET /events/{id}

Arguments:
  <id>  Id of the event, prefixed evt_.

```
yesly events retrieve <id> [flags]
```

### Examples

```
  yesly events retrieve evt_…
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

* [yesly events](yesly_events.md)	 - Everything that happened, as delivered to webhook endpoints

