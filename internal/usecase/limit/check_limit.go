package limit

func (s *UseCase) CheckLimit(input *CheckLimitInput) (Result, error) {
	// Adopter: implement rate limiting with a Redis-backed port.
	if input == nil {
		return Result{}, nil
	}
	return Result{UserID: input.UserID, LimitAvailable: 0}, nil
}

func (s *UseCase) ResetLimit(input *CheckLimitInput) (Result, error) {
	// Adopter: implement rate limit reset with a Redis-backed port.
	if input == nil {
		return Result{}, nil
	}
	return Result{UserID: input.UserID, LimitAvailable: 0}, nil
}
