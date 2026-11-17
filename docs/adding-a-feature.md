# Adding a feature

This walkthrough adds `GET /api/v1/orders/{id}`, touching each layer once. Work inside out: domain first, HTTP last. Replace "order" with your own feature.

```
internal/domain/order/order.go           1. entity
internal/infrastructure/repo/repo.go     2. repository interface
internal/infrastructure/repo/persistent/ 2. implementation
internal/usecase/order/                  3. use case
internal/api/dto/order.go                4. request/response types
internal/api/order.go                    4. handler
server/resolver/resolver.go              5. wiring
server/router/router.go                  6. route
```

## 1. Domain

```go
// internal/domain/order/order.go
package order

import "errors"

var ErrNotFound = errors.New("order not found")

type Order struct {
    ID         int
    UserID     int
    TotalCents int64
}

func (o Order) IsFree() bool { return o.TotalCents == 0 }
```

Keep this package free of HTTP, JSON tags and database code. Rules about an order belong here.

## 2. Repository interface and implementation

Declare what the use case needs, not what the database offers:

```go
// internal/infrastructure/repo/repo.go
type OrderRepo interface {
    Fetch(ctx context.Context, id int) (order.Order, error)
}
```

Implement it in a subpackage. Start with a stub, then replace it with real storage:

```go
// internal/infrastructure/repo/persistent/order.go
type orderRepo struct{ /* db *sql.DB */ }

func NewOrderRepo() repo.OrderRepo { return &orderRepo{} }

func (r *orderRepo) Fetch(_ context.Context, id int) (order.Order, error) {
    return order.Order{ID: id}, nil
}
```

## 3. Use case

Give the use case its own input type so it doesn't depend on the API layer:

```go
// internal/usecase/order/order.go
package order

type GetOrderInput struct{ ID int }

type IOrderUseCase interface {
    GetOrder(ctx context.Context, in GetOrderInput) (domain.Order, error)
}

type UseCase struct{ orders repo.OrderRepo }

func NewOrderUseCase(orders repo.OrderRepo) *UseCase { return &UseCase{orders: orders} }

func (u *UseCase) GetOrder(ctx context.Context, in GetOrderInput) (domain.Order, error) {
    return u.orders.Fetch(ctx, in.ID)
}
```

(`domain` is `internal/domain/order` imported under an alias, as the user use case does.)

## 4. DTO and handler

```go
// internal/api/dto/order.go
type OrderResponse struct {
    ID         int   `json:"id" example:"42" description:"Order identity"`
    UserID     int   `json:"userID" example:"1" description:"Owner"`
    TotalCents int64 `json:"totalCents" example:"1999" description:"Total in cents"`
}
```

```go
// internal/api/order.go

// GetOrder returns one order.
// @Title Get order
// @Produce json
// @Param id path int true "Order identity"
// @Success 200 {object} dto.OrderResponse "Order found"
// @Failure 400 {object} dto.ErrorResponse "Invalid order ID"
// @Failure 404 {object} dto.ErrorResponse "Order not found"
// @Resource orders
// @Router /api/v1/orders/{id} [get]
func (h *orderHandler) GetOrder(ctx *ginContext.GinContext) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid order ID"})
        return
    }

    o, err := h.orders.GetOrder(ctx.Request().Context(), order.GetOrderInput{ID: id})
    switch {
    case errors.Is(err, domain.ErrNotFound):
        ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "order not found"})
        return
    case err != nil:
        logger.Error(logger.GetLogContext(ctx.Context), "get order failed", logger.ErrorField(logger.FieldError, err))
        ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal error"})
        return
    }

    ctx.JSON(http.StatusOK, dto.OrderResponse{ID: o.ID, UserID: o.UserID, TotalCents: o.TotalCents})
}
```

The handler maps between HTTP and the use case: parse and validate input, call the use case, map domain errors to status codes and the domain result to a DTO.

## 5. Wire it

```go
// server/resolver/resolver.go
type ServerContext struct {
    // ...
    OrderHandler api.IOrderHandler
}

func (r *resolver) createServerContext() *resolver {
    // ...
    r.OrderHandler = api.NewOrderHandler(order.NewOrderUseCase(persistent.NewOrderRepo()))
    return r
}
```

## 6. Route

```go
// server/router/router.go, in RegisterRoutes
r.GET("/api/v1/orders/:id", WrapContext(serverContext.OrderHandler.GetOrder))
```

## 7. Docs and mocks

```bash
make mock      # mocks for IOrderUseCase and OrderRepo
make swagger   # adds the endpoint to docs/openapi/oas.json
```

Add the endpoint to [API](api.md) if it's part of your public contract.

## 8. Tests

| Test | File | Replaces |
|------|------|----------|
| Domain rules | `internal/domain/order/order_test.go` | nothing |
| Use case | `internal/usecase/order/order_test.go` | `OrderRepo` with its mock |
| Handler: 200, 400, 404, 500 | `internal/api/order_test.go` | `IOrderUseCase` with its mock |
| End to end | `integrationtests/order_test.go` | nothing; calls the running app |

Run `make pre-commit` before you open a pull request. See [Testing](testing.md) for the patterns.
