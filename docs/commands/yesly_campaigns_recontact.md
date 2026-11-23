## yesly campaigns recontact

Launch a re-contact

### Synopsis

Launch a re-contact.

Starts one new campaign per group of Get a campaign's re-contact groups, with the same agent, channel and plan, and each person's context from their conversation. The groups' people are worked out again when it is called, and who talked to the agent less than 20 hours before the group's start is left out. Each group of a campaign is launched once; a second time gets `409` `recontact_launched`, and a group nobody can be contacted in gets `409` `recontact_empty`. The agent remembers what each contact said the time before.

API: POST /campaigns/{id}/recontact

Arguments:
  <id>  Id of the campaign, prefixed cmp_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly campaigns recontact <id> [flags]
```

### Examples

```
  yesly campaigns recontact cmp_5Tn8Wq3Zr1Vb7K --data '{"groups":[{"group":"no_answer"}]}'
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
      --groups string            (required) JSON, inline or @file; a plain word, such as an id, is sent as a string.
  -h, --help                     help for recontact
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --rate int                 Conversations a minute. By default the campaign's.
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

