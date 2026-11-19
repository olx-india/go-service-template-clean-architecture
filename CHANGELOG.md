# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Added

- Configurable CORS allowed origins via `CORS_ALLOWED_ORIGINS`; credentialed requests are enabled only for explicit origins
- Guides in `docs/`: getting started, architecture, configuration, API, observability, testing, deployment, adding a feature
- Social preview image in `docs/assets/`
- Documentation site on GitHub Pages (MkDocs Material), with a hosted Swagger UI, deployed by `.github/workflows/docs.yml`

### Changed

- README reorganized around a quick start and a documentation index
- `make run` exports `.env` when present
- Default `REDIS_HOST` is now `localhost:6379` (was `localhost`, which has no port and never connected)
- `.env.example` lists every setting

## 1.1.0 - 2026-11-17

### Added

- OpenAPI 3 generation with [go-swagger3](https://github.com/parvez3019/go-swagger3) and Swagger UI at `/swagger`
- CONTRIBUTING.md, GitHub issue/PR templates, and CODEOWNERS
- docs/CUSTOMIZING.md, docs/EXTENDING.md, and ADR for clean architecture layers
- `/live` and `/ready` probe endpoints
- Graceful HTTP shutdown and `gin.Recovery()` middleware
- Optional OpenTelemetry via `OTEL_ENABLED` (disabled by default)
- CI jobs for govulncheck, Docker build, and CodeQL
- `.editorconfig` and `.pre-commit-config.yaml`
- Makefile targets: `format`, `vuln`, `swagger`, fixed `pre-commit` and `mock`

### Changed

- Go module path is now `github.com/olx-india/go-service-template-clean-architecture`
- Go 1.26.11; upgraded golang.org/x/net, OpenTelemetry (v1.47.0), gRPC (v1.85.0) and Gin (v1.12.0) to clear govulncheck findings
- `make vuln` runs `go tool govulncheck`, so no separate install is needed
- README: badges, comparison with other Go templates, direct "Use this template" link
- README rewritten as an honest scaffold (Apache 2.0 license, badges, Mermaid diagram)
- SECURITY.md and CODE_OF_CONDUCT.md completed with private reporting contacts
- `FetchUser` handler reads user ID from path parameter (`GET /api/v1/user/:id`)
- Domain layer no longer imports API DTOs (`domain.NewUser` with `CreateUserInput`)
- Removed orphaned migrate Makefile targets and evrone/postgres stub dependency
- Unified naming to `go-service-template` in `.env.example`

### Removed

- Broken `mockgen` Makefile target paths and unused golang-migrate tool dependency

## 1.0.1 - 2025-09-26

### Added

- Add Open Telemetry instrumentation to the application

## 1.0.0 - 2025-09-16

### Added

- Setup initial project structure with Clean Architecture for Go services
