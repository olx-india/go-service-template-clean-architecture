package limit

import "go-service-template/internal/api/dto"

func (s *UseCase) CheckLimit(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
	if req == nil {
		return dto.CheckLimitResponse{}, nil
	}
	return dto.CheckLimitResponse{UserID: req.UserID, LimitAvailable: 0}, nil
}

func (s *UseCase) ResetLimit(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
	if req == nil {
		return dto.CheckLimitResponse{}, nil
	}
	return dto.CheckLimitResponse{UserID: req.UserID, LimitAvailable: 0}, nil
}
