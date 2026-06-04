package user

// User represents a domain user entity.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// CreateUserInput holds domain input for creating a user.
type CreateUserInput struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// NewUser builds a domain user from input parameters.
func NewUser(input CreateUserInput) *User {
	return &User{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}
}
