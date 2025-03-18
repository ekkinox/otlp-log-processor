package internal_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/ekkinox/otlp-log-processor/internal"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestOTel(t *testing.T) {
	t.Parallel()

	stop, err := internal.SetupOTel(context.Background())
	assert.NoError(t, err)
	defer stop()

	t.Run("test tracer", func(t *testing.T) {
		t.Parallel()

		tracer := internal.Tracer()

		assert.NotNil(t, tracer)
		assert.Implements(t, (*trace.Tracer)(nil), tracer)
	})

	t.Run("test logger", func(t *testing.T) {
		t.Parallel()

		logger := internal.Logger()

		assert.NotNil(t, logger)
		assert.IsType(t, &slog.Logger{}, logger)
	})
}
