package app

import (
	"context"
	"net/http"
	"os"

	"go-service-template/internal/infrastructure/config"
	"go-service-template/internal/infrastructure/logger"
	"go-service-template/internal/infrastructure/tracer"
	"go-service-template/server/resolver"
	"go-service-template/server/router"
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

	shutdown := tracer.InitTracer(ctx, app.config)
	defer shutdown()

	serverContext := resolver.NewResolver(app.config).ResolveServerContext()
	r := router.NewRouter(app.config).
		RegisterRoutes(serverContext).
		Get()

	logger.Info(ctx, "Starting HTTP server",
		logger.String("host", app.config.GetServerHost()),
		logger.String("port", app.config.GetServerPort()),
	)

	server := &http.Server{
		Addr:         app.config.GetServerHost() + ":" + app.config.GetServerPort(),
		Handler:      r,
		ReadTimeout:  app.config.GetServerReadTimeout(),
		WriteTimeout: app.config.GetServerWriteTimeout(),
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error(ctx, "Error starting server",
			logger.String("error", err.Error()),
		)
	}
	os.Exit(0)
}
