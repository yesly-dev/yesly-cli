## yesly contacts merge

Merge contacts

### Synopsis

Merge contacts.

The contact in the body joins this one: its identifiers, conversations and memory move here, and it goes away. `contact.merged` fires. When both have an `external_id`, this one's stays.

API: POST /contacts/{id}/merge

Arguments:
  <id>  Id of the contact, prefixed ctc_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly contacts merge <id> [flags]
```

### Examples

```
  yesly contacts merge ctc_3Kf9Lq2Wm7Rt4B --contact ctc_…
```

### Options

```
      --contact string           (required) Id of the contact that joins this one and goes away.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for merge
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

* [yesly contacts](yesly_contacts.md)	 - Your customers, their identifiers, and what your agents remember about them

