package internal

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	collecctorpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logpb "go.opentelemetry.io/proto/otlp/logs/v1"
)

type Service struct {
	collecctorpb.UnimplementedLogsServiceServer
	storage   *Storage
	attribute string
	workers   int
}

func NewService(storage *Storage, attribute string, workers int) *Service {
	return &Service{
		storage:   storage,
		attribute: attribute,
		workers:   workers,
	}
}

func (s *Service) Export(ctx context.Context, req *collecctorpb.ExportLogsServiceRequest) (*collecctorpb.ExportLogsServiceResponse, error) {
	resourceLogsLen := len(req.ResourceLogs)

	ctx, span := Tracer().Start(ctx, "Service::Export")
	span.SetAttributes(attribute.Int("resource.logs", resourceLogsLen))
	defer span.End()

	Logger().InfoContext(ctx, "Service::Export")

	jobs := make(chan *logpb.ResourceLogs, resourceLogsLen)
	var wg sync.WaitGroup

	for i := 0; i < s.workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for resourceLogs := range jobs {
				s.processResourceLogs(ctx, resourceLogs)
			}
		}()
	}

	for _, resourceLogs := range req.ResourceLogs {
		jobs <- resourceLogs
	}
	close(jobs)

	wg.Wait()

	return &collecctorpb.ExportLogsServiceResponse{}, nil
}

func (s *Service) processResourceLogs(ctx context.Context, resourceLogs *logpb.ResourceLogs) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("Service::processResourceLogs")

	s.processAttributes(resourceLogs.Resource.Attributes)

	var wg sync.WaitGroup
	for _, scopeLogs := range resourceLogs.ScopeLogs {
		wg.Add(1)
		go func(scopeLogs *logpb.ScopeLogs) {
			defer wg.Done()

			s.processScopeLogs(ctx, scopeLogs)
		}(scopeLogs)
	}

	wg.Wait()
}

func (s *Service) processScopeLogs(ctx context.Context, scopeLogs *logpb.ScopeLogs) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("Service::processScopeLogs")

	s.processAttributes(scopeLogs.Scope.Attributes)

	var wg sync.WaitGroup
	for _, logRecord := range scopeLogs.LogRecords {
		wg.Add(1)
		go func(logRecord *logpb.LogRecord) {
			defer wg.Done()

			s.processLogRecord(ctx, logRecord)
		}(logRecord)
	}

	wg.Wait()
}

func (s *Service) processLogRecord(ctx context.Context, logRecord *logpb.LogRecord) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("Service::processLogRecord")

	s.processAttributes(logRecord.Attributes)
}

func (s *Service) processAttributes(attrs []*commonpb.KeyValue) {
	m := make(map[string]string, len(attrs))

	for _, attr := range attrs {
		m[attr.Key] = attr.Value.String()
	}

	if v, ok := m[s.attribute]; ok {
		s.storage.Store(v)
	}
}
