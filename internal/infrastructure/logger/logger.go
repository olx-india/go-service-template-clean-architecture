package logger

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger implements the Logger interface using zap.
type zapLogger struct {
	logger *zap.Logger
	fields []Field
}

// NewLogger creates a new logger instance.
func NewLogger(serviceName string) Logger {
	config := zap.NewProductionConfig()

	// Configure JSON output with required fields
	config.EncoderConfig.TimeKey = FieldTimeStamp
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.LevelKey = FieldSeverity
	config.EncoderConfig.MessageKey = FieldMessage
	config.EncoderConfig.CallerKey = FieldCaller
	config.EncoderConfig.StacktraceKey = FieldStackTrace

	// Set log level based on environment
	env := strings.ToLower(os.Getenv(LogLevelKey))
	switch env {
	case DebugKey:
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case InfoKey:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case WarnKey:
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case ErrorKey:
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// Build logger with service name
	logger, err := config.Build(
		zap.AddCallerSkip(1), // Skip this function in caller stack
		zap.Fields(zap.String(FieldService, serviceName)),
	)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return &zapLogger{
		logger: logger,
		fields: []Field{},
	}
}

// Info logs an info level message.
func (l *zapLogger) Info(ctx context.Context, message string, fields ...Field) {
	zapFields := l.buildZapFields(ctx, fields...)
	l.logger.Info(message, zapFields...)
}

// Infof logs an info level message with formatting.
func (l *zapLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	zapFields := l.buildZapFields(ctx)
	message := fmt.Sprintf(format, args...)
	l.logger.Info(message, zapFields...)
}

// Warn logs a warning level message.
func (l *zapLogger) Warn(ctx context.Context, message string, fields ...Field) {
	zapFields := l.buildZapFields(ctx, fields...)
	l.logger.Warn(message, zapFields...)
}

// Warnf logs a warning level message with formatting.
func (l *zapLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	zapFields := l.buildZapFields(ctx)
	message := fmt.Sprintf(format, args...)
	l.logger.Warn(message, zapFields...)
}

// Error logs an error level message.
func (l *zapLogger) Error(ctx context.Context, message string, fields ...Field) {
	zapFields := l.buildZapFields(ctx, fields...)
	l.logger.Error(message, zapFields...)
}

// Errorf logs an error level message with formatting.
func (l *zapLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	zapFields := l.buildZapFields(ctx)
	message := fmt.Sprintf(format, args...)
	l.logger.Error(message, zapFields...)
}

// Debug logs a debug level message.
func (l *zapLogger) Debug(ctx context.Context, message string, fields ...Field) {
	zapFields := l.buildZapFields(ctx, fields...)
	l.logger.Debug(message, zapFields...)
}

// Debugf logs a debug level message with formatting.
func (l *zapLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	zapFields := l.buildZapFields(ctx)
	message := fmt.Sprintf(format, args...)
	l.logger.Debug(message, zapFields...)
}

// Fatal logs a fatal level message and exits the program.
func (l *zapLogger) Fatal(ctx context.Context, message string, fields ...Field) {
	zapFields := l.buildZapFields(ctx, fields...)
	l.logger.Fatal(message, zapFields...)
}

// Fatalf logs a fatal level message with formatting and exits the program.
func (l *zapLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	zapFields := l.buildZapFields(ctx)
	message := fmt.Sprintf(format, args...)
	l.logger.Fatal(message, zapFields...)
}

// buildZapFields converts Field slice to zap.Field slice and adds context fields.
func (l *zapLogger) buildZapFields(ctx context.Context, fields ...Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(l.fields)+len(fields)+TwoForRequestIDAndTraceID) // +2 for request-id and trace-id

	// Add existing fields.
	for _, field := range l.fields {
		zapFields = append(zapFields, l.convertField(field))
	}

	// Add context fields (request-id and trace-id).
	if requestID := l.getRequestID(ctx); requestID != "" {
		zapFields = append(zapFields, zap.String(FieldRequestID, requestID))
	}

	if traceID := l.getTraceID(ctx); traceID != "" {
		zapFields = append(zapFields, zap.String(FieldTraceID, traceID))
	}

	// Add new fields.
	for _, field := range fields {
		zapFields = append(zapFields, l.convertField(field))
	}

	return zapFields
}

// convertField converts a Field to a zap.Field.
func (l *zapLogger) convertField(field Field) zap.Field {
	switch v := field.Value.(type) {
	case string:
		return zap.String(field.Key, v)
	case int:
		return zap.Int(field.Key, v)
	case int64:
		return zap.Int64(field.Key, v)
	case float64:
		return zap.Float64(field.Key, v)
	case bool:
		return zap.Bool(field.Key, v)
	case error:
		return zap.Error(v)
	default:
		return zap.Any(field.Key, v)
	}
}

// getRequestID extracts request ID from context.
func (l *zapLogger) getRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	// Try to get request ID from context.
	if requestID, ok := ctx.Value("request-id").(string); ok {
		return requestID
	}

	// Try common request ID header names
	if requestID, ok := ctx.Value("X-Request-ID").(string); ok {
		return requestID
	}

	if requestID, ok := ctx.Value("Request-ID").(string); ok {
		return requestID
	}

	return ""
}

// getTraceID extracts trace ID from context.
func (l *zapLogger) getTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	// Extract trace ID from OpenTelemetry context
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		return span.SpanContext().TraceID().String()
	}

	// Try to get trace ID from context.
	if traceID, ok := ctx.Value("trace-id").(string); ok {
		return traceID
	}

	// Try common trace ID header names.
	if traceID, ok := ctx.Value("X-Trace-ID").(string); ok {
		return traceID
	}

	if traceID, ok := ctx.Value("Trace-ID").(string); ok {
		return traceID
	}

	return ""
}

// Sync flushes any buffered log entries.
func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}

const (
	DebugKey    = "debug"
	InfoKey     = "info"
	WarnKey     = "warn"
	ErrorKey    = "error"
	LogLevelKey = "LOG_LEVEL"
)

const TwoForRequestIDAndTraceID = 2
