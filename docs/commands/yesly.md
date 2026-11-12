## yesly

Work with the Yesly API from your terminal

### Synopsis

yesly calls every operation of the Yesly API, signs and checks webhook
payloads, and forwards your account's events to a local server.

Authenticate once with "yesly login", or set YESLY_API_KEY.

### Options

```
      --api-key string       API key; overrides YESLY_API_KEY and the profile
      --api-version string   Yesly-Version header (default "2026-11-11")
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

* [yesly agents](yesly_agents.md)	 - The assistants Yesly built for you, and the settings you control
* [yesly analytics](yesly_analytics.md)	 - The figures behind the dashboard's Analytics page
* [yesly api](yesly_api.md)	 - Send any request to the API
* [yesly balance](yesly_balance.md)	 - The prepaid balance of your organization, shared by all its accounts, and every charge and top-up
* [yesly campaigns](yesly_campaigns.md)	 - One agent contacting many people at a steady pace
* [yesly completion](yesly_completion.md)	 - Generate the autocompletion script for the specified shell
* [yesly contacts](yesly_contacts.md)	 - Your customers, their identifiers, and what your agents remember about them
* [yesly conversations](yesly_conversations.md)	 - Calls and threads of messages between an agent and a contact
* [yesly events](yesly_events.md)	 - Everything that happened, as delivered to webhook endpoints
* [yesly listen](yesly_listen.md)	 - Receive your account's events locally, without a public URL
* [yesly login](yesly_login.md)	 - Save an API key in a profile
* [yesly logout](yesly_logout.md)	 - Remove a saved profile
* [yesly messages](yesly_messages.md)	 - Each SMS, WhatsApp message, email and call; and one message outside a conversation
* [yesly opt-outs](yesly_opt-outs.md)	 - Addresses that receive nothing more on a channel
* [yesly senders](yesly_senders.md)	 - The numbers, SMS sender names and addresses your agents speak from
* [yesly templates](yesly_templates.md)	 - WhatsApp templates, needed to write outside the 24-hour window
* [yesly version](yesly_version.md)	 - Print the CLI and API versions
* [yesly webhook-endpoints](yesly_webhook-endpoints.md)	 - URLs that receive events
* [yesly webhooks](yesly_webhooks.md)	 - Sign and verify webhook payloads locally
* [yesly whoami](yesly_whoami.md)	 - Show which key and base URL commands use, and check the key

