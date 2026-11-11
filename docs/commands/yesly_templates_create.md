## yesly templates create

Create a template

### Synopsis

Create a template.

Sends a WhatsApp template to Meta for review. It starts `pending`; Meta usually answers within 48 hours. In test mode it is `approved` at once.

API: POST /templates

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly templates create [flags]
```

### Examples

```
  yesly templates create --name delivery_reminder --language es --category utility \
    --body "Hola, tu pedido {{1}} llega mañana." --examples 8841
```

### Options

```
      --body string              (required) 
      --category string          (required) One of: utility, marketing, authentication.
  -d, --data string              JSON body, inline or @file (@- for stdin)
      --examples strings         A sample value for each variable; Meta rejects templates without them.
  -h, --help                     help for create
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --language string          (required) Such as es.
      --name string              (required) Lowercase letters, digits and _.
      --sender string            The WhatsApp sender it is for, prefixed snd_. Default: the account's first.
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

* [yesly templates](yesly_templates.md)	 - WhatsApp templates, needed to write outside the 24-hour window

