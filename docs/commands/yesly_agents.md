## yesly agents

The assistants Yesly built for you, and the settings you control

### Synopsis

The assistants Yesly built for you, and the settings you control.

```
yesly agents [flags]
```

### Options

```
  -h, --help   help for agents
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
* [yesly agents add-knowledge](yesly_agents_add-knowledge.md)	 - Add a document to an agent
* [yesly agents apply-variant](yesly_agents_apply-variant.md)	 - Apply a variant
* [yesly agents archive-variant](yesly_agents_archive-variant.md)	 - Archive a variant
* [yesly agents create-variant](yesly_agents_create-variant.md)	 - Create a variant
* [yesly agents list](yesly_agents_list.md)	 - List agents
* [yesly agents list-knowledge](yesly_agents_list-knowledge.md)	 - List an agent's documents
* [yesly agents list-variants](yesly_agents_list-variants.md)	 - List an agent's variants
* [yesly agents pause](yesly_agents_pause.md)	 - Pause an agent
* [yesly agents remove-knowledge](yesly_agents_remove-knowledge.md)	 - Remove a document from an agent
* [yesly agents resume](yesly_agents_resume.md)	 - Resume an agent
* [yesly agents retrieve](yesly_agents_retrieve.md)	 - Retrieve an agent
* [yesly agents retrieve-knowledge](yesly_agents_retrieve-knowledge.md)	 - Retrieve an agent's document
* [yesly agents revert-change](yesly_agents_revert-change.md)	 - Revert a change
* [yesly agents update](yesly_agents_update.md)	 - Update an agent's settings

