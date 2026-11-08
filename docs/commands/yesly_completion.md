## yesly completion

Generate the autocompletion script for the specified shell

### Synopsis

Generate the autocompletion script for yesly for the specified shell.
See each sub-command's help for details on how to use the generated script.


### Options

```
  -h, --help   help for completion
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
* [yesly completion bash](yesly_completion_bash.md)	 - Generate the autocompletion script for bash
* [yesly completion fish](yesly_completion_fish.md)	 - Generate the autocompletion script for fish
* [yesly completion powershell](yesly_completion_powershell.md)	 - Generate the autocompletion script for powershell
* [yesly completion zsh](yesly_completion_zsh.md)	 - Generate the autocompletion script for zsh

