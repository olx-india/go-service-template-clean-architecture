package context

import (
    stdctx "context"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGinContext_Wrap(t *testing.T) {
    gin.SetMode(gin.TestMode)
    c, _ := gin.CreateTestContext(httptest.NewRecorder())
    
    gctx, err := NewGinContext(c)
    
    assert.NoError(t, err)
    assert.NotNil(t, gctx) 
    assert.NotNil(t, gctx.Context)
}

func TestGinContext_Methods(t *testing.T) {
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("GET", "/", nil)
    
    gctx, _ := NewGinContext(c)

    gctx.JSON(200, gin.H{"ok": true})
    assert.Equal(t, http.StatusOK, w.Code)
    gctx.AbortWithStatus(204)
    gctx.Next()
    assert.NotNil(t, gctx.Request())
    assert.NotNil(t, gctx.Writer())
    var v struct{ A string }
    if err := gctx.ShouldBindJSON(&v); err == nil {
        _ = stdctx.Background()
    }
}


