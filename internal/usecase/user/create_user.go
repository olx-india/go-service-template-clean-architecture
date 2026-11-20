package user

import (
	domain "github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"
)

func (s *UseCase) CreateUserRequest(input *CreateUserInput) (*domain.User, error) {
	// Adopter: implement business logic and use s.userStore / s.userProvider.
	if input == nil {
		return &domain.User{}, nil
	}
	return &domain.User{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}, nil
}
