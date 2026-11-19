# Testing

The template ships unit tests for every layer and integration tests that run the real HTTP server. CI runs them on every pull request, along with lint and a vulnerability scan.

## Commands

| Command | What it does |
|---------|--------------|
| `make test` | All tests with the race detector; writes `coverage.txt` |
| `make lint` | `golangci-lint` with the rules in `.golangci.yml` |
| `make format` | Formats code with `gofumpt` |
| `make vuln` | `govulncheck` on reachable code |
| `make mock` | Regenerates interface mocks with Mockery |
| `make pre-commit` | `format`, `lint` and `test` together |

Install the pinned tool versions first with `make bin-deps`. They're declared as `tool` dependencies in `go.mod`, so everyone uses the same versions.

## Unit tests

Unit tests sit next to the code (`*_test.go`) and test one layer at a time, replacing the next layer with a mock:

| Layer | Example | Mocks |
|-------|---------|-------|
| Handler | `internal/api/user_test.go` | Use case (`usecase/user/mocks`) |
| Use case | `internal/usecase/user/user_test.go` | Use-case ports (`usecase/user/mocks`) |
| Domain | `internal/domain/user/user_test.go` | none |
| Config, logger, Redis | `internal/infrastructure/config/config_test.go` and siblings | Config provider |

Tests use [testify](https://github.com/stretchr/testify) for assertions and mocks. The generated mocks support the typed `EXPECT()` style as well as the classic `On(...)` style used in `internal/api/limiter_test.go`:

```go
useCase := mocks.NewIUserUseCase(t)
useCase.EXPECT().FetchUser(&user.FetchUserInput{ID: 1}).Return(&domain.User{ID: 1}, nil)
```

`NewIUserUseCase(t)` fails the test if an expected call never happens.

### Regenerating mocks

Mocks are generated for every interface in the module (`.mockery.yml`, `all: true`) into a `mocks/` folder next to the interface. After you add or change an interface, run:

```bash
make mock
```

and commit the generated files.

## Integration tests

`integrationtests/` starts the whole app in-process on `127.0.0.1:18080`, waits for `/health`, and calls the endpoints over HTTP with [baloo](https://github.com/h2non/baloo). Expected JSON responses are stored in `integrationtests/test_data/`.

They run as part of `make test`. Redis is optional: without it the app starts without a Redis client and the tests still pass. Port 18080 must be free.

To add one:

1. Add a `<feature>_test.go` in `integrationtests/` using the shared `TestClient`.
2. Put the expected response in `integrationtests/test_data/<feature>/`.
3. Load it with `LoadJSON` from `file_loader.go` and compare it with the response body.

## Coverage

`make test` writes `coverage.txt` (ignored by git). View it in a browser:

```bash
go tool cover -html=coverage.txt
```

CI uploads the same file to Codecov.

## Continuous integration

`.github/workflows/ci.yml` runs on every push to `main` and every pull request:

| Job | Runs |
|-----|------|
| lint | `golangci-lint`, at the version pinned in `go.mod` |
| test | `make test`, then uploads coverage |
| vuln | `make vuln`; fails on any reachable vulnerability |
| docker-build | Builds the production image |

CodeQL code scanning runs through GitHub's default setup in the repository settings.
