## yesly webhooks sign

Print the Yesly-Signature header for a payload

### Synopsis

Reads the payload from --payload (inline or @file) or from stdin, and prints the Yesly-Signature value the API would send with it.

```
yesly webhooks sign [flags]
```

### Examples

```
  yesly webhooks sign --secret whsec_… --payload @event.json
  curl localhost:3000/webhook -H "Yesly-Signature: $(yesly webhooks sign --secret whsec_… --payload @event.json)" \
    -H 'Content-Type: application/json' --data-binary @event.json
```

### Options

```
  -h, --help             help for sign
      --payload string   payload, inline or @file (default: stdin)
      --secret string    (required) the endpoint's signing secret, whsec_…
      --timestamp int    Unix time to sign with (default: now)
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

