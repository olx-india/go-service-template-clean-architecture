package logger

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoggingMiddleware creates a middleware that adds structured logging to all requests.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID, traceID := fetchRequestAndTraceIDs(c)
		addTraceIDsInResponseHeaders(c, requestID, traceID)
		ctx := createContextWithTraceIDs(requestID, traceID)
		setLoggerInContext(ctx, c)

		requestStartLog(ctx, c, requestID, traceID)

		c.Next()
		requestCompleteLog(ctx, c, start, requestID, traceID)
	}
}

func fetchRequestAndTraceIDs(c *gin.Context) (requestID, traceID string) {
	requestID = c.GetHeader(XRequestID)
	if requestID == "" {
		requestID = uuid.New().String()
	}
	traceID = c.GetHeader(XTraceID)
	if traceID == "" {
		traceID = uuid.New().String()
	}
	return requestID, traceID
}

func addTraceIDsInResponseHeaders(c *gin.Context, requestID, traceID string) {
	c.Header(XRequestID, requestID)
	c.Header(XTraceID, traceID)
}

func createContextWithTraceIDs(requestID, traceID string) context.Context {
	ctx := context.WithValue(context.Background(), RequestID, requestID)
	ctx = context.WithValue(ctx, TraceID, traceID)
	return ctx
}

func setLoggerInContext(ctx context.Context, c *gin.Context) {
	c.Set(LogContext, ctx)
}

func requestStartLog(ctx context.Context, c *gin.Context, requestID, traceID string) {
	GetGlobalLogger().Info(ctx, "Request started",
		String(FieldRequestID, requestID),
		String(FieldTraceID, traceID),
		String(FieldMethod, c.Request.Method),
		String(FieldPath, c.Request.URL.Path),
		String(FieldIP, c.ClientIP()),
		String(FieldUserAgent, c.Request.UserAgent()),
	)
}

func requestCompleteLog(ctx context.Context, c *gin.Context, start time.Time, requestID, traceID string) {
	duration := time.Since(start)
	GetGlobalLogger().Info(ctx, "Request completed",
		String(FieldRequestID, requestID),
		String(FieldTraceID, traceID),
		String(FieldMethod, c.Request.Method),
		String(FieldPath, c.Request.URL.Path),
		Int(FieldStatusCode, c.Writer.Status()),
		Int64(FieldDuration, duration.Milliseconds()),
		Int(FieldResponseSize, c.Writer.Size()),
	)
}

// GetLogContext extracts the logging context from Gin context.
func GetLogContext(c *gin.Context) context.Context {
	if ctx, exists := c.Get(LogContext); exists {
		if logCtx, ok := ctx.(context.Context); ok {
			return logCtx
		}
	}
	return context.Background()
}

// Custom types for context keys to avoid using basic types.
type contextKey string

const (
	XRequestID = "X-Request-ID"
	XTraceID   = "X-Trace-ID"
	RequestID  = contextKey("request-id")
	TraceID    = contextKey("trace-id")
	LogContext = "log-context"
)
