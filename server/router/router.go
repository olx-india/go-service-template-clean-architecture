package router

import (
	"context"
	"go-service-template/internal/api"
	"go-service-template/internal/infrastructure/config"
	gincontext "go-service-template/internal/infrastructure/context"
	"go-service-template/internal/infrastructure/logger"
	"go-service-template/server/resolver"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	RegisterRoutes(*resolver.ServerContext) *Router
	Get() *gin.Engine
}

type Router struct {
	*gin.Engine
	config config.Provider
}

func NewRouter(cfg config.Provider) *Router {
	ctx := context.Background()
	logger.Info(ctx, "Setting up endpoints...")

	router := &Router{
		Engine: gin.New(),
		config: cfg,
	}

	router.Use(logger.LoggingMiddleware())
	router.Use(corsMiddleware())
	return router
}

func (r *Router) RegisterRoutes(serverContext *resolver.ServerContext) *Router {
	healthHandler := api.NewHealthHandler()
	r.GET("/health", WrapContext(healthHandler.Check))
	r.POST("/api/v1/limit/check", WrapContext(serverContext.LimiterHandler.CheckLimit))
	r.POST("/api/v1/limit/reset", WrapContext(serverContext.LimiterHandler.ResetLimit))
	r.POST("/api/v1/user", WrapContext(serverContext.UserHandler.CreateUser))
	r.GET("/api/v1/user/:id", WrapContext(serverContext.UserHandler.FetchUser))
	return r
}

func (r *Router) Get() *gin.Engine {
	return r.Engine
}

func WrapContext(handler func(ctx *gincontext.GinContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, err := gincontext.NewGinContext(c)
		if err != nil {
			logger.Error(context.Background(), "gin context failed",
				logger.String("error", err.Error()),
			)
			return
		}
		handler(ctx)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
