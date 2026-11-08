## yesly appointments create

Create an appointment

### Synopsis

Create an appointment.

Creates an appointment and schedules the first contact. By default Yesly
writes the day before at 10:00 and moves through your channels if there
is no answer; if the appointment is booked after that moment, contact
starts at the next allowed time within your contact hours.

WhatsApp and SMS need a mobile number. If `customer.phone` is a
landline, those channels are skipped and only the call runs.

Idempotent by `external_id`: the same `external_id` with the same data
returns the existing appointment with `200`; with different data, `409`.

API: POST /appointments

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly appointments create [flags]
```

### Examples

```
  yesly appointments create --external-id ORD-48213 --type delivery \
    --customer.name Martha --customer.phone +34612345678 \
    --slot.start 2026-11-13T09:00:00+01:00 --slot.end 2026-11-13T11:00:00+01:00 \
    --metadata order_id=PV-18342 --metadata route_id=BIO-07
  yesly appointments create --data @appointment.json
```

### Options

```
      --channels strings        Override the channel order for this appointment, such as ["whatsapp", "voice"]. Only channels enabled in your account. One of: whatsapp, sms, voice, email.
      --customer.email string   Email for the email channel. Optional.
      --customer.name string    (required) Name used in the messages. A first name is enough.
      --customer.phone string   (required) Phone number in E.164 format, such as +34612345678.
  -d, --data string             JSON body, inline or @file (@- for stdin)
      --external-id string      (required) Your id for this appointment, such as the order or ticket number. Retrying with the same external_id is safe.
  -h, --help                    help for create
      --location string         Location id. Required when the type happens at one of your locations; otherwise optional, and used to pass calls to that location's phone.
      --metadata stringArray    Your own key-value pairs, returned unchanged. Up to 20 keys, keys up to 40 characters, values up to 500 characters, strings only. Do not put personal data here. Repeat as key=value.
      --slot.end string         End of the window. Leave it out for exact-time types; it is computed from the type's duration.
      --slot.start string       (required) Start of the appointment or of the time window, with offset.
      --type string             Appointment type id. Required if your account has more than one type.
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

