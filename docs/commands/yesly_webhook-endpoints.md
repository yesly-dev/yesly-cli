## yesly webhook-endpoints

URLs that receive events

### Synopsis

URLs that receive events.

```
yesly webhook-endpoints [flags]
```

### Options

```
  -h, --help   help for webhook-endpoints
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
* [yesly webhook-endpoints create](yesly_webhook-endpoints_create.md)	 - Create an endpoint
* [yesly webhook-endpoints delete](yesly_webhook-endpoints_delete.md)	 - Delete an endpoint
* [yesly webhook-endpoints list](yesly_webhook-endpoints_list.md)	 - List endpoints
* [yesly webhook-endpoints roll-secret](yesly_webhook-endpoints_roll-secret.md)	 - Roll the signing secret
* [yesly webhook-endpoints test](yesly_webhook-endpoints_test.md)	 - Send a test event
* [yesly webhook-endpoints update](yesly_webhook-endpoints_update.md)	 - Update an endpoint

