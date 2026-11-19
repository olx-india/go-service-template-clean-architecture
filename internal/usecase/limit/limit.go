package limit

// CheckLimitInput is the use-case input for checking a user's limit.
type CheckLimitInput struct {
	UserID int
}

// Result is the use-case output for a limit operation.
type Result struct {
	UserID         int
	LimitAvailable int
}

// Store is the storage port owned by the limit use case.
type Store interface {
	LimitAvailable(userID int) (int, error)
	ResetLimit(userID int) error
}

type UseCase struct {
	store Store
}

func NewLimitUseCase(store Store) *UseCase {
	return &UseCase{store: store}
}

type ILimitUseCase interface {
	CheckLimit(input *CheckLimitInput) (Result, error)
	ResetLimit(input *CheckLimitInput) (Result, error)
}
