package airbyte

import (
	"bufio"
	"io"
	"sync"
)

type safeWriter struct {
	w  io.Writer
	mu sync.Mutex
}

func newSafeWriter(w io.Writer) io.Writer {
	return &safeWriter{
		w: bufio.NewWriter(w),
	}
}

func (sw *safeWriter) Write(p []byte) (int, error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	return sw.w.Write(p)
}
