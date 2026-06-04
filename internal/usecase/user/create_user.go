package user

import (
	"go-service-template/internal/api/dto"
	domain "go-service-template/internal/domain/user"
)

func (s *UseCase) CreateUserRequest(_ *dto.CreateUserRequest) (*domain.User, error) {
	// Adopter: implement business logic and wire repositories.
	return &domain.User{}, nil
}
