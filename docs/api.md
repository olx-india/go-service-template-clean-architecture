# API

The service exposes health probes, two sample feature APIs (users and limits) and Swagger UI. The sample use cases are stubs: they validate input and return placeholder values until you implement them.

The full machine-readable contract is the OpenAPI 3 spec in [`openapi/oas.json`](openapi/oas.json), served at **`/swagger/index.html`**.

## Endpoints

| Method | Path | Purpose | Success |
|--------|------|---------|---------|
| GET | `/live` | Liveness probe: the process is up | 200 |
| GET | `/ready` | Readiness probe: dependencies are reachable | 200, or 503 |
| GET | `/health` | Legacy health check, same response as `/live` | 200 |
| POST | `/api/v1/user` | Create a user | 201 |
| GET | `/api/v1/user/{id}` | Fetch a user by numeric ID | 200 |
| POST | `/api/v1/limit/check` | Check a user's remaining rate limit | 200 |
| POST | `/api/v1/limit/reset` | Reset a user's rate limit | 200 |
| GET | `/swagger/*` | Swagger UI and the OpenAPI document | 200 |

## Health probes

```http
GET /ready
```
```json
{"status": "ok", "service": "go-service-template"}
```

When a Redis client is configured and its ping fails (5-second timeout), `/ready` returns **503**:

```json
{"status": "not_ready", "service": "go-service-template", "error": "redis unavailable"}
```

`service` is the `APP_NAME` setting. Use `/live` for liveness and `/ready` for readiness probes ([Deployment](deployment.md#health-probes)).

## Users

### Create a user

```http
POST /api/v1/user
Content-Type: application/json

{"id": 1, "name": "Ada Lovelace", "email": "ada@example.com", "age": 36}
```

| Field | Type | Rules |
|-------|------|-------|
| `id` | integer | required, non-zero |
| `name` | string | required |
| `email` | string | required, valid email |
| `age` | integer | required, 0–130 |

Response `201 Created` with a `UserResponse` (`id`, `name`, `email`, `age`). The stub returns zero values: `{"id":0,"name":"","email":"","age":0}`.

### Fetch a user

```http
GET /api/v1/user/1
```

Response `200 OK` with a `UserResponse`. A non-numeric ID returns 400.

## Limits

Both endpoints take the same body:

```http
POST /api/v1/limit/check
Content-Type: application/json

{"userID": 1}
```

Response `200 OK`:

```json
{"userID": 1, "limitAvailable": 0}
```

`userID` is required and non-zero. The stub always reports `limitAvailable: 0`; implement the Redis logic in `internal/usecase/limit/`.

## Errors

Every error uses the same shape, `dto.ErrorResponse`:

```json
{"error": "Invalid request body: Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"}
```

| Status | When |
|--------|------|
| 400 | Malformed JSON, failed validation, or a bad path parameter |
| 500 | The use case returned an error |
| 503 | `/ready` only: a dependency is down |

The message includes the underlying Go error, which helps during development but exposes internals. Map errors to stable, client-safe messages before going to production ([Deployment](deployment.md#production-checklist)).

Validation uses Gin's `binding` tags on the DTOs in `internal/api/dto/`, backed by [go-playground/validator](https://github.com/go-playground/validator).

## Headers

| Header | Direction | Behaviour |
|--------|-----------|-----------|
| `X-Request-ID` | request and response | Echoed back if sent; otherwise a UUID is generated. Included in every log line for the request. |
| `X-Trace-ID` | request and response | Same as above, for correlating calls across services |
| `Access-Control-Allow-*` | response | CORS allows the configured `CORS_ALLOWED_ORIGINS`; wildcard mode sends no credentials. `OPTIONS` preflight requests get 204. |

## OpenAPI and Swagger UI

The spec is generated from comments on the handlers by [go-swagger3](https://github.com/parvez3019/go-swagger3) and embedded into the binary (`docs/openapi/doc.go`), so Swagger UI works without extra files at runtime.

To change the API documentation:

1. Edit the annotations above a handler in `internal/api/`, for example:
   ```go
   // @Title Create user
   // @Param request body dto.CreateUserRequest true "Create user payload"
   // @Success 201 {object} dto.UserResponse "User created"
   // @Failure 400 {object} dto.ErrorResponse "Invalid request"
   // @Resource users
   // @Router /api/v1/user [post]
   ```
2. Add `example`, `description`, `minimum` / `maximum` tags to DTO fields in `internal/api/dto/`.
3. Run `make swagger` to regenerate `docs/openapi/oas.json`.
4. Commit the regenerated file; Docker and CI builds embed it as is.

Service-wide metadata (title, version, license, servers) lives in the comment above `main()` in `cmd/main.go`.
