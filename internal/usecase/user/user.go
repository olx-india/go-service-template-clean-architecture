package user

import (
	"go-service-template/internal/api/dto"
	domain "go-service-template/internal/domain/user"
	"go-service-template/internal/infrastructure/provider/redis"
	"go-service-template/internal/infrastructure/repo"
)

type UseCase struct {
	redisProvider      *redis.Provider
	userRepository     repo.UserRepo
	userWebAPIProvider repo.UserWebAPI
}

type IUserUseCase interface {
	CreateUserRequest(req *dto.CreateUserRequest) (*domain.User, error)
	FetchUser(req *dto.FetchUserRequest) (*domain.User, error)
}

func NewUserUseCase(redisProvider *redis.Provider, userRepository repo.UserRepo, userWebAPIProvider repo.UserWebAPI) *UseCase {
	return &UseCase{
		redisProvider:      redisProvider,
		userRepository:     userRepository,
		userWebAPIProvider: userWebAPIProvider,
	}
}
