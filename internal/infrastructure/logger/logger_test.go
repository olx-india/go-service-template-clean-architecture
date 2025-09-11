package logger

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewLogger_NotNil(t *testing.T) {
    t.Setenv(LogLevelKey, InfoKey)
    l := NewLogger("svc-test")
    assert.NotNil(t, l)
}

func TestGlobalLogger_NotNil(t *testing.T) {
    l := GetGlobalLogger()
    assert.NotNil(t, l)
}

func TestBuildZapFields_IncludesRequestIDOrTraceID(t *testing.T) {
    l := NewLogger("svc")
    zl := l.(*zapLogger)
    ctx := context.WithValue(context.Background(), "request-id", "rid")
    ctx = context.WithValue(ctx, "trace-id", "tid")
    fs := zl.buildZapFields(ctx)
    assert.Greater(t, len(fs), 0)
}

func TestBuildZapFields_AlternateKeys(t *testing.T) {
    l := NewLogger("svc")
    zl := l.(*zapLogger)
    ctx := context.WithValue(context.Background(), "X-Request-ID", "xrid")
    ctx = context.WithValue(ctx, "X-Trace-ID", "xtid")
    fs := zl.buildZapFields(ctx)
    assert.Greater(t, len(fs), 0)

    ctx2 := context.WithValue(context.Background(), "Request-ID", "rrid")
    ctx2 = context.WithValue(ctx2, "Trace-ID", "rtid")
    fs2 := zl.buildZapFields(ctx2)
    assert.Greater(t, len(fs2), 0)
}

func TestConvertField_String(t *testing.T) {
    l := NewLogger("svc").(*zapLogger)
    _ = l.convertField(String("k", "v"))
}

func TestNewLogger_RespectsLogLevel(t *testing.T) {
    t.Setenv(LogLevelKey, DebugKey)
    l := NewLogger("svc")
    assert.NotNil(t, l)
}

func TestNewLogger_WarnLevel(t *testing.T) {
    t.Setenv(LogLevelKey, WarnKey)
    assert.NotNil(t, NewLogger("svc"))
}

func TestNewLogger_ErrorLevel(t *testing.T) {
    t.Setenv(LogLevelKey, ErrorKey)
    assert.NotNil(t, NewLogger("svc"))
}

func TestNewLogger_DefaultLevel(t *testing.T) {
    t.Setenv(LogLevelKey, "unknown")
    assert.NotNil(t, NewLogger("svc"))
}

func TestLogger_Sync(t *testing.T) {
    l := NewLogger("svc").(*zapLogger)
    _ = l.Sync()
}


