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
- Document required headers/tokens in your API docs

## Metrics (Prometheus)

- Add `github.com/prometheus/client_golang` and expose `GET /metrics`
- Register HTTP middleware for request duration and status counters
- Keep metrics registration in `server/` or `internal/infrastructure/metrics/`

## OpenAPI / Swagger

- Add `swaggo/swag` or `oapi-codegen` annotations on handlers
- Serve generated spec at `/swagger` or publish as a static file
- Generate clients only if your team needs them

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
