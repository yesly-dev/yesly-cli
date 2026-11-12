## yesly balance retrieve

Retrieve the balance

### Synopsis

Retrieve the balance.

Yesly is prepaid. The balance belongs to your organization: every account in it spends from the same one, and this endpoint returns it whichever account's key you use. Each SMS, WhatsApp message and email, and each started minute of a call that was picked up, is taken from it at your account's prices; calls nobody picks up are free. A message that could not be sent is refunded with an `adjustment`. A live call holds its longest cost (the agent's `max_call_minutes`) until it ends. Top-ups are made by card in the dashboard and cannot be made through the API.

When the balance does not cover the next message, it is refused (`402`, or the conversation ends `failed`). Watch `balance.low`. Test mode always returns a balance of 0 and never blocks.

API: GET /balance

```
yesly balance retrieve [flags]
```

### Examples

```
  yesly balance retrieve
```

### Options

```
  -h, --help   help for retrieve
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

* [yesly balance](yesly_balance.md)	 - The prepaid balance of your organization, shared by all its accounts, and every charge and top-up

