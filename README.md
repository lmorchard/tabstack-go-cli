# tabstack

A Go command-line interface for the [Tabstack API](https://docs.tabstack.ai/) — extract clean markdown or structured JSON from web pages, run AI-powered web research, and drive AI browser-automation tasks.

`tabstack` wraps the Stainless-generated Go SDK at [`github.com/stainless-sdks/tabstack-go`](https://github.com/stainless-sdks/tabstack-go) and exposes every public API operation as a subcommand.

## Status

Early. The SDK is pinned by commit (no published semver tags yet); see [SDK pin note](#sdk-pin) below. All six API operations are wired and have been smoke-tested against `api.tabstack.ai`. Pretty event rendering, screenshot extraction, and a raw HTTP escape hatch are intentionally deferred to follow-up work.

## Installation

No prebuilt binaries yet — build from source.

### Prerequisites

- Go 1.22 or newer
- A Tabstack API key (set `TABSTACK_API_KEY` in your environment)

### Build

```sh
git clone https://github.com/lmorchard/tabstack-go-cli.git
cd tabstack-go-cli
make setup    # installs gofumpt + golangci-lint (one-time)
make build    # produces ./tabstack
```

`make build` injects version, commit, and build-date metadata via ldflags, so `tabstack version` will report the build provenance.

## Configuration

Three layers, highest precedence first:

1. **CLI flags** — e.g. `--api-key`, `--base-url`, `--effort`, `--nocache`
2. **Environment variables** — `TABSTACK_API_KEY`, `TABSTACK_BASE_URL`
3. **Config file** — `tabstack.yaml` in the current directory, or `--config <path>`

A starter config is in `tabstack.yaml.example`. Copy it and edit:

```sh
cp tabstack.yaml.example tabstack.yaml
```

The example file deliberately leaves `api_key` commented out — prefer the `TABSTACK_API_KEY` environment variable so the secret never lands in a config file. The `*.yaml` glob is already in `.gitignore` to reduce accidental commits, but env-first is still the safer default.

## Quick start

```sh
export TABSTACK_API_KEY=sk_…

# Fetch a page as markdown
tabstack extract markdown https://example.com

# Pull a structured field per a JSON Schema (schema piped via stdin)
echo '{"type":"object","properties":{"title":{"type":"string"}}}' \
  | tabstack extract json https://example.com --schema -

# Stream a research run, watch the events scroll by
tabstack research "What is the capital of France?" --mode fast
```

## Output conventions

- **Results** go to **stdout** — JSON objects (pretty-printed for one-shot calls, one-per-line for streaming calls).
- **Logs / progress / errors** go to **stderr**.

This means you can pipe stdout into `jq` or redirect to a file without contamination:

```sh
tabstack research "Effects of microplastics on marine life" \
  | jq -c 'select(.event == "complete") | .data.report'
```

## Commands

### `tabstack extract markdown <url>`

Fetch a URL and convert its content to clean markdown.

| Flag | Description |
|---|---|
| `--metadata` | Return metadata as a structured object (title, description, etc.) instead of as YAML frontmatter inside the content |
| `--nocache` | Bypass server-side cache |
| `--effort {min\|standard\|max}` | Speed vs. capability tradeoff |
| `--geo CC` | ISO 3166-1 alpha-2 country code for proxy geotargeting (e.g. `US`, `GB`, `JP`) |

```sh
tabstack extract markdown https://example.com --metadata
```

### `tabstack extract json <url> --schema FILE`

Fetch a URL and extract structured data conforming to your JSON Schema.

| Flag | Description |
|---|---|
| `--schema FILE` | Path to a JSON Schema file. Use `-` to read from stdin. **(required)** |
| `--nocache` | Bypass cache |
| `--effort {min\|standard\|max}` | Speed vs. capability tradeoff |
| `--geo CC` | Country code for geotargeting |

```sh
cat <<'EOF' > schema.json
{
  "type": "object",
  "properties": {
    "title": {"type": "string"},
    "summary": {"type": "string"}
  }
}
EOF

tabstack extract json https://example.com --schema schema.json
```

### `tabstack generate json <url> --schema FILE --instructions TEXT`

Fetch a URL, AI-transform its content according to free-form instructions, return JSON conforming to your schema.

| Flag | Description |
|---|---|
| `--schema FILE` | JSON Schema file or `-` for stdin **(required)** |
| `--instructions TEXT` | Inline instructions (max 20,000 chars) |
| `--instructions-file PATH` | Read instructions from a file (mutually exclusive with `--instructions`) |
| `--nocache` | Bypass cache |
| `--effort {min\|standard\|max}` | Speed vs. capability tradeoff |
| `--geo CC` | Country code for geotargeting |

```sh
echo '{"type":"object","properties":{"summary":{"type":"string"}}}' \
  | tabstack generate json https://example.com \
      --schema - \
      --instructions "Summarize the page in one sentence."
```

### `tabstack research <query>`

Stream a multi-source AI research run. Output is one JSON event per line on stdout. Events include `start`, `planning:start/end`, `searching:start/end`, `writing:start/end`, and a final `complete` event with the report and citation metadata.

| Flag | Description |
|---|---|
| `--mode {fast\|balanced}` | `fast` is single-iteration (~10–30s); `balanced` runs multiple iterations |
| `--nocache` | Bypass cache |
| `--fetch-timeout N` | Per-fetch timeout in seconds (0 = SDK default) |
| `--output {json\|pretty}` | `json` (default) emits one JSON object per line; `pretty` renders human-readable lines like `[12s] searching:end — found 9 URL(s), 3 new`. The final `complete` event prints its full report verbatim on continuation lines |
| `--color {auto\|always\|never}` | Color in pretty output. `auto` (default) enables color when stdout is a TTY and `NO_COLOR` is not set |

```sh
tabstack research "Effects of microplastics on marine life" --mode fast
tabstack research "deep-sea ecosystems" --output pretty
```

The `complete` event carries the final report in `.data.report` (currently HTML-flavored — strip tags downstream if needed) and the cited sources in `.data.metadata.citedPages`.

### `tabstack automate <task>`

Stream an AI browser-automation run that interprets a natural-language task. Output is one JSON event per line. Common event types: `cdp:endpoint_connected`, `browser:navigated`, `agent:step`, `agent:reasoned`, `agent:action`, `task:completed`, `complete`.

| Flag | Description |
|---|---|
| `--url URL` | Starting URL for the task |
| `--guardrails TEXT` | Safety constraints for execution |
| `--data FILE` | Path to JSON file passed as `data` context for the task (use `-` for stdin) |
| `--interactive` | Enable interactive form-data callbacks (TTY prompts; combine with `--input-from` for non-TTY) |
| `--input-from FILE` | Path to JSON file with answers for interactive form-data requests (see [interactive automation](#interactive-automation)) |
| `--geo CC` | Country code for geotargeting |
| `--max-iterations N` | Max task iterations (0 = SDK default) |
| `--max-validation-attempts N` | Max validation attempts (0 = SDK default) |
| `--output {json\|pretty}` | `json` (default) emits one JSON object per line; `pretty` renders human-readable lines like `[12s] browser:navigated https://example.com — "Example Domain"`. The final `complete` event and `agent:reasoned` events print their full content verbatim on continuation lines |
| `--color {auto\|always\|never}` | Color in pretty output. `auto` (default) enables color when stdout is a TTY and `NO_COLOR` is not set |

```sh
tabstack automate "Visit https://example.com and report the page title" \
  --url https://example.com
```

#### Interactive automation

When the agent encounters a form during execution, it can pause and request user input via an `interactive:form_data:request` event. To handle these:

```sh
# TTY mode: tabstack will prompt you for each field
tabstack automate "Sign in to my account at https://example.com/login" \
  --url https://example.com/login \
  --interactive
```

For scripted runs, supply answers in a JSON file keyed by request ID:

```sh
cat <<'EOF' > answers.json
{
  "<requestId-from-the-event>": {
    "E1": "alice@example.com",
    "E2": "hunter2"
  }
}
EOF

tabstack automate "Log in" --interactive --input-from answers.json
```

If `--interactive` is set but stdin is not a TTY and no `--input-from` is provided, form-data requests are auto-cancelled rather than blocking forever — a warning is printed to stderr.

### `tabstack automate input <request-id>`

Standalone command for replying to (or cancelling) an in-flight `automate` form-data request, identified by the `requestId` emitted in an `interactive:form_data:request` SSE event. Useful for resuming after the stream has already closed, or for submitting from a different shell.

| Flag | Description |
|---|---|
| `--values FILE` | Path to JSON file mapping field ref → value, e.g. `{"E1":"alice","E2":"bob"}` |
| `--cancel` | Cancel the request instead of providing values |

`--values` and `--cancel` are mutually exclusive; one is required.

```sh
echo '{"E1":"alice","E2":"hunter2"}' > /tmp/answers.json
tabstack automate input <request-id> --values /tmp/answers.json
```

The 2-minute callback window applies — the API returns `410 Gone` if the request has expired or already been answered.

### `tabstack version`

Print the binary's version, commit hash, and build date.

```sh
tabstack version
# → tabstack v0.1.0 (commit: a94a977, built: 2026-04-27T18:14:32Z)
```

## Composing with `jq`

Because streaming output is one JSON object per line, `jq` and similar tools work naturally:

```sh
# Print only event types from a research run
tabstack research "deep-sea ecosystems" \
  | jq -r '.event'

# Capture just the final report
tabstack research "deep-sea ecosystems" \
  | jq -r 'select(.event == "complete") | .data.report'

# Extract all browser actions during an automate run
tabstack automate "buy a hat at example-store.com" --url … \
  | jq -c 'select(.event == "agent:action") | .data'
```

## SDK pin

The Tabstack Go SDK is currently pinned to commit `020acb44b47c` (2026-04-21). HEAD on the SDK at the time of this writing (`47997ed`) had broken codegen — undefined `*FileIDUnion` types referenced from `agent.go`. The CLI rolls back to the most recent buildable commit. When the upstream regression is fixed, the pin will move forward and the streaming event types will become typed unions; the CLI's user-facing surface will not change.

## Project layout

```
.
├── main.go
├── cmd/                 # Cobra command definitions (one file per top-level command)
│   ├── root.go
│   ├── extract.go       # extract markdown|json
│   ├── generate.go      # generate json
│   ├── automate.go      # automate (with `input` subcommand)
│   ├── research.go      # research
│   └── version.go
└── internal/
    ├── client/          # SDK client constructor
    ├── config/          # Config struct
    ├── schema/          # JSON Schema loader (file or stdin)
    ├── sse/             # Generic JSON-lines writer for typed event streams
    └── interactive/     # Form-data prompter (TTY + file source) and SDK submitter
```

Commands stay thin: parse flags, build SDK params, call into `internal/`. The `internal/interactive` package defines neutral `FormDataRequest`/`FormDataField` types decoded from the SDK's untyped event payload, insulating the rest of the code from upcoming SDK schema changes.

## Make targets

| Target | What it does |
|---|---|
| `make setup` | Install `gofumpt` and `golangci-lint` |
| `make build` | Build `./tabstack` with version ldflags |
| `make run` | Build and run |
| `make lint` | `golangci-lint run --timeout 5m` |
| `make format` | `go fmt ./...` + `gofumpt -l -w .` |
| `make test` | `go test ./...` |
| `make clean` | Remove `./tabstack` and any `*.db` files |

## Releasing

Tagged releases are cross-compiled by GitHub Actions for linux/darwin/windows × amd64/arm64. Tag with `git tag vX.Y.Z && git push --tags`. The rolling-release workflow publishes a `latest` prerelease on every push to `main`.

## License

TBD.
