# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Purpose

`tabstack` is a Go CLI that wraps the **Tabstack API** (https://docs.tabstack.ai/) via the Stainless-generated Go SDK at `github.com/stainless-sdks/tabstack-go` (pinned by commit; no semver tags published).

Module path: `github.com/lmorchard/tabstack-go-cli`. Binary: `tabstack`.

The CLI exposes one top-level command per API capability (no `agent` parent group):

| Command | Notes |
|---|---|
| `extract markdown <url>` | Fetch URL → clean markdown |
| `extract json <url> --schema FILE` | Fetch URL → JSON per schema |
| `generate json <url> --schema FILE --instructions TEXT` | Fetch URL → AI-transformed JSON |
| `research <query>` | SSE-streamed cited research |
| `automate <task>` | SSE-streamed browser automation; supports interactive form-data callbacks |
| `automate input <request-id>` | Reply to an in-flight `automate` form-data request (2-min window) |

## Common Commands

```sh
make setup    # install gofumpt + golangci-lint (one-time)
make build    # build ./tabstack with version ldflags
make run      # build and run
make lint     # golangci-lint run --timeout 5m
make format   # go fmt + gofumpt -l -w .
make test     # go test ./...
make clean    # remove ./tabstack and *.db
```

`make lint` and `make format` resolve the binaries from `$HOME/go/bin` directly — if they fail with "not found", run `make setup`. The Makefile does not pass `-race` to `go test` despite the skill docs suggesting it; use `go test -race ./...` directly when racing matters.

Run a single test: `go test ./internal/foo -run TestName -v`.

## Architecture

**Cobra + Viper, three-tier config.** `cmd/root.go` is the only place that wires flags to Viper keys. Precedence is CLI flag → env var (Viper `AutomaticEnv`) → `tabstack.yaml` (current dir) → defaults. Config can be pointed elsewhere via `--config <path>`. Missing config is silently tolerated *unless* `--config` was passed explicitly.

**Logging.** A package-level `log` (logrus) is configured in `cmd/root.go`'s `setupLogging()` based on `verbose`/`debug`/`log_json` Viper keys. `PersistentPreRun` on the root command runs `initConfig()` then `setupLogging()` before any subcommand executes — new subcommands inherit this automatically as long as they're added under `rootCmd`.

**Config struct.** `internal/config/config.go` defines a `Config` struct that mirrors a subset of Viper keys. `cmd.GetConfig()` lazily builds it from Viper. The convention is: business logic in `internal/` packages takes a `*config.Config` rather than calling Viper directly. Add new fields to both the struct and `GetConfig()` when introducing new config keys.

**Adding a command.** Use `python3 ~/.claude/skills/go-cli-builder/scripts/add_command.py <name>` to scaffold `cmd/<name>.go` with the standard boilerplate. Keep command bodies thin — call into `internal/<domain>` packages for real work.

**Version injection.** `cmd/version.go` declares package-private `version`/`commit`/`date` vars. The Makefile and release workflows inject them via `-ldflags "-X github.com/lmorchard/tabstack-go-cli/cmd.version=..."` (and similar for commit/date). If you rename the module path or move these vars, update all three places.

## Package Layout

- `cmd/` — Cobra command definitions; one file per top-level command (`extract.go`, `generate.go`, `automate.go`, `research.go`). Commands stay thin: parse flags, build SDK params, call into `internal/`.
- `internal/client` — constructs the Tabstack SDK client from `*config.Config`. Layers `--api-key`/`--base-url` on top of the SDK's env defaults. Fails fast when no key resolves.
- `internal/schema` — loads JSON Schema documents from a path or `-` (stdin). Returns `any` for direct use as the SDK's `JsonSchema` field.
- `internal/sse` — generic JSON-lines writer for any `*ssestream.Stream[T]`-like source. Used by `research`. (`automate` inlines its own drain loop because it interleaves event emit with interactive callbacks.)
- `internal/interactive` — `Prompter` interface (TTY or file-source) and `Submitter` interface (wraps `Agent.AutomateInput`). The `--interactive` flag on `automate` engages this; `--input-from FILE` substitutes the file prompter for non-TTY runs.

## CI / Releases

Three workflows in `.github/workflows/`:
- `ci.yml` — lint + test on PRs to main; skipped if commit message starts with `[noci]`.
- `release.yml` — fires on `v*` tags; cross-compiles linux/darwin/windows × amd64/arm64, uploads tarballs/zips + checksums, creates a GitHub release. Docker steps are commented out.
- `rolling-release.yml` — same build, on every push to main, published as a "latest" prerelease.

Tag a release with `git tag vX.Y.Z && git push --tags`. The Docker username placeholder (`yourusername`) in `release.yml` only matters if you uncomment the Docker job.

## Development Notes

- This is **not a git repo yet** — `git init` before the first feature branch.
- Viper key convention in this scaffold uses snake_case for multi-word keys (`log_json`, not `logJson` or `log-json`); flags use kebab-case and are bound to snake_case keys explicitly. Stay consistent when adding new ones.
- `viper.BindPFlag` calls are intentionally prefixed with `_ =` to satisfy `errcheck`; preserve that pattern.
- The scaffold was generated by the `go-cli-builder` skill at `~/.claude/skills/go-cli-builder/`; its `references/` directory has deeper notes on Cobra/Viper integration, internal package layout, and template patterns if needed.
