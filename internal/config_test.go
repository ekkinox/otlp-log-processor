package internal_test

import (
	"testing"

	"github.com/ekkinox/otlp-log-processor/internal"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg := internal.NewConfig()

	assert.Equal(t, "foo", cfg.Attribute())
	assert.Equal(t, 1000, cfg.Interval())
	assert.Equal(t, 10, cfg.Workers())
}
