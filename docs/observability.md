# Observability

The service writes structured JSON logs for every request and can export OpenTelemetry traces. Metrics are not included; [Extending](EXTENDING.md#metrics-prometheus) explains how to add Prometheus.

## Logs

Logging uses [Zap](https://github.com/uber-go/zap) through the `internal/infrastructure/logger` package. Output is one JSON object per line on stdout:

```json
{"severity":"info","timestamp":"2026-11-17T10:33:31.459+0530","caller":"logger/middleware.go:67","message":"Request completed","service":"go-service-template","request-id":"bbc0173c-1943-47b9-9627-b388d128af3a","trace-id":"72748af1-2e75-40dd-8a16-4f066aa4c69e","method":"GET","path":"/health","status_code":200,"duration":0,"response_size":47}
```

| Field | Meaning |
|-------|---------|
| `severity` | `debug`, `info`, `warn`, `error` |
| `timestamp` | ISO 8601 with offset |
| `caller` | Source file and line |
| `message` | What happened |
| `service` | Fixed to `go-service-template`; change `serviceName` in `logger/global.go` when you rename the service |
| `request-id`, `trace-id` | Correlation IDs for the request (see below) |
| `method`, `path`, `status_code`, `duration` (ms), `response_size` (bytes) | Added on "Request completed" |

Set the level with `LOG_LEVEL` (`debug`, `info`, `warn`, `error`; default `info`).

### Logging from your code

Use the package-level helpers with the request's log context, so request and trace IDs are attached:

```go
logCtx := logger.GetLogContext(ctx.Context)
logger.Info(logCtx, "Creating order", logger.Int(logger.FieldUserID, req.UserID))
logger.Error(logCtx, "Payment failed", logger.ErrorField(logger.FieldError, err))
```

Field helpers (`String`, `Int`, `Int64`, `ErrorField`, …) and field-name constants live in `internal/infrastructure/logger`. Prefer the constants over ad-hoc field names so logs stay queryable.

## Request and trace IDs

`LoggingMiddleware` runs for every request:

1. Reads `X-Request-ID` and `X-Trace-ID` from the request, generating UUIDs for any that are missing.
2. Returns both as response headers.
3. Stores them in the log context, so every log line for that request carries them.
4. Logs "Request started" and "Request completed".

Pass these headers on when you call other services so you can follow one request across logs.

> These IDs are independent of OpenTelemetry trace IDs. When tracing is on, use the trace view in Jaeger for timing and the log IDs for log search.

## Tracing

Tracing is off by default. Turn it on with:

```bash
OTEL_ENABLED=true
OTLP_ENDPOINT=localhost:4317   # OTLP gRPC collector
```

When on, `server/telemetry/tracer.go`:

- exports spans over OTLP gRPC (insecure connection) with a batching exporter
- samples **every** request (`AlwaysSample`)
- names the service from `APP_NAME` and tags `deployment.environment` from `ENV`
- propagates W3C `traceparent` and `baggage` headers
- flushes remaining spans on shutdown

`otelgin` middleware creates a span per HTTP request. If the exporter can't be created, the service logs an error and keeps running without tracing.

### Viewing traces locally

`make compose-up` starts Jaeger with tracing already on. Send a few requests, open http://localhost:16686, pick the `go-service-template` service and click **Find Traces**.

### Adding spans

Create child spans around slow or important work:

```go
tracer := otel.Tracer("internal/usecase/order")
ctx, span := tracer.Start(ctx, "ChargeCard")
defer span.End()
```

Pass `ctx` down to clients that support OpenTelemetry (HTTP, gRPC, database drivers) so their spans join the same trace.

### Before production

- Replace `AlwaysSample()` with a ratio or parent-based sampler; sampling everything is costly at high traffic.
- Use TLS to your collector instead of `WithInsecure()`.
- The resource reports `service.version` as `1.0.0`; set it from your build version.
