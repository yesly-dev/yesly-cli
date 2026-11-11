## yesly messages list

List messages

### Synopsis

List messages.

Calls and messages, both ways, newest first. With `conversation`, that conversation's, oldest first.

API: GET /messages

```
yesly messages list [flags]
```

### Examples

```
  yesly messages list --conversation cnv_7Qp4Ns2Xk9Lm3D
```

### Options

```
      --all                     follow every page and print all items as one JSON array
      --contact string          Only this contact's.
      --conversation string     Only this conversation's.
  -h, --help                    help for list
      --limit int               Between 1 and 100. Default 20.
      --starting-after string   Cursor: the id of the last item of the previous page.
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

