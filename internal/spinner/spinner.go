// Package spinner provides a small in-line progress spinner for TTY output.
// It's a cosmetic device for streaming commands ('research', 'automate') so
// the user sees something happening between SSE events.
//
// The spinner ticks on a goroutine and rewrites a single line via "\r" + the
// ANSI "erase to end of line" escape. Callers must call ClearLine before
// printing their own output so the spinner frame is wiped first.
//
// When disabled (Spinner.New(out, false)), every method is a no-op — useful
// for non-TTY stdout where ANSI escapes and "\r" tricks would corrupt output.
package spinner

import (
	"fmt"
	"io"
	"sync"
	"time"
)

// Braille dot frames; widely supported in terminal fonts and visually compact.
var frames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// Spinner is a cancellable line-oriented progress indicator.
type Spinner struct {
	out      io.Writer
	enabled  bool
	interval time.Duration
	started  time.Time

	mu      sync.Mutex
	stop    chan struct{}
	done    chan struct{}
	visible bool
}

// New returns a Spinner that writes to out. When enabled is false, all
// methods are no-ops.
func New(out io.Writer, enabled bool) *Spinner {
	return &Spinner{
		out:      out,
		enabled:  enabled,
		interval: 100 * time.Millisecond,
		started:  time.Now(),
	}
}

// Start kicks off the ticker goroutine. Calling Start on an already-started
// spinner is a no-op.
func (s *Spinner) Start() {
	if !s.enabled {
		return
	}
	s.mu.Lock()
	if s.stop != nil {
		s.mu.Unlock()
		return
	}
	s.stop = make(chan struct{})
	s.done = make(chan struct{})
	stop, done := s.stop, s.done
	s.mu.Unlock()
	// Pass the channels directly: Stop nils the fields when shutting down, so
	// the goroutine can't read them safely from s.
	go s.run(stop, done)
}

// Stop halts the ticker, clears the spinner line, and waits for the goroutine
// to exit. Safe to call multiple times.
func (s *Spinner) Stop() {
	if !s.enabled {
		return
	}
	s.mu.Lock()
	stop, done := s.stop, s.done
	s.stop, s.done = nil, nil
	s.mu.Unlock()
	if stop == nil {
		return
	}
	close(stop)
	<-done
	s.ClearLine()
}

// ClearLine wipes the current spinner frame so the caller can write its own
// output. Idempotent.
func (s *Spinner) ClearLine() {
	if !s.enabled {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.visible {
		return
	}
	// \r moves to column 0; ESC[K erases from cursor to end of line.
	_, _ = fmt.Fprint(s.out, "\r\033[K")
	s.visible = false
}

func (s *Spinner) run(stop, done chan struct{}) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()
	defer close(done)

	var i int
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			s.mu.Lock()
			elapsed := int(time.Since(s.started).Seconds())
			_, _ = fmt.Fprintf(s.out, "\r\033[K%s %ds", frames[i%len(frames)], elapsed)
			s.visible = true
			i++
			s.mu.Unlock()
		}
	}
}
