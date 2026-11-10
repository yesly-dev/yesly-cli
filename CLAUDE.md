# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

The official command-line interface, `yesly`, in Go (`go 1.26`). Released as binaries on GitHub releases (GoReleaser), not through a registry. Usage is in `README.md` and the generated `docs/commands/`.

## Commands

```bash
go test ./...
go test ./internal/cli -run '<TestName>'          # one test
go run ./cmd/yesly --help
node ../yesly-openapi/sdk/generate.mjs cli .        # regenerate internal/spec/zz_generated.go
go run ./cmd/yesly gen-docs docs/commands           # regenerate the command docs (CI checks they match)
```

## How it is built

`internal/spec/zz_generated.go` is generated from `yesly-openapi/openapi.yaml` and lists every operation with its flags; the commands are built from it at run time, so a new API operation needs no hand-written command. Never edit the generated file: change the spec, regenerate with both repos side by side. Each operation also needs an example in `internal/cli/examples.go`; a test fails until every operation has one. Hand-written: `internal/api` (HTTP client), `internal/cli` (command tree, output, examples), `internal/config` (keys and profiles), `internal/webhook`, `internal/version`.

## Releasing

Bump `Version` in `internal/version/version.go` and the status line in the README, then push a tag `v<version>`. The release workflow runs the tests and GoReleaser in the `release` environment, which can require reviewers.
