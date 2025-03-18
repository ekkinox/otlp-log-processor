package internal

import (
	"context"
	"fmt"

	pb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
)

type LogServiceServer struct {
	pb.UnimplementedLogsServiceServer
}

func (s *LogServiceServer) Export(ctx context.Context, req *pb.ExportLogsServiceRequest) (*pb.ExportLogsServiceResponse, error) {
	for _, resourceLogs := range req.ResourceLogs {
		for _, scopeLogs := range resourceLogs.ScopeLogs {
			for _, logRecord := range scopeLogs.LogRecords {
				fmt.Println(logRecord)
			}
		}
	}

	return &pb.ExportLogsServiceResponse{}, nil
}
