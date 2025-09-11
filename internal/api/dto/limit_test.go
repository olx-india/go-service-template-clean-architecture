package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLimitRequest_ValidRequest_ReturnsNoError(t *testing.T) {
	req := CheckLimitRequest{
		UserID: 123,
	}

	assert.Equal(t, 123, req.UserID)
}

func TestCheckLimitRequest_ZeroUserID_ReturnsZeroValue(t *testing.T) {
	req := CheckLimitRequest{}

	assert.Equal(t, 0, req.UserID)
}

func TestCheckLimitResponse_ValidResponse_ReturnsCorrectValues(t *testing.T) {
	response := CheckLimitResponse{
		UserID:         123,
		LimitAvailable: 100,
	}

	assert.Equal(t, 123, response.UserID)
	assert.Equal(t, 100, response.LimitAvailable)
}

func TestCheckLimitResponse_ZeroValues_ReturnsZeroValues(t *testing.T) {
	response := CheckLimitResponse{}

	assert.Equal(t, 0, response.UserID)
	assert.Equal(t, 0, response.LimitAvailable)
}

func TestCheckLimitResponse_NegativeValues_ReturnsNegativeValues(t *testing.T) {
	response := CheckLimitResponse{
		UserID:         -1,
		LimitAvailable: -50,
	}

	assert.Equal(t, -1, response.UserID)
	assert.Equal(t, -50, response.LimitAvailable)
}

func TestCheckLimitResponse_MaxIntValues_ReturnsMaxIntValues(t *testing.T) {
	response := CheckLimitResponse{
		UserID:         2147483647,
		LimitAvailable: 2147483647,
	}

	assert.Equal(t, 2147483647, response.UserID)
	assert.Equal(t, 2147483647, response.LimitAvailable)
}
