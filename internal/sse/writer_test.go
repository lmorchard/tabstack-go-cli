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
func (f *fakeStream[T]) Current() T   { return f.items[f.i-1] }
func (f *fakeStream[T]) Err() error   { return f.err }
func (f *fakeStream[T]) Close() error { f.closed = true; return nil }

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
