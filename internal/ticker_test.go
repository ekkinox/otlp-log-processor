package internal_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/ekkinox/otlp-log-processor/internal"
	"github.com/stretchr/testify/assert"
)

func TestTicker(t *testing.T) {
	t.Parallel()

	str := internal.NewStorage()
	str.Store("foo")
	str.Store("bar")
	str.Store("foo")

	var buf bytes.Buffer

	ticker := internal.NewTicker(&buf, str, "attr", 1)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()

	go ticker.Start(ctx)

	<-ctx.Done()

	assert.Contains(t, buf.String(), "foo: 2")
	assert.Contains(t, buf.String(), "bar: 1")
}
