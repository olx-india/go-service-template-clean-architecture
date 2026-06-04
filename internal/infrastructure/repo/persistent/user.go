package persistent

import (
	"go-service-template/internal/domain/user"
	"go-service-template/internal/infrastructure/repo"
)

// userRepo is a stub repository. Replace with your database implementation.
type userRepo struct {
	// Adopter: inject your database connection here.
}

// NewUserRepo creates a stub user repository.
func NewUserRepo() repo.UserRepo {
	return &userRepo{}
}

// Fetch returns a stub user. Implement your query logic here.
func (r *userRepo) Fetch(_ int) (user.User, error) {
	return user.User{}, nil
}

// Save persists a stub user. Implement your save logic here.
func (r *userRepo) Save(_ user.User) (user.User, error) {
	return user.User{}, nil
}
