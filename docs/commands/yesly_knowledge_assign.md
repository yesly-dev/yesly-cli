## yesly knowledge assign

Assign a document to an agent

### Synopsis

Assign a document to an agent.

The agent gets a copy of the document. Assigning it again changes nothing. Up to 30 documents per agent. Keys that can write, and owners and admins in the dashboard; keys limited to one agent get `403` `api_key_agent_scope`.

API: POST /knowledge/{id}/agents

Arguments:
  <id>  Id of the document, prefixed doc_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly knowledge assign <id> [flags]
```

### Examples

```
  yesly knowledge assign doc_4Rk7Tn2Pq9Ws1L --agent agt_8Hd2kQ4mZ1LpXy
```

### Options

```
      --agent string             (required) The agent, prefixed agt_.
  -d, --data string              JSON body, inline or @file (@- for stdin)
  -h, --help                     help for assign
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

* [yesly knowledge](yesly_knowledge.md)	 - Your company's documents, and which agents know each one

