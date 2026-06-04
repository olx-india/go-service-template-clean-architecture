# ADR 001: Clean Architecture layers

## Status

Accepted

## Context

Go microservices benefit from a consistent structure that separates business rules from delivery mechanisms and infrastructure. This template targets teams who want a **lightweight starting point** they can customize—not a batteries-included framework.

## Decision

Organize code into concentric layers with dependency inversion:

1. **Domain** — pure business entities; no imports from API, HTTP, or database packages
2. **Use case** — orchestrates domain logic; depends on repository interfaces
3. **Infrastructure** — implements repositories, config, logging, Redis
4. **API** — HTTP handlers and DTOs; maps transport types to use case inputs
5. **Server** — composition root (resolver), routing, app lifecycle, telemetry

Dependency flow points inward: outer layers depend on inner abstractions, not the reverse.

## Included technology choices

| Choice | Rationale |
|--------|-----------|
| Gin | Widely used, fast HTTP router with middleware support |
| Zap | Structured JSON logging with minimal overhead |
| Redis client | Common cache/session/rate-limit backend; optional at runtime |
| OpenTelemetry | Optional tracing via OTLP; disabled by default for simple local runs |
| Mockery | Interface mocks for unit testing across layers |

## Intentionally omitted

- Database drivers and migrations
- Authentication / authorization
- Prometheus metrics and OpenAPI generation
- Kubernetes / Helm / Terraform

These are documented in [EXTENDING.md](../EXTENDING.md).

## Adding a new HTTP feature

1. Define domain types in `internal/domain/<name>/`
2. Add repository interface in `internal/infrastructure/repo/`
3. Implement use case in `internal/usecase/<name>/`
4. Add DTOs and handler in `internal/api/`
5. Wire dependencies in `server/resolver/resolver.go`
6. Register route in `server/router/router.go`
7. Add unit and integration tests

## Consequences

- Adopters must implement persistence and business rules themselves
- The template stays small and easy to understand
- Correct dependency direction is enforced by example (DTO mapping stays in handlers)
