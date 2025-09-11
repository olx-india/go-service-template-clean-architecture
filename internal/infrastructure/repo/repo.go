package repo

import "go-service-template/internal/domain/user"

// UserRepo interface for user repository operations.
type UserRepo interface {
	Save(user.User) (user.User, error)
	Fetch(int) (user.User, error)
}

// UserWebAPI interface for user web API operations.
type UserWebAPI interface {
	Save(user.User) (user.User, error)
	Fetch(int) (user.User, error)
}
