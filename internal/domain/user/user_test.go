package user

import (
    "testing"

    "go-service-template/internal/api/dto"
)

func TestCreateNewUser_BuildsFromDTO(t *testing.T) {
    req := dto.CreateUserRequest{ID: 42, Name: "Alice", Email: "alice@example.com", Age: 30}
    u := CreateNewUser(req)

    if u == nil {
        t.Fatalf("expected non-nil user")
    }
    if u.ID != req.ID || u.Name != req.Name || u.Email != req.Email || u.Age != req.Age {
        t.Fatalf("user fields mismatch: got %+v, want %+v", *u, req)
    }
}


