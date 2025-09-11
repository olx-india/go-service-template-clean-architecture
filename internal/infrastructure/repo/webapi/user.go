package webapi

import (
	"go-service-template/internal/domain/user"
	"go-service-template/internal/infrastructure/repo"
)

// UserWebAPI -.
type userWebAPI struct{}

// NewUserWebAPI -.
func NewUserWebAPI() repo.UserWebAPI {
	return &userWebAPI{}
}

// Fetch -.
func (r *userWebAPI) Fetch(_ int) (user.User, error) {
	return user.User{}, nil
}

// Save -.
func (r *userWebAPI) Save(_ user.User) (user.User, error) {
	return user.User{}, nil
}
