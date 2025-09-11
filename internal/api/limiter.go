package api

import (
	"net/http"

	"go-service-template/internal/api/dto"
	"go-service-template/internal/infrastructure/context"
	"go-service-template/internal/infrastructure/logger"
	limitPkg "go-service-template/internal/usecase/limit"

	"github.com/gin-gonic/gin"
)

type ILimiterHandler interface {
	CheckLimit(ctx *context.GinContext)
	ResetLimit(ctx *context.GinContext)
}

type limiterHandler struct {
	limit limitPkg.ILimitUseCase
}

func NewLimiterHandler(limiter limitPkg.ILimitUseCase) ILimiterHandler {
	return &limiterHandler{
		limit: limiter,
	}
}

func (api *limiterHandler) CheckLimit(ctx *context.GinContext) {
	api.handleLimit(ctx, "Checking limit", func(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
		return api.limit.CheckLimit(req)
	}, "Limit checked successfully", "Failed to fetch limit")
}

func (api *limiterHandler) ResetLimit(ctx *context.GinContext) {
	api.handleLimit(ctx, "Resetting limit", func(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
		return api.limit.ResetLimit(req)
	}, "Limit reset successfully", "Failed to reset limit")
}

func (api *limiterHandler) handleLimit(
	ctx *context.GinContext,
	startMsg string,
	usecaseFn func(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error),
	successMsg string,
	errorPrefix string,
) {
	logCtx := logger.GetLogContext(ctx.Context)
	logger.Info(logCtx, startMsg)

	var req dto.CheckLimitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(logCtx, "Invalid request body", logger.ErrorField(logger.FieldError, err))
		api.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body: ", err)
		return
	}

	response, err := usecaseFn(&req)
	if err != nil {
		logger.Error(logCtx, errorPrefix, logger.ErrorField(logger.FieldError, err), logger.Int(logger.FieldUserID, req.UserID))
		api.sendErrorResponse(ctx, http.StatusInternalServerError, errorPrefix+": ", err)
		return
	}

	logger.Info(logCtx, successMsg,
		logger.Int(logger.FieldUserID, req.UserID),
		logger.Int(logger.FieldStatusCode, http.StatusOK),
		logger.Int("limit_available", response.LimitAvailable),
	)
	api.sendSuccessResponse(ctx, http.StatusOK, response)
}

func (api *limiterHandler) sendErrorResponse(ctx *context.GinContext, statusCode int, message string, err error) {
	ctx.JSON(statusCode, gin.H{
		"error": message + err.Error(),
	})
}

func (api *limiterHandler) sendSuccessResponse(ctx *context.GinContext, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}
