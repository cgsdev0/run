package run

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/amonks/run/internal/mutex"
)

func newPrinter(run *Run) UI {
	return &printer{mu: mutex.New("printer"), run: run}
}

type printer struct {
	mu        *mutex.Mutex
	run       *Run
	stdout    io.Writer
	keyLength int
	lastKey   string
}

// *printer implements MultiWriter
var _ MultiWriter = &printer{}

func (p *printer) Writer(id string) io.Writer {
	return printerWriter{p, id}
}

var _ io.Writer = printerWriter{}

type printerWriter struct {
	printer *printer
	id      string
}

func (w printerWriter) Write(bs []byte) (int, error) {
	w.printer.Write(w.id, string(bs))
	return len(bs), nil
}

func (p *printer) Start(ctx context.Context, ready chan<- struct{}, _ io.Reader, stdout io.Writer) error {
	p.mu.Lock("Write")
	p.stdout = stdout
	p.keyLength = 0
	for _, id := range p.run.IDs() {
		if len(id) > p.keyLength {
			p.keyLength = len(id)
		}
	}
	p.mu.Unlock()

	ready <- struct{}{}

	<-ctx.Done()

	return nil
}

func (p *printer) Write(key, message string) {
	defer p.mu.Lock("Write").Unlock()

	if p.stdout == nil {
		panic("nil stdout in printer")
	}

	lines := strings.Split(message, "\n")
	for _, l := range lines {
		if l == "" {
			continue
		}
		space := ""
		if key != p.lastKey {
			if p.lastKey != "" {
				space = "\n"
			}
			p.lastKey = key
		}
		if p.stdout == nil {
			panic("nil stdout")
		}
		fmt.Fprintln(p.stdout, space+l)
	}
}
