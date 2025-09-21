package main

import (
	"context"
	"log"
	"os"

	"go-service-template/internal/infrastructure/config"
	"go-service-template/internal/infrastructure/logger"
	appPkg "go-service-template/server/app"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func initTracer(ctx context.Context, serviceName string, env string) func() {
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:4317"
	}

	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
	)
	if err != nil {
		log.Fatalf("failed to create trace exporter: %v", err)
	}

	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(serviceName),
		semconv.DeploymentEnvironment(env),
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

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	logger.InitGlobalLogger()
	logger.Info(ctx, "Starting application", LogAppName(cfg), LogEnv(cfg))

	shutdown := initTracer(ctx, cfg.GetAppName(), cfg.GetEnv())
	defer shutdown()

	app := appPkg.NewApp()
	app.Start()
}

func LogEnv(cfg *config.Config) logger.Field {
	return logger.String("environment", cfg.GetEnv())
}

func LogAppName(cfg *config.Config) logger.Field {
	return logger.String("app_name", cfg.GetAppName())
}
