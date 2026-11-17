package user

import (
	"github.com/olx-india/go-service-template-clean-architecture/internal/api/dto"
	domain "github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/provider/redis"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/repo"
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
