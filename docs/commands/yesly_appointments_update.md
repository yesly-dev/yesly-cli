## yesly appointments update

Update an appointment

### Synopsis

Update an appointment.

Changes an appointment. Send only what changes. `metadata` is merged:
keys you send are set, a key sent as an empty string is removed.

If you change the slot or the customer's phone, Yesly starts the
confirmation again: the status goes back to `scheduled`, any previous
answer is cleared and a new `first_contact_at` is computed. Changes to
metadata or the name never contact the customer.

API: PATCH /appointments/{id}

Arguments:
  <id>  Id of the appointment, prefixed apt_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly appointments update <id> [flags]
```

### Examples

```
  yesly appointments update apt_… --slot.start 2026-11-14T09:00:00+01:00 --slot.end 2026-11-14T11:00:00+01:00
  yesly appointments update apt_… --metadata route_id=BIO-09
```

### Options

```
      --channels strings        New channel order for this appointment. One of: whatsapp, sms, voice, email.
      --customer.email string   New email; null removes it.
      --customer.name string    New name.
      --customer.phone string   Phone number in E.164 format, such as +34612345678.
  -d, --data string             JSON body, inline or @file (@- for stdin)
  -h, --help                    help for update
      --location string         New location id.
      --metadata stringArray    Keys to set, merged with the existing ones. An empty string removes the key. Repeat as key=value.
      --slot.end string         New end of the window.
      --slot.start string       New start, with offset.
      --type string             New appointment type id.
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

* [yesly appointments](yesly_appointments.md)	 - Visits, deliveries or meetings Yesly confirms with your customers

