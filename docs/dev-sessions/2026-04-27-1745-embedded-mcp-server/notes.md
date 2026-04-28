# Embedded MCP Server — Session Notes & Retrospective

## Summary

Shipped `tabstack mcp` — a Cobra subcommand that runs the CLI as a Model Context Protocol server (stdio for Claude Code, optional HTTP on localhost). Five tools registered with the `tabstack_*` prefix, all reusing the existing `internal/client` SDK constructor: `extract_markdown`, `extract_json`, `generate_json`, `research`, `automate`. Streaming tools emit MCP `ProgressNotificationParams` for the same curated event subset the `internal/sse/pretty` renderer specializes; the final result is returned as a single `TextContent` block.

Closes #15.

PRs:
- #20 — deps split: vendor `github.com/modelcontextprotocol/go-sdk` v1.5.0 + empty `internal/mcp` skeleton (merged manually; vendored SDK > Copilot 20k-line cap)
- #16 — feature PR: 17 files / 1,757 lines after rebasing the vendor commits out

Follow-ups filed:
- #17 — MCP elicitation support so `tabstack_automate` can prompt for form-data fields mid-tool-call
- #18 — HTTP transport: bearer auth, external bind, CORS
- #19 — `structuredContent` on tool results for programmatic access alongside the TextContent fallback
- #7 — extended to cover screenshot attachment for both CLI and MCP

## What went well

- **SDK boundary held.** `internal/mcp` is a thin protocol adapter — no business logic, no duplicated error paths, no separate config plumbing. Each tool handler is small (build params, call SDK, format result), and the streaming tools reuse the same event-discriminator pattern as `internal/sse/pretty`.
- **Spec → plan → execute flow paid off.** The spec called out the elicitation/screenshots/HTTP-auth follow-ups before any code was written, so scope drift never happened. The plan's per-task SDK API quick-reference saved a lot of doc-spelunking.
- **Progress event filter sharing.** Curated event subsets for research/automate live alongside the existing pretty renderer's filters, so the two views (CLI pretty, MCP progress) stay aligned by construction.
- **Live smoke test caught nothing.** All 5 tools returned expected types against staging on first try (see `smoke-results.md`). The tests we wrote during the build (`progress_test.go`, `result_test.go`) covered the formatting helpers — actual API integration was exercised by the smoke run.

## What didn't go well

- **Copilot's 20k-line cap, again.** Same problem as the v1 SDK PR: vendoring the MCP SDK alone added ~33k lines, blowing past Copilot's review limit. Had to cherry-pick the three vendor commits onto a separate `deps-mcp-sdk` branch, merge that PR manually (Copilot couldn't review it either), then rebase `feat/embedded-mcp-server` to drop the now-redundant commits via patch-id matching. Two PRs for what was conceptually one feature.
- **Vendor-check false positive.** Initial `make vendor-check` failed because `go mod vendor` strips the local SDK patch (`vendor/github.com/stainless-sdks/tabstack-go/local_patch_file_id_unions.go`). Fixed by extending the `vendor-check` target to apply the same `git restore` step that `make vendor` does. Should have caught this when the Makefile target was first added — adding a new vendored dep was the trigger that exposed it.
- **Lint-on-commit fragility.** Bundling tasks 8/9/10 into a single commit was necessary because Task 8 added helper functions (`automateProgress`, `researchProgress`) that only got their first call site in tasks 9/10 — committing them separately tripped `unused` lint. The plan should have anticipated this and merged the tasks at plan-time.
- **`textResult` helper churn.** Wrote it in `result.go` for streaming tools, then realized streaming handlers build their own `CallToolResult` inline because they need to interleave progress notifications. Removed it in a later commit. Would have been cleaner to defer the streaming-helper API decision until the streaming tools actually existed.

## Surprises / things to remember

- **Copilot reviewer login is `Copilot`** (capitalized), not `copilot-pull-request-reviewer` or `github-copilot[bot]`. The `gh pr edit --add-reviewer` flag silently no-ops on the wrong name; the API call returns 422. Direct API call with `Copilot` works.
- **The official MCP Go SDK is at `github.com/modelcontextprotocol/go-sdk`**, maintained in collaboration with Google. v1.5.0 is the current stable. `mcp.AddTool` uses reflection on the input struct's `json:` and `jsonschema:` tags to derive the tool's input schema — no manual JSON Schema authoring needed for tool inputs.
- **HTTP transport** is `mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return server }, nil)`. The factory pattern is so multiple HTTP requests can hit the same server instance (since the SDK isolates session state internally).
- **`automate` non-interactive in MCP** is enforced by always sending `interactive: false`. If the agent emits an `interactive:form_data:request` event anyway, the handler calls `Agent.AutomateInput(ctx, requestID, AgentAutomateInputParams{Cancelled: param.NewOpt(true)})` to free up the agent's 2-minute timer instead of letting it wait. Worth double-checking after the elicitation follow-up lands — the cancel-on-request behavior should change to "prompt the client" in that path.

## Acceptance-criteria checklist

All boxes ticked, verified during execution:

- [x] `tabstack mcp --help` lists `--transport`, `--listen` and inherits root flags
- [x] Stdio + HTTP transports both register all 5 tools
- [x] Default HTTP bind is localhost-only (`127.0.0.1:8080`)
- [x] All 5 tools verified against staging
- [x] Progress notifications fire during streaming tools
- [x] API errors surface as `IsError: true` `CallToolResult` (not protocol errors)
- [x] Tests cover result helpers + per-event-type progress filters
- [x] `make build` / `make lint` / `make test` / `make vendor-check` all clean
- [x] README + CLAUDE.md updated
