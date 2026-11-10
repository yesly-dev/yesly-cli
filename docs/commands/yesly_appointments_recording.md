## yesly appointments recording

Get a call recording

### Synopsis

Get a call recording.

A link to the audio of a conversational call (attempts with
`has_recording` `true`): a WAV file, customer on the left channel and
Yesly's agent on the right. The link works for 15 minutes; ask again
for a new one. Recordings are kept as long as the account's retention
allows and stored in the EU. Every link is logged (who and when, a
member or an API key by name); owners and admins see the log in the
dashboard. Needs a key with full access: a read-only key gets `403`
`api_key_scope`.

API: GET /appointments/{id}/attempts/{attempt}/recording

Arguments:
  <id>  Id of the appointment, prefixed apt_.
  <attempt>  Id of the call attempt, prefixed att_.

```
yesly appointments recording <id> <attempt> [flags]
```

### Examples

```
  yesly appointments recording apt_… att_…
```

### Options

```
  -h, --help   help for recording
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

