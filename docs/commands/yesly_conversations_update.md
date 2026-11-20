## yesly conversations update

Reschedule a conversation

### Synopsis

Reschedule a conversation.

Changes the `start_at`, contact plan (`then`) or `context` of a `scheduled` conversation, before it starts. `conversation.rescheduled` fires. Once it started, ended or was cancelled it gets `409` `conversation_not_editable`. Fields you leave out stay as they were; `context: null` removes it.

API: PATCH /conversations/{id}

Arguments:
  <id>  Id of the conversation, prefixed cnv_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly conversations update <id> [flags]
```

### Examples

```
  yesly conversations update cnv_7Qp4Ns2Xk9Lm3D --start-at 2026-12-03T10:00:00+01:00
  yesly conversations update cnv_7Qp4Ns2Xk9Lm3D --context '{"order": "8841", "delivery_window": "2026-12-03T09:00/13:00"}'
```

### Options

```
      --context string           Your data for the agent, checked against its context_schema; null removes it. JSON, inline or @file; a plain word, such as an id, is sent as a string.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for update
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --start-at string          The new time, with an offset. Not before it, and always inside contact hours. At most 90 days ahead.
      --then string              The new contact plan, or null for none. JSON, inline or @file; a plain word, such as an id, is sent as a string.
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

