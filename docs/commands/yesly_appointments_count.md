## yesly appointments count

Count appointments by status

### Synopsis

Count appointments by status.

How many appointments List appointments would return for each status,
with the same filters (all of them but `status`, `limit` and
`starting_after`). Without `date`, `from` or `to` it counts every
appointment the account still keeps. Counts with whole days (`date`,
or `from` and `to` as dates) or no dates, by type and location, come
from counters kept as appointments change, a few seconds behind.

API: GET /appointments/counts

```
yesly appointments count [flags]
```

### Examples

```
  yesly appointments count --date 2026-11-13
  yesly appointments count --from 2026-11-01 --to 2026-11-30 --type delivery
```

### Options

```
      --date string            Only appointments whose slot starts on this day, YYYY-MM-DD in your account's time zone.
      --external-id string     The appointment with this external_id.
      --from string            Slot starts on or after this day or time. A day starts at 00:00 in your account's time zone.
  -h, --help                   help for count
      --location string        Only this location.
      --metadata stringArray   Only appointments whose metadata has this value, such as metadata[route_id]=BIO-07. Repeat as key=value.
      --to string              Slot starts on or before this day or time. A day ends at 23:59:59 in your account's time zone.
      --type string            Only this appointment type.
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

