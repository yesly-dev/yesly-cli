## yesly knowledge

Your company's documents, and which agents know each one

### Synopsis

Your company's documents, and which agents know each one.

```
yesly knowledge [flags]
```

### Options

```
  -h, --help   help for knowledge
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

* [yesly](yesly.md)	 - Work with the Yesly API from your terminal
* [yesly knowledge assign](yesly_knowledge_assign.md)	 - Assign a document to an agent
* [yesly knowledge create](yesly_knowledge_create.md)	 - Create a document
* [yesly knowledge delete](yesly_knowledge_delete.md)	 - Delete a document
* [yesly knowledge list](yesly_knowledge_list.md)	 - List documents
* [yesly knowledge retrieve](yesly_knowledge_retrieve.md)	 - Retrieve a document
* [yesly knowledge unassign](yesly_knowledge_unassign.md)	 - Unassign a document from an agent
* [yesly knowledge update](yesly_knowledge_update.md)	 - Update a document

