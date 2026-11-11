## yesly conversations cancel

Cancel a conversation

### Synopsis

Cancel a conversation.

Stops a `scheduled` or `in_progress` conversation: nothing more is sent, and a call in progress is hung up. `conversation.cancelled` fires. Anything else gets `409` `conversation_not_cancellable`.

API: POST /conversations/{id}/cancel

Arguments:
  <id>  Id of the conversation, prefixed cnv_.

```
yesly conversations cancel <id> [flags]
```

### Examples

```
  yesly conversations cancel cnv_7Qp4Ns2Xk9Lm3D
```

### Options

```
  -h, --help                     help for cancel
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
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

