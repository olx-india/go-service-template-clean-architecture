package app

import (
	"context"
	"net/http"

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
	shutdown := telemetry.InitTracer(ctx, app.config)
	defer shutdown()

	serverContext := resolver.NewResolver(app.config).ResolveServerContext()
	router := app.createRouterAndRegisterRoutes(serverContext)

	logger.Info(ctx, "Starting HTTP server",
		logger.String("host", app.config.GetServerHost()),
		logger.String("port", app.config.GetServerPort()),
	)

	server := &http.Server{
		Addr:         app.config.GetServerHost() + ":" + app.config.GetServerPort(),
		Handler:      router,
		ReadTimeout:  app.config.GetServerReadTimeout(),
		WriteTimeout: app.config.GetServerWriteTimeout(),
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error(ctx, "Error starting server",
			logger.String("error", err.Error()),
		)
	}
}

func (app *App) createRouterAndRegisterRoutes(serverContext *resolver.ServerContext) *gin.Engine {
	r := routerPkg.NewRouter(app.config).
		RegisterRoutes(serverContext).
		Get()
	return r
}
