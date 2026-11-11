## yesly contacts delete-memory

Delete a memory record

### Synopsis

Delete a memory record.

API: DELETE /contacts/{id}/memory/{record}

Arguments:
  <id>  Id of the contact, prefixed ctc_.
  <record>  Id of the memory record.

```
yesly contacts delete-memory <id> <record> [flags]
```

### Examples

```
  yesly contacts delete-memory ctc_3Kf9Lq2Wm7Rt4B mem-7f3a9c2e4b1d4e8a
```

### Options

```
  -h, --help   help for delete-memory
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

