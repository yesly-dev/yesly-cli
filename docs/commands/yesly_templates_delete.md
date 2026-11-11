## yesly templates delete

Delete a template

### Synopsis

Delete a template.

Deletes it at Meta (through Telnyx) and here. Keys that can write, and owners and admins in the dashboard.

API: DELETE /templates/{id}

Arguments:
  <id>  Id of the template, prefixed tpl_.

```
yesly templates delete <id> [flags]
```

### Examples

```
  yesly templates delete tpl_5Rk2Wd8Mq3Lx7B
```

### Options

```
  -h, --help   help for delete
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

