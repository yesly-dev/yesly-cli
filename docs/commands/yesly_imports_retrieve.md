## yesly imports retrieve

Retrieve an import

### Synopsis

Retrieve an import.

Shows how an import went. Items that fail do not stop the others; each error points to the item by its position and `external_id`.

API: GET /imports/{id}

Arguments:
  <id>  Id of the import, prefixed imp_.

```
yesly imports retrieve <id> [flags]
```

### Examples

```
  yesly imports retrieve imp_…
```

### Options

```
  -h, --help   help for retrieve
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

* [yesly imports](yesly_imports.md)	 - Results of appointment imports

