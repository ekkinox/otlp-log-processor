package internal

import (
	"context"
	"fmt"
	"io"
	"time"
)

type Ticker struct {
	writer    io.Writer
	storage   *Storage
	attribute string
	interval  int
}

func NewTicker(writer io.Writer, storage *Storage, attribute string, interval int) *Ticker {
	return &Ticker{
		writer:    writer,
		storage:   storage,
		attribute: attribute,
		interval:  interval,
	}
}

func (t *Ticker) Start(ctx context.Context) {
	tick := time.NewTicker(time.Duration(t.interval) * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			t.displayStorage()
		case <-ctx.Done():
			fmt.Fprintln(t.writer, "stopping ticker...")
			return
		}
	}
}

func (t *Ticker) displayStorage() {
	fmt.Fprintf(t.writer, "%s occurrences:\n", t.attribute)

	for k, v := range t.storage.Dump() {
		fmt.Fprintf(t.writer, "\t%s: %d\n", k, v)
	}

	fmt.Fprintln(t.writer, "")
}
