package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	ginContext "go-service-template/internal/infrastructure/context"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHealthHandler_ValidInput_ReturnsHealthHandler(t *testing.T) {
	handler := NewHealthHandler()

	assert.NotNil(t, handler)
	assert.IsType(t, &HealthHandler{}, handler)
}

func TestHealthHandler_Check_ValidContext_ReturnsOKStatus(t *testing.T) {
	handler := NewHealthHandler()
	w, ginCtx := setupHealthTestContext(t)
	
	handler.Check(ginCtx)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHealthHandler_Check_ResponseFields_ReturnsCorrectValues(t *testing.T) {
	tests := []struct {
		name           string
		field          string
		expectedValue  string
		setupContext   func(t *testing.T) (*httptest.ResponseRecorder, *ginContext.GinContext)
	}{
		{
			name:          "StatusField",
			field:         "status",
			expectedValue: "ok",
			setupContext:  setupHealthTestContext,
		},
		{
			name:          "ServiceField",
			field:         "service",
			expectedValue: "go-service-template",
			setupContext:  setupHealthTestContext,
		},
		{
			name:          "StatusFieldWithRequest",
			field:         "status",
			expectedValue: "ok",
			setupContext:  setupHealthTestContextWithRequest,
		},
		{
			name:          "ServiceFieldWithRequest",
			field:         "service",
			expectedValue: "go-service-template",
			setupContext:  setupHealthTestContextWithRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHealthHandler()
			w, ginCtx := tt.setupContext(t)
			
			handler.Check(ginCtx)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			
			assert.Equal(t, tt.expectedValue, response[tt.field])
		})
	}
}

func TestHealthHandler_Check_ResponseHeaders_ReturnsCorrectContentType(t *testing.T) {
	tests := []struct {
		name         string
		setupContext func(t *testing.T) (*httptest.ResponseRecorder, *ginContext.GinContext)
	}{
		{
			name:         "BasicContext",
			setupContext: setupHealthTestContext,
		},
		{
			name:         "ContextWithRequest",
			setupContext: setupHealthTestContextWithRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHealthHandler()
			w, ginCtx := tt.setupContext(t)
			
			handler.Check(ginCtx)

			assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
		})
	}
}

func TestHealthHandler_Check_ResponseFormat_ReturnsValidJSON(t *testing.T) {
	tests := []struct {
		name         string
		setupContext func(t *testing.T) (*httptest.ResponseRecorder, *ginContext.GinContext)
	}{
		{
			name:         "BasicContext",
			setupContext: setupHealthTestContext,
		},
		{
			name:         "ContextWithRequest",
			setupContext: setupHealthTestContextWithRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHealthHandler()
			w, ginCtx := tt.setupContext(t)
			
			handler.Check(ginCtx)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			
			assert.NotNil(t, response)
		})
	}
}

func TestHealthHandler_Check_CompleteResponse_ReturnsExpectedStructure(t *testing.T) {
	handler := NewHealthHandler()
	w, ginCtx := setupHealthTestContext(t)
	
	handler.Check(ginCtx)

	expectedResponse := map[string]interface{}{
		"status":  "ok",
		"service": "go-service-template",
	}
	
	var actualResponse map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	require.NoError(t, err)
	
	assert.Equal(t, expectedResponse, actualResponse)
}

func setupHealthTestContext(t *testing.T) (*httptest.ResponseRecorder, *ginContext.GinContext) {
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	ginCtx, err := ginContext.NewGinContext(c)
	require.NoError(t, err)
	
	return w, ginCtx
}

func setupHealthTestContextWithRequest(t *testing.T) (*httptest.ResponseRecorder, *ginContext.GinContext) {
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/health", nil)
	require.NoError(t, err)
	c.Request = req
	
	ginCtx, err := ginContext.NewGinContext(c)
	require.NoError(t, err)
	
	return w, ginCtx
}