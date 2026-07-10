package api

import (
	"context"
	"net/http"
	"time"

	"go-service-template/internal/infrastructure/config"
	gincontext "go-service-template/internal/infrastructure/context"
	"go-service-template/internal/infrastructure/logger"
	"go-service-template/internal/infrastructure/provider/redis"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	appName string
	redis   *redis.Provider
}

func NewHealthHandler(cfg config.Provider, redisProvider *redis.Provider) *HealthHandler {
	return &HealthHandler{
		appName: cfg.GetAppName(),
		redis:   redisProvider,
	}
}

// Check returns a legacy health status payload.
// @Title Health check
// @Description Legacy health check endpoint.
// @Produce json
// @Success 200 {object} dto.HealthResponse "Service is healthy"
// @Resource health
// @Router /health [get]
func (h *HealthHandler) Check(ctx *gincontext.GinContext) {
	h.respondOK(ctx)
}

// Live returns a liveness probe response.
// @Title Liveness probe
// @Description Always returns OK when the process is running.
// @Produce json
// @Success 200 {object} dto.HealthResponse "Process is alive"
// @Resource health
// @Router /live [get]
func (h *HealthHandler) Live(ctx *gincontext.GinContext) {
	h.respondOK(ctx)
}

// Ready returns a readiness probe response, checking Redis when configured.
// @Title Readiness probe
// @Description Returns OK when dependencies are available; 503 when Redis is down.
// @Produce json
// @Success 200 {object} dto.HealthResponse "Service is ready"
// @Failure 503 {object} dto.NotReadyResponse "Dependency unavailable"
// @Resource health
// @Router /ready [get]
func (h *HealthHandler) Ready(ctx *gincontext.GinContext) {
	logCtx := logger.GetLogContext(ctx.Context)

	if h.redis != nil {
		pingCtx, cancel := context.WithTimeout(ctx.Context, redisPingTimeout)
		defer cancel()

		if err := h.redis.Ping(pingCtx); err != nil {
			logger.Error(logCtx, "Readiness check failed", logger.ErrorField(logger.FieldError, err))
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "not_ready",
				"service": h.appName,
				"error":   "redis unavailable",
			})
			return
		}
	}

	h.respondOK(ctx)
}

func (h *HealthHandler) respondOK(ctx *gincontext.GinContext) {
	logCtx := logger.GetLogContext(ctx.Context)
	logger.Info(logCtx, "Health check requested", logger.Int(logger.FieldStatusCode, http.StatusOK))

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": h.appName,
	})
}

const redisPingTimeout = 5 * time.Second
