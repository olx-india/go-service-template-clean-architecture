package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/config"
	gincontext "github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/context"
	"github.com/olx-india/go-service-template-clean-architecture/server/resolver"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORSHeaders_Options_NoContent(t *testing.T) {
	tests := []struct {
		name            string
		method          string
		origin          string
		allowedOrigins  []string
		wantStatus      int
		wantOrigin      string
		wantCredentials string
		wantMethods     bool
	}{
		{
			name:           "wildcard allows any origin",
			method:         http.MethodGet,
			origin:         "https://app.example.com",
			allowedOrigins: []string{"*"},
			wantStatus:     http.StatusOK,
			wantOrigin:     "*",
		},
		{
			name:            "allowed origin echoes origin and credentials",
			method:          http.MethodGet,
			origin:          "https://app.example.com",
			allowedOrigins:  []string{"https://app.example.com"},
			wantStatus:      http.StatusOK,
			wantOrigin:      "https://app.example.com",
			wantCredentials: "true",
		},
		{
			name:           "disallowed origin omits CORS headers",
			method:         http.MethodGet,
			origin:         "https://evil.example.com",
			allowedOrigins: []string{"https://app.example.com"},
			wantStatus:     http.StatusOK,
		},
		{
			name:            "allowed origin preflight",
			method:          http.MethodOptions,
			origin:          "https://app.example.com",
			allowedOrigins:  []string{"https://app.example.com"},
			wantStatus:      http.StatusNoContent,
			wantOrigin:      "https://app.example.com",
			wantCredentials: "true",
			wantMethods:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := runCORSMiddleware(t, tt.method, tt.origin, tt.allowedOrigins)

			assert.Equal(t, tt.wantStatus, rr.Code)
			assert.Equal(t, tt.wantOrigin, rr.Header().Get("Access-Control-Allow-Origin"))
			assert.Equal(t, tt.wantCredentials, rr.Header().Get("Access-Control-Allow-Credentials"))
			assert.Equal(t, "Origin", rr.Header().Get("Vary"))
			if tt.wantMethods {
				assert.NotEmpty(t, rr.Header().Get("Access-Control-Allow-Methods"))
			}
		})
	}
}

func runCORSMiddleware(t *testing.T, method, origin string, allowedOrigins []string) *httptest.ResponseRecorder {
	t.Helper()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(method, "/any", nil)
	if origin != "" {
		req.Header.Set("Origin", origin)
	}
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	corsMiddleware(allowedOrigins)(c)

	return rr
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

func TestRegisterRoutes_SwaggerUI_OK(t *testing.T) {
	cfg := config.NewConfig()
	r := NewRouter(cfg)
	srvCtx := resolver.NewResolver(cfg).ResolveServerContext()
	r.RegisterRoutes(srvCtx)
	engine := r.Get()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
	engine.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "SwaggerUIBundle")

	rr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/swagger/doc.json", nil)
	engine.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"openapi": "3.0.0"`)
}
