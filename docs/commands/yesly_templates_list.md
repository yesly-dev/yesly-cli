## yesly templates list

List templates

### Synopsis

List templates.

Your WhatsApp templates and Meta's review of each, in one page.

API: GET /templates

```
yesly templates list [flags]
```

### Examples

```
  yesly templates list --status approved
```

### Options

```
      --all             follow every page and print all items as one JSON array
  -h, --help            help for list
      --status string   Only templates in this state. One of: pending, approved, rejected, paused, disabled.
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

