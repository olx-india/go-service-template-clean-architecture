package user

import (
	"testing"

	domain "github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"

	"github.com/stretchr/testify/assert"
)

type fakeStore struct{}

func (fakeStore) Save(user domain.User) (domain.User, error) { return user, nil }

func (fakeStore) Fetch(id int) (domain.User, error) { return domain.User{ID: id}, nil }

type fakeProvider struct{}

func (fakeProvider) Save(user domain.User) (domain.User, error) { return user, nil }

func (fakeProvider) Fetch(id int) (domain.User, error) { return domain.User{ID: id}, nil }

func TestNewUserUseCase_ReturnsConfiguredUseCase(t *testing.T) {
	store := fakeStore{}
	provider := fakeProvider{}
	useCase := NewUserUseCase(store, provider)

	assert.NotNil(t, useCase)
	assert.Equal(t, store, useCase.userStore)
	assert.Equal(t, provider, useCase.userProvider)
}

func TestUseCase_CreateUserRequest_MapsInputToDomainUser(t *testing.T) {
	useCase := NewUserUseCase(nil, nil)
	input := &CreateUserInput{ID: 123, Name: "John Doe", Email: "john@example.com", Age: 30}

	user, err := useCase.CreateUserRequest(input)

	assert.NoError(t, err)
	assert.Equal(t, &domain.User{ID: 123, Name: "John Doe", Email: "john@example.com", Age: 30}, user)
}

func TestUseCase_CreateUserRequest_NilInput_ReturnsEmptyUser(t *testing.T) {
	useCase := NewUserUseCase(nil, nil)

	user, err := useCase.CreateUserRequest(nil)

	assert.NoError(t, err)
	assert.Equal(t, &domain.User{}, user)
}

func TestUseCase_FetchUser_ReturnsDomainUserWithID(t *testing.T) {
	useCase := NewUserUseCase(nil, nil)

	user, err := useCase.FetchUser(&FetchUserInput{ID: 123})

	assert.NoError(t, err)
	assert.Equal(t, &domain.User{ID: 123}, user)
}

func TestUseCase_FetchUser_NilInput_ReturnsEmptyUser(t *testing.T) {
	useCase := NewUserUseCase(nil, nil)

	user, err := useCase.FetchUser(nil)

	assert.NoError(t, err)
	assert.Equal(t, &domain.User{}, user)
}

func TestUseCase_InterfaceCompliance(t *testing.T) {
	var _ IUserUseCase = NewUserUseCase(nil, nil)
}
