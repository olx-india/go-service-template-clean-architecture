package dto

// CreateUserRequest represents the request for creating a new user.
type CreateUserRequest struct {
	ID    int    `json:"id" binding:"required" example:"1" description:"User identity"`
	Name  string `json:"name" binding:"required" example:"Ada Lovelace" description:"Display name"`
	Email string `json:"email" binding:"required,email" example:"ada@example.com" description:"Email address"`
	Age   int    `json:"age" binding:"required,gte=0,lte=130" example:"36" minimum:"0" maximum:"130" description:"Age in years"`
}

// FetchUserRequest represents the request for fetching user details.
type FetchUserRequest struct {
	ID int `json:"id" binding:"required" example:"1" description:"User identity"`
}

// UserResponse is the API representation of a user.
type UserResponse struct {
	ID    int    `json:"id" example:"1" description:"User identity"`
	Name  string `json:"name" example:"Ada Lovelace" description:"Display name"`
	Email string `json:"email" example:"ada@example.com" description:"Email address"`
	Age   int    `json:"age" example:"36" description:"Age in years"`
}
