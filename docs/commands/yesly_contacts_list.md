## yesly contacts list

List contacts

### Synopsis

List contacts.

Newest first. The filters match exactly.

API: GET /contacts

```
yesly contacts list [flags]
```

### Examples

```
  yesly contacts list --external-id customer-3391
  yesly contacts list --phone +34600000001
  yesly contacts list --all
```

### Options

```
      --all                     follow every page and print all items as one JSON array
      --email string            The contact with this email.
      --external-id string      The contact with this external_id.
  -h, --help                    help for list
      --limit int               Between 1 and 100. Default 20.
      --name string             Contacts whose name starts with this (two characters at least), ignoring case and accents, in name order.
      --phone string            The contact with this phone.
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

* [yesly contacts](yesly_contacts.md)	 - Your customers, their identifiers, and what your agents remember about them

