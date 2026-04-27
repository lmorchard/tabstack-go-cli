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
