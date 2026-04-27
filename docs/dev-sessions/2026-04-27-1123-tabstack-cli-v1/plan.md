# Tabstack CLI v1 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use `superpowers:subagent-driven-development` (recommended) or `superpowers:executing-plans` to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Ship a `tabstack` CLI that exercises every public Tabstack API capability end-to-end via the Stainless-generated Go SDK.

**Architecture:** Cobra command tree mirrors the API's flat action structure (`extract markdown|json`, `generate json`, `agent automate|research|input`). `internal/client` constructs the SDK client from Viper-backed config. Streaming endpoints emit SSE events to stdout as one JSON object per line (raw mode only in v1 — pretty rendering deferred). Interactive automation prompts on TTY for form-data callbacks while the SSE stream is still open.

**Tech Stack:** Cobra + Viper + logrus (already scaffolded), `github.com/stainless-sdks/tabstack-go` pinned to commit `020acb44b47c` (2026-04-21). Go 1.22 (per SDK go.mod).

> **SDK pin note:** Originally we tried `47997ed` (HEAD on 2026-04-24) but Stainless emitted broken codegen there: four `*FileIDUnion` types referenced from `agent.go` are undefined. Rolled back to the most recent buildable commit. At this version the streaming event types are still untyped — `AutomateEvent`/`ResearchEvent` are `{ Event string; Data any }` and we decode payloads ourselves. When the upstream regression is fixed we can re-pin and migrate to the union types in a follow-up dev session.

**Scope (v1):** all six API operations wired end-to-end with raw JSON output, TTY-interactive automation, schema input from file or stdin. **Deferred to later sessions:** pretty SSE event rendering, screenshot extraction to disk, `raw` HTTP escape hatch using `Client.Execute`/`Get`/`Post`.

**Decisions locked in:**
- Output discipline: results to **stdout**; logs/progress to **stderr** (logrus default). Streaming events are results.
- Schema input: `--schema FILE` where FILE may be `-` for stdin. No inline schemas in v1.
- Interactive: TTY prompt by default; `--input-from FILE` for non-TTY/scripted runs; if non-TTY and no `--input-from`, fail closed by sending `cancelled: true`.
- API key precedence: `--api-key` flag > config file `api_key` > `TABSTACK_API_KEY` env (the SDK reads env; we layer flag/config on top via `option.WithAPIKey`).
- Per-command flags for `--effort`, `--nocache`, `--geo` (only on commands that accept them — Cobra persistent flags don't fit since not every command takes all three).

**Testing approach:** TDD for genuine logic packages (`internal/schema`, `internal/sse`, `internal/interactive`). Cobra command wiring is verified by **manual smoke test against the live API** in Task 14 — automated tests of thin SDK wrappers would mostly verify mocks, not behavior. Each task ends in a single commit.

---

## File Structure

**Created:**
- `internal/client/client.go` — SDK client constructor reading `*config.Config`
- `internal/schema/schema.go` — JSON Schema loader (file or stdin → `any`)
- `internal/schema/schema_test.go`
- `internal/sse/writer.go` — generic SSE-stream → JSON-lines writer
- `internal/sse/writer_test.go`
- `internal/interactive/interactive.go` — form-data prompter (TTY + file-source); SDK input submitter
- `internal/interactive/interactive_test.go`
- `cmd/extract.go` — parent + `markdown` + `json` subcommands
- `cmd/generate.go` — parent + `json` subcommand
- `cmd/agent.go` — parent + `automate` + `research` + `input` subcommands

**Modified:**
- `go.mod`, `go.sum` — add SDK pin
- `cmd/root.go` — `--api-key`, `--base-url` persistent flags; bind to Viper
- `internal/config/config.go` — add `APIKey`, `BaseURL` fields
- `tabstack.yaml.example` — show `base_url` (api_key stays env-only)
- `cmd/constants.go` — (only if new constants needed)
- `CLAUDE.md` — document new packages and command tree

---

## Pre-flight

Confirm starting state before Task 1:

```sh
cd /Users/lorchard/devel/tabs-project/tabstack-go-cli
git status                # clean working tree, branch main
go build ./...            # builds clean
./tabstack version        # prints version + commit + date
```

---

## Phase 1 — Foundation

### Task 1: Pin the Tabstack Go SDK

**Files:**
- Modify: `go.mod`, `go.sum`

- [ ] **Step 1: Pin SDK to current main HEAD**

```sh
go get github.com/stainless-sdks/tabstack-go@47997ed
go mod tidy
```

- [ ] **Step 2: Verify build still passes**

```sh
go build ./...
```

Expected: no output, exit 0.

- [ ] **Step 3: Verify SDK is importable**

Create a throwaway file `/tmp/sdkcheck.go`:

```go
package main

import (
	"fmt"
	tabstack "github.com/stainless-sdks/tabstack-go"
)

func main() {
	c := tabstack.NewClient()
	fmt.Println(c.Agent, c.Extract, c.Generate)
}
```

```sh
cd /Users/lorchard/devel/tabs-project/tabstack-go-cli
go run /tmp/sdkcheck.go
```

Expected: prints three non-zero service struct values without error. Delete `/tmp/sdkcheck.go` after.

- [ ] **Step 4: Commit**

```sh
git add go.mod go.sum
git commit -m "deps: pin tabstack-go SDK at 47997ed"
```

---

### Task 2: Add `api_key` / `base_url` to config and root command

**Files:**
- Modify: `internal/config/config.go`
- Modify: `cmd/root.go`
- Modify: `tabstack.yaml.example`

- [ ] **Step 1: Extend the Config struct**

Replace the body of `internal/config/config.go` with:

```go
package config

// Config holds application configuration.
type Config struct {
	// Core settings
	Verbose bool
	Debug   bool
	LogJSON bool

	// Tabstack API client settings
	APIKey  string
	BaseURL string
}
```

- [ ] **Step 2: Add persistent flags and Viper bindings in `cmd/root.go`**

In `cmd/root.go` `init()`, after the existing `Logging flags` block, add:

```go
	// Tabstack API flags
	rootCmd.PersistentFlags().String("api-key", "", "Tabstack API key (env: TABSTACK_API_KEY)")
	rootCmd.PersistentFlags().String("base-url", "", "Tabstack API base URL (env: TABSTACK_BASE_URL)")

	_ = viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	_ = viper.BindPFlag("base_url", rootCmd.PersistentFlags().Lookup("base-url"))
	_ = viper.BindEnv("api_key", "TABSTACK_API_KEY")
	_ = viper.BindEnv("base_url", "TABSTACK_BASE_URL")
```

In `GetConfig()` in the same file, extend the struct literal to include the new fields:

```go
		cfg = &config.Config{
			Verbose: viper.GetBool("verbose"),
			Debug:   viper.GetBool("debug"),
			LogJSON: viper.GetBool("log_json"),
			APIKey:  viper.GetString("api_key"),
			BaseURL: viper.GetString("base_url"),
		}
```

- [ ] **Step 3: Update `tabstack.yaml.example`**

Replace the file contents with:

```yaml
# Example tabstack configuration.
#
# Place this at ./tabstack.yaml or pass --config <path>.
# CLI flags > env vars > this file > defaults.

# Tabstack API base URL (defaults to https://api.tabstack.ai/v1/).
# base_url: https://api.tabstack.ai/v1/

# DO NOT put your API key in this file unless the file is .gitignored.
# Prefer the TABSTACK_API_KEY environment variable.
# api_key: ""

# Logging
verbose: false
debug: false
log_json: false
```

- [ ] **Step 4: Verify build and that `--api-key` flag is visible**

```sh
go build ./...
./tabstack --help 2>&1 | grep -E 'api-key|base-url'
```

Expected: both flags listed in help output.

- [ ] **Step 5: Commit**

```sh
git add cmd/root.go internal/config/config.go tabstack.yaml.example
git commit -m "feat(config): add api_key and base_url config keys"
```

---

### Task 3: Add `internal/client` package

**Files:**
- Create: `internal/client/client.go`

- [ ] **Step 1: Write `internal/client/client.go`**

```go
// Package client constructs a Tabstack SDK client from CLI configuration.
package client

import (
	"fmt"
	"os"

	"github.com/lmorchard/tabstack-go-cli/internal/config"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/option"
)

// New constructs a Tabstack SDK client. Values from cfg are layered on top of
// the SDK's environment defaults (TABSTACK_API_KEY, TABSTACK_BASE_URL), so a
// non-empty cfg.APIKey or cfg.BaseURL overrides whatever was in the env.
//
// Returns an error if no API key resolves from cfg or env — the SDK itself
// would happily make unauthenticated requests and fail with a 401 mid-stream,
// which is a worse experience than failing fast here.
func New(cfg *config.Config) (*tabstack.Client, error) {
	if cfg.APIKey == "" && os.Getenv("TABSTACK_API_KEY") == "" {
		return nil, fmt.Errorf("no Tabstack API key set: provide --api-key, set api_key in config, or export TABSTACK_API_KEY")
	}

	var opts []option.RequestOption
	if cfg.APIKey != "" {
		opts = append(opts, option.WithAPIKey(cfg.APIKey))
	}
	if cfg.BaseURL != "" {
		opts = append(opts, option.WithBaseURL(cfg.BaseURL))
	}

	c := tabstack.NewClient(opts...)
	return &c, nil
}
```

- [ ] **Step 2: Verify build**

```sh
go build ./...
```

Expected: clean build.

- [ ] **Step 3: Commit**

```sh
git add internal/client/client.go
git commit -m "feat(client): add internal/client SDK constructor"
```

---

## Phase 2 — Schema loader (TDD)

### Task 4: `internal/schema` package — JSON Schema loader

**Files:**
- Create: `internal/schema/schema.go`
- Create: `internal/schema/schema_test.go`

- [ ] **Step 1: Write the failing test file**

Create `internal/schema/schema_test.go`:

```go
package schema

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad_FromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "schema.json")
	body := []byte(`{"type":"object","properties":{"name":{"type":"string"}}}`)
	if err := os.WriteFile(path, body, 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	m, ok := got.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", got)
	}
	if m["type"] != "object" {
		t.Errorf("type: got %v, want object", m["type"])
	}
}

func TestLoad_FromStdin(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	orig := os.Stdin
	os.Stdin = r
	t.Cleanup(func() {
		os.Stdin = orig
		_ = r.Close()
	})

	go func() {
		_, _ = w.Write([]byte(`{"type":"string"}`))
		_ = w.Close()
	}()

	got, err := Load("-")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	m, ok := got.(map[string]any)
	if !ok || m["type"] != "string" {
		t.Errorf("got %#v", got)
	}
}

func TestLoad_MissingFile(t *testing.T) {
	_, err := Load("/nonexistent/does/not/exist.json")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
	if !strings.Contains(err.Error(), "open schema") {
		t.Errorf("error should mention 'open schema': %v", err)
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.json")
	if err := os.WriteFile(path, []byte(`{not json`), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := Load(path)
	if err == nil {
		t.Fatal("expected parse error")
	}
	if !strings.Contains(err.Error(), "parse schema") {
		t.Errorf("error should mention 'parse schema': %v", err)
	}
}
```

- [ ] **Step 2: Run tests; expect failure**

```sh
go test ./internal/schema/... -run TestLoad -v 2>&1 | head -30
```

Expected: build error (`undefined: Load`) — the package file does not exist yet.

- [ ] **Step 3: Write the implementation**

Create `internal/schema/schema.go`:

```go
// Package schema loads JSON Schema documents for Tabstack extract/generate
// requests. Schemas are passed through verbatim — no validation here.
package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Load reads a JSON document from the given path and returns it as a generic
// value (typically map[string]any) suitable for the Tabstack SDK's JsonSchema
// field. The path "-" reads from stdin.
func Load(path string) (any, error) {
	var r io.Reader
	if path == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open schema %q: %w", path, err)
		}
		defer func() { _ = f.Close() }()
		r = f
	}

	var schema any
	if err := json.NewDecoder(r).Decode(&schema); err != nil {
		return nil, fmt.Errorf("parse schema %q as JSON: %w", path, err)
	}
	return schema, nil
}
```

- [ ] **Step 4: Run tests; expect pass**

```sh
go test ./internal/schema/... -v
```

Expected: `PASS` for all four tests.

- [ ] **Step 5: Commit**

```sh
git add internal/schema/
git commit -m "feat(schema): add JSON Schema loader (file or stdin)"
```

---

## Phase 3 — Extract commands

### Task 5: `cmd/extract.go` parent and `markdown` subcommand

**Files:**
- Create: `cmd/extract.go`

- [ ] **Step 1: Write the parent command and `markdown` subcommand**

Create `cmd/extract.go`:

```go
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract content from a URL",
	Long:  `Extract markdown or structured JSON from a fetched URL.`,
}

var (
	extractMarkdownMetadata bool
	extractMarkdownNocache  bool
	extractMarkdownEffort   string
	extractMarkdownGeo      string
)

var extractMarkdownCmd = &cobra.Command{
	Use:   "markdown <url>",
	Short: "Fetch a URL and convert to clean markdown",
	Args:  cobra.ExactArgs(1),
	RunE:  runExtractMarkdown,
}

func runExtractMarkdown(_ *cobra.Command, args []string) error {
	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.ExtractMarkdownParams{
		URL: args[0],
	}
	if extractMarkdownMetadata {
		body.Metadata = param.NewOpt(true)
	}
	if extractMarkdownNocache {
		body.Nocache = param.NewOpt(true)
	}
	if extractMarkdownEffort != "" {
		body.Effort = tabstack.ExtractMarkdownParamsEffort(extractMarkdownEffort)
	}
	if extractMarkdownGeo != "" {
		body.GeoTarget = tabstack.ExtractMarkdownParamsGeoTarget{
			Country: param.NewOpt(extractMarkdownGeo),
		}
	}

	resp, err := c.Extract.Markdown(context.Background(), body)
	if err != nil {
		return fmt.Errorf("extract markdown: %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(resp)
}

func init() {
	extractMarkdownCmd.Flags().BoolVar(&extractMarkdownMetadata, "metadata", false, "include extracted metadata")
	extractMarkdownCmd.Flags().BoolVar(&extractMarkdownNocache, "nocache", false, "bypass cache")
	extractMarkdownCmd.Flags().StringVar(&extractMarkdownEffort, "effort", "", "effort level: min, standard, or max")
	extractMarkdownCmd.Flags().StringVar(&extractMarkdownGeo, "geo", "", "ISO 3166-1 alpha-2 country code (e.g. US, GB)")

	extractCmd.AddCommand(extractMarkdownCmd)
	rootCmd.AddCommand(extractCmd)
}
```

- [ ] **Step 2: Verify build**

```sh
go build ./...
```

Expected: clean build.

- [ ] **Step 3: Verify command surface**

```sh
./tabstack extract --help
./tabstack extract markdown --help
```

Expected: both print help text including the four flags on `markdown`.

- [ ] **Step 4: Commit**

```sh
git add cmd/extract.go
git commit -m "feat(cmd): add 'extract markdown' command"
```

---

### Task 6: `extract json` subcommand

**Files:**
- Modify: `cmd/extract.go`

- [ ] **Step 1: Append the `json` subcommand to `cmd/extract.go`**

At the end of `cmd/extract.go`, append:

```go
var (
	extractJsonSchemaPath string
	extractJsonNocache    bool
	extractJsonEffort     string
	extractJsonGeo        string
)

var extractJsonCmd = &cobra.Command{
	Use:   "json <url>",
	Short: "Fetch a URL and extract structured data per a JSON Schema",
	Args:  cobra.ExactArgs(1),
	RunE:  runExtractJson,
}

func runExtractJson(_ *cobra.Command, args []string) error {
	if extractJsonSchemaPath == "" {
		return fmt.Errorf("--schema is required")
	}
	sch, err := schema.Load(extractJsonSchemaPath)
	if err != nil {
		return err
	}

	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.ExtractJsonParams{
		URL:        args[0],
		JsonSchema: sch,
	}
	if extractJsonNocache {
		body.Nocache = param.NewOpt(true)
	}
	if extractJsonEffort != "" {
		body.Effort = tabstack.ExtractJsonParamsEffort(extractJsonEffort)
	}
	if extractJsonGeo != "" {
		body.GeoTarget = tabstack.ExtractJsonParamsGeoTarget{
			Country: param.NewOpt(extractJsonGeo),
		}
	}

	resp, err := c.Extract.Json(context.Background(), body)
	if err != nil {
		return fmt.Errorf("extract json: %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(resp)
}

func initExtractJson() {
	extractJsonCmd.Flags().StringVar(&extractJsonSchemaPath, "schema", "", `path to JSON Schema file (use "-" for stdin)`)
	extractJsonCmd.Flags().BoolVar(&extractJsonNocache, "nocache", false, "bypass cache")
	extractJsonCmd.Flags().StringVar(&extractJsonEffort, "effort", "", "effort level: min, standard, or max")
	extractJsonCmd.Flags().StringVar(&extractJsonGeo, "geo", "", "ISO 3166-1 alpha-2 country code")

	_ = extractJsonCmd.MarkFlagRequired("schema")
	extractCmd.AddCommand(extractJsonCmd)
}
```

- [ ] **Step 2: Add the schema import and register the subcommand**

At the top of `cmd/extract.go`, change the import block to include:

```go
	"github.com/lmorchard/tabstack-go-cli/internal/schema"
```

(Place it next to the existing `internal/client` import.)

Inside the existing `init()` function in `cmd/extract.go`, add a single line at the end (just before the closing brace):

```go
	initExtractJson()
```

- [ ] **Step 3: Verify build**

```sh
go build ./...
```

Expected: clean.

- [ ] **Step 4: Verify command surface**

```sh
./tabstack extract json --help
```

Expected: shows `--schema` (required), `--nocache`, `--effort`, `--geo`.

```sh
./tabstack extract json https://example.com 2>&1 | head -3
```

Expected: error message about missing required `--schema` flag.

- [ ] **Step 5: Commit**

```sh
git add cmd/extract.go
git commit -m "feat(cmd): add 'extract json' command with schema input"
```

---

## Phase 4 — Generate command

### Task 7: `cmd/generate.go` with `json` subcommand

**Files:**
- Create: `cmd/generate.go`

- [ ] **Step 1: Write `cmd/generate.go`**

```go
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	"github.com/lmorchard/tabstack-go-cli/internal/schema"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate AI-transformed content from a URL",
}

var (
	generateJsonSchemaPath       string
	generateJsonInstructions     string
	generateJsonInstructionsFile string
	generateJsonNocache          bool
	generateJsonEffort           string
	generateJsonGeo              string
)

var generateJsonCmd = &cobra.Command{
	Use:   "json <url>",
	Short: "Fetch a URL and AI-transform it into JSON per a schema and instructions",
	Args:  cobra.ExactArgs(1),
	RunE:  runGenerateJson,
}

func runGenerateJson(_ *cobra.Command, args []string) error {
	instructions, err := resolveInstructions(generateJsonInstructions, generateJsonInstructionsFile)
	if err != nil {
		return err
	}
	sch, err := schema.Load(generateJsonSchemaPath)
	if err != nil {
		return err
	}

	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.GenerateJsonParams{
		URL:          args[0],
		Instructions: instructions,
		JsonSchema:   sch,
	}
	if generateJsonNocache {
		body.Nocache = param.NewOpt(true)
	}
	if generateJsonEffort != "" {
		body.Effort = tabstack.GenerateJsonParamsEffort(generateJsonEffort)
	}
	if generateJsonGeo != "" {
		body.GeoTarget = tabstack.GenerateJsonParamsGeoTarget{
			Country: param.NewOpt(generateJsonGeo),
		}
	}

	resp, err := c.Generate.Json(context.Background(), body)
	if err != nil {
		return fmt.Errorf("generate json: %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(resp)
}

// resolveInstructions returns the instructions string from either the inline
// flag or a file path. Exactly one must be set.
func resolveInstructions(inline, path string) (string, error) {
	if inline == "" && path == "" {
		return "", fmt.Errorf("provide --instructions or --instructions-file")
	}
	if inline != "" && path != "" {
		return "", fmt.Errorf("--instructions and --instructions-file are mutually exclusive")
	}
	if inline != "" {
		return inline, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read instructions: %w", err)
	}
	return strings.TrimSpace(string(b)), nil
}

func init() {
	generateJsonCmd.Flags().StringVar(&generateJsonSchemaPath, "schema", "", `path to JSON Schema file (use "-" for stdin)`)
	generateJsonCmd.Flags().StringVar(&generateJsonInstructions, "instructions", "", "transformation instructions (max 20000 chars)")
	generateJsonCmd.Flags().StringVar(&generateJsonInstructionsFile, "instructions-file", "", "path to a file containing the instructions")
	generateJsonCmd.Flags().BoolVar(&generateJsonNocache, "nocache", false, "bypass cache")
	generateJsonCmd.Flags().StringVar(&generateJsonEffort, "effort", "", "effort level: min, standard, or max")
	generateJsonCmd.Flags().StringVar(&generateJsonGeo, "geo", "", "ISO 3166-1 alpha-2 country code")

	_ = generateJsonCmd.MarkFlagRequired("schema")

	generateCmd.AddCommand(generateJsonCmd)
	rootCmd.AddCommand(generateCmd)
}
```

- [ ] **Step 2: Verify build**

```sh
go build ./...
```

Expected: clean.

- [ ] **Step 3: Verify command surface**

```sh
./tabstack generate json --help
```

Expected: shows `--schema` (required), `--instructions`, `--instructions-file`, plus the three common flags.

- [ ] **Step 4: Commit**

```sh
git add cmd/generate.go
git commit -m "feat(cmd): add 'generate json' command"
```

---

## Phase 5 — SSE writer (TDD)

### Task 8: `internal/sse` — JSON-line writer

**Files:**
- Create: `internal/sse/writer.go`
- Create: `internal/sse/writer_test.go`

- [ ] **Step 1: Write the failing test**

Create `internal/sse/writer_test.go`:

```go
package sse

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

type fakeStream[T any] struct {
	items  []T
	i      int
	err    error
	closed bool
}

func (f *fakeStream[T]) Next() bool {
	if f.i >= len(f.items) {
		return false
	}
	f.i++
	return true
}
func (f *fakeStream[T]) Current() T    { return f.items[f.i-1] }
func (f *fakeStream[T]) Err() error    { return f.err }
func (f *fakeStream[T]) Close() error  { f.closed = true; return nil }

type ev struct {
	Type string `json:"type"`
	Msg  string `json:"msg,omitempty"`
}

func TestWriteJSONLines_AllEvents(t *testing.T) {
	s := &fakeStream[ev]{items: []ev{
		{Type: "start", Msg: "go"},
		{Type: "complete", Msg: "done"},
	}}
	var buf bytes.Buffer
	if err := WriteJSONLines[ev](&buf, s); err != nil {
		t.Fatalf("WriteJSONLines: %v", err)
	}
	if !s.closed {
		t.Error("stream not closed")
	}
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("got %d lines, want 2: %q", len(lines), buf.String())
	}
	if !strings.Contains(lines[0], `"type":"start"`) {
		t.Errorf("line 0: %s", lines[0])
	}
}

func TestWriteJSONLines_PropagatesStreamError(t *testing.T) {
	s := &fakeStream[ev]{
		items: []ev{{Type: "start"}},
		err:   errors.New("network died"),
	}
	var buf bytes.Buffer
	err := WriteJSONLines[ev](&buf, s)
	if err == nil || !strings.Contains(err.Error(), "network died") {
		t.Fatalf("expected 'network died', got %v", err)
	}
	if !s.closed {
		t.Error("stream not closed after error")
	}
}

func TestWriteJSONLines_EmptyStream(t *testing.T) {
	s := &fakeStream[ev]{}
	var buf bytes.Buffer
	if err := WriteJSONLines[ev](&buf, s); err != nil {
		t.Fatalf("WriteJSONLines: %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("expected empty output, got %q", buf.String())
	}
	if !s.closed {
		t.Error("stream not closed")
	}
}
```

- [ ] **Step 2: Run tests; expect failure**

```sh
go test ./internal/sse/... -v 2>&1 | head -20
```

Expected: build error (`undefined: WriteJSONLines`).

- [ ] **Step 3: Write the implementation**

Create `internal/sse/writer.go`:

```go
// Package sse renders Tabstack streaming-event responses to plain output.
//
// In v1 the only renderer is WriteJSONLines, which emits one JSON object per
// event. A typed pretty renderer is intentionally deferred — JSON-lines
// composes well with jq and is unambiguous about the underlying event shape.
package sse

import (
	"encoding/json"
	"fmt"
	"io"
)

// Streamer is the minimal interface needed from a Tabstack SDK SSE stream.
// *ssestream.Stream[T] from github.com/stainless-sdks/tabstack-go satisfies
// this interface.
type Streamer[T any] interface {
	Next() bool
	Current() T
	Err() error
	Close() error
}

// WriteJSONLines drains s, encoding each event as one JSON object on its own
// line to w. The stream is always closed before returning. Returns the first
// error from the stream (after Next() returns false) or any encode error.
func WriteJSONLines[T any](w io.Writer, s Streamer[T]) error {
	defer func() { _ = s.Close() }()
	enc := json.NewEncoder(w)
	for s.Next() {
		if err := enc.Encode(s.Current()); err != nil {
			return fmt.Errorf("encode event: %w", err)
		}
	}
	return s.Err()
}
```

- [ ] **Step 4: Run tests; expect pass**

```sh
go test ./internal/sse/... -v
```

Expected: three PASS lines.

- [ ] **Step 5: Commit**

```sh
git add internal/sse/
git commit -m "feat(sse): add JSON-lines writer for typed event streams"
```

---

## Phase 6 — Streaming commands

### Task 9: `cmd/agent.go` parent and `research` subcommand

**Files:**
- Create: `cmd/agent.go`

- [ ] **Step 1: Write `cmd/agent.go`**

```go
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	"github.com/lmorchard/tabstack-go-cli/internal/sse"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "AI-driven browser automation and research",
}

var (
	agentResearchMode         string
	agentResearchNocache      bool
	agentResearchFetchTimeout int64
)

var agentResearchCmd = &cobra.Command{
	Use:   "research <query>",
	Short: "Stream a multi-source research run",
	Args:  cobra.ExactArgs(1),
	RunE:  runAgentResearch,
}

func runAgentResearch(_ *cobra.Command, args []string) error {
	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.AgentResearchParams{
		Query: args[0],
	}
	if agentResearchMode != "" {
		body.Mode = tabstack.AgentResearchParamsMode(agentResearchMode)
	}
	if agentResearchNocache {
		body.Nocache = param.NewOpt(true)
	}
	if agentResearchFetchTimeout > 0 {
		body.FetchTimeout = param.NewOpt(agentResearchFetchTimeout)
	}

	stream := c.Agent.ResearchStreaming(context.Background(), body)
	if err := sse.WriteJSONLines[tabstack.ResearchEvent](os.Stdout, stream); err != nil {
		return fmt.Errorf("research stream: %w", err)
	}
	return nil
}

func init() {
	agentResearchCmd.Flags().StringVar(&agentResearchMode, "mode", "", "research mode: fast or balanced")
	agentResearchCmd.Flags().BoolVar(&agentResearchNocache, "nocache", false, "bypass cache")
	agentResearchCmd.Flags().Int64Var(&agentResearchFetchTimeout, "fetch-timeout", 0, "per-fetch timeout in seconds (0 = SDK default)")

	agentCmd.AddCommand(agentResearchCmd)
	rootCmd.AddCommand(agentCmd)
}
```

- [ ] **Step 2: Verify build**

```sh
go build ./...
```

Expected: clean.

- [ ] **Step 3: Verify command surface**

```sh
./tabstack agent research --help
```

Expected: shows `--mode`, `--nocache`, `--fetch-timeout`.

- [ ] **Step 4: Commit**

```sh
git add cmd/agent.go
git commit -m "feat(cmd): add 'agent research' streaming command"
```

---

### Task 10: `agent automate` subcommand (no interactive yet)

**Files:**
- Modify: `cmd/agent.go`

This task wires the streaming `automate` command but **does not** handle interactive form-data callbacks; if the user passes `--interactive`, we still send `interactive: true` to the API and let the events stream out, but we don't reply to them. Task 13 wires the interactive loop in. This split keeps the streaming-plumbing change separate from the input-collection logic.

- [ ] **Step 1: Append the `automate` subcommand**

Append to `cmd/agent.go`:

```go
var (
	agentAutomateURL         string
	agentAutomateGuardrails  string
	agentAutomateSchemaPath  string
	agentAutomateInteractive bool
	agentAutomateGeo         string
	agentAutomateMaxIter     int64
	agentAutomateMaxValid    int64
)

var agentAutomateCmd = &cobra.Command{
	Use:   "automate <task>",
	Short: "Stream an AI browser-automation run",
	Args:  cobra.ExactArgs(1),
	RunE:  runAgentAutomate,
}

func runAgentAutomate(_ *cobra.Command, args []string) error {
	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.AgentAutomateParams{
		Task: args[0],
	}
	if agentAutomateURL != "" {
		body.URL = param.NewOpt(agentAutomateURL)
	}
	if agentAutomateGuardrails != "" {
		body.Guardrails = param.NewOpt(agentAutomateGuardrails)
	}
	if agentAutomateInteractive {
		body.Interactive = param.NewOpt(true)
	}
	if agentAutomateMaxIter > 0 {
		body.MaxIterations = param.NewOpt(agentAutomateMaxIter)
	}
	if agentAutomateMaxValid > 0 {
		body.MaxValidationAttempts = param.NewOpt(agentAutomateMaxValid)
	}
	if agentAutomateGeo != "" {
		body.GeoTarget = tabstack.AgentAutomateParamsGeoTarget{
			Country: param.NewOpt(agentAutomateGeo),
		}
	}
	if agentAutomateSchemaPath != "" {
		sch, err := schema.Load(agentAutomateSchemaPath)
		if err != nil {
			return err
		}
		body.Data = sch
	}

	stream := c.Agent.AutomateStreaming(context.Background(), body)
	if err := sse.WriteJSONLines[tabstack.AutomateEvent](os.Stdout, stream); err != nil {
		return fmt.Errorf("automate stream: %w", err)
	}
	return nil
}

func initAgentAutomate() {
	agentAutomateCmd.Flags().StringVar(&agentAutomateURL, "url", "", "starting URL for the task")
	agentAutomateCmd.Flags().StringVar(&agentAutomateGuardrails, "guardrails", "", "safety constraints for execution")
	agentAutomateCmd.Flags().StringVar(&agentAutomateSchemaPath, "data-schema", "", `path to JSON file passed as 'data' context (use "-" for stdin)`)
	agentAutomateCmd.Flags().BoolVar(&agentAutomateInteractive, "interactive", false, "enable interactive form-data callbacks (no auto-reply yet — see Task 13)")
	agentAutomateCmd.Flags().StringVar(&agentAutomateGeo, "geo", "", "ISO 3166-1 alpha-2 country code")
	agentAutomateCmd.Flags().Int64Var(&agentAutomateMaxIter, "max-iterations", 0, "max task iterations (0 = SDK default)")
	agentAutomateCmd.Flags().Int64Var(&agentAutomateMaxValid, "max-validation-attempts", 0, "max validation attempts (0 = SDK default)")

	agentCmd.AddCommand(agentAutomateCmd)
}
```

- [ ] **Step 2: Add the schema import and register**

At the top of `cmd/agent.go`, add to the import block:

```go
	"github.com/lmorchard/tabstack-go-cli/internal/schema"
```

Inside the existing `init()` in `cmd/agent.go`, before `rootCmd.AddCommand(agentCmd)`, add:

```go
	initAgentAutomate()
```

- [ ] **Step 3: Verify build**

```sh
go build ./...
```

Expected: clean.

- [ ] **Step 4: Verify command surface**

```sh
./tabstack agent automate --help
```

Expected: shows all seven flags.

- [ ] **Step 5: Commit**

```sh
git add cmd/agent.go
git commit -m "feat(cmd): add 'agent automate' streaming command (no interactive reply yet)"
```

---

## Phase 7 — Interactive automation

### Task 11: `internal/interactive` package — prompter and submitter

**Files:**
- Create: `internal/interactive/interactive.go`
- Create: `internal/interactive/interactive_test.go`

**SDK shape note:** At our pinned SDK commit (`020acb44`), the streaming event types are not yet typed unions — they're plain `tabstack.AutomateEvent { Event string; Data any }`. So this package defines its own neutral `FormDataRequest`/`FormDataField` types and a `DecodeFormDataRequest(any) (FormDataRequest, error)` helper that JSON-roundtrips the `Data` field into the local type. This isolates us from upcoming SDK schema churn.

The package has three public collaborators:

- **`FormDataRequest`/`FormDataField`** — neutral data types (no SDK dep)
- **Prompter** — given a `FormDataRequest`, returns either field values or "cancel". Two implementations: TTY (uses `bufio.Scanner` on os.Stdin and prints to os.Stderr) and File (loads a JSON map keyed by request ID).
- **Submitter** — wraps the SDK's `Agent.AutomateInput` call. Defining it as an interface lets us test the prompter logic without a live SDK.

Plus a `DecodeFormDataRequest` helper that callers use before invoking `HandleRequest`.

The TTY prompter is hard to unit-test in isolation; we test the file prompter, the decoder, and the wiring function directly.

- [ ] **Step 1: Write the failing test file**

Create `internal/interactive/interactive_test.go`:

```go
package interactive

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func makeRequest(reqID string, fieldRefs ...string) FormDataRequest {
	fields := make([]FormDataField, 0, len(fieldRefs))
	for _, ref := range fieldRefs {
		fields = append(fields, FormDataField{
			Ref:       ref,
			Label:     "Field " + ref,
			FieldType: "text",
			Required:  true,
		})
	}
	return FormDataRequest{
		RequestID: reqID,
		Fields:    fields,
	}
}

func TestDecodeFormDataRequest_RoundTrip(t *testing.T) {
	raw := map[string]any{
		"requestId":       "req-1",
		"pageTitle":       "Sign in",
		"pageUrl":         "https://example.com/login",
		"formDescription": "Login form",
		"fields": []any{
			map[string]any{
				"ref":       "E1",
				"label":     "Email",
				"fieldType": "email",
				"required":  true,
			},
			map[string]any{
				"ref":       "E2",
				"label":     "Password",
				"fieldType": "password",
				"required":  true,
			},
		},
	}
	got, err := DecodeFormDataRequest(raw)
	if err != nil {
		t.Fatalf("DecodeFormDataRequest: %v", err)
	}
	if got.RequestID != "req-1" {
		t.Errorf("RequestID: %q", got.RequestID)
	}
	if got.PageTitle != "Sign in" || got.PageURL != "https://example.com/login" {
		t.Errorf("page: %q / %q", got.PageTitle, got.PageURL)
	}
	if len(got.Fields) != 2 {
		t.Fatalf("fields: %d", len(got.Fields))
	}
	if got.Fields[0].Ref != "E1" || got.Fields[0].FieldType != "email" {
		t.Errorf("field 0: %+v", got.Fields[0])
	}
}

func TestFilePrompter_AnswersByRequestID(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "answers.json")
	body := []byte(`{
		"req-1": {"E1": "alice", "E2": "alice@example.com"},
		"req-2": {"E1": "bob"}
	}`)
	if err := os.WriteFile(path, body, 0o600); err != nil {
		t.Fatal(err)
	}

	p, err := NewFilePrompter(path)
	if err != nil {
		t.Fatalf("NewFilePrompter: %v", err)
	}
	got, cancelled, err := p.Prompt(makeRequest("req-1", "E1", "E2"))
	if err != nil {
		t.Fatalf("Prompt: %v", err)
	}
	if cancelled {
		t.Fatal("unexpected cancel")
	}
	if len(got) != 2 {
		t.Fatalf("got %d fields, want 2", len(got))
	}
	values := map[string]string{}
	for _, f := range got {
		values[f.Ref] = f.Value
	}
	if values["E1"] != "alice" || values["E2"] != "alice@example.com" {
		t.Errorf("values: %+v", values)
	}
}

func TestFilePrompter_MissingRequestIDCancels(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "answers.json")
	if err := os.WriteFile(path, []byte(`{}`), 0o600); err != nil {
		t.Fatal(err)
	}
	p, err := NewFilePrompter(path)
	if err != nil {
		t.Fatal(err)
	}
	_, cancelled, err := p.Prompt(makeRequest("req-missing", "E1"))
	if err != nil {
		t.Fatalf("Prompt: %v", err)
	}
	if !cancelled {
		t.Error("expected cancel for missing request id")
	}
}

type fakeSubmitter struct {
	calls []submitCall
	err   error
}

type submitCall struct {
	requestID string
	fields    []FieldValue
	cancelled bool
}

func (f *fakeSubmitter) Submit(_ context.Context, requestID string, fields []FieldValue, cancelled bool) error {
	f.calls = append(f.calls, submitCall{requestID, fields, cancelled})
	return f.err
}

func TestHandleRequest_PromptThenSubmit(t *testing.T) {
	prompter := stubPrompter{
		fields:    []FieldValue{{Ref: "E1", Value: "hello"}},
		cancelled: false,
	}
	sub := &fakeSubmitter{}
	err := HandleRequest(context.Background(), sub, prompter, makeRequest("req-x", "E1"))
	if err != nil {
		t.Fatalf("HandleRequest: %v", err)
	}
	if len(sub.calls) != 1 {
		t.Fatalf("submits: %d", len(sub.calls))
	}
	if sub.calls[0].requestID != "req-x" {
		t.Errorf("requestID: %s", sub.calls[0].requestID)
	}
	if sub.calls[0].cancelled {
		t.Error("expected non-cancelled submit")
	}
}

func TestHandleRequest_CancelledPath(t *testing.T) {
	prompter := stubPrompter{cancelled: true}
	sub := &fakeSubmitter{}
	err := HandleRequest(context.Background(), sub, prompter, makeRequest("req-y", "E1"))
	if err != nil {
		t.Fatalf("HandleRequest: %v", err)
	}
	if !sub.calls[0].cancelled {
		t.Error("expected cancelled submit")
	}
	if len(sub.calls[0].fields) != 0 {
		t.Errorf("expected no fields on cancel, got %v", sub.calls[0].fields)
	}
}

func TestHandleRequest_PromptError(t *testing.T) {
	prompter := stubPrompter{err: errors.New("prompt blew up")}
	sub := &fakeSubmitter{}
	err := HandleRequest(context.Background(), sub, prompter, makeRequest("req-z", "E1"))
	if err == nil {
		t.Fatal("expected error")
	}
	if len(sub.calls) != 0 {
		t.Error("submitter should not be called on prompt error")
	}
}

type stubPrompter struct {
	fields    []FieldValue
	cancelled bool
	err       error
}

func (s stubPrompter) Prompt(_ FormDataRequest) ([]FieldValue, bool, error) {
	return s.fields, s.cancelled, s.err
}
```

- [ ] **Step 2: Run tests; expect failure**

```sh
go test ./internal/interactive/... -v 2>&1 | head -30
```

Expected: build errors — none of the symbols exist yet.

- [ ] **Step 3: Write the implementation**

Create `internal/interactive/interactive.go`:

```go
// Package interactive handles Tabstack agent.automate's mid-stream
// "interactive:form_data:request" events. The package defines neutral
// FormDataRequest/FormDataField types (decoded from the SDK's untyped
// event payload), a Prompter abstraction (TTY or file-source), and a
// thin Submitter wrapper around the SDK's Agent.AutomateInput call.
package interactive

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

// FormDataRequest is the decoded payload of an
// "interactive:form_data:request" event. The struct tags match the API's
// JSON shape so DecodeFormDataRequest can JSON-roundtrip the SDK's
// untyped event Data into this type.
type FormDataRequest struct {
	RequestID       string          `json:"requestId"`
	IterationID     string          `json:"iterationId"`
	PageTitle       string          `json:"pageTitle"`
	PageURL         string          `json:"pageUrl"`
	FormDescription string          `json:"formDescription"`
	Fields          []FormDataField `json:"fields"`
}

// FormDataField is a single field the agent needs data for.
type FormDataField struct {
	Ref          string   `json:"ref"`
	Label        string   `json:"label"`
	FieldType    string   `json:"fieldType"`
	Required     bool     `json:"required"`
	CurrentValue string   `json:"currentValue"`
	Description  string   `json:"description"`
	Options      []string `json:"options"`
}

// FieldValue is a single {ref, value} answer for a form field.
type FieldValue struct {
	Ref   string
	Value string
}

// DecodeFormDataRequest converts an SDK event Data payload (typed as `any`,
// usually a map[string]any from JSON unmarshalling) into a FormDataRequest.
func DecodeFormDataRequest(data any) (FormDataRequest, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return FormDataRequest{}, fmt.Errorf("marshal event data: %w", err)
	}
	var req FormDataRequest
	if err := json.Unmarshal(b, &req); err != nil {
		return FormDataRequest{}, fmt.Errorf("unmarshal event data: %w", err)
	}
	return req, nil
}

// Prompter resolves form-data requests into field values.
type Prompter interface {
	// Prompt returns either a list of field values or a cancellation. An
	// error means the prompt itself failed (e.g. read error); the caller
	// must not submit anything in that case.
	Prompt(req FormDataRequest) ([]FieldValue, bool, error)
}

// Submitter posts a response to /automate/{requestId}/input.
type Submitter interface {
	Submit(ctx context.Context, requestID string, fields []FieldValue, cancelled bool) error
}

// HandleRequest runs the prompter and forwards its result to the submitter.
// On prompt error the submitter is NOT called.
func HandleRequest(ctx context.Context, s Submitter, p Prompter, req FormDataRequest) error {
	fields, cancelled, err := p.Prompt(req)
	if err != nil {
		return fmt.Errorf("prompt for request %s: %w", req.RequestID, err)
	}
	if cancelled {
		fields = nil
	}
	return s.Submit(ctx, req.RequestID, fields, cancelled)
}

// --- TTY prompter ---

// TTYPrompter prompts on stderr and reads from stdin.
type TTYPrompter struct {
	in  io.Reader
	out io.Writer
}

// NewTTYPrompter constructs a prompter that reads from os.Stdin and writes
// prompts to os.Stderr (so stdout stays reserved for SSE events).
func NewTTYPrompter() *TTYPrompter {
	return &TTYPrompter{in: os.Stdin, out: os.Stderr}
}

func (t *TTYPrompter) Prompt(req FormDataRequest) ([]FieldValue, bool, error) {
	fmt.Fprintf(t.out, "\n=== Tabstack form-data request %s ===\n", req.RequestID)
	fmt.Fprintf(t.out, "Page: %s (%s)\n", req.PageTitle, req.PageURL)
	if req.FormDescription != "" {
		fmt.Fprintf(t.out, "Form: %s\n", req.FormDescription)
	}
	fmt.Fprintln(t.out, `Type "/cancel" at any prompt to cancel the request.`)

	scanner := bufio.NewScanner(t.in)
	values := make([]FieldValue, 0, len(req.Fields))
	for _, f := range req.Fields {
		label := f.Label
		if label == "" {
			label = f.Ref
		}
		marker := "(required)"
		if !f.Required {
			marker = "(optional)"
		}
		extra := ""
		if f.Description != "" {
			extra = " — " + f.Description
		}
		if len(f.Options) > 0 {
			extra += " options: " + strings.Join(f.Options, ", ")
		}
		fmt.Fprintf(t.out, "  %s %s [%s]%s\n  > ", label, marker, f.FieldType, extra)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return nil, false, fmt.Errorf("read input: %w", err)
			}
			return nil, true, nil // EOF — treat as cancel
		}
		v := scanner.Text()
		if v == "/cancel" {
			return nil, true, nil
		}
		values = append(values, FieldValue{Ref: f.Ref, Value: v})
	}
	return values, false, nil
}

// --- File prompter ---

// FilePrompter looks up answers in a JSON file keyed by request ID.
// File shape:
//
//	{
//	  "<requestId>": { "<fieldRef>": "<value>", ... },
//	  ...
//	}
//
// A missing request ID is treated as a cancel.
type FilePrompter struct {
	answers map[string]map[string]string
}

func NewFilePrompter(path string) (*FilePrompter, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	var m map[string]map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("parse input file: %w", err)
	}
	return &FilePrompter{answers: m}, nil
}

func (f *FilePrompter) Prompt(req FormDataRequest) ([]FieldValue, bool, error) {
	answers, ok := f.answers[req.RequestID]
	if !ok {
		return nil, true, nil
	}
	values := make([]FieldValue, 0, len(req.Fields))
	for _, fld := range req.Fields {
		if v, ok := answers[fld.Ref]; ok {
			values = append(values, FieldValue{Ref: fld.Ref, Value: v})
		}
	}
	return values, false, nil
}

// --- SDK submitter ---

// SDKSubmitter implements Submitter against the live Tabstack client.
type SDKSubmitter struct {
	Client *tabstack.Client
}

func (s SDKSubmitter) Submit(ctx context.Context, requestID string, fields []FieldValue, cancelled bool) error {
	body := tabstack.AgentAutomateInputParams{}
	if cancelled {
		body.Cancelled = param.NewOpt(true)
	} else {
		body.Fields = make([]tabstack.AgentAutomateInputParamsField, 0, len(fields))
		for _, f := range fields {
			body.Fields = append(body.Fields, tabstack.AgentAutomateInputParamsField{
				Ref:   param.NewOpt(f.Ref),
				Value: param.NewOpt(f.Value),
			})
		}
	}
	_, err := s.Client.Agent.AutomateInput(ctx, requestID, body)
	return err
}
```

- [ ] **Step 4: Run tests; expect pass**

```sh
go test ./internal/interactive/... -v
```

Expected: five PASS lines.

- [ ] **Step 5: Verify the package builds in the wider tree**

```sh
go build ./...
```

Expected: clean.

- [ ] **Step 6: Commit**

```sh
git add internal/interactive/
git commit -m "feat(interactive): add TTY/file prompter and SDK submitter for form-data callbacks"
```

---

### Task 12: `agent input` standalone subcommand

**Files:**
- Modify: `cmd/agent.go`

The standalone subcommand exists for after-the-fact resume (e.g. you saw the requestID scroll past, want to answer it from a different terminal) and for plain testing.

- [ ] **Step 1: Append the `input` subcommand**

Append to `cmd/agent.go`:

```go
var (
	agentInputValuesPath string
	agentInputCancel     bool
)

var agentInputCmd = &cobra.Command{
	Use:   "input <request-id>",
	Short: "Submit a response to an interactive form-data request",
	Long: `Reply to (or cancel) an in-flight automation form-data request, identified
by the requestId emitted in an interactive:form_data:request SSE event.

Provide either --values <file> (JSON: {"E1":"alice","E2":"bob"}) or --cancel.`,
	Args: cobra.ExactArgs(1),
	RunE: runAgentInput,
}

func runAgentInput(_ *cobra.Command, args []string) error {
	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.AgentAutomateInputParams{}
	switch {
	case agentInputCancel && agentInputValuesPath != "":
		return fmt.Errorf("--cancel and --values are mutually exclusive")
	case agentInputCancel:
		body.Cancelled = param.NewOpt(true)
	case agentInputValuesPath != "":
		raw, err := os.ReadFile(agentInputValuesPath)
		if err != nil {
			return fmt.Errorf("read values: %w", err)
		}
		var values map[string]string
		if err := json.Unmarshal(raw, &values); err != nil {
			return fmt.Errorf("parse values as JSON object: %w", err)
		}
		body.Fields = make([]tabstack.AgentAutomateInputParamsField, 0, len(values))
		for ref, v := range values {
			body.Fields = append(body.Fields, tabstack.AgentAutomateInputParamsField{
				Ref:   param.NewOpt(ref),
				Value: param.NewOpt(v),
			})
		}
	default:
		return fmt.Errorf("provide --values <file> or --cancel")
	}

	_, err = c.Agent.AutomateInput(context.Background(), args[0], body)
	if err != nil {
		return fmt.Errorf("submit input: %w", err)
	}
	fmt.Fprintln(os.Stderr, "ok")
	return nil
}

func initAgentInput() {
	agentInputCmd.Flags().StringVar(&agentInputValuesPath, "values", "", `path to JSON file mapping field ref -> value`)
	agentInputCmd.Flags().BoolVar(&agentInputCancel, "cancel", false, "cancel the request instead of providing values")
	agentCmd.AddCommand(agentInputCmd)
}
```

- [ ] **Step 2: Add the `encoding/json` import and register**

In `cmd/agent.go` imports, add `"encoding/json"`.

In the existing `init()` in `cmd/agent.go`, add before `rootCmd.AddCommand(agentCmd)`:

```go
	initAgentInput()
```

- [ ] **Step 3: Verify build**

```sh
go build ./...
```

Expected: clean.

- [ ] **Step 4: Verify command surface**

```sh
./tabstack agent input --help
```

Expected: shows `--values` and `--cancel`.

- [ ] **Step 5: Commit**

```sh
git add cmd/agent.go
git commit -m "feat(cmd): add 'agent input' standalone resume command"
```

---

### Task 13: Wire the interactive prompter into `agent automate`

**Files:**
- Modify: `cmd/agent.go`

Up to now `agent automate` has just dumped events. Now we make it react to `AutomateEventInteractiveFormDataRequest` events when `--interactive` is set, choosing a Prompter based on whether `--input-from FILE` was given and whether stdin is a TTY.

Note: the SSE writer in `internal/sse` doesn't expose a per-event hook, and rebuilding it for one caller would be over-engineering. We replace the simple `sse.WriteJSONLines` call in the automate path with an inline drain loop that does both the JSON-line emit and the interactive handling. The other streaming command (`agent research`) keeps using `sse.WriteJSONLines`.

- [ ] **Step 1: Add the `--input-from` flag**

In `initAgentAutomate()`, add a new flag declaration alongside the others:

```go
	agentAutomateCmd.Flags().StringVar(&agentAutomateInputFrom, "input-from", "", `path to JSON file with answers for interactive form-data requests`)
```

And declare the variable in the existing `var (...)` block at the top of `cmd/agent.go`:

```go
	agentAutomateInputFrom   string
```

- [ ] **Step 2: Replace `runAgentAutomate` with an interactive-aware drain loop**

Replace the body of `runAgentAutomate` with:

```go
func runAgentAutomate(_ *cobra.Command, args []string) error {
	ctx := context.Background()

	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.AgentAutomateParams{
		Task: args[0],
	}
	if agentAutomateURL != "" {
		body.URL = param.NewOpt(agentAutomateURL)
	}
	if agentAutomateGuardrails != "" {
		body.Guardrails = param.NewOpt(agentAutomateGuardrails)
	}
	if agentAutomateInteractive {
		body.Interactive = param.NewOpt(true)
	}
	if agentAutomateMaxIter > 0 {
		body.MaxIterations = param.NewOpt(agentAutomateMaxIter)
	}
	if agentAutomateMaxValid > 0 {
		body.MaxValidationAttempts = param.NewOpt(agentAutomateMaxValid)
	}
	if agentAutomateGeo != "" {
		body.GeoTarget = tabstack.AgentAutomateParamsGeoTarget{
			Country: param.NewOpt(agentAutomateGeo),
		}
	}
	if agentAutomateSchemaPath != "" {
		sch, err := schema.Load(agentAutomateSchemaPath)
		if err != nil {
			return err
		}
		body.Data = sch
	}

	prompter, err := selectPrompter(agentAutomateInteractive, agentAutomateInputFrom)
	if err != nil {
		return err
	}
	submitter := interactive.SDKSubmitter{Client: c}

	stream := c.Agent.AutomateStreaming(ctx, body)
	defer func() { _ = stream.Close() }()

	enc := json.NewEncoder(os.Stdout)
	for stream.Next() {
		ev := stream.Current()
		if err := enc.Encode(ev); err != nil {
			return fmt.Errorf("encode event: %w", err)
		}
		if prompter == nil {
			continue
		}
		if ev.Event != "interactive:form_data:request" {
			continue
		}
		req, err := interactive.DecodeFormDataRequest(ev.Data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "interactive: decode form-data event: %v\n", err)
			continue
		}
		if err := interactive.HandleRequest(ctx, submitter, prompter, req); err != nil {
			fmt.Fprintf(os.Stderr, "interactive: %v\n", err)
		}
	}
	if err := stream.Err(); err != nil {
		return fmt.Errorf("automate stream: %w", err)
	}
	return nil
}

// selectPrompter returns nil when no prompter is needed (interactive mode off
// or no callback expected). It only returns a prompter when --interactive was
// set; otherwise the API will not emit form-data events and there's nothing
// to handle.
func selectPrompter(interactiveMode bool, inputFrom string) (interactive.Prompter, error) {
	if !interactiveMode {
		return nil, nil
	}
	if inputFrom != "" {
		return interactive.NewFilePrompter(inputFrom)
	}
	if isTerminal(os.Stdin) {
		return interactive.NewTTYPrompter(), nil
	}
	// Interactive mode requested but no TTY and no --input-from. Fail-safe by
	// auto-cancelling any callbacks rather than blocking forever on stdin.
	fmt.Fprintln(os.Stderr, "warning: --interactive set but stdin is not a TTY and --input-from is empty; form-data requests will be auto-cancelled")
	return autoCancelPrompter{}, nil
}

type autoCancelPrompter struct{}

func (autoCancelPrompter) Prompt(_ interactive.FormDataRequest) ([]interactive.FieldValue, bool, error) {
	return nil, true, nil
}

func isTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}
```

- [ ] **Step 3: Update imports in `cmd/agent.go`**

Add to the imports:

```go
	"encoding/json"
	"github.com/lmorchard/tabstack-go-cli/internal/interactive"
```

(The `encoding/json` import was added in Task 12; if it's already there, leave it.)

Remove the `internal/sse` import only if no other use of it remains in `cmd/agent.go` (research still uses it, so leave it in).

- [ ] **Step 4: Verify build**

```sh
go build ./...
```

Expected: clean.

- [ ] **Step 5: Verify command surface**

```sh
./tabstack agent automate --help
```

Expected: shows `--input-from` alongside the existing flags.

- [ ] **Step 6: Commit**

```sh
git add cmd/agent.go
git commit -m "feat(cmd): wire interactive form-data prompter into 'agent automate'"
```

---

## Phase 8 — Smoke test and docs

### Task 14: End-to-end smoke against the live API

This task is a **manual** verification, not code. The intent is to confirm every wired endpoint actually works against `api.tabstack.ai`. Findings inform whether to iterate further before declaring v1 done.

**Prerequisite:** `export TABSTACK_API_KEY=...` (a live key).

**Files:**
- Create: `docs/dev-sessions/2026-04-27-1123-tabstack-cli-v1/smoke-results.md` (capture findings)

- [ ] **Step 1: Build a fresh binary**

```sh
make build
```

- [ ] **Step 2: Smoke each command**

```sh
# 1. extract markdown — small, public page
./tabstack extract markdown https://example.com

# 2. extract markdown with metadata
./tabstack extract markdown https://example.com --metadata

# 3. extract json — using a tiny inline schema piped via stdin
echo '{"type":"object","properties":{"title":{"type":"string"}}}' | \
  ./tabstack extract json https://example.com --schema -

# 4. generate json
echo '{"type":"object","properties":{"summary":{"type":"string"}}}' | \
  ./tabstack generate json https://example.com \
    --schema - \
    --instructions "Summarize the page in one sentence."

# 5. agent research — fast mode
./tabstack agent research "What is the capital of France?" --mode fast | head -20

# 6. agent automate — non-interactive (just confirm the stream opens)
./tabstack agent automate "Visit example.com and tell me the page title" --url https://example.com | head -20
```

- [ ] **Step 3: Smoke interactive automation (optional, requires a task that asks for form data)**

Pick a task likely to trigger an interactive callback (e.g. "Sign in to <site>" against a public sandbox). Run with `--interactive` and answer prompts.

Skip if no convenient sandbox is available. The standalone `agent input` command can be smoke-tested separately by capturing a `requestId` from a stream and then in another shell running:

```sh
echo '{"E1":"test"}' > /tmp/values.json
./tabstack agent input <request-id> --values /tmp/values.json
```

- [ ] **Step 4: Record findings in `smoke-results.md`**

Write a short summary: which commands worked first try, which had surprises, anything to fix in a follow-up. No template — be honest about gaps.

- [ ] **Step 5: Commit findings**

```sh
git add docs/dev-sessions/2026-04-27-1123-tabstack-cli-v1/smoke-results.md
git commit -m "docs: add v1 smoke-test results"
```

---

### Task 15: Update CLAUDE.md

**Files:**
- Modify: `CLAUDE.md`

- [ ] **Step 1: Update the `Project Purpose` section**

Replace the `## Project Purpose` section with:

```markdown
## Project Purpose

`tabstack` is a Go CLI that wraps the **Tabstack API** (https://docs.tabstack.ai/) via the Stainless-generated Go SDK at `github.com/stainless-sdks/tabstack-go` (pinned by commit; no semver tags published).

Module path: `github.com/lmorchard/tabstack-go-cli`. Binary: `tabstack`.

The API surface is small and action-oriented (no resource hierarchy):

| Service | Operation | Notes |
|---|---|---|
| `extract` | `markdown`, `json` | One-shot: URL → markdown / URL + schema → JSON |
| `generate` | `json` | URL + schema + instructions → AI-transformed JSON |
| `agent` | `automate` | SSE-streamed browser automation; supports interactive form-data callbacks |
| `agent` | `research` | SSE-streamed cited research |
| `agent` | `input` | Reply to an in-flight `automate` form-data request (2-min window) |
```

- [ ] **Step 2: Add a `## Package Layout` section after `## Architecture`**

```markdown
## Package Layout

- `cmd/` — Cobra command definitions; one file per top-level command (`extract.go`, `generate.go`, `agent.go`). Commands stay thin: parse flags, build SDK params, call into `internal/`.
- `internal/client` — constructs the Tabstack SDK client from `*config.Config`. Layers `--api-key`/`--base-url` on top of the SDK's env defaults. Fails fast when no key resolves.
- `internal/schema` — loads JSON Schema documents from a path or `-` (stdin). Returns `any` for direct use as the SDK's `JsonSchema` field.
- `internal/sse` — generic JSON-lines writer for any `*ssestream.Stream[T]`-like source. Used by `agent research`. (`agent automate` inlines its own drain loop because it interleaves event emit with interactive callbacks.)
- `internal/interactive` — `Prompter` interface (TTY or file-source) and `Submitter` interface (wraps `Agent.AutomateInput`). The `--interactive` flag on `agent automate` engages this; `--input-from FILE` substitutes the file prompter for non-TTY runs.
```

- [ ] **Step 3: Verify CLAUDE.md still renders sensibly**

Read it through; no broken section ordering, no contradictory claims with the older content.

- [ ] **Step 4: Commit**

```sh
git add CLAUDE.md
git commit -m "docs: update CLAUDE.md with v1 architecture and command tree"
```

---

## Self-review checklist

After execution, the following should be true:

- `go build ./...` clean
- `go test ./...` clean (schema, sse, interactive packages have tests)
- `make lint` clean (`make setup` first if golangci-lint isn't installed)
- `./tabstack --help` shows: `extract`, `generate`, `agent`, `version` subcommands
- Each subcommand's `--help` lists every flag mentioned in its task
- Smoke test in Task 14 confirms each command actually reaches the API

## Out of scope (future dev sessions)

- Pretty-printed SSE event renderer (typed switch over event union, with progress timestamps and color)
- Screenshot extraction to `--screenshot-dir`
- `tabstack raw <method> <path>` escape hatch using `Client.Execute`/`Get`/`Post`
- `--output {pretty|json|raw}` global flag (right now extract/generate emit pretty JSON, streaming emits JSON-lines — no choice yet)
- Output to file via `-o` (currently always stdout)
- Config file API-key warning if it's set in a non-`.gitignore`d file
- Tests for the Cobra command wiring layer (currently smoke-only)
