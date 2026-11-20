## yesly knowledge retrieve

Retrieve a document

### Synopsis

Retrieve a document.

The document with its parts (`content`) and, when it was written, its `text`, to edit it again. Keys limited to one agent get `403` `api_key_agent_scope`.

API: GET /knowledge/{id}

Arguments:
  <id>  Id of the document, prefixed doc_.

```
yesly knowledge retrieve <id> [flags]
```

### Examples

```
  yesly knowledge retrieve doc_4Rk7Tn2Pq9Ws1L
```

### Options

```
  -h, --help   help for retrieve
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

