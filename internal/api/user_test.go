package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go-service-template/internal/domain/user"
	"go-service-template/internal/usecase/user/mocks"
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

func TestNewUserHandler_ValidInput_ReturnsUserHandler(t *testing.T) {
	mockUseCase := &mocks.IUserUseCase{}
	handler := NewUserHandler(mockUseCase)

	assert.NotNil(t, handler)
	assert.IsType(t, &userHandler{}, handler)
}

func TestUserHandler_CreateUser_ValidRequest_ReturnsCreatedResponse(t *testing.T) {
	requestBody := createUserRequestDTO()
	mockUseCase, handler, expectedUser := mockCreateUserUsecase(requestBody)
	w, ginCtx := setupUserTestContextWithJSON(t, requestBody)

	handler.CreateUser(ginCtx)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response user.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, *expectedUser, response)
	mockUseCase.AssertExpectations(t)
}

func mockCreateUserUsecase(requestBody dto.CreateUserRequest) (*mocks.IUserUseCase, IUserHandler, *user.User) {
	mockUseCase := &mocks.IUserUseCase{}
	handler := NewUserHandler(mockUseCase)
	expectedUser := user.CreateNewUser(requestBody)
	mockUseCase.On("CreateUserRequest", mock.AnythingOfType("*dto.CreateUserRequest")).Return(expectedUser, nil)
	return mockUseCase, handler, expectedUser
}

func createUserRequestDTO() dto.CreateUserRequest {
	return dto.CreateUserRequest{
		ID:    123,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}
}

func TestUserHandler_CreateUser_InvalidJSON_ReturnsBadRequest(t *testing.T) {
	tests := []struct {
		name        string
		requestBody string
		errorMsg    string
	}{
		{
			name:        "InvalidIDType",
			requestBody: `{"id": "invalid", "name": "John", "email": "john@example.com", "age": 30}`,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "MissingRequiredFields",
			requestBody: `{"name": "John"}`,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "InvalidEmail",
			requestBody: `{"id": 123, "name": "John", "email": "invalid-email", "age": 30}`,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "InvalidAge",
			requestBody: `{"id": 123, "name": "John", "email": "john@example.com", "age": -1}`,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "EmptyRequestBody",
			requestBody: ``,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "MalformedJSON",
			requestBody: `{"id": 123, "name": "John",}`,
			errorMsg:    "Invalid request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mocks.IUserUseCase{}
			handler := NewUserHandler(mockUseCase)

			w, ginCtx := setupUserTestContextWithRawJSON(t, tt.requestBody)

			handler.CreateUser(ginCtx)

			assert.Equal(t, http.StatusBadRequest, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Contains(t, response["error"], tt.errorMsg)
			mockUseCase.AssertNotCalled(t, "CreateUserRequest")
		})
	}
}

func TestUserHandler_CreateUser_UseCaseError_ReturnsInternalServerError(t *testing.T) {
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
			name:          "ValidationError",
			useCaseError:  errors.New("user already exists"),
			expectedError: "user already exists",
		},
		{
			name:          "GenericError",
			useCaseError:  errors.New("internal server error"),
			expectedError: "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mocks.IUserUseCase{}
			handler := NewUserHandler(mockUseCase)

			mockUseCase.On("CreateUserRequest", mock.AnythingOfType("*dto.CreateUserRequest")).Return((*user.User)(nil), tt.useCaseError)

			requestBody := dto.CreateUserRequest{
				ID:    123,
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   30,
			}
			w, ginCtx := setupUserTestContextWithJSON(t, requestBody)

			handler.CreateUser(ginCtx)

			assert.Equal(t, http.StatusInternalServerError, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Contains(t, response["error"], "Failed to create user")
			assert.Contains(t, response["error"], tt.expectedError)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestUserHandler_CreateUser_ValidUserData_ReturnsSuccessResponse(t *testing.T) {
	tests := []struct {
		name             string
		requestBody      dto.CreateUserRequest
		expectedResponse map[string]interface{}
	}{
		{
			name: "NormalUser",
			requestBody: dto.CreateUserRequest{
				ID:    123,
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   30,
			},
			expectedResponse: map[string]interface{}{
				"id":    float64(123),
				"name":  "John Doe",
				"email": "john@example.com",
				"age":   float64(30),
			},
		},
		{
			name: "MinValidValues",
			requestBody: dto.CreateUserRequest{
				ID:    1,
				Name:  "A",
				Email: "test@domain.co.uk",
				Age:   1,
			},
			expectedResponse: map[string]interface{}{
				"id":    float64(1),
				"name":  "A",
				"email": "test@domain.co.uk",
				"age":   float64(1),
			},
		},
		{
			name: "MaxAgeValue",
			requestBody: dto.CreateUserRequest{
				ID:    1,
				Name:  "Max Age User",
				Email: "max@example.com",
				Age:   130,
			},
			expectedResponse: map[string]interface{}{
				"id":    float64(1),
				"name":  "Max Age User",
				"email": "max@example.com",
				"age":   float64(130),
			},
		},
		{
			name: "ComplexEmail",
			requestBody: dto.CreateUserRequest{
				ID:    1,
				Name:  "Complex Email User",
				Email: "user.name+tag@subdomain.example.com",
				Age:   25,
			},
			expectedResponse: map[string]interface{}{
				"id":    float64(1),
				"name":  "Complex Email User",
				"email": "user.name+tag@subdomain.example.com",
				"age":   float64(25),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mocks.IUserUseCase{}
			handler := NewUserHandler(mockUseCase)

			expectedUser := user.CreateNewUser(dto.CreateUserRequest{
				ID:    int(tt.expectedResponse["id"].(float64)),
				Name:  tt.expectedResponse["name"].(string),
				Email: tt.expectedResponse["email"].(string),
				Age:   int(tt.expectedResponse["age"].(float64)),
			})
			mockUseCase.On("CreateUserRequest", mock.AnythingOfType("*dto.CreateUserRequest")).Return(expectedUser, nil)

			w, ginCtx := setupUserTestContextWithJSON(t, tt.requestBody)

			handler.CreateUser(ginCtx)

			assert.Equal(t, http.StatusCreated, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedResponse, response)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestUserHandler_FetchUser_ValidRequest_ReturnsOKResponse(t *testing.T) {
	mockUseCase := &mocks.IUserUseCase{}
	handler := NewUserHandler(mockUseCase)

	expectedResponse := map[string]interface{}{
		"id":    float64(123),
		"name":  "John Doe",
		"email": "john@example.com",
		"age":   float64(30),
	}

	expectedUser := user.CreateNewUser(dto.CreateUserRequest{
		ID:    int(expectedResponse["id"].(float64)),
		Name:  expectedResponse["name"].(string),
		Email: expectedResponse["email"].(string),
		Age:   int(expectedResponse["age"].(float64)),
	})
	mockUseCase.On("FetchUser", mock.AnythingOfType("*dto.FetchUserRequest")).Return(expectedUser, nil)

	requestBody := dto.FetchUserRequest{ID: 123}
	w, ginCtx := setupUserTestContextWithJSON(t, requestBody)

	handler.FetchUser(ginCtx)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expectedResponse, response)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_FetchUser_InvalidJSON_ReturnsBadRequest(t *testing.T) {
	tests := []struct {
		name        string
		requestBody string
		errorMsg    string
	}{
		{
			name:        "InvalidIDType",
			requestBody: `{"id": "invalid"}`,
			errorMsg:    "Invalid request body",
		},
		{
			name:        "MissingID",
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
			requestBody: `{"id": 123,}`,
			errorMsg:    "Invalid request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mocks.IUserUseCase{}
			handler := NewUserHandler(mockUseCase)

			w, ginCtx := setupUserTestContextWithRawJSON(t, tt.requestBody)

			handler.FetchUser(ginCtx)

			assert.Equal(t, http.StatusBadRequest, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Contains(t, response["error"], tt.errorMsg)
			mockUseCase.AssertNotCalled(t, "FetchUser")
		})
	}
}

func TestUserHandler_FetchUser_UseCaseError_ReturnsInternalServerError(t *testing.T) {
	tests := []struct {
		name          string
		useCaseError  error
		expectedError string
	}{
		{
			name:          "UserNotFound",
			useCaseError:  errors.New("user not found"),
			expectedError: "user not found",
		},
		{
			name:          "DatabaseError",
			useCaseError:  errors.New("database connection failed"),
			expectedError: "database connection failed",
		},
		{
			name:          "GenericError",
			useCaseError:  errors.New("internal server error"),
			expectedError: "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mocks.IUserUseCase{}
			handler := NewUserHandler(mockUseCase)

			mockUseCase.On("FetchUser", mock.AnythingOfType("*dto.FetchUserRequest")).Return((*user.User)(nil), tt.useCaseError)

			requestBody := dto.FetchUserRequest{ID: 123}
			w, ginCtx := setupUserTestContextWithJSON(t, requestBody)

			handler.FetchUser(ginCtx)

			assert.Equal(t, http.StatusInternalServerError, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Contains(t, response["error"], "Failed to fetch user")
			assert.Contains(t, response["error"], tt.expectedError)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestUserHandler_FetchUser_ValidUserIDs_ReturnsSuccessResponse(t *testing.T) {
	tests := []struct {
		name             string
		userID           int
		expectedResponse map[string]interface{}
	}{
		{
			name:   "MinValidID",
			userID: 1,
			expectedResponse: map[string]interface{}{
				"id":    float64(1),
				"name":  "Min User",
				"email": "min@example.com",
				"age":   float64(18),
			},
		},
		{
			name:   "PositiveID",
			userID: 123,
			expectedResponse: map[string]interface{}{
				"id":    float64(123),
				"name":  "John Doe",
				"email": "john@example.com",
				"age":   float64(30),
			},
		},
		{
			name:   "MaxIntID",
			userID: 2147483647,
			expectedResponse: map[string]interface{}{
				"id":    float64(2147483647),
				"name":  "Max ID User",
				"email": "max@example.com",
				"age":   float64(25),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mocks.IUserUseCase{}
			handler := NewUserHandler(mockUseCase)

			expectedUser := user.CreateNewUser(dto.CreateUserRequest{
				ID:    int(tt.expectedResponse["id"].(float64)),
				Name:  tt.expectedResponse["name"].(string),
				Email: tt.expectedResponse["email"].(string),
				Age:   int(tt.expectedResponse["age"].(float64)),
			})
			mockUseCase.On("FetchUser", mock.AnythingOfType("*dto.FetchUserRequest")).Return(expectedUser, nil)

			requestBody := dto.FetchUserRequest{ID: tt.userID}
			w, ginCtx := setupUserTestContextWithJSON(t, requestBody)

			handler.FetchUser(ginCtx)

			assert.Equal(t, http.StatusOK, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedResponse, response)
			mockUseCase.AssertExpectations(t)
		})
	}
}

func setupUserTestContextWithJSON(t *testing.T, requestBody interface{}) (*httptest.ResponseRecorder, *ginContext.GinContext) {
	gin.SetMode(gin.TestMode)

	jsonBody, err := json.Marshal(requestBody)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ginCtx, err := ginContext.NewGinContext(c)
	require.NoError(t, err)

	return w, ginCtx
}

func setupUserTestContextWithRawJSON(t *testing.T, requestBody string) (*httptest.ResponseRecorder, *ginContext.GinContext) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/users", bytes.NewBufferString(requestBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ginCtx, err := ginContext.NewGinContext(c)
	require.NoError(t, err)

	return w, ginCtx
}
