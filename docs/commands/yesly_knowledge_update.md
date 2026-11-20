## yesly knowledge update

Update a document

### Synopsis

Update a document.

Renames the document (`name` alone) or replaces its content (`data` or `text`, and `name` if it changes too). Every agent that has it gets the new version. Keys that can write, and owners and admins in the dashboard; keys limited to one agent get `403` `api_key_agent_scope`.

API: PATCH /knowledge/{id}

Arguments:
  <id>  Id of the document, prefixed doc_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly knowledge update <id> [flags]
```

### Examples

```
  yesly knowledge update doc_4Rk7Tn2Pq9Ws1L --text "$(cat opening-hours.md)"
```

### Options

```
  -d, --data string        JSON body, inline or @file (@- for stdin)
      --file-data string   The new file, in base64.
  -h, --help               help for update
      --name string        
      --text string        The new text, in Markdown.
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

