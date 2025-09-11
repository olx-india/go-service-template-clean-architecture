package webapi

import (
    "testing"
    domain "go-service-template/internal/domain/user"
    "github.com/stretchr/testify/assert"
)

func TestNewUserWebAPI_ReturnsImpl(t *testing.T) {
    w := NewUserWebAPI()
    assert.NotNil(t, w)
}

func TestUserWebAPI_Save_Defaults(t *testing.T) {
    w := &userWebAPI{}
    _, err := w.Save(domain.User{})
    assert.NoError(t, err)
}

func TestUserWebAPI_Fetch_Defaults(t *testing.T) {
    w := &userWebAPI{}
    _, err := w.Fetch(1)
    assert.NoError(t, err)
}


