package internal_test

import (
	"testing"

	"github.com/ekkinox/otlp-log-processor/internal"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	t.Run("test config defaults", func(t *testing.T) {
		t.Parallel()

		cfg := internal.NewConfig()

		assert.Equal(t, "foo", cfg.Attribute())
		assert.Equal(t, 1000, cfg.Interval())
		assert.Equal(t, 10, cfg.Workers())
	})

	t.Run("test config with flags", func(t *testing.T) {
		t.Parallel()

		// todo
	})
}
