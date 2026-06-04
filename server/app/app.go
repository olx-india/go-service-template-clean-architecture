package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-service-template/internal/infrastructure/config"
	"go-service-template/internal/infrastructure/logger"
	"go-service-template/server/resolver"
	routerPkg "go-service-template/server/router"
	"go-service-template/server/telemetry"

	"github.com/gin-gonic/gin"
)

type App struct {
	config config.Provider
}

func NewApp() *App {
	return &App{
		config: config.NewConfig(),
	}
}

func (app *App) Start() {
	ctx := context.Background()
	shutdownTracer := telemetry.InitTracer(ctx, app.config)
	defer shutdownTracer()

	serverContext := resolver.NewResolver(app.config).ResolveServerContext()
	router := app.createRouterAndRegisterRoutes(serverContext)

	server := &http.Server{
		Addr:         app.config.GetServerHost() + ":" + app.config.GetServerPort(),
		Handler:      router,
		ReadTimeout:  app.config.GetServerReadTimeout(),
		WriteTimeout: app.config.GetServerWriteTimeout(),
	}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info(ctx, "Starting HTTP server",
			logger.String("host", app.config.GetServerHost()),
			logger.String("port", app.config.GetServerPort()),
		)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		logger.Error(ctx, "Error starting server", logger.String("error", err.Error()))
	case sig := <-quit:
		logger.Info(ctx, "Shutdown signal received", logger.String("signal", sig.String()))
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error(ctx, "Server shutdown failed", logger.String("error", err.Error()))
	}

	if serverContext.RedisProvider != nil {
		if err := serverContext.RedisProvider.Close(); err != nil {
			logger.Error(ctx, "Redis shutdown failed", logger.String("error", err.Error()))
		}
	}
}

func (app *App) createRouterAndRegisterRoutes(serverContext *resolver.ServerContext) *gin.Engine {
	r := routerPkg.NewRouter(app.config).
		RegisterRoutes(serverContext).
		Get()
	return r
}

const shutdownTimeout = 30 * time.Second
