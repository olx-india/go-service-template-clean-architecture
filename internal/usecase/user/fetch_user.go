package user

import (
	"github.com/olx-india/go-service-template-clean-architecture/internal/api/dto"
	domain "github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"
)

func (s *UseCase) FetchUser(_ *dto.FetchUserRequest) (*domain.User, error) {
	// Adopter: implement fetch logic and wire repositories.
	return &domain.User{}, nil
}
