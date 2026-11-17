# Getting started

Run the service, call its endpoints and find your way around the code in about 10 minutes.

## Prerequisites

| Tool | Version | Needed for |
|------|---------|------------|
| Go | 1.26.11+ | Building and running the service |
| Make | any | The shortcuts in the `Makefile` |
| Docker + Docker Compose | recent | Redis and Jaeger, or the full stack |

## 1. Get the code

Create your own repository with **[Use this template](https://github.com/olx-india/go-service-template-clean-architecture/generate)**, then clone it. To try the template without creating a repository:

```bash
git clone https://github.com/olx-india/go-service-template-clean-architecture.git
cd go-service-template-clean-architecture
make deps        # go mod tidy + verify
make bin-deps    # install pinned dev tools into ./bin
```

## 2. Run it

### Option A: everything in Docker

```bash
make compose-up
```

This builds the app image and starts three containers:

| Service | URL | Purpose |
|---------|-----|---------|
| app | http://localhost:8080 | The service, with tracing turned on |
| redis | localhost:6379 | Backs the readiness probe and the sample use cases |
| jaeger | http://localhost:16686 | Trace viewer |

Stop everything with `make compose-down`.

### Option B: the app on your machine

Start Redis (optional, see the note below), copy the example config and run:

```bash
docker run -d --name redis -p 6379:6379 redis:8-alpine
cp .env.example .env
make run
```

`make run` exports the variables in `.env` before starting the app. If you run the binary another way (`go run`, your IDE), set the variables yourself — the app reads only real environment variables, it does not load `.env` files.

> **Redis is optional.** If Redis can't be reached at startup the service logs an error and keeps running without it. The sample endpoints still respond, and `/ready` reports OK because no Redis client was configured. See [Configuration → Redis](configuration.md#redis).

## 3. Try it

```bash
curl -s localhost:8080/health
# {"status":"ok","service":"go-service-template"}

curl -s -X POST localhost:8080/api/v1/user \
  -H 'Content-Type: application/json' \
  -d '{"id":1,"name":"Ada Lovelace","email":"ada@example.com","age":36}'

curl -s localhost:8080/api/v1/user/1

curl -s -X POST localhost:8080/api/v1/limit/check \
  -H 'Content-Type: application/json' -d '{"userID":1}'
```

The sample use cases are stubs, so they return empty or zero values. That's expected: they exist to show the wiring, not business logic.

Open **http://localhost:8080/swagger/index.html** to browse and call the API from Swagger UI. With Docker Compose, open **http://localhost:16686** and pick the `go-service-template` service to see a trace for each request.

## 4. Find your way around

```
cmd/main.go              Entry point: loads config, starts the app
server/
  app/                   HTTP server lifecycle and graceful shutdown
  resolver/              Composition root: builds every dependency
  router/                Middleware and route table
  telemetry/             OpenTelemetry tracer setup
internal/
  domain/                Entities and business rules
  usecase/               Application logic, one package per feature
  api/                   HTTP handlers; api/dto holds request/response types
  infrastructure/        Config, logging, Redis, repository implementations
integrationtests/        Starts the real app and calls it over HTTP
docs/                    These guides and the generated OpenAPI spec
```

[Architecture](architecture.md) explains how these fit together.

## 5. Check your setup

```bash
make test   # unit + integration tests with the race detector
make lint   # golangci-lint
make vuln   # govulncheck
```

## Common problems

| Symptom | Cause and fix |
|---------|---------------|
| `Failed to create redis provider … connection refused` in the logs | Redis isn't running. Start it (step 2) or ignore it if you don't need Redis. |
| `address already in use` | Something else is on port 8080. Run with `PORT=8081 make run`. |
| Settings in `.env` are ignored | You started the app without `make run`. Export the variables in your shell or IDE run configuration. |
| `make swagger` fails with `binary operator expected` | The repository path contains a space. Move the checkout to a path without spaces. |
| No traces in Jaeger | Tracing is off by default. Set `OTEL_ENABLED=true` and point `OTLP_ENDPOINT` at Jaeger's port 4317. |

## Next steps

- [Customizing](CUSTOMIZING.md): make the template your own service
- [Adding a feature](adding-a-feature.md): build your first endpoint
