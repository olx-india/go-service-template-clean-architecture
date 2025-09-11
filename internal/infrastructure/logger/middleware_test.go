package logger

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware_SetsTracingHeaders(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.New()
    r.Use(LoggingMiddleware())
    r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

    w := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/ping", nil)
    r.ServeHTTP(w, req)

    assert.NotEmpty(t, w.Header().Get(XRequestID))
}

func TestLoggingMiddleware_ResponseStatusOK(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.New()
    r.Use(LoggingMiddleware())
    r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

    w := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/ping", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestFetchRequestID_GeneratesWhenMissing(t *testing.T) {
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("GET", "/", nil)
    
    rid, _ := fetchRequestAndTraceIDs(c)
    
    assert.NotEmpty(t, rid)
}

func TestFetchTraceID_GeneratesWhenMissing(t *testing.T) {
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("GET", "/", nil)
    
    _, tid := fetchRequestAndTraceIDs(c)
    
    assert.NotEmpty(t, tid)
}

func TestAddTraceIDsInResponseHeaders_SetsRequestID(t *testing.T) {
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    
    addTraceIDsInResponseHeaders(c, "r", "t")
    
    assert.Equal(t, "r", w.Header().Get(XRequestID))
}

func TestAddTraceIDsInResponseHeaders_SetsTraceID(t *testing.T) {
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    
    addTraceIDsInResponseHeaders(c, "r", "t")
    
    assert.Equal(t, "t", w.Header().Get(XTraceID))
}

func TestCreateContextWithTraceIDs_SetsRequestID(t *testing.T) {
    ctx := createContextWithTraceIDs("r", "t")
    assert.NotNil(t, ctx.Value(RequestID))
}

func TestCreateContextWithTraceIDs_SetsTraceID(t *testing.T) {
    ctx := createContextWithTraceIDs("r", "t")
    assert.NotNil(t, ctx.Value(TraceID))
}

func TestSetLoggerInContext_AndGetLogContext(t *testing.T) {
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    ctx := createContextWithTraceIDs("r", "t")
    
    setLoggerInContext(ctx, c)
    
    got := GetLogContext(c)
    assert.NotNil(t, got)
}


