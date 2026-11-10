## yesly analytics retrieve-summary

Summary

### Synopsis

Summary.

The figures behind the dashboard's Analytics page, for your own reports:
outcomes, which channel got the answer, how long customers took, usage
and spend, in total and day by day. Appointments count on the day of
their slot, in your account's time zone.

API: GET /analytics/summary

```
yesly analytics retrieve-summary [flags]
```

### Examples

```
  yesly analytics retrieve-summary --range 30d --type delivery
```

### Options

```
  -h, --help              help for retrieve-summary
      --location string   Only this location.
      --range string      Days covered, ending today. One of: 7d, 30d, 90d.
      --type string       Only this appointment type.
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

* [yesly analytics](yesly_analytics.md)	 - The figures behind the dashboard's Analytics page

