## yesly messages retrieve

Retrieve a message

### Synopsis

Retrieve a message.

API: GET /messages/{id}

Arguments:
  <id>  Id of the message, prefixed msg_.

```
yesly messages retrieve <id> [flags]
```

### Examples

```
  yesly messages retrieve msg_…
```

### Options

```
  -h, --help   help for retrieve
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

* [yesly messages](yesly_messages.md)	 - Each SMS, WhatsApp message, email and call; and one message outside a conversation

