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
