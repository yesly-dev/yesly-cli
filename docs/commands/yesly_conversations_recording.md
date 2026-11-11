## yesly conversations recording

Retrieve a call recording

### Synopsis

Retrieve a call recording.

A link, valid 15 minutes, to a call's audio (WAV). Without `message`, the conversation's last recorded call. Needs a key with full access (`403` `api_key_scope` otherwise); `404` when there is no recording, or it expired. Every link is in your account's access log.

API: GET /conversations/{id}/recording

Arguments:
  <id>  Id of the conversation, prefixed cnv_.

```
yesly conversations recording <id> [flags]
```

### Examples

```
  yesly conversations recording cnv_7Qp4Ns2Xk9Lm3D --message msg_…
```

### Options

```
  -h, --help             help for recording
      --message string   The call, prefixed msg_.
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

* [yesly conversations](yesly_conversations.md)	 - Calls and threads of messages between an agent and a contact

