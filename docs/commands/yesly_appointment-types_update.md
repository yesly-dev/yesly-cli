## yesly appointment-types update

Update an appointment type

### Synopsis

Update an appointment type.

Changes a type. Appointments already contacted keep the text they were sent; the next messages use the new one.

API: PATCH /appointment_types/{id}

Arguments:
  <id>  Id of the appointment type, such as delivery.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly appointment-types update <id> [flags]
```

### Examples

```
  yesly appointment-types update delivery --instructions "Have the order number at hand" --reschedulable=false
```

### Options

```
      --cancellable            Whether the customer may cancel.
  -d, --data string            JSON body, inline or @file (@- for stdin)
      --first-contact string   Its own first message; null goes back to your account's setting. JSON, inline or @file.
  -h, --help                   help for update
      --instructions string    New instructions. An empty string clears them.
      --minutes int            New window length or duration.
      --name string            New name.
      --reschedulable          Whether the customer may ask for another date.
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

