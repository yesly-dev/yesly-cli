## yesly agents revert-change

Revert a change

### Synopsis

Revert a change.

Puts the agent's voice and instructions back as they were before the change. Once only (`409` `change_reverted`), and only if nothing changed them since (`409` `change_superseded`).

API: POST /agents/{id}/changes/{change}/revert

Arguments:
  <id>  Id of the agent, prefixed agt_.
  <change>  Id of the change, prefixed chg_.

```
yesly agents revert-change <id> <change> [flags]
```

### Examples

```
  yesly agents revert-change agt_8Hd2kQ4mZ1LpXy chg_2Kd8Rq5Wn1Lz3P
```

### Options

```
  -h, --help                     help for revert-change
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

* [yesly agents](yesly_agents.md)	 - The assistants Yesly built for you, and the settings you control

