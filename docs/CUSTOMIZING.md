# Customizing this template

Use this guide after creating a repository from the **Use this template** button.

## 1. Rename the Go module

```bash
go mod edit -module github.com/your-org/your-service
```

Find and replace the old module path across the codebase:

```bash
# Example: replace go-service-template with your module path in imports
grep -r "go-service-template" --include="*.go" .
```

Update these non-Go references:

- `APP_NAME` in `.env` / `.env.example`
- Docker image name in `docker-compose.yml` and README
- Health check `service` field (via `APP_NAME`)
- GitHub URLs in README, CONTRIBUTING, and SECURITY

## 2. Update GitHub metadata

- Enable **Issues** and **Discussions** if desired
- Replace `.github/CODEOWNERS` with your team
- Update issue/PR template URLs if the repo name changed
- Mark the repository as a template again if you want others to fork yours

## 3. Strip sample features (optional)

Remove the example user and limit flows if you do not need them:

| Remove | Files / locations |
|--------|-------------------|
| User API | `internal/api/user.go`, `internal/usecase/user/`, routes in `server/router/router.go` |
| Limit API | `internal/api/limiter.go`, `internal/usecase/limit/`, routes in `server/router/router.go` |
| Sample domain | `internal/domain/user/` |
| Sample repos | `internal/infrastructure/repo/persistent/`, `webapi/` |
| Tests | Matching `*_test.go` and `integrationtests/user_test.go`, `limit_test.go` |
| Resolver wiring | `server/resolver/resolver.go` — remove handler/use case construction |

Keep the health endpoints, logging, config, and resolver pattern as your starting point.

## 4. Add your first feature

Follow the existing pattern:

1. **Domain** — define entities in `internal/domain/<feature>/`
2. **Repository interface** — add to `internal/infrastructure/repo/`
3. **Use case** — implement logic in `internal/usecase/<feature>/`
4. **DTO + handler** — add request/response types in `internal/api/dto/` and handler in `internal/api/`
5. **Wire** — register in `server/resolver/resolver.go`
6. **Route** — add in `server/router/router.go`
7. **Test** — unit tests per layer; integration test in `integrationtests/`

Look for `Adopter:` comments in the resolver and use cases for extension points.

## 5. Verify locally

```bash
make pre-commit
make compose-up
```

## 6. Database migrations (when you add Postgres)

This scaffold does not include a database. When you add one:

1. Add `golang-migrate` (or your tool of choice) to `go.mod`
2. Create a `migrations/` directory
3. Add Makefile targets for `migrate-up` / `migrate-down`
4. Wire the repository implementation in `internal/infrastructure/repo/persistent/`

See [EXTENDING.md](EXTENDING.md) for more optional additions.
