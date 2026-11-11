## yesly opt-outs delete

Delete an opt-out

### Synopsis

Delete an opt-out.

Only when the person asked to be contacted again. `contact.opted_in` fires.

API: DELETE /opt_outs/{id}

Arguments:
  <id>  Id of the opt-out, prefixed opt_.

```
yesly opt-outs delete <id> [flags]
```

### Examples

```
  yesly opt-outs delete opt_…
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

* [yesly opt-outs](yesly_opt-outs.md)	 - Addresses that receive nothing more on a channel

