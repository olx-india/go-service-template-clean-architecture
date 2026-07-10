package main

import (
	"context"

	"go-service-template/internal/infrastructure/config"
	"go-service-template/internal/infrastructure/logger"
	appPkg "go-service-template/server/app"
)

// @Title Go Service Template API
// @Version 1.0.0
// @Description Clean Architecture Go service scaffold with sample user and limit endpoints.
// @ContactName API Support
// @LicenseName Apache 2.0
// @LicenseURL https://www.apache.org/licenses/LICENSE-2.0
// @Server http://localhost:8080 Local
func main() {
	cfg := config.NewConfig()

	logger.InitGlobalLogger()
	logger.Info(context.Background(), "Starting application", LogAppName(cfg), LogEnv(cfg))

	app := appPkg.NewApp()
	app.Start()
}

func LogEnv(cfg *config.Config) logger.Field {
	return logger.String("environment", cfg.GetEnv())
}

func LogAppName(cfg *config.Config) logger.Field {
	return logger.String("app_name", cfg.GetAppName())
}
