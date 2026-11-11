## yesly messages create

Send a message

### Synopsis

Send a message.

One SMS, WhatsApp message or email to a contact, from an agent's sender, outside any conversation. It follows the same rules as an agent's messages: the opt-out (`409` `contact_opted_out`), your contact hours (`422` `outside_contact_hours`), WhatsApp's 24-hour rule (`422` `template_required`), the agent's spend limit and the balance (`402`). A reply from the contact opens a conversation with the agent.

API: POST /messages

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly messages create [flags]
```

### Examples

```
  yesly messages create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B \
    --channel sms --text "Hola Marta, tu pedido 8841 sale hoy."
  yesly messages create --agent agt_8Hd2kQ4mZ1LpXy --contact ctc_3Kf9Lq2Wm7Rt4B --channel whatsapp \
    --template.name delivery_reminder --template.language es --template.variables 8841
```

### Options

```
      --agent string                 (required) The agent it goes from, prefixed agt_: its sender for the channel is used.
      --channel string               (required) A channel that carries text. One of: sms, whatsapp, email.
      --contact string               (required) A contact id, or the data to find or create one. JSON, inline or @file; a plain word, such as an id, is sent as a string.
  -d, --data string                  JSON body, inline or @file (@- for stdin)
  -h, --help                         help for create
      --idempotency-key string       Idempotency-Key header (default: a new UUID, reused on retries)
      --subject string               Email subject. Default: your company's name.
      --template.language string     
      --template.name string         
      --template.variables strings   Values for {{1}}, {{2}}… in order.
      --text string                  The text. SMS up to 1,600 characters, WhatsApp 4,096, email 20,000.
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

* [yesly messages](yesly_messages.md)	 - Each SMS, WhatsApp message, email and call; and one message outside a conversation

