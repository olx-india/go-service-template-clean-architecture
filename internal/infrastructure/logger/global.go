package logger

import (
	"context"
	"sync"
)

//nolint:gochecknoglobals // Lazy singleton for package-level logger API; sync.OnceValue must be shared across calls.
var globalLoggerFn = sync.OnceValue(func() Logger {
	return NewLogger(serviceName)
})

// InitGlobalLogger initializes the global logger singleton.
// Safe to call from main or tests; subsequent calls are no-ops.
func InitGlobalLogger() {
	_ = globalLoggerFn()
}

// GetGlobalLogger returns the global logger singleton, lazily initializing it on first use
// if InitGlobalLogger was not called earlier.
func GetGlobalLogger() Logger {
	return globalLoggerFn()
}

// Info Global logging functions for convenience.
func Info(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Info(ctx, message, fields...)
}

func Warn(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Warn(ctx, message, fields...)
}

func Error(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Error(ctx, message, fields...)
}

func Debug(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Debug(ctx, message, fields...)
}

func Fatal(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Fatal(ctx, message, fields...)
}

// Infof Global logging functions for convenience.
func Infof(ctx context.Context, format string, args ...interface{}) {
	GetGlobalLogger().Infof(ctx, format, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	GetGlobalLogger().Warnf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	GetGlobalLogger().Errorf(ctx, format, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	GetGlobalLogger().Debugf(ctx, format, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	GetGlobalLogger().Fatalf(ctx, format, args...)
}

const serviceName = "go-service-template"
