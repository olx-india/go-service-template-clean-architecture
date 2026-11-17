# Configuration

All configuration comes from environment variables, read once at startup by `internal/infrastructure/config`. The app does not load `.env` files itself; `make run` exports `.env` for you, and Docker Compose sets variables in `docker-compose.yml`.

## Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `HOST` | `0.0.0.0` | Interface the HTTP server binds to |
| `PORT` | `8080` | HTTP port |
| `READ_TIMEOUT` | `60s` | Maximum time to read a request, as a Go duration (`500ms`, `30s`, `2m`) |
| `WRITE_TIMEOUT` | `60s` | Maximum time to write a response |
| `APP_NAME` | `go-service-template` | Service name in `/health` responses and in traces |
| `ENV` | `local` | Environment name, attached to traces as `deployment.environment` |
| `LOG_LEVEL` | `info` | `debug`, `info`, `warn` or `error`; anything else means `info` |
| `REDIS_HOST` | `localhost:6379` | Redis address as `host:port` |
| `OTEL_ENABLED` | `false` | `true` or `1` turns on OpenTelemetry tracing |
| `OTLP_ENDPOINT` | `localhost:4317` | OTLP gRPC collector address, e.g. Jaeger |

Invalid values fall back to the default without an error: a `READ_TIMEOUT` of `abc` becomes `60s`. Validate stricter requirements at startup in your fork (see [Extending](EXTENDING.md#configuration-validation)).

## Example files

- [`.env.example`](https://github.com/olx-india/go-service-template-clean-architecture/blob/main/.env.example) — local development values; copy it to `.env`
- [`docker-compose.yml`](https://github.com/olx-india/go-service-template-clean-architecture/blob/main/docker-compose.yml) — values for the Compose stack (Redis and Jaeger by service name, tracing on)

## Redis

At startup the service pings `REDIS_HOST` with a 5-second timeout.

| Redis at startup | Result |
|------------------|--------|
| Reachable | The client is shared by use cases; `/ready` pings it on every call and returns 503 when it stops responding |
| Unreachable | An error is logged and the service runs **without** Redis. Use cases receive a nil provider and `/ready` returns 200 |

If your service can't work without Redis, make startup fail instead: return an error from `resolveProviders()` in `server/resolver/resolver.go` and stop in `cmd/main.go`.

## Adding a setting

1. Add an `Env…` name and `Default…` value constant in `internal/infrastructure/config/config.go`.
2. Read it in `NewConfig()` with `getEnv`, `getEnvAsDuration` or `getEnvAsBool`.
3. Add a getter to the `Provider` interface in `provider.go` and implement it on `*Config`.
4. Regenerate the config mock with `make mock` so tests can stub the new getter.
5. Document it in this table and in `.env.example`.

Depend on `config.Provider` rather than `*config.Config` so tests can pass a mock.

## Secrets

Don't commit secrets to `.env.example` or `docker-compose.yml`. Inject them at deploy time from your platform's secret store as environment variables.
