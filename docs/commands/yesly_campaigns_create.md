## yesly campaigns create

Create a campaign

### Synopsis

Create a campaign.

Asks an agent to contact many people. Every recipient is checked before anything is stored; the first problem is the error, and `param` points at the row (`recipients[12].phone`, `recipients[3].email`, `recipients[7]`, `recipients[0].context`, `recipients[5].contact`). Once a minute, for each running campaign, Yesly starts the next `rate` conversations as ordinary outbound conversations, only inside the agent's contact hours in live mode (test mode ignores hours). The campaign is `scheduled` until `start_at`. A recipient is either an existing contact (`contact`) or the person's details, which find or create the contact when its conversation starts. Live phone numbers must be Spanish mobiles or landlines. Send an `Idempotency-Key`: a retry then never creates two campaigns.

API: POST /campaigns

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly campaigns create [flags]
```

### Examples

```
  yesly campaigns create --name "Delivery confirmations December" --agent agt_8Hd2kQ4mZ1LpXy \
    --channels whatsapp,call --start-at 2026-12-01T10:00:00+01:00 --rate 5 \
    --recipients @recipients.json --idempotency-key december-deliveries
  yesly campaigns create --data @campaign.json
```

### Options

```
      --agent string             (required) Id of the agent, prefixed agt_.
      --channels strings         (required) The order to try for every recipient, with the same rules as Create a conversation. One of: call, sms, whatsapp, email.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for create
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --name string              (required) 
      --rate int                 Conversations started per minute.
      --recipients string        (required) Who to contact, in order. JSON, inline or @file; a plain word, such as an id, is sent as a string.
      --start-at string          Not before this time, and always inside contact hours in live mode. Default now. At most 90 days ahead.
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

* [yesly campaigns](yesly_campaigns.md)	 - One agent contacting many people at a steady pace

