## yesly knowledge list

List documents

### Synopsis

List documents.

Every document of the knowledge library, newest first, in one page, with the agents each is assigned to. Keys limited to one agent get `403` `api_key_agent_scope`: they list their agent's documents instead.

API: GET /knowledge

```
yesly knowledge list [flags]
```

### Examples

```
  yesly knowledge list --all
```

### Options

```
      --all    follow every page and print all items as one JSON array
  -h, --help   help for list
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

