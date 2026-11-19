package limit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeStore struct{}

func (fakeStore) LimitAvailable(int) (int, error) { return 0, nil }

func (fakeStore) ResetLimit(int) error { return nil }

func TestNewLimitUseCase_ReturnsConfiguredUseCase(t *testing.T) {
	store := fakeStore{}
	useCase := NewLimitUseCase(store)

	assert.NotNil(t, useCase)
	assert.Equal(t, store, useCase.store)
}

func TestUseCase_CheckLimit_ReturnsUserResult(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	result, err := useCase.CheckLimit(&CheckLimitInput{UserID: 123})

	assert.NoError(t, err)
	assert.Equal(t, Result{UserID: 123, LimitAvailable: 0}, result)
}

func TestUseCase_CheckLimit_NilInput_ReturnsEmptyResult(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	result, err := useCase.CheckLimit(nil)

	assert.NoError(t, err)
	assert.Equal(t, Result{}, result)
}

func TestUseCase_ResetLimit_ReturnsUserResult(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	result, err := useCase.ResetLimit(&CheckLimitInput{UserID: 123})

	assert.NoError(t, err)
	assert.Equal(t, Result{UserID: 123, LimitAvailable: 0}, result)
}

func TestUseCase_ResetLimit_NilInput_ReturnsEmptyResult(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	result, err := useCase.ResetLimit(nil)

	assert.NoError(t, err)
	assert.Equal(t, Result{}, result)
}

func TestUseCase_InterfaceCompliance(t *testing.T) {
	var _ ILimitUseCase = NewLimitUseCase(nil)
}
