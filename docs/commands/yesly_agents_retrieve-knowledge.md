## yesly agents retrieve-knowledge

Retrieve an agent's document

### Synopsis

Retrieve an agent's document.

The document with its parts, as the agent reads them.

API: GET /agents/{id}/knowledge/{document}

Arguments:
  <id>  Id of the agent, prefixed agt_.
  <document>  Id of the document, prefixed doc_.

```
yesly agents retrieve-knowledge <id> <document> [flags]
```

### Examples

```
  yesly agents retrieve-knowledge agt_8Hd2kQ4mZ1LpXy doc_4Rk7Tn2Pq9Ws1L
```

### Options

```
  -h, --help   help for retrieve-knowledge
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

