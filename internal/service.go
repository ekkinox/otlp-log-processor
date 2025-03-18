package internal

import (
	"context"
	"sync"

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
	jobs := make(chan *logpb.ResourceLogs, len(req.ResourceLogs))
	var wg sync.WaitGroup

	for i := 0; i < s.workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for resourceLogs := range jobs {
				s.processResourceLogs(resourceLogs)
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

func (s *Service) processResourceLogs(resourceLogs *logpb.ResourceLogs) {
	s.processAttributes(resourceLogs.Resource.Attributes)

	var wg sync.WaitGroup
	for _, scopeLogs := range resourceLogs.ScopeLogs {
		wg.Add(1)
		go func(scopeLogs *logpb.ScopeLogs) {
			defer wg.Done()

			s.processScopeLogs(scopeLogs)
		}(scopeLogs)
	}

	wg.Wait()
}

func (s *Service) processScopeLogs(scopeLogs *logpb.ScopeLogs) {
	s.processAttributes(scopeLogs.Scope.Attributes)

	var wg sync.WaitGroup
	for _, logRecord := range scopeLogs.LogRecords {
		wg.Add(1)
		go func(logRecord *logpb.LogRecord) {
			defer wg.Done()

			s.processLogRecord(logRecord)
		}(logRecord)
	}

	wg.Wait()
}

func (s *Service) processLogRecord(logRecord *logpb.LogRecord) {
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
