package user

import (
	"go-service-template/internal/api/dto"
	domain "go-service-template/internal/domain/user"
)

func (s *UseCase) FetchUser(_ *dto.FetchUserRequest) (*domain.User, error) {
	// Adopter: implement fetch logic and wire repositories.
	return &domain.User{}, nil
}
