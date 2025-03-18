package internal

import (
	"context"
	"fmt"
	"time"
)

type Ticker struct {
	storage   *Storage
	attribute string
	interval  int
}

func NewTicker(storage *Storage, attribute string, interval int) *Ticker {
	return &Ticker{
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
			fmt.Println("stopping ticker...")
			return
		}
	}
}

func (t *Ticker) displayStorage() {
	fmt.Printf("attribute %s values:\n", t.attribute)

	for k, v := range t.storage.Dump() {
		fmt.Println("\t", k, ":", v)
	}

	fmt.Println("")
}
