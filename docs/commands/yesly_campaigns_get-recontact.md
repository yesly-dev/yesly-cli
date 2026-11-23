## yesly campaigns get-recontact

Get a campaign's re-contact groups

### Synopsis

Get a campaign's re-contact groups.

Who of a completed or cancelled campaign did not convert, grouped by why, from each conversation's status and signals: `no_answer` (no reply, voicemail), `other_person` (did not reach who decides), `call_later` (asked to be called at another time), `objection` (price, another provider, no slot), `unanswered` (asked something the agent could not answer) and `cut_off` (the conversation broke off). Never in a group: who converted or reached the goal, who said no (an outcome whose result is `failure`, or not interested), a wrong contact, who opted out on the campaign's channel, and who already had 3 conversations with the agent in the last 30 days. A campaign still running gets `409` `campaign_not_finished`. The recipients carry the contact's own words about when to call, so a read-only key gets `403`.

API: GET /campaigns/{id}/recontact

Arguments:
  <id>  Id of the campaign, prefixed cmp_.

```
yesly campaigns get-recontact <id> [flags]
```

### Examples

```
  yesly campaigns get-recontact cmp_5Tn8Wq3Zr1Vb7K
```

### Options

```
  -h, --help   help for get-recontact
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

