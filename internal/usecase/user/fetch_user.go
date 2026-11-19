package user

import (
	domain "github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"
)

func (s *UseCase) FetchUser(input *FetchUserInput) (*domain.User, error) {
	// Adopter: implement fetch logic and use s.userStore / s.userProvider.
	if input == nil {
		return &domain.User{}, nil
	}
	return &domain.User{ID: input.ID}, nil
}
