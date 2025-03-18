package internal

import (
	pb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer(svc *Service, opts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opts...)

	pb.RegisterLogsServiceServer(srv, svc)
	reflection.Register(srv)

	return srv
}
