# Embedded MCP Server — Implementation Plan

Working from `spec.md` in this directory. Execution is inline (single developer working in this worktree), commits at phase boundaries.

## SDK API quick reference (verified against go-sdk v1.5.0)

```go
import "github.com/modelcontextprotocol/go-sdk/mcp"

server := mcp.NewServer(&mcp.Implementation{Name, Version}, nil)

// Register a tool with reflection-derived input schema:
type ExtractMarkdownInput struct {
    URL  string `json:"url" jsonschema:"URL to fetch and convert to markdown"`
    // optional fields use omitempty so the schema marks them not-required
    Metadata bool `json:"metadata,omitempty" jsonschema:"include metadata as a separate field"`
}
mcp.AddTool(server, &mcp.Tool{Name: "...", Description: "..."},
    func(ctx context.Context, req *mcp.CallToolRequest, args *ExtractMarkdownInput) (*mcp.CallToolResult, any, error) { ... })

// Stdio:
server.Run(ctx, &mcp.StdioTransport{})

// HTTP:
handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server { return server }, nil)
http.ListenAndServe(addr, handler)

// Progress notification (only if client sent a progress token):
if token := req.Params.GetProgressToken(); token != nil {
    _ = req.Session.NotifyProgress(ctx, &mcp.ProgressNotificationParams{
        Message: "...", ProgressToken: token, Progress: float64(n),
    })
}

// Tool errors (API/runtime, NOT protocol):
return &mcp.CallToolResult{
    Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
    IsError: true,
}, nil, nil
```

---

## Phase 1: Dependency + cmd skeleton

### Task 1: Add MCP SDK dep, vendor with patch restore

```sh
cd <worktree>
GOPRIVATE=github.com/stainless-sdks/tabstack-go go get github.com/modelcontextprotocol/go-sdk@latest
make vendor    # runs go mod vendor && git restores the SDK patch
go build ./...   # confirm compiles
make vendor-check  # verify clean
```

Expected: go.mod gets `github.com/modelcontextprotocol/go-sdk vX.Y.Z` direct require. `vendor/github.com/modelcontextprotocol/go-sdk/...` populated.

Commit: `deps: vendor github.com/modelcontextprotocol/go-sdk`

### Task 2: Scaffold `cmd/mcp.go`

New file. Cobra subcommand, persistent flags inherited from root, two local flags (`--transport`, `--listen`), `RunE` initially routes to a function that just logs the chosen transport and returns.

```go
package cmd

import (
    "github.com/spf13/cobra"
)

var (
    mcpTransport string
    mcpListen    string
)

var mcpCmd = &cobra.Command{
    Use:   "mcp",
    Short: "Run the Tabstack API as an MCP (Model Context Protocol) server",
    Long: `Expose every Tabstack API operation as an MCP tool, usable from
Claude Code, Cursor, or any other MCP-aware client. The server reuses the
same TABSTACK_API_KEY env / --api-key flag / tabstack.yaml config as the
other CLI subcommands.`,
    RunE: runMCP,
}

func init() {
    mcpCmd.Flags().StringVar(&mcpTransport, "transport", "stdio",
        "wire protocol: stdio (for Claude Code etc.) or http")
    mcpCmd.Flags().StringVar(&mcpListen, "listen", "127.0.0.1:8080",
        "HTTP bind address (used only when --transport http). Localhost-only by default.")
    rootCmd.AddCommand(mcpCmd)
}
```

Add `validMCPTransports = []string{"stdio", "http"}` to `cmd/enums.go`. Validate at the top of `runMCP`.

`runMCP` body for now:

```go
func runMCP(_ *cobra.Command, _ []string) error {
    if err := validateEnum("transport", mcpTransport, validMCPTransports); err != nil {
        return err
    }
    return fmt.Errorf("not yet implemented (Task 4 wires the server)")
}
```

Smoke: `./tabstack mcp --help` works; `./tabstack mcp --transport yolo` errors cleanly.

Commit: `feat(cmd): scaffold 'mcp' subcommand with --transport and --listen flags`

---

## Phase 2: Server scaffolding

### Task 3: `internal/mcp/server.go` — empty server constructor

New file. Just the constructor; no tools yet.

```go
// Package mcp embeds an MCP (Model Context Protocol) server inside the
// tabstack CLI binary. The server exposes every Tabstack API operation as
// an MCP tool, reusing the existing internal/client + vendored SDK.
package mcp

import (
    sdk "github.com/modelcontextprotocol/go-sdk/mcp"
    tabstack "github.com/stainless-sdks/tabstack-go"
)

// NewServer constructs an MCP server preloaded with all Tabstack tools.
// `version` is the binary's version string (e.g. set via ldflags) — the MCP
// client uses this to render the server identity.
func NewServer(client *tabstack.Client, version string) *sdk.Server {
    s := sdk.NewServer(&sdk.Implementation{
        Name:    "tabstack",
        Version: version,
    }, nil)
    // Tool registrations land here in Phase 3+.
    return s
}
```

Commit: `feat(mcp): add internal/mcp package skeleton`

### Task 4: Wire `cmd/mcp.go` to actually run the server

Replace the not-implemented stub. Build a `*tabstack.Client` via `client.New(GetConfig())`, hand it to `mcp.NewServer`, then dispatch on transport:

```go
func runMCP(_ *cobra.Command, _ []string) error {
    if err := validateEnum("transport", mcpTransport, validMCPTransports); err != nil {
        return err
    }
    c, err := client.New(GetConfig())
    if err != nil {
        return err
    }
    srv := mcppkg.NewServer(c, version)
    ctx := context.Background()

    switch mcpTransport {
    case "stdio":
        GetLogger().Infof("MCP server starting on stdio")
        return srv.Run(ctx, &sdk.StdioTransport{})
    case "http":
        GetLogger().Infof("MCP server starting on http://%s", mcpListen)
        handler := sdk.NewStreamableHTTPHandler(func(*http.Request) *sdk.Server { return srv }, nil)
        return http.ListenAndServe(mcpListen, handler)
    }
    return nil
}
```

Imports (alias the SDK as `sdk` and our package as `mcppkg`):
```go
sdk "github.com/modelcontextprotocol/go-sdk/mcp"
mcppkg "github.com/lmorchard/tabstack-go-cli/internal/mcp"
```

Where does `version` come from? `cmd/version.go` has the package-private `version` var. Pass it directly: `mcppkg.NewServer(c, version)`.

Smoke test (with no tools registered yet):

```sh
TABSTACK_API_KEY=dummy ./tabstack mcp --transport http --listen 127.0.0.1:7777 &
sleep 1
curl -s -X POST http://127.0.0.1:7777 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}'
kill %1
```

Expected: a JSON-RPC response with an empty tools list (since we haven't added any yet).

Commit: `feat(mcp): wire stdio and http transports into 'tabstack mcp'`

---

## Phase 3: One-shot tools

Each task: input struct in its file, registration call in `server.go`, smoke-test the tool via a tools/list response.

### Task 5: `tabstack_extract_markdown`

New file `internal/mcp/tools_extract.go`:

```go
package mcp

import (
    "context"
    "encoding/json"
    "fmt"

    sdk "github.com/modelcontextprotocol/go-sdk/mcp"
    tabstack "github.com/stainless-sdks/tabstack-go"
    "github.com/stainless-sdks/tabstack-go/packages/param"
)

type ExtractMarkdownInput struct {
    URL      string `json:"url" jsonschema:"URL to fetch and convert to clean markdown" jsonschema_extras:"format=uri"`
    Metadata bool   `json:"metadata,omitempty" jsonschema:"return metadata as a structured field instead of YAML frontmatter"`
    Nocache  bool   `json:"nocache,omitempty" jsonschema:"bypass server-side cache"`
    Effort   string `json:"effort,omitempty" jsonschema:"speed/capability tradeoff: min, standard, or max"`
    Geo      string `json:"geo,omitempty" jsonschema:"ISO 3166-1 alpha-2 country code (e.g. US, GB, JP)"`
}

func registerExtractMarkdown(s *sdk.Server, c *tabstack.Client) {
    sdk.AddTool(s, &sdk.Tool{
        Name: "tabstack_extract_markdown",
        Description: "Fetch a URL and return its main content as clean markdown. " +
            "Use when you need the readable text of a web page without HTML or layout boilerplate.",
    }, func(ctx context.Context, req *sdk.CallToolRequest, in *ExtractMarkdownInput) (*sdk.CallToolResult, any, error) {
        body := tabstack.ExtractMarkdownParams{URL: in.URL}
        if in.Metadata {
            body.Metadata = param.NewOpt(true)
        }
        if in.Nocache {
            body.Nocache = param.NewOpt(true)
        }
        if in.Effort != "" {
            body.Effort = tabstack.ExtractMarkdownParamsEffort(in.Effort)
        }
        if in.Geo != "" {
            body.GeoTarget = tabstack.ExtractMarkdownParamsGeoTarget{Country: param.NewOpt(in.Geo)}
        }
        resp, err := c.Extract.Markdown(ctx, body)
        if err != nil {
            return toolError(fmt.Errorf("extract markdown: %w", err)), nil, nil
        }
        return jsonResult(resp), nil, nil
    })
}
```

Two helpers in a new `internal/mcp/result.go`:

```go
package mcp

import (
    "encoding/json"

    sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// jsonResult serializes v as pretty-printed JSON and wraps it in a single
// TextContent block. Falls back to a stringified Go value if marshaling
// fails (shouldn't happen for SDK response types).
func jsonResult(v any) *sdk.CallToolResult {
    b, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return toolError(err)
    }
    return &sdk.CallToolResult{
        Content: []sdk.Content{&sdk.TextContent{Text: string(b)}},
    }
}

// toolError wraps a runtime error as a non-protocol tool error: IsError=true,
// message in TextContent. The MCP SDK distinguishes this from a returned
// `error` value, which becomes a protocol-level error.
func toolError(err error) *sdk.CallToolResult {
    return &sdk.CallToolResult{
        Content: []sdk.Content{&sdk.TextContent{Text: err.Error()}},
        IsError: true,
    }
}
```

Add `registerExtractMarkdown(s, client)` to `NewServer` in `server.go`.

Smoke (HTTP transport): tools/list shows `tabstack_extract_markdown` with the right schema. tools/call invokes against api.tabstack.ai (manual smoke; deferred to Task 12).

Commit: `feat(mcp): add tabstack_extract_markdown tool`

### Task 6: `tabstack_extract_json`

Append to `tools_extract.go` (same pattern). Input struct has an additional `JsonSchema map[string]any`json:"json_schema"`` field marked required via the absence of `omitempty`. Handler calls `c.Extract.Json(ctx, body)`.

```go
type ExtractJsonInput struct {
    URL        string         `json:"url" jsonschema:"URL to fetch and extract structured data from"`
    JsonSchema map[string]any `json:"json_schema" jsonschema:"JSON Schema describing the structure of data to extract"`
    Nocache    bool           `json:"nocache,omitempty"`
    Effort     string         `json:"effort,omitempty" jsonschema:"min, standard, or max"`
    Geo        string         `json:"geo,omitempty" jsonschema:"ISO 3166-1 alpha-2 country code"`
}
```

Description: *"Fetch a URL and extract structured data conforming to a JSON Schema you provide. Use when you need specific fields out of a page (e.g. title, price, author)."*

Commit: `feat(mcp): add tabstack_extract_json tool`

### Task 7: `tabstack_generate_json`

New file `internal/mcp/tools_generate.go`. Same pattern; input has `Instructions` field (required), calls `c.Generate.Json`.

```go
type GenerateJsonInput struct {
    URL          string         `json:"url"`
    JsonSchema   map[string]any `json:"json_schema"`
    Instructions string         `json:"instructions" jsonschema:"transformation instructions (max 20000 chars)"`
    Nocache      bool           `json:"nocache,omitempty"`
    Effort       string         `json:"effort,omitempty"`
    Geo          string         `json:"geo,omitempty"`
}
```

Description: *"Fetch a URL and AI-transform its content per your instructions, returning JSON that conforms to a schema you provide. Use when you need a synthesized or summarized version of a page in a structured shape."*

Register in `server.go`.

Commit: `feat(mcp): add tabstack_generate_json tool`

---

## Phase 4: Streaming tools with progress

### Task 8: Progress helper

New file `internal/mcp/progress.go`:

```go
package mcp

import (
    "context"
    "fmt"

    sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// progressEmitter sends MCP progress notifications during a streaming tool
// call. Created at the top of a streaming handler; emit() is called per
// curated event. If the client didn't supply a progress token, all calls
// are no-ops (progress notifications would be unreachable anyway).
type progressEmitter struct {
    ctx     context.Context
    session *sdk.ServerSession
    token   any
    count   float64
}

func newProgressEmitter(ctx context.Context, req *sdk.CallToolRequest) *progressEmitter {
    return &progressEmitter{
        ctx:     ctx,
        session: req.Session,
        token:   req.Params.GetProgressToken(),
    }
}

func (p *progressEmitter) emit(format string, args ...any) {
    if p.token == nil {
        return
    }
    p.count++
    _ = p.session.NotifyProgress(p.ctx, &sdk.ProgressNotificationParams{
        Message:       fmt.Sprintf(format, args...),
        ProgressToken: p.token,
        Progress:      p.count,
    })
}
```

Plus a per-event-type filter that decides which automate/research events deserve a progress emission. The filter mirrors the pretty renderer's curation but produces a short progress-line string (no time prefix; the client adds its own UI).

```go
import tabstack "github.com/stainless-sdks/tabstack-go"

// automateProgress returns the progress message for the given event, or "" to skip.
func automateProgress(ev tabstack.AutomateEventUnion) string {
    switch v := ev.AsAny().(type) {
    case tabstack.AutomateEventCdpEndpointConnected:
        return "Connected to browser"
    case tabstack.AutomateEventAgentStatus:
        return v.Data.Message
    case tabstack.AutomateEventAgentStep:
        return fmt.Sprintf("Step %d", int(v.Data.CurrentIteration))
    case tabstack.AutomateEventBrowserNavigated:
        return fmt.Sprintf("Navigated to %s", v.Data.URL)
    case tabstack.AutomateEventAgentAction:
        return fmt.Sprintf("Action: %s", v.Data.Action)
    case tabstack.AutomateEventAgentReasoned:
        // Reasoning can be many paragraphs; the progress message is for a
        // status-line UI, so collapse to a single line and cap length. Full
        // reasoning is still in the SSE stream and would be available in a
        // hypothetical full-events tool result.
        msg := strings.Join(strings.Fields(v.Data.Reasoning), " ")
        if len(msg) > 120 {
            msg = msg[:119] + "…"
        }
        return "Reasoning: " + msg
    case tabstack.AutomateEventTaskStarted:
        return "Task started"
    case tabstack.AutomateEventTaskValidated:
        return "Task validated"
    case tabstack.AutomateEventTaskCompleted:
        return "Task completed"
    case tabstack.AutomateEventTaskAborted:
        return "Task aborted"
    case tabstack.AutomateEventComplete:
        return "Done"
    case tabstack.AutomateEventError:
        return "Error: " + v.Data.Error.Message
    default:
        return ""
    }
}

// researchProgress similarly.
func researchProgress(ev tabstack.ResearchEventUnion) string {
    switch v := ev.AsAny().(type) {
    case tabstack.ResearchEventPlanningStart:
        return "Planning research"
    case tabstack.ResearchEventPlanningEnd:
        return fmt.Sprintf("Plan ready (%s, %d queries)", v.Data.Complexity, len(v.Data.Queries))
    case tabstack.ResearchEventIterationStart:
        return fmt.Sprintf("Iteration %d/%d", int(v.Data.Iteration), int(v.Data.MaxIterations))
    case tabstack.ResearchEventSearchingStart:
        return fmt.Sprintf("Searching (%d quer%s)", len(v.Data.Queries), pluralYY(len(v.Data.Queries)))
    case tabstack.ResearchEventSearchingEnd:
        return fmt.Sprintf("Found %d URL(s), %d new", int(v.Data.URLsFound), int(v.Data.URLsNew))
    case tabstack.ResearchEventWritingStart:
        return fmt.Sprintf("Writing report (attempt %d/%d)", int(v.Data.Attempt), int(v.Data.MaxAttempts))
    case tabstack.ResearchEventWritingEnd:
        return "Report draft complete"
    case tabstack.ResearchEventComplete:
        return "Done"
    case tabstack.ResearchEventError:
        return "Error: " + v.Data.Error.Message
    default:
        return ""
    }
}

func pluralYY(n int) string {
    if n == 1 {
        return "y"
    }
    return "ies"
}
```

Tests in `internal/mcp/progress_test.go`: construct a few typed events via JSON unmarshal (same trick as `internal/sse/pretty_test.go`), assert message strings.

Commit: `feat(mcp): add progress notification helper and event-to-message filter`

### Task 9: `tabstack_research`

New file `internal/mcp/tools_research.go`. Handler:
1. Build params from input
2. Open the SSE stream via `c.Agent.ResearchStreaming(ctx, body)`
3. Drain the stream:
   - For each event, ask `researchProgress(ev)` for a message; if non-empty, `emitter.emit(msg)`
   - Capture the `complete` event's `report` field
   - Return early on `error` event
4. After the loop, return `&sdk.CallToolResult{Content: []sdk.Content{&sdk.TextContent{Text: report}}}, nil, nil`

Skeleton:

```go
type ResearchInput struct {
    Query        string `json:"query"`
    Mode         string `json:"mode,omitempty" jsonschema:"fast or balanced"`
    Nocache      bool   `json:"nocache,omitempty"`
    FetchTimeout int64  `json:"fetch_timeout,omitempty" jsonschema:"per-fetch timeout in seconds (0 = SDK default)"`
}

func registerResearch(s *sdk.Server, c *tabstack.Client) {
    sdk.AddTool(s, &sdk.Tool{
        Name: "tabstack_research",
        Description: "Run a multi-source AI research query and return a cited markdown report. " +
            "Use for questions that require synthesizing information from multiple web sources.",
    }, func(ctx context.Context, req *sdk.CallToolRequest, in *ResearchInput) (*sdk.CallToolResult, any, error) {
        body := tabstack.AgentResearchParams{Query: in.Query}
        if in.Mode != "" {
            body.Mode = tabstack.AgentResearchParamsMode(in.Mode)
        }
        if in.Nocache {
            body.Nocache = param.NewOpt(true)
        }
        if in.FetchTimeout > 0 {
            body.FetchTimeout = param.NewOpt(in.FetchTimeout)
        }

        emitter := newProgressEmitter(ctx, req)
        stream := c.Agent.ResearchStreaming(ctx, body)
        defer func() { _ = stream.Close() }()

        var report string
        for stream.Next() {
            if err := ctx.Err(); err != nil {
                return toolError(err), nil, nil
            }
            ev := stream.Current()
            if msg := researchProgress(ev); msg != "" {
                emitter.emit("%s", msg)
            }
            if v, ok := ev.AsAny().(tabstack.ResearchEventComplete); ok {
                report = v.Data.Report
            }
            if v, ok := ev.AsAny().(tabstack.ResearchEventError); ok {
                return toolError(fmt.Errorf("research: %s", v.Data.Error.Message)), nil, nil
            }
        }
        if err := stream.Err(); err != nil {
            return toolError(fmt.Errorf("research stream: %w", err)), nil, nil
        }
        if report == "" {
            return toolError(fmt.Errorf("research stream ended without a complete event")), nil, nil
        }
        return &sdk.CallToolResult{
            Content: []sdk.Content{&sdk.TextContent{Text: report}},
        }, nil, nil
    })
}
```

Register in `server.go`.

Commit: `feat(mcp): add tabstack_research tool with progress notifications`

### Task 10: `tabstack_automate`

New file `internal/mcp/tools_automate.go`. Same shape as research; differences:

- Input has `task`, `url`, `guardrails`, `data`, `geo`, `max_iterations`, `max_validation_attempts`. The `data` field is `map[string]any`.
- ALWAYS sets `body.Interactive = false` (not surfaced in tool input).
- Captures `complete.Data.FinalAnswer` as the result.
- Mid-stream: if an `interactive:form_data:request` event arrives, log a warning to stderr (the agent will sit waiting for 2 minutes then expire — fine for v1; alternative is to call `c.Agent.AutomateInput(ctx, requestID, AgentAutomateInputParams{Cancelled: param.NewOpt(true)})` to immediately decline. **Decision: auto-decline** — the SDK call is cheap and the alternative is the agent waiting 2 minutes for nothing).

```go
type AutomateInput struct {
    Task                  string         `json:"task"`
    URL                   string         `json:"url,omitempty"`
    Guardrails            string         `json:"guardrails,omitempty"`
    Data                  map[string]any `json:"data,omitempty" jsonschema:"JSON context the agent can use during the task"`
    Geo                   string         `json:"geo,omitempty"`
    MaxIterations         int64          `json:"max_iterations,omitempty"`
    MaxValidationAttempts int64          `json:"max_validation_attempts,omitempty"`
}
```

Description: *"Run an autonomous AI browser-automation task described in natural language. The agent navigates web pages, interacts with elements, and returns a final answer or report. Tasks requiring user-typed form data (passwords, etc.) are not supported in this MCP version — use the `tabstack` CLI directly for those."*

The handler closely mirrors `runAutomate` in `cmd/automate.go` but without the prompter logic. When an `interactive:form_data:request` event arrives:

```go
case tabstack.AutomateEventInteractiveFormDataRequest:
    // We can't prompt the user mid-tool-call (yet — see issue for elicitation
    // follow-up). Auto-decline so the agent doesn't hang for 2 minutes.
    _, _ = c.Agent.AutomateInput(ctx, v.Data.RequestID, tabstack.AgentAutomateInputParams{
        Cancelled: param.NewOpt(true),
    })
    emitter.emit("Form-data request auto-declined (use the CLI for interactive mode)")
```

Register in `server.go`.

Commit: `feat(mcp): add tabstack_automate tool with progress and auto-decline of form-data callbacks`

---

## Phase 5: Polish

### Task 11: Tests

Targets:
- `internal/mcp/result_test.go` — `jsonResult` and `toolError` produce the right `CallToolResult` shape (TextContent count, IsError flag, content text).
- `internal/mcp/progress_test.go` — `automateProgress` and `researchProgress` return expected strings for representative events; default for unknown variants returns `""`. Use the JSON-unmarshal pattern from `internal/sse/pretty_test.go`.
- `internal/mcp/server_test.go` — `NewServer(nil, "test").Implementation` reports `Name: "tabstack"`, `Version: "test"`. Validate that `NewServer` registers exactly the 5 expected tool names. (Use `sdk.Server`'s introspection if it exposes registered tools; otherwise call against an in-process client and compare `ListTools` results.)

For the tool-handler tests, mock the SDK client minimally — handlers depend on `*tabstack.Client.Extract.Markdown` etc. We can't easily mock these struct fields; test handler logic via integration where possible, otherwise trust the smoke (Task 12).

If full handler tests are too costly: just unit-test the helper functions (jsonResult, toolError, progress filters) and rely on smoke-test for end-to-end coverage. **Take this path** — it's where we landed for the v1 cmd-level tests too.

Commit: `test(mcp): cover result helpers and progress filters`

### Task 12: Live smoke test

Capture results in `docs/dev-sessions/2026-04-27-1745-embedded-mcp-server/smoke-results.md`.

```sh
make build
export TABSTACK_API_KEY=...

# Start in HTTP mode for manual probing
./tabstack mcp --transport http --listen 127.0.0.1:7777 &
SVR=$!

# tools/list — should show 5 tools
curl -s -X POST http://127.0.0.1:7777 \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}'

# call extract_markdown
curl -s -X POST http://127.0.0.1:7777 \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"tabstack_extract_markdown","arguments":{"url":"https://example.com"}}}'

# call research (no progress token — verify it works without progress)
curl -s --max-time 60 -X POST http://127.0.0.1:7777 \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"tabstack_research","arguments":{"query":"capital of France","mode":"fast"}}}'

kill $SVR
```

Also: wire it into Claude Code via a temporary config to confirm the stdio transport works end-to-end. Document that flow in smoke-results.md.

Commit: `docs(dev-session): add smoke-test results for MCP server`

### Task 13: README + CLAUDE.md

**README** — new `## MCP server` section after the `## Commands` section:

```markdown
## MCP server

`tabstack mcp` runs the CLI itself as a [Model Context Protocol](https://modelcontextprotocol.io) server, exposing every API operation as an MCP tool.

```sh
tabstack mcp                                       # stdio (for Claude Code etc.)
tabstack mcp --transport http --listen 127.0.0.1:8080
```

### Wiring it into Claude Code

Add to your `~/.claude/mcp.json` (or equivalent):

```json
{
  "mcpServers": {
    "tabstack": {
      "command": "/path/to/tabstack",
      "args": ["mcp"],
      "env": { "TABSTACK_API_KEY": "..." }
    }
  }
}
```

Then in any Claude Code session you'll see `tabstack_extract_markdown`, `tabstack_extract_json`, `tabstack_generate_json`, `tabstack_research`, `tabstack_automate` available as tools.

(...existing limitations note: interactive automation, screenshots…)
```

**CLAUDE.md** — append `internal/mcp` to the Package Layout section:

```markdown
- `internal/mcp` — embedded MCP server. `NewServer(*tabstack.Client, version)` constructs an MCP server with all 5 Tabstack tools registered. Used by `cmd/mcp.go`. Reuses the SSE-curated event filter for progress notifications during streaming tools.
```

Update the command table at the top of CLAUDE.md to add a `mcp` row.

Commit: `docs: README + CLAUDE.md sections for the MCP server`

### Task 14: File follow-up issues

After PR is open (so the issues can `Refs #N`), file:

1. **MCP elicitation for `tabstack_automate` interactive mode** — when Claude Code/MCP-clients support elicitation, plug it into the form-data callback flow. Body references the auto-decline behavior in v1.
2. **MCP HTTP transport: bearer auth and external bind** — for v2 if anyone wants to expose the server over the network. Body covers `--auth-token`, `--bind 0.0.0.0`, CORS.
3. **MCP `structuredContent` for extract/generate JSON tools** — low priority; lets clients programmatically access the JSON without re-parsing TextContent.

Note: existing #7 (screenshots) already covers MCP screenshot attachment — no new issue needed there.

(No commit; these are GitHub-side actions.)

---

## Self-review checklist

After execution:

- [ ] `make build` clean
- [ ] `make lint` 0 issues
- [ ] `make test` clean
- [ ] `make vendor-check` clean
- [ ] Smoke results captured
- [ ] README + CLAUDE.md updated
- [ ] Follow-up issues filed (Task 14)

## PR

Open against `main` with title `feat: embedded MCP server (closes #15)`. Body: spec summary + test plan + reference to smoke-results.md. Add Copilot reviewer per the standard `pr` flow.

Squash-merge or leave commits granular — pick at merge time.
