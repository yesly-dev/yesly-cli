## yesly senders list

List senders

### Synopsis

List senders.

The numbers, SMS sender names and email addresses of your account, in one page. Yesly sets them up with you.

API: GET /senders

```
yesly senders list [flags]
```

### Examples

```
  yesly senders list --agent agt_8Hd2kQ4mZ1LpXy --channel whatsapp
```

### Options

```
      --agent string     Only this agent's senders.
      --all              follow every page and print all items as one JSON array
      --channel string   Only this channel. One of: call, sms, whatsapp, email.
  -h, --help             help for list
```

### Options inherited from parent commands

```
      --api-key string       API key; overrides YESLY_API_KEY and the profile
      --api-version string   Yesly-Version header (default "2026-11-11")
      --base-url string      API base URL; overrides YESLY_BASE_URL and the profile (default https://api.yesly.dev/v1)
      --compact              print JSON on one line
      --dry-run              print the request instead of sending it
      --live                 allow requests that change data with a live key
      --max-retries int      retries after a network error, 408, 429 or 5xx (default 2)
  -p, --profile string       profile saved by "yesly login" (default "default", or YESLY_PROFILE)
      --timeout duration     time limit for each attempt (default 30s)
```

### SEE ALSO

* [yesly senders](yesly_senders.md)	 - The numbers, SMS sender names and addresses your agents speak from

