# Extending the template

This scaffold intentionally omits several production concerns. Add them in your fork when needed.

## Database and migrations

- Add Postgres (or your store) to `docker-compose.yml`
- Implement repository interfaces under `internal/infrastructure/repo/`
- Use [golang-migrate](https://github.com/golang-migrate/migrate) or your preferred tool
- Extend `/ready` to check database connectivity

## Authentication and authorization

- Add middleware in `server/router/router.go`
- Keep auth logic out of handlers—validate in use cases or a dedicated auth package
- Document required headers/tokens with go-swagger3 `@SecurityScheme` / `@Security` on `cmd/main.go` and regenerate with `make swagger`

## Metrics (Prometheus)

- Add `github.com/prometheus/client_golang` and expose `GET /metrics`
- Register HTTP middleware for request duration and status counters
- Keep metrics registration in `server/` or `internal/infrastructure/metrics/`

## OpenAPI / Swagger

OpenAPI 3 is already wired via [go-swagger3](https://github.com/parvez3019/go-swagger3):

1. Annotate handlers under `internal/api/` (`@Title`, `@Router`, `@Param`, `@Success`, …)
2. Add `json` / OAS tags on DTOs in `internal/api/dto/`
3. Run `make swagger` to refresh `docs/openapi/oas.json`
4. Browse the UI at `http://localhost:8080/swagger/index.html`

Service-level metadata lives on `cmd/main.go`. Commit the regenerated `oas.json` so Docker/CI builds embed the latest spec without running the CLI.

## Configuration validation

- Validate required env vars in `config.NewConfig()` or at startup in `cmd/main.go`
- Fail fast with clear error messages for missing secrets/URLs

## Deployment

- Add Kubernetes manifests, Helm, or Terraform in your service repository—not in this template
- Use `/live` and `/ready` for probe endpoints
- Run as non-root in production Docker images

## Security hardening

- Restrict CORS origins (currently `*` for local dev)
- Add request size limits and security headers middleware
- Run `govulncheck` locally (`make vuln`) and enable GitHub Code scanning (default setup) in repository settings
