## yesly agents add-knowledge

Add a document to an agent

### Synopsis

Add a document to an agent.

Adds a document to the knowledge library and assigns it to this agent, as Create a document with `agents` set to this one. A key limited to this agent can use it. Keys that can write, and owners and admins in the dashboard.

API: POST /agents/{id}/knowledge

Arguments:
  <id>  Id of the agent, prefixed agt_.

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly agents add-knowledge <id> [flags]
```

### Examples

```
  yesly agents add-knowledge agt_8Hd2kQ4mZ1LpXy --name opening-hours.md --text "$(cat opening-hours.md)"
```

### Options

```
  -d, --data string              JSON body, inline or @file (@- for stdin)
      --file-data string         The file in base64 (a data: URL prefix is accepted): PDF, Word (.docx), text, Markdown, CSV or HTML, up to 3 MB. A scanned PDF has no text: send it as text or Word.
  -h, --help                     help for add-knowledge
      --idempotency-key string   Idempotency-Key header (default: a new UUID, reused on retries)
      --name string              (required) With the file's extension when it is a file: it helps tell the format. Kept to 120 characters; / and \ become _.
      --text string              The document written in Markdown; # headings head its parts. Up to 100,000 characters.
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

