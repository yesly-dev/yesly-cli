## yesly balance list-transactions

List balance transactions

### Synopsis

List balance transactions.

Every charge, top-up and adjustment, newest first. Charges point to the agent, the conversation and the message, so you can pass the cost on to an order or a customer.

API: GET /balance/transactions

```
yesly balance list-transactions [flags]
```

### Examples

```
  yesly balance list-transactions --type usage --agent agt_8Hd2kQ4mZ1LpXy --from 2026-12-01 --all
```

### Options

```
      --agent string            Only charges of this agent.
      --all                     follow every page and print all items as one JSON array
      --conversation string     Only charges of this conversation.
      --from string             On or after this day or time. A day alone means 00:00 UTC.
  -h, --help                    help for list-transactions
      --limit int               Between 1 and 100. Default 20.
      --starting-after string   Cursor: the id of the last item of the previous page.
      --to string               On or before this day or time. A day alone means 23:59:59 UTC.
      --type string             usage, top_up or adjustment. One of: usage, top_up, adjustment.
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

* [yesly balance](yesly_balance.md)	 - The prepaid balance and every charge and top-up

