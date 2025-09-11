package limit

import (
	"testing"

	"go-service-template/internal/api/dto"

	"github.com/stretchr/testify/assert"
)


func TestNewLimitUseCase_ValidInput_ReturnsLimitUseCase(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	assert.NotNil(t, useCase)
	assert.IsType(t, &UseCase{}, useCase)
	assert.Nil(t, useCase.redisProvider)
}

func TestNewLimitUseCase_NilInput_ReturnsLimitUseCase(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	assert.NotNil(t, useCase)
	assert.IsType(t, &UseCase{}, useCase)
	assert.Nil(t, useCase.redisProvider)
}

func TestUseCase_CheckLimit_ValidRequest_ReturnsResponse(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	request := &dto.CheckLimitRequest{
		UserID: 123,
	}

	response, err := useCase.CheckLimit(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, dto.CheckLimitResponse{}, response)
}

func TestUseCase_CheckLimit_NilRequest_ReturnsResponse(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	response, err := useCase.CheckLimit(nil)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, dto.CheckLimitResponse{}, response)
}

func TestUseCase_CheckLimit_EmptyRequest_ReturnsResponse(t *testing.T) {
	useCase := NewLimitUseCase(nil)

	request := &dto.CheckLimitRequest{}
	response, err := useCase.CheckLimit(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, dto.CheckLimitResponse{}, response)
}

func TestUseCase_CheckLimit_EdgeCaseUserIDs_ReturnsResponse(t *testing.T) {
	tests := []struct {
		name    string
		request *dto.CheckLimitRequest
	}{
		{
			name: "MinValidUserID",
			request: &dto.CheckLimitRequest{
				UserID: 1,
			},
		},
		{
			name: "PositiveUserID",
			request: &dto.CheckLimitRequest{
				UserID: 123,
			},
		},
		{
			name: "MaxIntUserID",
			request: &dto.CheckLimitRequest{
				UserID: 2147483647,
			},
		},
		{
			name: "ZeroUserID",
			request: &dto.CheckLimitRequest{
				UserID: 0,
			},
		},
		{
			name: "NegativeUserID",
			request: &dto.CheckLimitRequest{
				UserID: -1,
			},
		},
		{
			name: "MinIntUserID",
			request: &dto.CheckLimitRequest{
				UserID: -2147483648,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewLimitUseCase(nil)

			response, err := useCase.CheckLimit(tt.request)

			assert.NoError(t, err)
			assert.NotNil(t, response)
			assert.IsType(t, dto.CheckLimitResponse{}, response)
		})
	}
}

func TestUseCase_CheckLimit_ResponseStructure_ReturnsCorrectType(t *testing.T) {
	useCase := NewLimitUseCase(nil)
	request := &dto.CheckLimitRequest{
		UserID: 123,
	}

	response, err := useCase.CheckLimit(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, 0, response.UserID)
	assert.IsType(t, 0, response.LimitAvailable)
}

func TestUseCase_CheckLimit_InterfaceCompliance(t *testing.T) {
	useCase := NewLimitUseCase(nil)
	var _ ILimitUseCase = useCase
	request := &dto.CheckLimitRequest{UserID: 123}
	
	_, err := useCase.CheckLimit(request)

	assert.NoError(t, err)
}

func TestUseCase_CheckLimit_ErrorHandling(t *testing.T) {
	useCase := NewLimitUseCase(nil)
	request := &dto.CheckLimitRequest{UserID: 123}

	_, err := useCase.CheckLimit(request)
	
	assert.NoError(t, err)
}