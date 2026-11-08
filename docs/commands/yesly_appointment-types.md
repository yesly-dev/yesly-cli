## yesly appointment-types

What each appointment is, how its time is shown and the instructions in the message

### Synopsis

What each appointment is, how its time is shown and the instructions in the message.

```
yesly appointment-types [flags]
```

### Options

```
  -h, --help   help for appointment-types
```

### Options inherited from parent commands

```
      --api-key string       API key; overrides YESLY_API_KEY and the profile
      --api-version string   Yesly-Version header (default "2026-10-01")
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
* [yesly appointment-types create](yesly_appointment-types_create.md)	 - Create an appointment type
* [yesly appointment-types list](yesly_appointment-types_list.md)	 - List appointment types
* [yesly appointment-types update](yesly_appointment-types_update.md)	 - Update an appointment type

