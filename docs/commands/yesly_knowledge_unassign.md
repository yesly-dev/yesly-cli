## yesly knowledge unassign

Unassign a document from an agent

### Synopsis

Unassign a document from an agent.

The agent's copy is deleted; the library keeps the document. Keys that can write, and owners and admins in the dashboard; keys limited to one agent get `403` `api_key_agent_scope`.

API: DELETE /knowledge/{id}/agents/{agent}

Arguments:
  <id>  Id of the document, prefixed doc_.
  <agent>  Id of the agent, prefixed agt_.

```
yesly knowledge unassign <id> <agent> [flags]
```

### Examples

```
  yesly knowledge unassign doc_4Rk7Tn2Pq9Ws1L agt_8Hd2kQ4mZ1LpXy
```

### Options

```
  -h, --help   help for unassign
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

* [yesly knowledge](yesly_knowledge.md)	 - Your company's documents, and which agents know each one

