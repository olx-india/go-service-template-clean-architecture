package logger

import (
	"context"
)

// Logger defines the interface for structured logging.
type Logger interface {
	// Info logs an info level message
	Info(ctx context.Context, message string, fields ...Field)

	// Infof logs an info level message with formatting
	Infof(ctx context.Context, format string, args ...interface{})

	// Warn logs a warning level message
	Warn(ctx context.Context, message string, fields ...Field)

	// Warnf logs a warning level message with formatting
	Warnf(ctx context.Context, format string, args ...interface{})

	// Error logs an error level message
	Error(ctx context.Context, message string, fields ...Field)

	// Errorf logs an error level message with formatting
	Errorf(ctx context.Context, format string, args ...interface{})

	// Debug logs a debug level message
	Debug(ctx context.Context, message string, fields ...Field)

	// Debugf logs a debug level message with formatting
	Debugf(ctx context.Context, format string, args ...interface{})

	// Fatal logs a fatal level message and exits the program
	Fatal(ctx context.Context, message string, fields ...Field)

	// Fatalf logs a fatal level message with formatting and exits the program
	Fatalf(ctx context.Context, format string, args ...interface{})
}

// The Field represents a key-value pair for structured logging.
type Field struct {
	Key   string
	Value interface{}
}

// String Field constructors for common field types.
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

func Int64(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

func Float64(key string, value float64) Field {
	return Field{Key: key, Value: value}
}

func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

func Any(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

func ErrorField(key string, err error) Field {
	return Field{Key: key, Value: err}
}

// Common field keys for structured logging.
const (
	FieldService      = "service"
	FieldRequestID    = "request-id"
	FieldTraceID      = "trace-id"
	FieldMessage      = "message"
	FieldSeverity     = "severity"
	FieldError        = "error"
	FieldDuration     = "duration"
	FieldMethod       = "method"
	FieldPath         = "path"
	FieldStatusCode   = "status_code"
	FieldUserID       = "user_id"
	FieldUserAgent    = "user_agent"
	FieldIP           = "ip"
	FieldTimeStamp    = "timestamp"
	FieldCaller       = "caller"
	FieldStackTrace   = "stacktrace"
	FieldResponseSize = "response_size"
)
