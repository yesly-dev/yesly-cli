## yesly slots

(deprecated) Free times published in Yesly

### Synopsis

Deprecated. Free times published in Yesly. Day changes now come from your own system (live availability) or the customer's proposal; these operations keep working for accounts that use them.

```
yesly slots [flags]
```

### Options

```
  -h, --help   help for slots
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

* [yesly](yesly.md)	 - Work with the Yesly API from your terminal
* [yesly slots create](yesly_slots_create.md)	 - (deprecated) Publish slots
* [yesly slots delete](yesly_slots_delete.md)	 - (deprecated) Delete a slot
* [yesly slots list](yesly_slots_list.md)	 - (deprecated) List slots

