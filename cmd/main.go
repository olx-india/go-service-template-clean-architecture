package main

import (
	"context"

	"go-service-template/internal/infrastructure/config"
	"go-service-template/internal/infrastructure/logger"
	appPkg "go-service-template/server/app"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	logger.InitGlobalLogger()
	logger.Info(ctx, "Starting application", LogAppName(cfg), LogEnv(cfg))

	app := appPkg.NewApp()
	app.Start()
}

func LogEnv(cfg *config.Config) logger.Field {
	return logger.String("environment", cfg.GetEnv())
}

func LogAppName(cfg *config.Config) logger.Field {
	return logger.String("app_name", cfg.GetAppName())
}
