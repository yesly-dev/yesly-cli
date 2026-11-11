## yesly conversations create

Create a conversation

### Synopsis

Create a conversation.

Asks an agent to contact someone. The conversation is `scheduled` until `start_at` (now by default) falls inside contact hours; then the agent tries each channel in order, waiting for the contact to engage before the next. Follow it with `conversation.*` events. Send an `Idempotency-Key`: a retry then never contacts anyone twice.

API: POST /conversations

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly conversations create [flags]
```

### Examples

```
  yesly conversations create --agent agt_8Hd2kQ4mZ1LpXy \
    --contact '{"external_id": "customer-3391", "phone": "+34600000001", "name": "Marta"}' \
    --channels whatsapp,call --start-at 2026-12-01T10:00:00+01:00 \
    --context '{"order": "8841", "delivery_window": "2026-12-02T09:00/13:00"}' \
    --idempotency-key order-8841-delivery
  yesly conversations create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B \
    --channels call --context '{"order": "8902"}'
```

### Options

```
      --agent string             (required) Id of the agent, prefixed agt_.
      --channels strings         (required) The order to try, one to four channels without repeats. One channel is a single call or message; with several, the next is tried when the contact has not engaged within the agent's wait. One of: call, sms, whatsapp, email.
      --contact string           (required) A contact id, or the data to find or create one. JSON, inline or @file; a plain word, such as an id, is sent as a string.
      --context string           Your data for the agent, checked against its context_schema. JSON, inline or @file; a plain word, such as an id, is sent as a string.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for create
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --start-at string          Not before this time, and always inside contact hours. Default now. At most 90 days ahead.
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

