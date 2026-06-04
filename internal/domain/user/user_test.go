package user

import (
	"testing"
)

func TestNewUser_BuildsFromInput(t *testing.T) {
	input := CreateUserInput{ID: 42, Name: "Alice", Email: "alice@example.com", Age: 30}
	u := NewUser(input)

	if u == nil {
		t.Fatalf("expected non-nil user")
	}
	if u.ID != input.ID || u.Name != input.Name || u.Email != input.Email || u.Age != input.Age {
		t.Fatalf("user fields mismatch: got %+v, want %+v", *u, input)
	}
}
