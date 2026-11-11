## yesly contacts update

Update a contact

### Synopsis

Update a contact.

`phones` and `emails` replace the lists. An identifier of another contact gets `409` `contact_identifier_taken`.

API: PATCH /contacts/{id}

Arguments:
  <id>  Id of the contact, prefixed ctc_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly contacts update <id> [flags]
```

### Examples

```
  yesly contacts update ctc_3Kf9Lq2Wm7Rt4B --phones +34600000001,+34600000002 --metadata segment=standard
```

### Options

```
  -d, --data string            JSON body, inline or @file (@- for stdin)
      --emails strings         Replaces the emails.
      --external-id string     
  -h, --help                   help for update
      --language string        
      --metadata stringArray   Keys to set; a null value removes the key. Repeat as key=value.
      --name string            
      --phones strings         Replaces the phones.
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

