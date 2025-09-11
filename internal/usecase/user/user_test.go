package user

import (
	"testing"

	"go-service-template/internal/api/dto"
	domain "go-service-template/internal/domain/user"

	"github.com/stretchr/testify/assert"
)


func TestNewUserUseCase_ValidInput_ReturnsUserUseCase(t *testing.T) {
	useCase := NewUserUseCase(nil, nil, nil)

	assert.NotNil(t, useCase)
	assert.IsType(t, &UseCase{}, useCase)
	assert.Nil(t, useCase.redisProvider)
	assert.Nil(t, useCase.userRepository)
	assert.Nil(t, useCase.userWebAPIProvider)
}

func TestNewUserUseCase_NilInputs_ReturnsUserUseCase(t *testing.T) {
	useCase := NewUserUseCase(nil, nil, nil)

	assert.NotNil(t, useCase)
	assert.IsType(t, &UseCase{}, useCase)
	assert.Nil(t, useCase.redisProvider)
	assert.Nil(t, useCase.userRepository)
	assert.Nil(t, useCase.userWebAPIProvider)
}

func TestUseCase_CreateUserRequest_ValidRequest_ReturnsUser(t *testing.T) {
	useCase := NewUserUseCase(nil, nil, nil)

	request := &dto.CreateUserRequest{
		ID:    123,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	user, err := useCase.CreateUserRequest(request)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &domain.User{}, user)
}

func TestUseCase_CreateUserRequest_NilRequest_ReturnsUser(t *testing.T) {
	useCase := NewUserUseCase(nil, nil, nil)

	user, err := useCase.CreateUserRequest(nil)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &domain.User{}, user)
}

func TestUseCase_CreateUserRequest_EmptyRequest_ReturnsUser(t *testing.T) {
	useCase := NewUserUseCase(nil, nil, nil)

	request := &dto.CreateUserRequest{}
	user, err := useCase.CreateUserRequest(request)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &domain.User{}, user)
}

func TestUseCase_CreateUserRequest_EdgeCaseValues_ReturnsUser(t *testing.T) {
	tests := []struct {
		name    string
		request *dto.CreateUserRequest
	}{
		{
			name: "MinValidValues",
			request: &dto.CreateUserRequest{
				ID:    1,
				Name:  "A",
				Email: "test@domain.co.uk",
				Age:   1,
			},
		},
		{
			name: "MaxAgeValue",
			request: &dto.CreateUserRequest{
				ID:    1,
				Name:  "Max Age User",
				Email: "max@example.com",
				Age:   130,
			},
		},
		{
			name: "ComplexEmail",
			request: &dto.CreateUserRequest{
				ID:    1,
				Name:  "Complex Email User",
				Email: "user.name+tag@subdomain.example.com",
				Age:   25,
			},
		},
		{
			name: "SpecialCharactersInName",
			request: &dto.CreateUserRequest{
				ID:    1,
				Name:  "José María O'Connor-Smith",
				Email: "jose@example.com",
				Age:   35,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
					useCase := NewUserUseCase(nil, nil, nil)

			user, err := useCase.CreateUserRequest(tt.request)

			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.IsType(t, &domain.User{}, user)
		})
	}
}

func TestUseCase_FetchUser_ValidRequest_ReturnsUser(t *testing.T) {
			useCase := NewUserUseCase(nil, nil, nil)

	request := &dto.FetchUserRequest{
		ID: 123,
	}

	user, err := useCase.FetchUser(request)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &domain.User{}, user)
}

func TestUseCase_FetchUser_NilRequest_ReturnsUser(t *testing.T) {
			useCase := NewUserUseCase(nil, nil, nil)

	user, err := useCase.FetchUser(nil)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &domain.User{}, user)
}

func TestUseCase_FetchUser_EmptyRequest_ReturnsUser(t *testing.T) {
			useCase := NewUserUseCase(nil, nil, nil)

	request := &dto.FetchUserRequest{}
	user, err := useCase.FetchUser(request)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &domain.User{}, user)
}

func TestUseCase_FetchUser_EdgeCaseIDs_ReturnsUser(t *testing.T) {
	tests := []struct {
		name    string
		request *dto.FetchUserRequest
	}{
		{
			name: "MinValidID",
			request: &dto.FetchUserRequest{
				ID: 1,
			},
		},
		{
			name: "PositiveID",
			request: &dto.FetchUserRequest{
				ID: 123,
			},
		},
		{
			name: "MaxIntID",
			request: &dto.FetchUserRequest{
				ID: 2147483647,
			},
		},
		{
			name: "ZeroID",
			request: &dto.FetchUserRequest{
				ID: 0,
			},
		},
		{
			name: "NegativeID",
			request: &dto.FetchUserRequest{
				ID: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
					useCase := NewUserUseCase(nil, nil, nil)

			user, err := useCase.FetchUser(tt.request)

			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.IsType(t, &domain.User{}, user)
		})
	}
}

func TestUseCase_InterfaceCompliance(t *testing.T) {
	useCase := NewUserUseCase(nil, nil, nil)
	var _ IUserUseCase = useCase
	request := &dto.CreateUserRequest{ID: 1, Name: "Test", Email: "test@example.com", Age: 25}

	_, err := useCase.CreateUserRequest(request)
	assert.NoError(t, err)

	fetchRequest := &dto.FetchUserRequest{ID: 1}
	_, err = useCase.FetchUser(fetchRequest)
	assert.NoError(t, err)
}