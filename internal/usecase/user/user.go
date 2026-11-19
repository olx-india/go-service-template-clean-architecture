package user

import (
	domain "github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"
)

// CreateUserInput is the use-case input for creating a user.
type CreateUserInput struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// FetchUserInput is the use-case input for fetching a user.
type FetchUserInput struct {
	ID int
}

// Store is the persistence port owned by the user use case.
type Store interface {
	Save(domain.User) (domain.User, error)
	Fetch(int) (domain.User, error)
}

// Provider is the external-service port owned by the user use case.
type Provider interface {
	Save(domain.User) (domain.User, error)
	Fetch(int) (domain.User, error)
}

type UseCase struct {
	userStore    Store
	userProvider Provider
}

type IUserUseCase interface {
	CreateUserRequest(input *CreateUserInput) (*domain.User, error)
	FetchUser(input *FetchUserInput) (*domain.User, error)
}

func NewUserUseCase(userStore Store, userProvider Provider) *UseCase {
	return &UseCase{
		userStore:    userStore,
		userProvider: userProvider,
	}
}
