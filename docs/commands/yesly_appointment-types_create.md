## yesly appointment-types create

Create an appointment type

### Synopsis

Create an appointment type.

Adds an appointment type. Its name and time format are used in the messages sent from then on.

API: POST /appointment_types

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly appointment-types create [flags]
```

### Examples

```
  yesly appointment-types create --id delivery --name "Home delivery" \
    --time-format window --place customer --first-contact '{"days_before": 1, "at": "18:00"}' \
    --reschedule.min-notice-hours 24 --reschedule.window-days 14
  yesly appointment-types create --id repair --name "Boiler repair" \
    --time-format exact --place customer --first-contact '{"hours_before": 3}'
```

### Options

```
      --cancellable                       Default true.
  -d, --data string                       JSON body, inline or @file (@- for stdin)
      --first-contact string              When the first message goes out. Default: the day before at 10:00. JSON, inline or @file.
  -h, --help                              help for create
      --id string                         (required) 1 to 40 lowercase letters, numbers, - and _. Cannot change later.
      --idempotency-key string            Idempotency-Key header (default: a new UUID, reused on retries)
      --instructions string               Optional instructions for the customer. Not included in messages today.
      --name string                       (required) Name shown to the customer.
      --place string                      (required) customer (at the customer's address) or location (at one of your locations). One of: customer, location.
      --reschedulable                     Default true.
      --reschedule.min-notice-hours int   The new time is at least this many hours later.
      --reschedule.window-days int        The new time is at most this many days later.
      --time-format string                (required) window ("between 9:00 and 11:00") or exact ("at 17:30"). One of: window, exact.
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

* [yesly appointment-types](yesly_appointment-types.md)	 - What each appointment is, how its time is shown and the instructions in the message

