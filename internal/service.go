package internal

import (
	"context"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	pb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
)

type LogServiceServer struct {
	pb.UnimplementedLogsServiceServer
}

func (s *LogServiceServer) Export(ctx context.Context, req *pb.ExportLogsServiceRequest) (*pb.ExportLogsServiceResponse, error) {
	ctx, span := otel.Tracer(Name).Start(ctx, "LogServiceServer::Export")
	defer span.End()

	logger := otelslog.NewLogger(Name)

	logger.InfoContext(ctx, "LogServiceServer::Export")

	for _, resourceLogs := range req.ResourceLogs {
		for _, scopeLogs := range resourceLogs.ScopeLogs {
			for _, logRecord := range scopeLogs.LogRecords {
				logger.InfoContext(ctx, logRecord.Attributes[0].Key)
			}
		}
	}

	return &pb.ExportLogsServiceResponse{}, nil
}
