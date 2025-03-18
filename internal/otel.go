package internal

import (
	"context"
	"log"
	"log/slog"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	otellog "go.opentelemetry.io/otel/sdk/log"
	otelsdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const name = "github.com/ekkinox/otlp-log-processor"

func Tracer(opts ...trace.TracerOption) trace.Tracer {
	return otel.Tracer(name, opts...)
}

func Logger() *slog.Logger {
	return otelslog.NewLogger(name)
}

func SetupOTel(ctx context.Context) (context.CancelFunc, error) {
	tmp := newTextMapPropagator()
	otel.SetTextMapPropagator(tmp)

	tp, err := newTracerProvider()
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tp)

	lp, err := newLoggerProvider()
	if err != nil {
		return nil, err
	}
	global.SetLoggerProvider(lp)

	return func() {
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = lp.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}, nil
}

func newTextMapPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider() (*otelsdktrace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tracerProvider := otelsdktrace.NewTracerProvider(otelsdktrace.WithBatcher(traceExporter, otelsdktrace.WithBatchTimeout(time.Second)))

	return tracerProvider, nil
}

func newLoggerProvider() (*otellog.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := otellog.NewLoggerProvider(otellog.WithProcessor(otellog.NewBatchProcessor(logExporter)))

	return loggerProvider, nil
}
