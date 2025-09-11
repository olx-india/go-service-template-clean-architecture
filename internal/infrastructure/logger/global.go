package logger

import (
	"context"
)

// InitGlobalLogger initializes the global logger instance.
// This function is kept for backward compatibility but doesn't do anything
// since we now create logger instances on demand.
func InitGlobalLogger() {
	// No-op: logger instances are created on demand
}

// GetGlobalLogger returns a new logger instance.
// This replaces the global logger pattern to avoid global variables.
func GetGlobalLogger() Logger {
	return NewLogger(serviceName)
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
