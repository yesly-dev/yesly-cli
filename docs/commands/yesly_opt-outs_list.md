## yesly opt-outs list

List opt-outs

### Synopsis

List opt-outs.

Newest first.

API: GET /opt_outs

```
yesly opt-outs list [flags]
```

### Examples

```
  yesly opt-outs list --channel sms
```

### Options

```
      --address string          Only this phone or email.
      --all                     follow every page and print all items as one JSON array
      --channel string          Only this channel. One of: call, sms, whatsapp, email.
      --contact string          The opt-outs of every phone and email of this contact.
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

* [yesly opt-outs](yesly_opt-outs.md)	 - Addresses that receive nothing more on a channel

