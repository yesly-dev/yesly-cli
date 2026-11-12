## yesly balance

The prepaid balance of your organization, shared by all its accounts, and every charge and top-up

### Synopsis

The prepaid balance of your organization, shared by all its accounts, and every charge and top-up.

```
yesly balance [flags]
```

### Options

```
  -h, --help   help for balance
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

* [yesly](yesly.md)	 - Work with the Yesly API from your terminal
* [yesly balance list-transactions](yesly_balance_list-transactions.md)	 - List balance transactions
* [yesly balance retrieve](yesly_balance_retrieve.md)	 - Retrieve the balance

