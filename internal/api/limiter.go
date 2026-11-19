package api

import (
	"net/http"

	"github.com/olx-india/go-service-template-clean-architecture/internal/api/dto"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/context"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/logger"
	limitPkg "github.com/olx-india/go-service-template-clean-architecture/internal/usecase/limit"
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

// CheckLimit checks the remaining rate limit for a user.
// @Title Check limit
// @Description Check the remaining rate limit for a user.
// @Accept json
// @Produce json
// @Param request body dto.CheckLimitRequest true "Limit check payload"
// @Success 200 {object} dto.CheckLimitResponse "Limit status"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal error"
// @Resource limits
// @Router /api/v1/limit/check [post]
func (api *limiterHandler) CheckLimit(ctx *context.GinContext) {
	api.handleLimit(ctx, "Checking limit", func(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
		result, err := api.limit.CheckLimit(&limitPkg.CheckLimitInput{UserID: req.UserID})
		return dto.CheckLimitResponse{
			UserID:         result.UserID,
			LimitAvailable: result.LimitAvailable,
		}, err
	}, "Limit checked successfully", "Failed to fetch limit")
}

// ResetLimit resets the rate limit for a user.
// @Title Reset limit
// @Description Reset the rate limit for a user.
// @Accept json
// @Produce json
// @Param request body dto.CheckLimitRequest true "Limit reset payload"
// @Success 200 {object} dto.CheckLimitResponse "Limit status after reset"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal error"
// @Resource limits
// @Router /api/v1/limit/reset [post]
func (api *limiterHandler) ResetLimit(ctx *context.GinContext) {
	api.handleLimit(ctx, "Resetting limit", func(req *dto.CheckLimitRequest) (dto.CheckLimitResponse, error) {
		result, err := api.limit.ResetLimit(&limitPkg.CheckLimitInput{UserID: req.UserID})
		return dto.CheckLimitResponse{
			UserID:         result.UserID,
			LimitAvailable: result.LimitAvailable,
		}, err
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
	ctx.JSON(statusCode, dto.ErrorResponse{Error: message + err.Error()})
}

func (api *limiterHandler) sendSuccessResponse(ctx *context.GinContext, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}
