## yesly agents create-variant

Create a variant

### Synopsis

Create a variant.

A named change on the agent to try in a campaign against the agent as it is: another voice, instruction sections added after the agent's, or both. Up to 10 per agent; an 11th gets `409` `limit_reached`.

API: POST /agents/{id}/variants

Arguments:
  <id>  Id of the agent, prefixed agt_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly agents create-variant <id> [flags]
```

### Examples

```
  yesly agents create-variant agt_8Hd2kQ4mZ1LpXy --name "Free first visit" --data '{"prompt":["Offer the first visit free when the contact hesitates."]}'
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for create-variant
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --name string              (required) 
      --prompt strings           
      --voice string             A voice id, as in the agent's voice.
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

