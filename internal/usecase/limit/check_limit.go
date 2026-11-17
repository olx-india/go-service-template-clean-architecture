package limit

import "github.com/olx-india/go-service-template-clean-architecture/internal/api/dto"

func (s *UseCase) CheckLimit(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
	// Adopter: implement rate limiting with Redis.
	if req == nil {
		return dto.CheckLimitResponse{}, nil
	}
	return dto.CheckLimitResponse{UserID: req.UserID, LimitAvailable: 0}, nil
}

func (s *UseCase) ResetLimit(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
	// Adopter: implement rate limit reset with Redis.
	if req == nil {
		return dto.CheckLimitResponse{}, nil
	}
	return dto.CheckLimitResponse{UserID: req.UserID, LimitAvailable: 0}, nil
}
