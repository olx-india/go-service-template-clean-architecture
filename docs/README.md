# Documentation

Guides for running, understanding and building on the Go Service Template.

These pages are also published at **[https://olx-india.github.io/go-service-template-clean-architecture/](https://olx-india.github.io/go-service-template-clean-architecture/)**, built from this folder by `.github/workflows/docs.yml`. Preview locally with `pip install -r docs/requirements.txt && mkdocs serve`.

## Start here

| Guide | Read it when you want to… |
|-------|---------------------------|
| [Getting started](getting-started.md) | Run the service locally or with Docker Compose and send your first requests |
| [Customizing](CUSTOMIZING.md) | Turn the template into your own service: rename the module, strip the samples |
| [Adding a feature](adding-a-feature.md) | Build a new endpoint end to end, layer by layer |

## Reference

| Guide | Covers |
|-------|--------|
| [Architecture](architecture.md) | Layers, the dependency rule, request lifecycle, the resolver |
| [Configuration](configuration.md) | Every environment variable, its default and what it does |
| [API](api.md) | Endpoints, request and response shapes, errors, headers, OpenAPI |
| [Observability](observability.md) | Structured logs, request and trace IDs, OpenTelemetry tracing |
| [Testing](testing.md) | Unit and integration tests, mocks, coverage, lint and vulnerability checks |
| [Deployment](deployment.md) | Docker image, health probes, graceful shutdown, production checklist |
| [Extending](EXTENDING.md) | Optional additions: database, auth, metrics, hardening |

## Decisions

- [ADR 001: Clean Architecture layers](adr/001-clean-architecture-layers.md)

## Generated

- [`openapi/oas.json`](openapi/oas.json) — OpenAPI 3 spec, served by the app at `/swagger/index.html`
