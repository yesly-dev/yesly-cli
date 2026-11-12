## yesly campaigns pause

Pause a campaign

### Synopsis

Pause a campaign.

Stops starting new conversations; the ones already started go on. Only a `scheduled`, `running` or `waiting` campaign: anything else gets `409` `campaign_not_pausable`. Keys that can write, and members, admins and owners in the dashboard.

API: POST /campaigns/{id}/pause

Arguments:
  <id>  Id of the campaign, prefixed cmp_.

```
yesly campaigns pause <id> [flags]
```

### Examples

```
  yesly campaigns pause cmp_5Tn8Wq3Zr1Vb7K
```

### Options

```
  -h, --help                     help for pause
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
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

* [yesly campaigns](yesly_campaigns.md)	 - One agent contacting many people at a steady pace

