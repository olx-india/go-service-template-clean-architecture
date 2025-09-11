package limit

import (
	"go-service-template/internal/api/dto"
	"go-service-template/internal/infrastructure/provider/redis"
)

type UseCase struct {
	redisProvider *redis.Provider
}

func NewLimitUseCase(redisProvider *redis.Provider) *UseCase {
	return &UseCase{
		redisProvider: redisProvider,
	}
}

type ILimitUseCase interface {
	CheckLimit(_ *dto.CheckLimitRequest) (dto.CheckLimitResponse, error)
	ResetLimit(_ *dto.CheckLimitRequest) (dto.CheckLimitResponse, error)
}
