## yesly contacts

Your customers, their identifiers, and what your agents remember about them

### Synopsis

Your customers, their identifiers, and what your agents remember about them.

```
yesly contacts [flags]
```

### Options

```
  -h, --help   help for contacts
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
* [yesly contacts create](yesly_contacts_create.md)	 - Create or update a contact
* [yesly contacts delete](yesly_contacts_delete.md)	 - Delete a contact
* [yesly contacts delete-memory](yesly_contacts_delete-memory.md)	 - Delete a memory record
* [yesly contacts list](yesly_contacts_list.md)	 - List contacts
* [yesly contacts list-memory](yesly_contacts_list-memory.md)	 - List memory records
* [yesly contacts merge](yesly_contacts_merge.md)	 - Merge contacts
* [yesly contacts retrieve](yesly_contacts_retrieve.md)	 - Retrieve a contact
* [yesly contacts update](yesly_contacts_update.md)	 - Update a contact

