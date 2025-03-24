package internal_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/ekkinox/otlp-log-processor/internal"
	"github.com/ekkinox/otlp-log-processor/internal/testdata"
	"github.com/stretchr/testify/assert"
	collecctorpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestServer(t *testing.T) {
	lis := bufconn.Listen(1024)

	conn, err := grpc.NewClient(
		fmt.Sprintf("passthrough://%s", lis.Addr().String()),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)

	str := internal.NewStorage()
	svc := internal.NewService(str, "foo", 10)
	srv := internal.NewServer(svc)
	defer srv.GracefulStop()

	go func() {
		//nolint:errcheck
		srv.Serve(lis)
	}()

	cli := collecctorpb.NewLogsServiceClient(conn)

	resp, err := cli.Export(context.Background(), testdata.TestReq)
	assert.NoError(t, err)

	assert.Equal(t, "<nil>", resp.PartialSuccess.String())
	assert.Equal(
		t,
		map[string]int64{
			"bool_value:true":      2,
			"double_value:1.23":    2,
			"string_value:\"bar\"": 4,
			"string_value:\"baz\"": 2,
		},
		str.Dump(),
	)
}
