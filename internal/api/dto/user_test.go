package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserRequest_ValidRequest_ReturnsCorrectValues(t *testing.T) {
	req := CreateUserRequest{
		ID:    123,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	assert.Equal(t, 123, req.ID)
	assert.Equal(t, "John Doe", req.Name)
	assert.Equal(t, "john@example.com", req.Email)
	assert.Equal(t, 30, req.Age)
}

func TestCreateUserRequest_ZeroValues_ReturnsZeroValues(t *testing.T) {
	req := CreateUserRequest{}

	assert.Equal(t, 0, req.ID)
	assert.Equal(t, "", req.Name)
	assert.Equal(t, "", req.Email)
	assert.Equal(t, 0, req.Age)
}

func TestCreateUserRequest_EdgeCaseValues_ReturnsCorrectValues(t *testing.T) {
	req := CreateUserRequest{
		ID:    0,
		Name:  "",
		Email: "test@domain.co.uk",
		Age:   0,
	}

	assert.Equal(t, 0, req.ID)
	assert.Equal(t, "", req.Name)
	assert.Equal(t, "test@domain.co.uk", req.Email)
	assert.Equal(t, 0, req.Age)
}

func TestCreateUserRequest_MaxAgeValue_ReturnsMaxAge(t *testing.T) {
	req := CreateUserRequest{
		ID:    1,
		Name:  "Max Age User",
		Email: "max@example.com",
		Age:   130,
	}

	assert.Equal(t, 130, req.Age)
}

func TestCreateUserRequest_ComplexEmail_ReturnsCorrectEmail(t *testing.T) {
	req := CreateUserRequest{
		ID:    1,
		Name:  "Complex Email User",
		Email: "user.name+tag@subdomain.example.com",
		Age:   25,
	}

	assert.Equal(t, "user.name+tag@subdomain.example.com", req.Email)
}

func TestCreateUserRequest_SpecialCharactersInName_ReturnsCorrectName(t *testing.T) {
	req := CreateUserRequest{
		ID:    1,
		Name:  "José María O'Connor-Smith",
		Email: "jose@example.com",
		Age:   35,
	}

	assert.Equal(t, "José María O'Connor-Smith", req.Name)
}

func TestFetchUserRequest_ValidRequest_ReturnsCorrectID(t *testing.T) {
	req := FetchUserRequest{
		ID: 456,
	}

	assert.Equal(t, 456, req.ID)
}

func TestFetchUserRequest_ZeroID_ReturnsZeroID(t *testing.T) {
	req := FetchUserRequest{}

	assert.Equal(t, 0, req.ID)
}

func TestFetchUserRequest_NegativeID_ReturnsNegativeID(t *testing.T) {
	req := FetchUserRequest{
		ID: -1,
	}

	assert.Equal(t, -1, req.ID)
}

func TestFetchUserRequest_MaxIntID_ReturnsMaxIntID(t *testing.T) {
	req := FetchUserRequest{
		ID: 2147483647,
	}

	assert.Equal(t, 2147483647, req.ID)
}
