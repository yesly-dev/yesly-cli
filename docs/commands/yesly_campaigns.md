## yesly campaigns

One agent contacting many people at a steady pace

### Synopsis

One agent contacting many people at a steady pace.

```
yesly campaigns [flags]
```

### Options

```
  -h, --help   help for campaigns
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
* [yesly campaigns cancel](yesly_campaigns_cancel.md)	 - Cancel a campaign
* [yesly campaigns create](yesly_campaigns_create.md)	 - Create a campaign
* [yesly campaigns list](yesly_campaigns_list.md)	 - List campaigns
* [yesly campaigns list-recipients](yesly_campaigns_list-recipients.md)	 - List a campaign's recipients
* [yesly campaigns pause](yesly_campaigns_pause.md)	 - Pause a campaign
* [yesly campaigns resume](yesly_campaigns_resume.md)	 - Resume a campaign
* [yesly campaigns retrieve](yesly_campaigns_retrieve.md)	 - Retrieve a campaign

