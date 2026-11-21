## yesly conversations record-conversion

Record a conversion

### Synopsis

Record a conversion.

Says whether the conversation led to what your company wanted, as only you know it: the appointment was kept, the order paid, the debt settled. The agent's outcome is what the agent understood; this is what happened. It replaces the one recorded before, and is counted in analytics and in the conversation's campaign (`verified`). Only once the conversation has closed (`ended`, `no_reply`, `failed` or `cancelled`); before that, `409` `conversation_open`. A read-only key gets `403`.

API: POST /conversations/{id}/conversion

Arguments:
  <id>  Id of the conversation, prefixed cnv_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly conversations record-conversion <id> [flags]
```

### Examples

```
  yesly conversations record-conversion cnv_7Qp4Ns2Xk9Lm3D --converted=true
```

### Options

```
      --converted                (required) true if the conversation led to what your company wanted, false if it did not.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for record-conversion
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

