package internal_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/ekkinox/otlp-log-processor/internal"
	"github.com/stretchr/testify/assert"
	collecctorpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logpb "go.opentelemetry.io/proto/otlp/logs/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
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

	resp, err := cli.Export(context.Background(), &collecctorpb.ExportLogsServiceRequest{
		ResourceLogs: []*logpb.ResourceLogs{
			{
				Resource: &resourcepb.Resource{
					Attributes: make([]*commonpb.KeyValue, 0),
				},
				ScopeLogs: []*logpb.ScopeLogs{
					{
						Scope: &commonpb.InstrumentationScope{
							Attributes: make([]*commonpb.KeyValue, 0),
						},
						LogRecords: []*logpb.LogRecord{
							{
								Attributes: []*commonpb.KeyValue{
									{
										Key: "foo",
										Value: &commonpb.AnyValue{
											Value: &commonpb.AnyValue_StringValue{StringValue: "bar"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	assert.NoError(t, err)

	assert.Equal(t, "<nil>", resp.PartialSuccess.String())
	assert.Equal(
		t,
		map[string]int64{
			"string_value:\"bar\"": 1,
		},
		str.Dump(),
	)
}
