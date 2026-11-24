## yesly agents list-insights

List an agent's findings

### Synopsis

List an agent's findings.

Once a day, after the first conversation it analyses that day, Yesly reads the signals of the agent's last 30 days and writes findings: a pattern (`objection`, `not_reached_decider`, `cut_point`, `knowledge_gap`, `low_quality`), its evidence (how many conversations, their share, conversion with and without it, and a few conversation ids) and a proposal: an instruction section to try as a variant, or for `knowledge_gap` the topics of the questions the agent could not answer. Highest score first: the conversations a pattern touches times the conversion it costs. Needs 20 analysed conversations. A proposal passes a review first: no figures or prices, promises, urgency, or anything about the AI notice or recording. Open and tried findings only. A read-only key gets `403`.

API: GET /agents/{id}/insights

Arguments:
  <id>  Id of the agent, prefixed agt_.

```
yesly agents list-insights <id> [flags]
```

### Examples

```
  yesly agents list-insights agt_8Hd2kQ4mZ1LpXy
```

### Options

```
      --all    follow every page and print all items as one JSON array
  -h, --help   help for list-insights
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

