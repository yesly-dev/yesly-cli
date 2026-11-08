## yesly slots delete

(deprecated) Delete a slot

### Synopsis

Deprecated. Delete a slot.

Stops offering a slot, for example when your own system fills it. Appointments already moved into it stay where they are.

API: DELETE /slots/{id}

Arguments:
  <id>  Id of the slot, prefixed slt_.

```
yesly slots delete <id> [flags]
```

### Examples

```
  yesly slots delete slot_…
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --api-key string       API key; overrides YESLY_API_KEY and the profile
      --api-version string   Yesly-Version header (default "2026-10-01")
      --base-url string      API base URL; overrides YESLY_BASE_URL and the profile (default https://api.yesly.dev/v1)
      --compact              print JSON on one line
      --dry-run              print the request instead of sending it
      --live                 allow requests that change data with a live key
      --max-retries int      retries after a network error, 408, 429 or 5xx (default 2)
  -p, --profile string       profile saved by "yesly login" (default "default", or YESLY_PROFILE)
      --timeout duration     time limit for each attempt (default 30s)
```

### SEE ALSO

* [yesly slots](yesly_slots.md)	 - (deprecated) Free times published in Yesly

