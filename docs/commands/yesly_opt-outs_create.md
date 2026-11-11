## yesly opt-outs create

Add an opt-out

### Synopsis

Add an opt-out.

Load the opt-outs you already have: nothing more goes to that address on that channel, from any agent. `contact.opted_out` fires.

API: POST /opt_outs

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly opt-outs create [flags]
```

### Examples

```
  yesly opt-outs create --channel sms --address +34600000001
```

### Options

```
      --address string           (required) A phone in E.164 or an email address.
      --channel string           (required) One of: call, sms, whatsapp, email.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for create
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

* [yesly opt-outs](yesly_opt-outs.md)	 - Addresses that receive nothing more on a channel

