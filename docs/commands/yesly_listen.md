## yesly listen

Receive your account's events locally, without a public URL

### Synopsis

Polls GET /events for new events and prints each one as a line of JSON, or,
with --forward-to, posts it to a local URL signed like a webhook delivery,
so you can test your handler without exposing it to the internet.

Deliveries are signed with --secret, or with a secret this command creates
and saves in the profile; verify them with that secret in development. Events
arrive with up to --interval of delay, from the moment listen starts. This
does not create a webhook endpoint, and real endpoints keep receiving events.

```
yesly listen [flags]
```

### Examples

```
  yesly listen
  yesly listen --forward-to http://localhost:3000/yesly/webhook
  yesly listen --events appointment.confirmed,appointment.no_reply
```

### Options

```
      --events strings      only these event types (default: all)
      --forward-to string   URL to post each event to, signed with Yesly-Signature
  -h, --help                help for listen
      --interval duration   time between polls (default 2s)
      --secret string       signing secret for forwarded events (default: one saved in the profile)
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

* [yesly](yesly.md)	 - Work with the Yesly API from your terminal

