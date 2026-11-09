## yesly slots create

(deprecated) Publish slots

### Synopsis

Deprecated. Publish slots.

Publishes free slots in bulk, up to 500 per request. A slot with the
same type, location and start as an existing one replaces it, so you can
send your full availability every few hours; a replaced slot keeps its
`booked` count.

API: POST /slots

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly slots create [flags]
```

### Examples

```
  yesly slots create --slots @slots.json
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for create
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --slots string             (required) Between 1 and 500 slots. JSON, inline or @file.
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

