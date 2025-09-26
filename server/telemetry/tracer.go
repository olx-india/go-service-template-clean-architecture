package telemetry

import (
	"context"
	"log"

	"go-service-template/internal/infrastructure/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// TracerBuilder holds the state during the fluent chain.
type TracerBuilder struct {
	ctx      context.Context
	cfg      config.Provider
	exporter *otlptrace.Exporter
	err      error
	resource *resource.Resource
	tp       *sdktrace.TracerProvider
}

func InitTracer(ctx context.Context, cfg config.Provider) func() {
	return (&TracerBuilder{}).
		createTelemetryExporter(ctx, cfg.GetOTLPEndpoint()).
		createApplicationResource(ctx, cfg).
		createTraceProvider().
		SetTracerProvider().
		setTextMapPropagator().
		shutdownFunction()
}

// createTelemetryExporter creates the OTLP trace exporter.
func (tb *TracerBuilder) createTelemetryExporter(ctx context.Context, endpoint string) *TracerBuilder {
	tb.ctx = ctx
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
	)
	if err != nil {
		log.Fatalf("failed to create trace exporter: %v", err)
	}
	tb.exporter = exp
	tb.err = err
	return tb
}

// createApplicationResource creates the application resource.
func (tb *TracerBuilder) createApplicationResource(ctx context.Context, cfg config.Provider) *TracerBuilder {
	if tb.existError() {
		return tb
	}

	tb.cfg = cfg
	res, err := createResource(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create resource: %v", err)
	}
	tb.resource = res
	return tb
}

// createTraceProvider creates the trace provider.
func (tb *TracerBuilder) createTraceProvider() *TracerBuilder {
	if tb.existError() {
		return tb
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(tb.exporter),
		sdktrace.WithResource(tb.resource),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	tb.tp = tp
	return tb
}

// SetTracerProvider sets the global tracer provider.
func (tb *TracerBuilder) SetTracerProvider() *TracerBuilder {
	if tb.existError() {
		return tb
	}

	otel.SetTracerProvider(tb.tp)
	return tb
}

// setTextMapPropagator sets the text map propagator.
func (tb *TracerBuilder) setTextMapPropagator() *TracerBuilder {
	if tb.existError() {
		return tb
	}
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{},
	))
	return tb
}

// shutdownFunction returns the shutdown function.
func (tb *TracerBuilder) shutdownFunction() func() {
	return func() {
		if err := tb.tp.Shutdown(tb.ctx); err != nil {
			log.Printf("error shutting down tracer provider: %v", err)
		}
	}
}

func createResource(ctx context.Context, cfg config.Provider) (*resource.Resource, error) {
	return resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(cfg.GetAppName()),
		semconv.DeploymentEnvironment(cfg.GetEnv()),
		semconv.ServiceVersion("1.0.0"),
	))
}

func (tb *TracerBuilder) existError() bool {
	return tb.err != nil
}
