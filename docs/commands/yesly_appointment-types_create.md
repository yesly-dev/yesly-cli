## yesly appointment-types create

Create an appointment type

### Synopsis

Create an appointment type.

Adds an appointment type. Message templates pick it up at once through `{type}`, `{time}` and `{instructions}`.

API: POST /appointment_types

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly appointment-types create [flags]
```

### Examples

```
  yesly appointment-types create --id delivery --name "Home delivery" \
    --time-format window --minutes 120 --place customer
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for create
      --id string                (required) Lowercase letters, numbers, - and _. Cannot change later.
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --instructions string      Optional instructions for the customer.
      --minutes int              (required) Window length or duration.
      --name string              (required) Name shown to the customer.
      --place string             (required) customer (at the customer's address) or location (at one of your locations). One of: customer, location.
      --reschedulable            Default true.
      --time-format string       (required) window ("between 9:00 and 11:00") or exact ("at 17:30"). One of: window, exact.
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

