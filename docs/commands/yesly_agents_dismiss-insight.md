## yesly agents dismiss-insight

Dismiss a finding

### Synopsis

Dismiss a finding.

It is no longer listed, and the same pattern is not proposed again for 30 days.

API: POST /agents/{id}/insights/{insight}/dismiss

Arguments:
  <id>  Id of the agent, prefixed agt_.
  <insight>  Id of the finding, prefixed ins_.

```
yesly agents dismiss-insight <id> <insight> [flags]
```

### Examples

```
  yesly agents dismiss-insight agt_8Hd2kQ4mZ1LpXy ins_7Rk2Pw9Lm4Tq1Z
```

### Options

```
  -h, --help                     help for dismiss-insight
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

