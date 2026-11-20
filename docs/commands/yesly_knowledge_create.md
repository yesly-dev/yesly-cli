## yesly knowledge create

Create a document

### Synopsis

Create a document.

Adds a document to the library: a file in `data` (base64) or Markdown in `text`, one of the two. Its text is cut into parts, each under the heading it falls in. Send `agents` to assign it at once. Up to 100 documents per library. Keys that can write, and owners and admins in the dashboard; keys limited to one agent get `403` `api_key_agent_scope`.

API: POST /knowledge

Flags set the body's fields; --data sends a whole JSON body, and flags given with it override its fields.

```
yesly knowledge create [flags]
```

### Examples

```
  yesly knowledge create --name opening-hours.md --text "$(cat opening-hours.md)" --agents agt_8Hd2kQ4mZ1LpXy
  yesly knowledge create --name prices.pdf --file-data "$(base64 < prices.pdf)"
```

### Options

```
      --agents strings           Agents to assign it to at once.
  -d, --data string              JSON body, inline or @file (@- for stdin)
      --file-data string         The file in base64 (a data: URL prefix is accepted): PDF, Word (.docx), text, Markdown, CSV or HTML, up to 3 MB. A scanned PDF has no text: send it as text or Word.
  -h, --help                     help for create
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

* [yesly knowledge](yesly_knowledge.md)	 - Your company's documents, and which agents know each one

