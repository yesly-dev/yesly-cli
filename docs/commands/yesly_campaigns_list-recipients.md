## yesly campaigns list-recipients

List a campaign's recipients

### Synopsis

List a campaign's recipients.

The rows in the order they were sent, row 1 first, with the conversation each one started. A row that started shows the contact's current details (`null` once the contact is erased). Unlike other lists, `starting_after` is a row number: pass the `row` of the last item received.

API: GET /campaigns/{id}/recipients

Arguments:
  <id>  Id of the campaign, prefixed cmp_.

```
yesly campaigns list-recipients <id> [flags]
```

### Examples

```
  yesly campaigns list-recipients cmp_5Tn8Wq3Zr1Vb7K --status failed --all
```

### Options

```
      --all                  follow every page and print all items as one JSON array
  -h, --help                 help for list-recipients
      --limit int            Between 1 and 100. Default 50.
      --starting-after int   Cursor: the row number of the last item of the previous page.
      --status string        Only rows in this status. One of: pending, launched, failed, skipped.
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

* [yesly campaigns](yesly_campaigns.md)	 - One agent contacting many people at a steady pace

