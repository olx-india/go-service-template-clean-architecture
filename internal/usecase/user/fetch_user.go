package user

import (
	"go-service-template/internal/api/dto"
	"go-service-template/internal/domain/user"
)

func (s *UseCase) FetchUser(_ *dto.FetchUserRequest) (*user.User, error) {
	// Placeholder for rate limiting logic
	return &user.User{}, nil
}
