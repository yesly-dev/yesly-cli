## yesly conversations list

List conversations

### Synopsis

List conversations.

Newest first, by creation.

API: GET /conversations

```
yesly conversations list [flags]
```

### Examples

```
  yesly conversations list --agent agt_8Hd2kQ4mZ1LpXy --status no_reply --from 2026-12-01
  yesly conversations list --contact ctc_3Kf9Lq2Wm7Rt4B --all
```

### Options

```
      --agent string            Only this agent's.
      --all                     follow every page and print all items as one JSON array
      --channel string          Only conversations whose current or last channel is this one. One of: call, sms, whatsapp, email.
      --contact string          Only this contact's.
      --direction string        Only outbound or only inbound. One of: outbound, inbound.
      --from string             Created on or after this day or time. A day starts at 00:00 in your account's time zone.
  -h, --help                    help for list
      --limit int               Between 1 and 100. Default 20.
      --outcome string          Only conversations that ended with this outcome code.
      --starting-after string   Cursor: the id of the last item of the previous page.
      --status string           Only in this status. One of: scheduled, in_progress, ended, no_reply, cancelled, failed.
      --to string               Created on or before this day or time. A day ends at 23:59:59 in your account's time zone.
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

* [yesly conversations](yesly_conversations.md)	 - Calls and threads of messages between an agent and a contact

