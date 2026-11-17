package webapi

import (
	"github.com/olx-india/go-service-template-clean-architecture/internal/domain/user"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/repo"
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
