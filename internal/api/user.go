package api

import (
	"context"
	"net/http"

	"go-service-template/internal/api/dto"
	ginContext "go-service-template/internal/infrastructure/context"
	"go-service-template/internal/infrastructure/logger"
	userPkg "go-service-template/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateUser(ctx *ginContext.GinContext)
	FetchUser(ctx *ginContext.GinContext)
}

type userHandler struct {
	user userPkg.IUserUseCase
}

func NewUserHandler(userUseCase userPkg.IUserUseCase) IUserHandler {
	return &userHandler{
		user: userUseCase,
	}
}

// CreateUser create user from the user request.
func (api *userHandler) CreateUser(ctx *ginContext.GinContext) {
	logCtx := logger.GetLogContext(ctx.Context)
	logger.Info(logCtx, "Creating user")

	var req dto.CreateUserRequest
	if !api.validateRequest(logCtx, ctx, &req) {
		return
	}

	response, err := api.user.CreateUserRequest(&req)
	if !api.handleUseCaseError(logCtx, ctx, err, "Failed to create user", req.ID) {
		return
	}

	api.logAndSendSuccess(logCtx, ctx, "User created successfully", req.ID, http.StatusCreated, response)
}

// FetchUser fetches user from a database.
func (api *userHandler) FetchUser(ctx *ginContext.GinContext) {
	logCtx := logger.GetLogContext(ctx.Context)
	logger.Info(logCtx, "Fetching user")

	var req dto.FetchUserRequest
	if !api.validateRequest(logCtx, ctx, &req) {
		return
	}

	user, err := api.user.FetchUser(&req)
	if !api.handleUseCaseError(logCtx, ctx, err, "Failed to fetch user", req.ID) {
		return
	}

	api.logAndSendSuccess(logCtx, ctx, "User fetched successfully", req.ID, http.StatusOK, user)
}

// validateRequest validates the JSON request body and handles binding errors.
func (api *userHandler) validateRequest(logCtx context.Context, ctx *ginContext.GinContext, req interface{}) bool {
	if err := ctx.ShouldBindJSON(req); err != nil {
		logger.Error(logCtx, "Invalid request body", logger.ErrorField(logger.FieldError, err))
		api.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body: ", err)
		return false
	}
	return true
}

// handleUseCaseError handles use case errors and returns false if an error occurred.
func (api *userHandler) handleUseCaseError(logCtx context.Context, ctx *ginContext.GinContext, err error, errorMessage string, userID int) bool {
	if err != nil {
		logger.Error(logCtx, errorMessage, logger.ErrorField(logger.FieldError, err), logger.Int(logger.FieldUserID, userID))
		api.sendErrorResponse(ctx, http.StatusInternalServerError, errorMessage+": ", err)
		return false
	}
	return true
}

// logAndSendSuccess logs success and sends the success response.
func (api *userHandler) logAndSendSuccess(logCtx context.Context, ctx *ginContext.GinContext, message string, userID, statusCode int, data interface{}) {
	logger.Info(logCtx, message,
		logger.Int(logger.FieldUserID, userID),
		logger.Int(logger.FieldStatusCode, statusCode),
	)
	api.sendSuccessResponse(ctx, statusCode, data)
}

// sendErrorResponse is a helper function to send error responses.
func (api *userHandler) sendErrorResponse(ctx *ginContext.GinContext, statusCode int, message string, err error) {
	ctx.JSON(statusCode, gin.H{
		"error": message + err.Error(),
	})
}

// sendSuccessResponse is a helper function to send success responses.
func (api *userHandler) sendSuccessResponse(ctx *ginContext.GinContext, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}
