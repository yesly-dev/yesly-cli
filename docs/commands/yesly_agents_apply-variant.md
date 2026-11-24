## yesly agents apply-variant

Apply a variant

### Synopsis

Apply a variant.

Changes the agent as the variant does: its voice, and its instruction sections after the agent's. The change is listed in the agent's `changes` and can be reverted with Revert a change. Usually the winning arm of a campaign's `experiment`.

API: POST /agents/{id}/variants/{variant}/apply

Arguments:
  <id>  Id of the agent, prefixed agt_.
  <variant>  Id of the variant, prefixed var_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly agents apply-variant <id> <variant> [flags]
```

### Examples

```
  yesly agents apply-variant agt_8Hd2kQ4mZ1LpXy var_4Hn2Lq8Wz5Rt1K --campaign cmp_5Tn8Wq3Zr1Vb7K
```

### Options

```
      --campaign string          The campaign whose experiment the variant won, kept with the change.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for apply-variant
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

