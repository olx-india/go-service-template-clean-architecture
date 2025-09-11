package user

import "go-service-template/internal/api/dto"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func CreateNewUser(request dto.CreateUserRequest) *User {
	return &User{
		ID:    request.ID,
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	}
}
