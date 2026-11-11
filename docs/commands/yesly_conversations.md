## yesly conversations

Calls and threads of messages between an agent and a contact

### Synopsis

Calls and threads of messages between an agent and a contact.

```
yesly conversations [flags]
```

### Options

```
  -h, --help   help for conversations
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
* [yesly conversations cancel](yesly_conversations_cancel.md)	 - Cancel a conversation
* [yesly conversations create](yesly_conversations_create.md)	 - Create a conversation
* [yesly conversations delete](yesly_conversations_delete.md)	 - Delete a conversation
* [yesly conversations list](yesly_conversations_list.md)	 - List conversations
* [yesly conversations list-messages](yesly_conversations_list-messages.md)	 - List a conversation's messages
* [yesly conversations recording](yesly_conversations_recording.md)	 - Retrieve a call recording
* [yesly conversations retrieve](yesly_conversations_retrieve.md)	 - Retrieve a conversation

