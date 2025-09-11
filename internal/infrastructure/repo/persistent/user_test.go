package persistent

import (
    "testing"

    "github.com/evrone/go-clean-template/pkg/postgres"
    domain "go-service-template/internal/domain/user"
    "github.com/stretchr/testify/assert"
)

func TestNewUserRepo_ReturnsImpl(t *testing.T) {
    var pg *postgres.Postgres // nil is acceptable for constructor
    r := NewUserRepo(pg)
    assert.NotNil(t, r)
}

func TestUserRepo_Save_Defaults(t *testing.T) {
    r := &userRepo{Postgres: nil}
    _, err := r.Save(domain.User{})
    assert.NoError(t, err)
}

func TestUserRepo_Fetch_Defaults(t *testing.T) {
    r := &userRepo{Postgres: nil}
    _, err := r.Fetch(1)
    assert.NoError(t, err)
}


