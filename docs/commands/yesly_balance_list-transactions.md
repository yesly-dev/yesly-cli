## yesly balance list-transactions

List balance transactions

### Synopsis

List balance transactions.

Every charge, top-up and adjustment, newest first. Charges point to the appointment and its `external_id`, so you can pass the cost on to an order, route or customer. Refunds of failed messages (`adjustment`) do not point to the appointment.

API: GET /balance/transactions

```
yesly balance list-transactions [flags]
```

### Examples

```
  yesly balance list-transactions --type usage --from 2026-11-01 --all
```

### Options

```
      --all                     follow every page and print all items as one JSON array
      --appointment string      Only charges for this appointment. Same as subject with an appointment id.
      --from string             On or after this day or time. A day alone means 00:00 UTC.
  -h, --help                    help for list-transactions
      --limit int               Between 1 and 100. Default 20.
      --starting-after string   Cursor: the id of the last item of the previous page.
      --subject string          Only charges for this object, by its id (today an appointment, apt_…).
      --to string               On or before this day or time. A day alone means 23:59:59 UTC.
      --type string             usage, top_up or adjustment. One of: usage, top_up, adjustment.
      --use-case string         Only charges of this use case. Top-ups and adjustments belong to the account and are left out. One of: appointments.
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

