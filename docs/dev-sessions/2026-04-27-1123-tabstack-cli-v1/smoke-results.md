# Smoke Test Results — Tabstack CLI v1

**Date:** 2026-04-27  
**Binary:** built fresh via `make build`  
**Target API:** `api.tabstack.ai`  
**API key:** temporary key (redacted)

---

## Summary

All 7 smoke commands attempted. 6 of 7 succeeded cleanly on the first try. Command 7 (`agent input` with a bogus request-id) returned a 410 Gone, which is the expected/correct behavior — the API correctly rejected an invalid/expired request ID. No retries needed.

---

## Per-Command Results

### 1. `extract markdown https://example.com`

**Status: PASS**

Returned a JSON object with `content`, `url`, and `metadata` fields. The `content` field included YAML frontmatter with `title` and `url` keys, followed by a brief page summary. The `metadata` object was present but mostly empty fields (author, description, etc. all blank strings/nulls; `title` blank). This is expected — the API returns metadata but the example.com page has minimal structured metadata.

```json
{
  "content": "---\ntitle: Example Domain\nurl: https://example.com\n---\n\nThis domain is for use ...",
  "url": "https://example.com",
  "metadata": { ... all blank ... }
}
```

### 2. `extract markdown https://example.com --metadata`

**Status: PASS — with an interesting behavioral difference**

With `--metadata`, the API response changed: the `content` field no longer included frontmatter — it returned only the body text. The title and URL moved into `metadata.title` and `metadata.url`. This suggests `--metadata` switches the API from embedding metadata as frontmatter into a structured metadata object.

The `--metadata` flag is correctly a pass-through that changes how the API returns data, not just a display toggle.

### 3. `extract json https://example.com --schema -`

**Schema:** `{"type":"object","properties":{"title":{"type":"string"}}}`  
**Status: PASS**

Returned `{"title": "Example Domain"}` — clean, schema-conformant JSON. Schema piped via stdin worked correctly.

### 4. `generate json https://example.com --schema - --instructions "Summarize the page in one sentence."`

**Schema:** `{"type":"object","properties":{"summary":{"type":"string"}}}`  
**Status: PASS**

Returned:
```json
{
  "summary": "This page describes 'Example Domain' as a domain intended for documentation examples without requiring special permission."
}
```

Clean response. The instructions were honored and the output matches the schema.

### 5. `agent research "What is the capital of France?" --mode fast`

**Status: PASS**

Streaming output was one JSON object per line throughout. Event sequence observed:

1. `start`
2. `planning:start`
3. `planning:end` — included complexity ("simple"), objective, and the generated queries
4. `iteration:start`
5. `searching:start`
6. `searching:end` — 9 URLs found
7. `iteration:end` — stop reason "max_iterations" (fast mode = 1 iteration)
8. `writing:start`
9. `writing:end`
10. `complete` — included the full report text and citation metadata

The `complete` event was present and contained the final `report` field. All lines were valid standalone JSON. Response time appeared to be roughly 10-15 seconds wall-clock. The report text was substantive and accurate (Paris, population figures, EU ranking, etc.).

**Observations:**
- The `complete` event's `data.report` contains HTML-ish fragments (e.g. `<sup>2</sup>` for superscripts, `&` as `&`). The report is not plain text — it's HTML-flavored markdown or raw HTML. Future display code should account for this.
- Citation metadata in `complete.data.metadata.citedPages` is rich (title, URL, id). Useful for building citation displays downstream.

### 6. `agent automate "Visit https://example.com and tell me the page title" --url https://example.com`

**Status: PASS**

Event sequence observed:

1. `cdp:endpoint_connected`
2. `agent:processing` (planning)
3. `agent:status` (plan created)
4. `browser:navigated` — included `title` and `url`
5. `task:started`
6. `agent:step`
7. `agent:processing` (thinking)
8. `agent:reasoned` — showed the model's reasoning text
9. `agent:action` — `{"action":"done","value":"The page title of https://example.com is \"Example Domain\"."}`
10. `agent:processing` (validating)
11. `task:validated`
12. `task:completed`
13. `complete` — included `finalAnswer`, `success: true`, and timing stats

All 13 lines were valid standalone JSON. A `complete` event was present with `success: true`. Total duration per the `complete` event stats: ~21 seconds.

**Observations:**
- The streaming event vocabulary for `automate` is richer than `research` — includes browser navigation events, per-step reasoning, and action events.
- The `agent:reasoned` event carries the model's visible chain-of-thought, which could be surfaced in a verbose display mode.
- `complete.data.stats.iterations` was 0 even though there was one `agent:step` event. The `actions` count was also 0. Looks like those stats count something more specific than I'd expect — possibly only explicit "action" tool calls, not the final "done" action.

### 7. `agent input bogus-request-id --cancel`

**Status: EXPECTED ERROR (PASS)**

The API returned `410 Gone` with body:
```
{"error":"input request has expired or already been answered"}
```

The CLI surfaced this as an error on both stderr and as the final line, and exited non-zero. The error message was clear and actionable. The usage help was printed alongside the error, which is fine.

The `|| true` in the test command confirmed the CLI correctly exits non-zero on API errors.

---

## Interactive Automate Test

**Deferred.** No obvious interactive-safe sandbox was available during this test run. Skipping per the task spec.

---

## Issues / Follow-up Items

1. **`agent research` report contains HTML fragments.** The `report` field in the `complete` event uses HTML tags (`<sup>`, `&amp;` encoded as `&`). If the CLI ever adds a pretty-print display mode or the report text is displayed directly, HTML stripping or rendering will be needed. Not a CLI bug — just something to document.

2. **`agent automate` `complete.stats.iterations` and `stats.actions` were 0** despite the agent clearly performing one planning+action step. This might be a server-side counting quirk or the stats only count non-"done" actions. Worth noting if someone relies on those fields for telemetry.

3. **`extract markdown` metadata behavior with/without `--metadata` is subtle.** Without the flag, title appears in frontmatter inside `content`. With the flag, title moves to `metadata.title` and `content` is frontmatter-free. This is correct API behavior but not obvious from the flag name alone. A future `--help` improvement could clarify this.

4. **`agent input` error display prints usage.** On API errors, the CLI prints the usage block before (or alongside) the error message. For user-facing errors like "request already answered," printing usage is noise. This is common Cobra default behavior — consider suppressing usage output on runtime (non-flag) errors.

---

## Overall Assessment

The CLI is production-functional. All six API endpoints are correctly wired. Streaming events are valid JSON, one per line, with appropriate terminator events. Auth plumbing via `TABSTACK_API_KEY` env var works. Error handling surfaces API error messages. No 5xx errors encountered.
