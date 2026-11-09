## yesly balance retrieve

Retrieve the balance

### Synopsis

Retrieve the balance.

Yesly is prepaid. Each WhatsApp message, SMS, email and started minute
of a call that was picked up is taken from the balance at the public
price; calls nobody picks up are free. A message that could not be sent
is refunded with an `adjustment`. When a customer asks for another day
on a call, the SMS with the link to pick a day is charged as an SMS.
Top-ups are made by card in the dashboard and cannot be made through
the API.

If the balance does not cover the next message, contact for that
appointment stops: it stays `in_progress` and is not resumed after a
top-up. Appointments can still be created. Keep a balance that covers
your usage, and watch `balance.low`. Test mode always returns a
balance of 0 and never blocks.

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

* [yesly balance](yesly_balance.md)	 - The prepaid balance and every charge and top-up

