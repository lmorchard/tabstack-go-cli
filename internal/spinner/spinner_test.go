package spinner

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

// When disabled, every method should be a no-op and the buffer untouched.
func TestDisabledSpinnerWritesNothing(t *testing.T) {
	var buf bytes.Buffer
	s := New(&buf, false)
	s.Start()
	time.Sleep(50 * time.Millisecond)
	s.ClearLine()
	s.Stop()
	if buf.Len() != 0 {
		t.Errorf("disabled spinner wrote %d bytes; want 0 (got %q)", buf.Len(), buf.String())
	}
}

// Enabled spinner should emit at least one frame within ~3 ticks, then leave
// the buffer in a "cleared" state after Stop (the trailing escape is the
// line-clear, not a frame).
func TestEnabledSpinnerEmitsFramesAndClearsOnStop(t *testing.T) {
	var buf bytes.Buffer
	s := New(&buf, true)
	s.interval = 10 * time.Millisecond // tighter for test
	s.Start()
	time.Sleep(50 * time.Millisecond)
	s.Stop()

	out := buf.String()
	if !strings.Contains(out, "\r\033[K") {
		t.Errorf("expected at least one clear sequence (\\r\\033[K) in output: %q", out)
	}
	// Last write is always the clear-on-stop, so output ends with the escape.
	if !strings.HasSuffix(out, "\r\033[K") {
		t.Errorf("expected output to end with clear sequence: %q", out)
	}
}

// ClearLine on a never-visible spinner should not write anything.
func TestClearLineBeforeFirstTickIsNoop(t *testing.T) {
	var buf bytes.Buffer
	s := New(&buf, true)
	s.ClearLine()
	if buf.Len() != 0 {
		t.Errorf("expected no output before first tick; got %q", buf.String())
	}
}

// Calling Stop without Start must not panic or hang.
func TestStopWithoutStartIsSafe(t *testing.T) {
	s := New(&bytes.Buffer{}, true)
	s.Stop()
}

// Regression: Stop used to nil out s.stop/s.done before signalling, but the
// goroutine read s.stop directly inside its select. After nilling, the
// `case <-s.stop:` arm became `case <-nil:` (blocks forever), so Stop
// deadlocked on <-done until go test's 10-minute timeout fired. This test
// runs many Start/Stop cycles tightly so any reintroduction of the race
// shows up as a test-timeout in seconds, not in 600s.
func TestStartStopCyclesDoNotDeadlock(t *testing.T) {
	s := New(&bytes.Buffer{}, true)
	s.interval = time.Millisecond // fast ticks to maximize the race window
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 20; i++ {
			s.Start()
			time.Sleep(2 * time.Millisecond)
			s.Stop()
		}
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Start/Stop cycles deadlocked")
	}
}
