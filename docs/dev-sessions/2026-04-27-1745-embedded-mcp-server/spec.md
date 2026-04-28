# Embedded MCP Server — v1 Spec

## Goal

Add a `tabstack mcp` subcommand that runs the CLI itself as a [Model Context Protocol](https://modelcontextprotocol.io) server, exposing the Tabstack API operations as MCP tools usable from Claude Code, Cursor, and other MCP-aware clients.

**Tracks:** [#15](https://github.com/lmorchard/tabstack-go-cli/issues/15).

## Motivation

- **Use Tabstack from any MCP-aware client** without each client implementing a separate Tabstack integration.
- **Local-first** — no proxy hop, the API key never leaves the user's machine, lower latency than the hosted `tabstack.stlmcp.com` endpoint.
- **Single binary** — drops into a Claude Code `mcp.json` config or similar. Same `make build`, same `TABSTACK_API_KEY` env, same `tabstack.yaml` config as the existing CLI subcommands.
- **Reuse the SDK we already vendor** — the MCP server is a thin protocol adapter on top of the existing `internal/client`/`internal/schema`/SDK integration. No duplicated code.

## User-facing surface

### Command

```sh
tabstack mcp                                      # default: stdio transport
tabstack mcp --transport http --listen 127.0.0.1:8080
```

Persistent flags inherited from `rootCmd`: `--api-key`, `--base-url`, `--config`, `--verbose`, `--debug`, `--log-json`.

Subcommand-specific flags:

| Flag | Default | Description |
|---|---|---|
| `--transport {stdio\|http}` | `stdio` | Wire protocol. `stdio` runs against the parent process's stdin/stdout (the way Claude Code launches MCP servers). `http` runs an HTTP server. |
| `--listen <addr>` | `127.0.0.1:8080` | HTTP bind address. Default is localhost-only — exposing externally requires a reverse proxy with auth (out of v1 scope). |

### Tools (5)

All tools are prefixed `tabstack_*` to avoid collision with other MCP servers a user has configured. Inputs mirror the existing CLI flag sets exactly (in snake_case to match the API JSON shapes).

#### `tabstack_extract_markdown`

Fetch a URL and convert to clean markdown.

| Input | Type | Required | Notes |
|---|---|---|---|
| `url` | string (uri) | yes | |
| `metadata` | bool | no | Return metadata as a structured field instead of as YAML frontmatter inside content. |
| `nocache` | bool | no | |
| `effort` | enum: `min`/`standard`/`max` | no | |
| `geo` | string (ISO-3166-1 alpha-2) | no | |

**Result:** single `TextContent` block — pretty-printed JSON of the API's `ExtractMarkdownResponse` (content + metadata + url).

#### `tabstack_extract_json`

Fetch a URL and extract structured data conforming to a JSON Schema.

| Input | Type | Required |
|---|---|---|
| `url` | string (uri) | yes |
| `json_schema` | object | yes |
| `nocache`, `effort`, `geo` | (same as above) | no |

**Result:** single `TextContent` block — pretty-printed JSON.

#### `tabstack_generate_json`

Fetch a URL, AI-transform its content per instructions, return JSON conforming to a schema.

| Input | Type | Required |
|---|---|---|
| `url` | string (uri) | yes |
| `json_schema` | object | yes |
| `instructions` | string (max 20k chars) | yes |
| `nocache`, `effort`, `geo` | | no |

**Result:** single `TextContent` block — pretty-printed JSON.

#### `tabstack_research`

Stream a multi-source AI research run. Returns the final report.

| Input | Type | Required |
|---|---|---|
| `query` | string | yes |
| `mode` | enum: `fast`/`balanced` | no |
| `nocache` | bool | no |
| `fetch_timeout` | int (seconds) | no |

**Streaming behavior:** the SSE event stream is consumed server-side. Curated progress notifications (planning:end, searching:end, writing:end, etc.) are sent to the MCP client during the run via `ProgressNotificationParams`. Noise events (system:debug_*, ai:generation chunks if any) are skipped.

**Result:** single `TextContent` block — the `complete` event's `report` field (markdown-flavored).

#### `tabstack_automate`

Stream an AI browser-automation run. Returns the final answer.

| Input | Type | Required |
|---|---|---|
| `task` | string | yes |
| `url` | string (uri) | no |
| `guardrails` | string | no |
| `data` | object | no | JSON context the agent can use during the task (was `--data` in CLI). |
| `geo` | string | no |
| `max_iterations` | int | no |
| `max_validation_attempts` | int | no |

**Note: `interactive` is always sent as `false`.** MCP-driven `automate` is non-interactive in v1 — tasks that require form-data input fail (or, if the agent emits a form-data request, we auto-cancel it as in the CLI's non-TTY fallback). Interactive support via MCP elicitation is filed as a follow-up.

**Streaming behavior:** same as research — curated progress notifications during the stream (browser:navigated, agent:step, agent:action, task:completed, etc.).

**Result:** single `TextContent` block — the `complete` event's `finalAnswer` field. Screenshots are NOT attached in v1 (folded into #7 follow-up).

### Progress notification mapping

The MCP `ProgressNotificationParams` carries a free-form message string and an optional progress count. We use the same curated event subset as `internal/sse`'s pretty renderer:

**Research:** `planning:start/end`, `iteration:start/end`, `searching:start/end`, `writing:start/end`, `complete`, `error`. The notification message uses the renderer's existing one-line format (e.g. "searching:end — found 9 URL(s), 3 new").

**Automate:** `cdp:endpoint_connected`, `agent:status`, `agent:step`, `agent:action`, `agent:reasoned` (truncated for progress — full text only in result), `browser:navigated`, `task:started/validated/completed/aborted`, `complete`, `error`.

Other event variants are ignored. The `progress` field counts emitted events (informational, not a percentage).

### Tool descriptions

Each tool needs an LLM-facing description (used by the model to decide when to call it). These are richer than the CLI's `Short:` strings — they explicitly tell the model when the tool is appropriate. Drafted here, finalized during implementation:

- `tabstack_extract_markdown` — *"Fetch a URL and return its main content as clean markdown. Use when you need the readable text of a web page without HTML or layout boilerplate."*
- `tabstack_extract_json` — *"Fetch a URL and extract structured data conforming to a JSON Schema you provide. Use when you need specific fields out of a page (e.g. title, price, author)."*
- `tabstack_generate_json` — *"Fetch a URL and AI-transform its content per your instructions, returning JSON that conforms to a schema you provide. Use when you need a synthesized or summarized version of a page in a structured shape (e.g. a one-paragraph summary, a sentiment classification, a key-points list)."*
- `tabstack_research` — *"Run a multi-source AI research query and return a cited markdown report. Use for questions that require synthesizing information from multiple web sources."*
- `tabstack_automate` — *"Run an autonomous AI browser-automation task described in natural language. The agent navigates web pages, interacts with elements, and returns a final answer or report. Tasks requiring user-typed form data (passwords, etc.) are not supported in this MCP version — use the `tabstack` CLI directly for those."*

### Error handling

API and protocol errors surface to the MCP client as `CallToolResult` values with `IsError: true` and a single `TextContent` block whose text is the wrapped error message (e.g. `"extract markdown: tabstack: 401 Unauthorized — invalid API key"`). The handler does NOT return an `error` to the SDK in this case — that would translate to a JSON-RPC protocol error, which is wrong for "tool ran but failed." Protocol errors (malformed input that fails schema validation) are returned as `error` so the SDK formats them per spec.

### Concurrency

The Tabstack SDK client is goroutine-safe (it just wraps `*http.Client`). Handlers are stateless beyond closures over the shared client, so multiple tool calls can run concurrently — no per-handler locking needed.

### Context cancellation

Each tool handler receives a `context.Context` from the MCP SDK. When the client cancels a tool call (or the connection drops), the context is cancelled and any in-flight SSE streaming loop returns from `stream.Next()` (the SDK's stream type respects context). Handlers must check `ctx.Err()` between events and exit cleanly when set.

### Schema generation

Tool input schemas are derived from Go input structs via the official SDK's reflection-based generator (struct tags like `json:"url"` → JSON Schema `properties.url`, plus annotation tags for `description`, `enum`, `format` etc. — exact tag conventions confirmed during implementation). One input struct per tool, defined in the corresponding `tools_*.go` file.

### `data` field semantics (automate)

The CLI's `--data <file>` flag reads a JSON file from disk. The MCP `tabstack_automate` tool's `data` field is an inline JSON object passed directly in the tool call — the LLM constructs it from the conversation context. Both end up in the same place (the API's `Data any` field).

## Configuration

Same precedence as the rest of the CLI:

1. Tool-input field (e.g. an `url` arg passed to `tabstack_extract_markdown`)
2. CLI flag at server start (`--api-key`, `--base-url`, `--config`)
3. Environment (`TABSTACK_API_KEY`, `TABSTACK_BASE_URL`)
4. Config file (`tabstack.yaml` at cwd, or `--config` path)

The MCP server uses `internal/client.New(GetConfig())` exactly like the other commands — same fail-fast behavior when no API key resolves.

## Logging

The server uses the existing logrus logger configured in `cmd/root.go`. All output goes to **stderr**:

- For `--transport stdio`: stderr is separate from the protocol channel (stdout). Safe to log freely.
- For `--transport http`: stderr is the standard server log destination.

The `--verbose` / `--debug` / `--log-json` flags work the same way as elsewhere.

## Implementation approach

- **Library:** [`github.com/modelcontextprotocol/go-sdk`](https://github.com/modelcontextprotocol/go-sdk) (official, maintained in collaboration with Google; latest stable v1.5.0 as of 2026-04-07).
- **Layout:**
  - `cmd/mcp.go` — Cobra subcommand, flag parsing, transport wiring.
  - `internal/mcp/server.go` — `NewServer(*tabstack.Client) *mcp.Server`. Constructs the MCP server, registers all five tools.
  - `internal/mcp/tools_extract.go`, `tools_generate.go`, `tools_research.go`, `tools_automate.go` — one file per tool family, each defining the tool's input struct, schema annotations, and handler function. Handlers call into the existing SDK client and `internal/sse` for stream draining.
  - `internal/mcp/progress.go` — helpers for emitting `ProgressNotificationParams` from a streaming event loop. Reuses the curated event filter from `internal/sse/pretty.go`.
- **Vendor refresh:** `go get github.com/modelcontextprotocol/go-sdk@latest` then `make vendor` (which restores the SDK patch from git).

## Acceptance criteria

- [ ] `tabstack mcp --help` lists `--transport`, `--listen` and inherits root flags.
- [ ] `tabstack mcp` starts on stdio, registers all 5 tools with correct names + schemas.
- [ ] `tabstack mcp --transport http --listen 127.0.0.1:8080` runs the same surface over HTTP.
- [ ] Default HTTP bind is localhost-only.
- [ ] Each tool returns expected result type when invoked against `api.tabstack.ai` with a valid key.
- [ ] Progress notifications fire during `tabstack_research` and `tabstack_automate` runs (verified via a smoke client or by capturing the protocol stream).
- [ ] API errors surface as MCP tool errors with the wrapped message readable client-side.
- [ ] Tests cover: tool registration, input-schema validation (required fields), one happy-path per tool against a mocked SDK client, progress-emission filter logic.
- [ ] `make build`, `make lint`, `make test` all clean.
- [ ] `make vendor-check` clean (no SDK drift after adding the MCP library).
- [ ] README has an `MCP server` section explaining how to wire it into a Claude Code config.
- [ ] CLAUDE.md's "Package Layout" updated.

## Out of scope (with follow-ups)

| Item | Where it lands |
|---|---|
| Interactive automation via MCP elicitation | New issue: "MCP elicitation support for `tabstack_automate` interactive form-data" |
| Screenshot attachment in `tabstack_automate` results | Folded into existing #7 (cover both CLI and MCP in one go) |
| HTTP authentication, external bind, CORS | New issue: "MCP HTTP transport: bearer auth and external-bind support" |
| Per-request API keys (multi-tenant HTTP) | Same issue as above |
| Hosted-MCP-style sandboxed JS `execute` tool | Out of project scope (different product) |
| MCP resources / prompts | Not needed; tools-only is sufficient |
| `structuredContent` on tool results | New issue (low priority): "Add MCP structuredContent for extract/generate JSON tools" |

## Risks

1. **Official Go SDK API stability.** v1.5.0 is stable, but we're on a fast-moving spec. Breaking changes between minor versions are possible. Mitigation: pin the version, vendor it, treat upgrades as deliberate.
2. **Progress notification client support.** Not all MCP clients render progress. Claude Code does. Worst case in a non-supporting client: progress is silently dropped, the user sees a long-running tool with no updates until the final result. Acceptable.
3. **Tool result size.** Research reports can be many KB; automate finalAnswers are usually small. Both fit within typical LLM context windows. We're not attaching screenshots in v1, so the worst case is a long markdown report — fine.
4. **Vendor patch interaction.** Adding the MCP SDK changes `go.sum` and `vendor/modules.txt`. The existing SDK patch (`local_patch_file_id_unions.go`) survives `make vendor` via the `git restore` step. Verified workflow — should still work.
