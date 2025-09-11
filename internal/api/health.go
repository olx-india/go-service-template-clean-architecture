package api

import (
	"net/http"

	"go-service-template/internal/infrastructure/context"
	"go-service-template/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(ctx *context.GinContext) {
	logCtx := logger.GetLogContext(ctx.Context)
	logger.Info(logCtx, "Health check requested", logger.Int(logger.FieldStatusCode, http.StatusOK))

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "go-service-template",
	})
}
