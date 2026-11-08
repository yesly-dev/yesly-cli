## yesly webhooks verify

Check a Yesly-Signature header against a payload

### Synopsis

Reads the payload from --payload (inline or @file) or from stdin. Exits 0 when the header matches and is within the tolerance, 1 otherwise.

```
yesly webhooks verify [flags]
```

### Examples

```
  yesly webhooks verify --secret whsec_… --header 't=1762531200,v1=5f0c…' --payload @body.json
```

### Options

```
      --header string        (required) the Yesly-Signature value
  -h, --help                 help for verify
      --payload string       raw payload, inline or @file (default: stdin)
      --secret string        (required) the endpoint's signing secret, whsec_…
      --tolerance duration   maximum age of the timestamp; 0 turns the check off (default 5m0s)
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

* [yesly webhooks](yesly_webhooks.md)	 - Sign and verify webhook payloads locally

