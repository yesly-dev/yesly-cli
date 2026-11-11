## yesly analytics summary

Summary

### Synopsis

Summary.

The figures behind the dashboard's Analytics page, for your own reports: conversations, outcomes, usage and spend, in total and day by day. Filter by agent, channel or direction.

API: GET /analytics/summary

```
yesly analytics summary [flags]
```

### Examples

```
  yesly analytics summary --range 30d --agent agt_8Hd2kQ4mZ1LpXy
```

### Options

```
      --agent string       Only this agent's conversations.
      --channel string     Only conversations whose last channel is this one. One of: call, sms, whatsapp, email.
      --direction string   Only outbound or only inbound. One of: outbound, inbound.
  -h, --help               help for summary
      --range string       Days covered, ending today. One of: 7d, 30d, 90d.
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

* [yesly analytics](yesly_analytics.md)	 - The figures behind the dashboard's Analytics page

