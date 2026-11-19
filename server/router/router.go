package router

import (
	"context"
	"net/http"

	"github.com/olx-india/go-service-template-clean-architecture/docs/openapi"
	"github.com/olx-india/go-service-template-clean-architecture/internal/api"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/config"
	gincontext "github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/context"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/logger"
	"github.com/olx-india/go-service-template-clean-architecture/server/resolver"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
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

	router.Use(gin.Recovery())
	router.Use(logger.LoggingMiddleware())
	router.Use(otelgin.Middleware(cfg.GetAppName()))
	router.Use(corsMiddleware(cfg.GetCORSAllowedOrigins()))
	return router
}

func (r *Router) RegisterRoutes(serverContext *resolver.ServerContext) *Router {
	healthHandler := api.NewHealthHandler(r.config, serverContext.RedisProvider)
	r.GET("/health", WrapContext(healthHandler.Check))
	r.GET("/live", WrapContext(healthHandler.Live))
	r.GET("/ready", WrapContext(healthHandler.Ready))
	r.POST("/api/v1/limit/check", WrapContext(serverContext.LimiterHandler.CheckLimit))
	r.POST("/api/v1/limit/reset", WrapContext(serverContext.LimiterHandler.ResetLimit))
	r.POST("/api/v1/user", WrapContext(serverContext.UserHandler.CreateUser))
	r.GET("/api/v1/user/:id", WrapContext(serverContext.UserHandler.FetchUser))
	r.Any("/swagger/*any", gin.WrapH(openapi.Handler()))
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

func corsMiddleware(allowedOrigins []string) gin.HandlerFunc {
	allowAll := contains(allowedOrigins, "*")
	allowed := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		if origin != "*" {
			allowed[origin] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		origin := c.GetHeader("Origin")
		if allowAll {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else if origin != "" {
			if _, ok := allowed[origin]; ok {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}
		c.Writer.Header().Add("Vary", "Origin")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
