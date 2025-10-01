package resolver

import (
	"testing"

	"go-service-template/internal/infrastructure/config"

	"github.com/stretchr/testify/assert"
)

func TestNewResolver_ReturnsResolver(t *testing.T) {
	cfg := config.NewConfig()

	r := NewResolver(cfg)

	assert.NotNil(t, r)
}

func TestResolveServerContext_ReturnsContext(t *testing.T) {
	cfg := config.NewConfig()
	r := NewResolver(cfg)

	ctx := r.ResolveServerContext()

	assert.NotNil(t, ctx)
}
