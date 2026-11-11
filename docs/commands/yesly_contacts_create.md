## yesly contacts create

Create or update a contact

### Synopsis

Create or update a contact.

Finds the contact by `external_id`, then by any of its phones, then by any of its emails, and adds what is new; creates it when there is none. Send your id and every identifier you know, so a call or a message from any of them reaches the same contact. An identifier that belongs to another contact gets `409` `contact_identifier_taken`.

API: POST /contacts

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly contacts create [flags]
```

### Examples

```
  yesly contacts create --external-id customer-3391 --name Marta \
    --phone +34600000001 --metadata segment=premium
  yesly contacts create --data @contact.json
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
      --email string             Email address.
      --emails strings           
      --external-id string       Your own id for the customer.
  -h, --help                     help for create
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --language string          
      --metadata stringArray     Your own key-value pairs, returned unchanged. Up to 20 keys, keys up to 40 characters, values up to 500 characters, strings only. Repeat as key=value.
      --name string              
      --phone string             Phone number in E.164 format, such as +34612345678.
      --phones strings           
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

* [yesly contacts](yesly_contacts.md)	 - Your customers, their identifiers, and what your agents remember about them

