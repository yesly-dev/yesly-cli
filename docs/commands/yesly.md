## yesly

Work with the Yesly API from your terminal

### Synopsis

yesly calls every operation of the Yesly API, signs and checks webhook
payloads, and forwards your account's events to a local server.

Authenticate once with "yesly login", or set YESLY_API_KEY.

### Options

```
      --api-key string       API key; overrides YESLY_API_KEY and the profile
      --api-version string   Yesly-Version header (default "2026-10-01")
      --base-url string      API base URL; overrides YESLY_BASE_URL and the profile (default https://api.yesly.dev/v1)
      --compact              print JSON on one line
      --dry-run              print the request instead of sending it
  -h, --help                 help for yesly
      --live                 allow requests that change data with a live key
      --max-retries int      retries after a network error, 408, 429 or 5xx (default 2)
  -p, --profile string       profile saved by "yesly login" (default "default", or YESLY_PROFILE)
      --timeout duration     time limit for each attempt (default 30s)
```

### SEE ALSO

* [yesly analytics](yesly_analytics.md)	 - The figures behind the dashboard's Analytics page
* [yesly api](yesly_api.md)	 - Send any request to the API
* [yesly appointment-types](yesly_appointment-types.md)	 - What each appointment is, how its time is shown and the instructions in the message
* [yesly appointments](yesly_appointments.md)	 - Visits, deliveries or meetings Yesly confirms with your customers
* [yesly balance](yesly_balance.md)	 - The prepaid balance and every charge and top-up
* [yesly completion](yesly_completion.md)	 - Generate the autocompletion script for the specified shell
* [yesly events](yesly_events.md)	 - Everything that happened, as delivered to webhook endpoints
* [yesly imports](yesly_imports.md)	 - Results of appointment imports
* [yesly listen](yesly_listen.md)	 - Receive your account's events locally, without a public URL
* [yesly locations](yesly_locations.md)	 - Your centres, shops or depots, and the phone calls are passed to
* [yesly login](yesly_login.md)	 - Save an API key in a profile
* [yesly logout](yesly_logout.md)	 - Remove a saved profile
* [yesly slots](yesly_slots.md)	 - (deprecated) Free times published in Yesly
* [yesly version](yesly_version.md)	 - Print the CLI and API versions
* [yesly webhook-endpoints](yesly_webhook-endpoints.md)	 - URLs that receive events
* [yesly webhooks](yesly_webhooks.md)	 - Sign and verify webhook payloads locally
* [yesly whoami](yesly_whoami.md)	 - Show which key and base URL commands use, and check the key

