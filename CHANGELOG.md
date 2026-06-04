# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Added

- CONTRIBUTING.md, GitHub issue/PR templates, and CODEOWNERS
- docs/CUSTOMIZING.md, docs/EXTENDING.md, and ADR for clean architecture layers
- `/live` and `/ready` probe endpoints
- Graceful HTTP shutdown and `gin.Recovery()` middleware
- Optional OpenTelemetry via `OTEL_ENABLED` (disabled by default)
- CI jobs for govulncheck, Docker build, and CodeQL
- `.editorconfig` and `.pre-commit-config.yaml`
- Makefile targets: `format`, `vuln`, fixed `pre-commit` and `mock`

### Changed

- README rewritten as an honest scaffold (Apache 2.0 license, badges, Mermaid diagram)
- SECURITY.md and CODE_OF_CONDUCT.md completed with private reporting contacts
- `FetchUser` handler reads user ID from path parameter (`GET /api/v1/user/:id`)
- Domain layer no longer imports API DTOs (`domain.NewUser` with `CreateUserInput`)
- Removed orphaned migrate Makefile targets and evrone/postgres stub dependency
- Unified naming to `go-service-template` in `.env.example`

### Removed

- Broken `mockgen` Makefile target paths and unused golang-migrate tool dependency

## 1.0.1 - 2026-09-26

### Added

- Add Open Telemetry instrumentation to the application

## 1.0.0 - 2026-09-16

### Added

- Setup initial project structure with Clean Architecture for Go services
