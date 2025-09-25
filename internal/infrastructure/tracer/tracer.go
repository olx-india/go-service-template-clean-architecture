package tracer

import (
	"context"
	"log"

	"go-service-template/internal/infrastructure/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// InitTracer initializes OpenTelemetry tracer with OTLP exporter
func InitTracer(ctx context.Context, cfg config.Provider) func() {
	endpoint := cfg.GetOTLPEndpoint()

	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
	)
	if err != nil {
		log.Fatalf("failed to create trace exporter: %v", err)
	}

	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(cfg.GetAppName()),
		semconv.DeploymentEnvironment(cfg.GetEnv()),
		semconv.ServiceVersion("1.0.0"),
	))
	if err != nil {
		log.Fatalf("failed to create resource: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{},
	))

	return func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("error shutting down tracer provider: %v", err)
		}
	}
}
