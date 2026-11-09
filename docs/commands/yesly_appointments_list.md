## yesly appointments list

List appointments

### Synopsis

List appointments.

Lists appointments. With `status`, `date`, `from`, `to` or `metadata`
the list is in slot order, soonest first; otherwise it is newest
first, by creation. A `date`, `from` or `to` filter without `status`
or `metadata` covers at most 62 days. A common use is to ask, each
afternoon, which of tomorrow's appointments ended with `no_reply` so
your team can call them.

API: GET /appointments

```
yesly appointments list [flags]
```

### Examples

```
  yesly appointments list --status no_reply --date 2026-11-13
  yesly appointments list --metadata route_id=BIO-07 --all
  yesly appointments list --external-id ORD-48213
```

### Options

```
      --all                     follow every page and print all items as one JSON array
      --date string             Only appointments whose slot starts on this day, YYYY-MM-DD in your account's time zone.
      --external-id string      The appointment with this external_id.
      --from string             Slot starts on or after this day or time. A day starts at 00:00 in your account's time zone.
  -h, --help                    help for list
      --limit int               Between 1 and 100. Default 20.
      --location string         Only this location.
      --metadata stringArray    Only appointments whose metadata has this value, such as metadata[route_id]=BIO-07. Repeat as key=value.
      --starting-after string   Cursor: the id of the last item of the previous page.
      --status string           Only appointments in this status. One of: scheduled, in_progress, confirmed, rescheduled, reschedule_requested, cancelled, no_reply.
      --to string               Slot starts on or before this day or time. A day ends at 23:59:59 in your account's time zone.
      --type string             Only this appointment type.
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

