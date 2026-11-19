# Deployment

The template builds a small, static Docker image and exposes the probes an orchestrator needs. It doesn't include deployment manifests; add Kubernetes, Helm or Terraform files in your own service repository.

## Docker image

The `Dockerfile` has two stages:

1. **Build** — `golang:1.26.11-alpine3.24` downloads modules and builds a static binary (`CGO_ENABLED=0`).
2. **Run** — `alpine:latest` with CA certificates and the binary, exposing port 8080.

```bash
docker build -t my-service .
docker run --rm -p 8080:8080 -e REDIS_HOST=redis.internal:6379 my-service
```

`IMAGE_PREFIX` is a build argument for pulling the base image through a private registry mirror:

```bash
docker build --build-arg IMAGE_PREFIX=registry.example.com/mirror/ -t my-service .
```

The OpenAPI spec is embedded in the binary, so the image needs no extra files.

## Health probes

| Probe | Endpoint | Fails when |
|-------|----------|------------|
| Liveness | `GET /live` | Never while the process serves HTTP |
| Readiness | `GET /ready` | A configured Redis client stops answering pings (503) |

Kubernetes example:

```yaml
livenessProbe:
  httpGet: { path: /live, port: 8080 }
  periodSeconds: 10
readinessProbe:
  httpGet: { path: /ready, port: 8080 }
  periodSeconds: 5
  failureThreshold: 3
terminationGracePeriodSeconds: 40
```

Add your own dependencies (database, downstream services) to the readiness check in `internal/api/health.go`.

## Graceful shutdown

On `SIGTERM` or `SIGINT` the service stops accepting connections, gives in-flight requests up to **30 seconds** to finish, closes the Redis client and flushes traces. Set your orchestrator's grace period above 30 seconds (40 in the example above) so the process isn't killed mid-shutdown.

## Configuration

Pass settings as environment variables ([Configuration](configuration.md)). A typical production set:

```bash
ENV=production
APP_NAME=orders-service
LOG_LEVEL=info
REDIS_HOST=redis.internal:6379
OTEL_ENABLED=true
OTLP_ENDPOINT=otel-collector.observability:4317
```

## Production checklist

The template favours a quick local start. Review these before you ship:

- [ ] **Run as non-root.** The runtime image runs as root. Add a user and `USER` it, or switch to a distroless base image.
- [ ] **Pin the runtime image.** Replace `alpine:latest` with a specific version or digest.
- [ ] **Restrict CORS.** Set `CORS_ALLOWED_ORIGINS` to your front-end origins. The default `*` allows any origin without credentials; an explicit list enables credentialed requests only for matching origins.
- [ ] **Hide internal errors.** Error responses include raw Go error text ([API → Errors](api.md#errors)). Return stable, client-safe messages and log the details.
- [ ] **Decide what happens without Redis.** By default the service starts without it and `/ready` passes ([Configuration → Redis](configuration.md#redis)).
- [ ] **Tune tracing.** Change the sampler from `AlwaysSample` and use TLS to the collector ([Observability](observability.md#before-production)).
- [ ] **Set timeouts** (`READ_TIMEOUT`, `WRITE_TIMEOUT`) to match your SLOs; the 60-second defaults are generous.
- [ ] **Add authentication, request size limits and security headers** ([Extending](EXTENDING.md)).
- [ ] **Keep dependencies patched.** Dependabot and the `vuln` CI job flag known vulnerabilities.
