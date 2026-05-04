package logger

import (
	"context"
	"sync"
)

var (
	globalLogger Logger
	globalOnce   sync.Once
)

// InitGlobalLogger initializes the global logger singleton.
// Safe to call from main or tests; subsequent calls are no-ops.
func InitGlobalLogger() {
	globalOnce.Do(func() {
		globalLogger = NewLogger(serviceName)
	})
}

// GetGlobalLogger returns the global logger singleton, lazily initialising it on first use
// if InitGlobalLogger was not called earlier.
func GetGlobalLogger() Logger {
	globalOnce.Do(func() {
		globalLogger = NewLogger(serviceName)
	})
	return globalLogger
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
