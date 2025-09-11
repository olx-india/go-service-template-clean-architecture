package router

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    gincontext "go-service-template/internal/infrastructure/context"
    "go-service-template/internal/infrastructure/config"
    "go-service-template/server/resolver"
)

func TestInitializeNewRelic_InvalidConfig(t *testing.T) {
    _, err := InitializeNewRelic("", "")
    assert.Error(t, err)
}

func TestCORSHeaders_Options_NoContent(t *testing.T) {
    mw := corsMiddleware()
    rr := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodOptions, "/any", nil)
    c, _ := gin.CreateTestContext(rr)
    c.Request = req
    mw(c)

    assert.Equal(t, http.StatusNoContent, rr.Code)
    assert.NotEmpty(t, rr.Header().Get("Access-Control-Allow-Origin"))
}

func TestWrapContext_CallsHandler(t *testing.T) {
    called := false
    wrapped := WrapContext(func(ctx *gincontext.GinContext) { called = true })
    rr := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(rr)
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    c.Request = req
    wrapped(c)
    assert.Equal(t, true, called)
}

func TestRegisterRoutes_Health_OK(t *testing.T) {
    cfg := config.NewConfig()
    r := NewRouter(cfg)
    srvCtx := resolver.NewResolver(cfg).ResolveServerContext()
    r.RegisterRoutes(srvCtx)
    engine := r.Get()
    rr := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    engine.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
}


