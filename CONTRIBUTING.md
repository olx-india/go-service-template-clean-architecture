# Contributing to Go Service Template

Thank you for contributing. This project is an opinionated scaffold—keep changes focused and aligned with the lightweight template philosophy.

## Getting started

1. Fork the repository and clone your fork.
2. Install Go 1.24.6+ and [Make](https://www.gnu.org/software/make/).
3. Copy environment variables: `cp .env.example .env`
4. Install tooling: `make bin-deps`
5. Run the stack locally: `make compose-up` or `make run`

## Development workflow

1. Create a branch from `main` (e.g. `feature/my-change` or `fix/issue-123`).
2. Make your changes with tests where behavior changes.
3. Run checks before opening a PR:

```bash
make pre-commit
```

This runs formatting (`gofumpt`), linting (`golangci-lint`), and tests with the race detector.

## Commit messages

Use clear, imperative commit messages:

- `fix: correct FetchUser path parameter handling`
- `docs: add customizing guide`
- `chore: update CI workflow`

## Pull requests

- Keep PRs small and focused on one concern.
- Update [CHANGELOG.md](CHANGELOG.md) under `Unreleased` for user-visible changes.
- Ensure CI passes (lint, tests, govulncheck, Docker build).
- Link related issues when applicable.

## Scope guidelines

**In scope for this template:**

- Scaffold improvements (DI, logging, tracing hooks, health probes)
- Documentation accuracy and adopter guides
- CI, linting, and contributor experience

**Out of scope (add in your fork after using the template):**

- Full database implementations
- Auth, metrics, OpenAPI, deployment manifests

See [docs/CUSTOMIZING.md](docs/CUSTOMIZING.md) for how adopters extend the template.

## Code of conduct

This project follows the [Contributor Covenant](CODE_OF_CONDUCT.md). Report unacceptable behavior to the contacts listed there.

## Questions

Open a [GitHub issue](https://github.com/olx-india/go-service-template-clean-architecture/issues) for bugs, feature requests, or questions.
