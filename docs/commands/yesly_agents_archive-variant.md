## yesly agents archive-variant

Archive a variant

### Synopsis

Archive a variant.

It can no longer be tried in new campaigns. Campaigns that already try it go on.

API: DELETE /agents/{id}/variants/{variant}

Arguments:
  <id>  Id of the agent, prefixed agt_.
  <variant>  Id of the variant, prefixed var_.

```
yesly agents archive-variant <id> <variant> [flags]
```

### Examples

```
  yesly agents archive-variant agt_8Hd2kQ4mZ1LpXy var_4Hn2Lq8Wz5Rt1K
```

### Options

```
  -h, --help   help for archive-variant
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

