package user

import (
	"go-service-template/internal/api/dto"
	"go-service-template/internal/domain/user"
)

func (s *UseCase) CreateUserRequest(_ *dto.CreateUserRequest) (*user.User, error) {
	// Placeholder for reset logic
	return &user.User{}, nil
}
