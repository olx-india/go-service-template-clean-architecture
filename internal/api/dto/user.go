package dto

// CreateUserRequest represents the request for creating a new user.
type CreateUserRequest struct {
	ID    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=0,lte=130"`
}

// FetchUserRequest represents the request for fetching user details.
type FetchUserRequest struct {
	ID int `json:"id" binding:"required"`
}
