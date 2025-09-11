package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-service-template/internal/api/dto"
	ginContext "go-service-template/internal/infrastructure/context"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockLimitUseCase struct {
	mock.Mock
}

func (m *MockLimitUseCase) CheckLimit(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
	args := m.Called(req)
	return args.Get(0).(dto.CheckLimitResponse), args.Error(1)
}

func (m *MockLimitUseCase) ResetLimit(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
	args := m.Called(req)
	return args.Get(0).(dto.CheckLimitResponse), args.Error(1)
}

func TestNewLimiterHandler_ValidInput_ReturnsLimiterHandler(t *testing.T) {
	mockUseCase := &MockLimitUseCase{}
	handler := NewLimiterHandler(mockUseCase)

	assert.NotNil(t, handler)
	assert.IsType(t, &limiterHandler{}, handler)
}

func TestLimiterHandler_CheckLimit_ValidRequest_ReturnsSuccessResponse(t *testing.T) {
	mockUseCase := &MockLimitUseCase{}
	handler := NewLimiterHandler(mockUseCase)
	expectedResponse := dto.CheckLimitResponse{
		UserID:         123,
		LimitAvailable: 100,
	}
	mockUseCase.On("CheckLimit", mock.AnythingOfType("*dto.CheckLimitRequest")).Return(expectedResponse, nil)
	w, ginCtx := setupLimiterTestContextWithJSON(t, dto.CheckLimitRequest{UserID: 123})
	
	handler.CheckLimit(ginCtx)

	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.CheckLimitResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	mockUseCase.AssertExpectations(t)
}

func TestLimiterHandler_CheckLimit_InvalidJSON_ReturnsBadRequest(t *testing.T) {
	tests := []struct {
		name        string
		requestBody string
		errorMsg    string
	}{
		{
			name:        "InvalidUserIDType",
			requestBody: `{"userID": "invalid"}`,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "MissingUserID",
			requestBody: `{}`,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "EmptyRequestBody",
			requestBody: ``,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "MalformedJSON",
			requestBody: `{"userID": 123,}`,
			errorMsg:    "Invalid request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockLimitUseCase{}
			handler := NewLimiterHandler(mockUseCase)
			
			w, ginCtx := setupLimiterTestContextWithRawJSON(t, tt.requestBody)
			
			handler.CheckLimit(ginCtx)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			
			assert.Contains(t, response["error"], tt.errorMsg)
			mockUseCase.AssertNotCalled(t, "CheckLimit")
		})
	}
}

func TestLimiterHandler_CheckLimit_UseCaseError_ReturnsInternalServerError(t *testing.T) {
	tests := []struct {
		name          string
		useCaseError  error
		expectedError string
	}{
		{
			name:          "DatabaseError",
			useCaseError:  errors.New("database connection failed"),
			expectedError: "database connection failed",
		},
		{
			name:          "RedisError",
			useCaseError:  errors.New("redis connection timeout"),
			expectedError: "redis connection timeout",
		},
		{
			name:          "GenericError",
			useCaseError:  errors.New("internal server error"),
			expectedError: "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockLimitUseCase{}
			handler := NewLimiterHandler(mockUseCase)
			
			mockUseCase.On("CheckLimit", mock.AnythingOfType("*dto.CheckLimitRequest")).Return(dto.CheckLimitResponse{}, tt.useCaseError)
			
			w, ginCtx := setupLimiterTestContextWithJSON(t, dto.CheckLimitRequest{UserID: 123})
			
			handler.CheckLimit(ginCtx)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			
			assert.Contains(t, response["error"], "Failed to fetch limit")
			assert.Contains(t, response["error"], tt.expectedError)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestLimiterHandler_CheckLimit_ValidUserIDs_ReturnsSuccessResponse(t *testing.T) {
	tests := []struct {
		name             string
		userID           int
		expectedResponse dto.CheckLimitResponse
	}{
		{

			name:   "MinValidUserID",
			userID: 1,
			expectedResponse: dto.CheckLimitResponse{
				UserID:         1,
				LimitAvailable: 50,
			},
		},
		{
			name:   "NegativeUserID",
			userID: -1,
			expectedResponse: dto.CheckLimitResponse{
				UserID:         -1,
				LimitAvailable: 0,
			},
		},
		{
			name:   "MaxIntUserID",
			userID: 2147483647,
			expectedResponse: dto.CheckLimitResponse{
				UserID:         2147483647,
				LimitAvailable: 1000,
			},
		},
		{
			name:   "PositiveUserID",
			userID: 123,
			expectedResponse: dto.CheckLimitResponse{
				UserID:         123,
				LimitAvailable: 100,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &MockLimitUseCase{}
			handler := NewLimiterHandler(mockUseCase)
			
			mockUseCase.On("CheckLimit", mock.AnythingOfType("*dto.CheckLimitRequest")).Return(tt.expectedResponse, nil)
			
			w, ginCtx := setupLimiterTestContextWithJSON(t, dto.CheckLimitRequest{UserID: tt.userID})
			
			handler.CheckLimit(ginCtx)

			assert.Equal(t, http.StatusOK, w.Code)
			
			var response dto.CheckLimitResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			
			assert.Equal(t, tt.expectedResponse, response)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func setupLimiterTestContextWithJSON(t *testing.T, requestBody interface{}) (*httptest.ResponseRecorder, *ginContext.GinContext) {
	gin.SetMode(gin.TestMode)
	
	jsonBody, err := json.Marshal(requestBody)
	require.NoError(t, err)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/check-limit", bytes.NewBuffer(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	
	ginCtx, err := ginContext.NewGinContext(c)
	require.NoError(t, err)
	
	return w, ginCtx
}

func setupLimiterTestContextWithRawJSON(t *testing.T, requestBody string) (*httptest.ResponseRecorder, *ginContext.GinContext) {
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/check-limit", bytes.NewBufferString(requestBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	
	ginCtx, err := ginContext.NewGinContext(c)
	require.NoError(t, err)
	
	return w, ginCtx
}
