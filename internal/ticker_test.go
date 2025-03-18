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
	str := internal.NewStorage()
	str.Store("foo")
	str.Store("bar")
	str.Store("foo")

	var buf bytes.Buffer

	ticker := internal.NewTicker(&buf, str, "attr", 1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	ticker.Start(ctx)

	<-ctx.Done()

	bufStr := buf.String()

	assert.Contains(t, bufStr, "foo: 2")
	assert.Contains(t, bufStr, "bar: 1")
}
