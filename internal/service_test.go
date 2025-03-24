package internal_test

import (
	"context"
	"testing"

	"github.com/ekkinox/otlp-log-processor/internal"
	"github.com/ekkinox/otlp-log-processor/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func BenchmarkService(b *testing.B) {
	str := internal.NewStorage()
	svc := internal.NewService(str, "foo", 10)

	for i := 0; i < b.N; i++ {
		_, err := svc.Export(context.Background(), testdata.TestReq)
		assert.NoError(b, err)
	}
}
