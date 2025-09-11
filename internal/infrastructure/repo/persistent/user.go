package persistent

import (
	"github.com/evrone/go-clean-template/pkg/postgres"

	"go-service-template/internal/domain/user"
	"go-service-template/internal/infrastructure/repo"
)

// UserRepoImpl -.
type userRepo struct {
	*postgres.Postgres
}

// NewUserRepo -.
func NewUserRepo(pg *postgres.Postgres) repo.UserRepo {
	return &userRepo{pg}
}

// Fetch -.
func (r *userRepo) Fetch(_ int) (user.User, error) {
	return user.User{}, nil
}

// Save -.
func (r *userRepo) Save(_ user.User) (user.User, error) {
	return user.User{}, nil
}
