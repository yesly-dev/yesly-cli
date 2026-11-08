## yesly login

Save an API key in a profile

### Synopsis

Saves an API key, after checking it against the API, in the profile named
by --profile ("default" if not given). Paste the key at the prompt, or pipe
it in: echo "$KEY" | yesly login. --base-url is saved with the profile.

Create keys in the dashboard, under API keys. Keys starting with ysk_test_
run in test mode; ysk_live_ keys contact real customers and spend balance.

```
yesly login [flags]
```

### Examples

```
  yesly login
  yesly login --profile mock --base-url http://localhost:4010
```

### Options

```
  -h, --help   help for login
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

