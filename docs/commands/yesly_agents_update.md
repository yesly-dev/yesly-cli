## yesly agents update

Update an agent's settings

### Synopsis

Update an agent's settings.

Changes the settings you own: voice, limits, contact hours, recording, the disclosure sentence and the transfer targets. The agent's code, models, outcomes and context are Yesly's: ask for a new version. Accounts with EU data residency cannot choose a voice that is processed outside the EU (`voice_not_in_eu`).

API: PATCH /agents/{id}

Arguments:
  <id>  Id of the agent, prefixed agt_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly agents update <id> [flags]
```

### Examples

```
  yesly agents update agt_8Hd2kQ4mZ1LpXy --limits.messages-per-conversation 20 --recording.enabled
  yesly agents update agt_8Hd2kQ4mZ1LpXy --contact-hours @hours.json
```

### Options

```
      --contact-hours string                   The agent's own contact hours, or null to use the account's. JSON, inline or @file; a plain word, such as an id, is sent as a string.
  -d, --data string                            JSON body, inline or @file (@- for stdin)
      --disclosure string                      Cannot be empty (disclosure_required).
  -h, --help                                   help for update
      --limits.daily-spend float               
      --limits.max-call-minutes int            
      --limits.messages-per-conversation int   
      --recording.enabled                      
      --transfer-targets string                Replaces the list. Up to 10. JSON, inline or @file; a plain word, such as an id, is sent as a string.
      --voice string                           A voice of the catalog, or null for the agent's default.
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

* [yesly agents](yesly_agents.md)	 - The assistants Yesly built for you, and the settings you control

