package persistent

import (
	"testing"

	domain "go-service-template/internal/domain/user"

	"github.com/stretchr/testify/assert"
)

func TestNewUserRepo_ReturnsImpl(t *testing.T) {
	r := NewUserRepo()
	assert.NotNil(t, r)
}

func TestUserRepo_Save_Defaults(t *testing.T) {
	r := &userRepo{}
	_, err := r.Save(domain.User{})
	assert.NoError(t, err)
}

func TestUserRepo_Fetch_Defaults(t *testing.T) {
	r := &userRepo{}
	_, err := r.Fetch(1)
	assert.NoError(t, err)
}
