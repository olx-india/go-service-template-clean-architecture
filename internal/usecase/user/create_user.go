package user

import (
	"github.com/olx-india/go-service-template-clean-architecture/internal/api/dto"
	domain "github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"
)

func (s *UseCase) CreateUserRequest(_ *dto.CreateUserRequest) (*domain.User, error) {
	// Adopter: implement business logic and wire repositories.
	return &domain.User{}, nil
}
