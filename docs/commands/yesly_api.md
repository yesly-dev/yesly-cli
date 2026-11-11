## yesly api

Send any request to the API

### Synopsis

Sends a request to any path of the API, relative to the base URL, with
your key, Yesly-Version, retries and, for POSTs the API makes idempotent, an
Idempotency-Key.

-f key=value adds a string; -F key=value adds a typed value (true, false,
null, numbers and JSON stay as they are, @file reads a file). For GET and
DELETE they become query parameters; otherwise they go in the JSON body,
where "contact.name" or "contact[name]" nests and "channels[]" appends.

```
yesly api <method> <path> [flags]
```

### Examples

```
  yesly api GET /conversations -f status=no_reply -F limit=5
  yesly api POST /conversations -f agent=agt_8Hd2kQ4mZ1LpXy -f contact.external_id=customer-3391 \
    -f contact.phone=+34600000001 -f channels[]=whatsapp -f channels[]=call -f context.order=8841
  yesly api PATCH /contacts/ctc_3Kf9Lq2Wm7Rt4B --data @changes.json
```

### Options

```
      --all                       follow every page of a list and print all items as one JSON array
  -d, --data string               JSON body, inline or @file (@- for stdin)
  -f, --field stringArray         add a string field, key=value
  -h, --help                      help for api
      --idempotency-key string    Idempotency-Key header
  -F, --typed-field stringArray   add a typed field, key=value
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

* [yesly](yesly.md)	 - Work with the Yesly API from your terminal

