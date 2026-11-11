## yesly conversations list-messages

List a conversation's messages

### Synopsis

List a conversation's messages.

What the contact and the agent said or wrote, in order, on every channel, read from the agent's memory. All of it, in one page. Empty once the transcript has expired or been deleted (the conversation's `transcript` says which). Every read is in your account's access log.

API: GET /conversations/{id}/messages

Arguments:
  <id>  Id of the conversation, prefixed cnv_.

```
yesly conversations list-messages <id> [flags]
```

### Examples

```
  yesly conversations list-messages cnv_7Qp4Ns2Xk9Lm3D
```

### Options

```
      --all    follow every page and print all items as one JSON array
  -h, --help   help for list-messages
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

