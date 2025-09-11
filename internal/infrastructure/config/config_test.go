package config

import (
    "os"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestNewConfig_DefaultServerHost(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultHost, c.Server.Host)
}

func TestNewConfig_DefaultServerPort(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultPort, c.Server.Port)
}

func TestNewConfig_DefaultReadTimeout(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultReadTimeout, c.Server.ReadTimeout)
}

func TestNewConfig_DefaultWriteTimeout(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultWriteTimeout, c.Server.WriteTimeout)
}

func TestNewConfig_DefaultAppName(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultAppName, c.Server.AppName)
}

func TestNewConfig_DefaultNRLicenseKey(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultNRLicenseKey, c.Server.NRLicenseKey)
}

func TestNewConfig_DefaultRedisHost(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultRedisHost, c.Redis.Host)
}

func TestNewConfig_DefaultEnv(t *testing.T) {
    os.Clearenv()
    c := NewConfig()
    assert.Equal(t, DefaultEnv, c.Env)
}

func TestNewConfig_ServerHostFromEnv(t *testing.T) {
    t.Setenv(EnvHost, "1.2.3.4")
    c := NewConfig()
    assert.Equal(t, "1.2.3.4", c.Server.Host)
}

func TestNewConfig_ServerPortFromEnv(t *testing.T) {
    t.Setenv(EnvPort, "9090")
    c := NewConfig()
    assert.Equal(t, "9090", c.Server.Port)
}

func TestNewConfig_ReadTimeoutFromEnv(t *testing.T) {
    t.Setenv(EnvReadTimeout, "5s")
    c := NewConfig()
    assert.Equal(t, 5*time.Second, c.Server.ReadTimeout)
}

func TestNewConfig_WriteTimeoutFromEnv(t *testing.T) {
    t.Setenv(EnvWriteTimeout, "7s")
    c := NewConfig()
    assert.Equal(t, 7*time.Second, c.Server.WriteTimeout)
}

func TestNewConfig_AppNameFromEnv(t *testing.T) {
    t.Setenv(EnvAppName, "svc")
    c := NewConfig()
    assert.Equal(t, "svc", c.Server.AppName)
}

func TestNewConfig_NRKeyFromEnv(t *testing.T) {
    t.Setenv(EnvNRLicenseKey, "abc")
    c := NewConfig()
    assert.Equal(t, "abc", c.Server.NRLicenseKey)
}

func TestNewConfig_RedisHostFromEnv(t *testing.T) {
    t.Setenv(EnvRedisHost, "redis:6379")
    c := NewConfig()
    assert.Equal(t, "redis:6379", c.Redis.Host)
}

func TestNewConfig_EnvFromEnv(t *testing.T) {
    t.Setenv(EnvEnvironment, "stg")
    c := NewConfig()
    assert.Equal(t, "stg", c.Env)
}

func TestProviderInterface(t *testing.T) {
    os.Clearenv()
    cfg := NewConfig()
    var p Provider = cfg
    assert.NotEmpty(t, p.GetServerHost())
}

func TestGetEnv_Fallback(t *testing.T) {
    os.Unsetenv("SOME_KEY_THAT_DOESNT_EXIST")
    v := getEnv("SOME_KEY_THAT_DOESNT_EXIST", "fallback")
    assert.Equal(t, "fallback", v)
}

func TestGetEnvAsDuration_Invalid_Fallback(t *testing.T) {
    t.Setenv(EnvReadTimeout, "not-a-duration")
    d := getEnvAsDuration(EnvReadTimeout, 3*time.Second)
    assert.Equal(t, 3*time.Second, d)
}

func TestProvider_GetServerHost_Value(t *testing.T) {
    t.Setenv(EnvHost, "h")
    c := NewConfig()
    assert.Equal(t, "h", c.GetServerHost())
}

func TestProvider_GetServerPort_Value(t *testing.T) {
    t.Setenv(EnvPort, "p")
    c := NewConfig()
    assert.Equal(t, "p", c.GetServerPort())
}

func TestProvider_GetServerReadTimeout_Value(t *testing.T) {
    t.Setenv(EnvReadTimeout, "2s")
    c := NewConfig()
    assert.Equal(t, 2*time.Second, c.GetServerReadTimeout())
}

func TestProvider_GetServerWriteTimeout_Value(t *testing.T) {
    t.Setenv(EnvWriteTimeout, "4s")
    c := NewConfig()
    assert.Equal(t, 4*time.Second, c.GetServerWriteTimeout())
}

func TestProvider_GetRedisHost_Value(t *testing.T) {
    t.Setenv(EnvRedisHost, "r")
    c := NewConfig()
    assert.Equal(t, "r", c.GetRedisHost())
}

func TestProvider_GetEnv_Value(t *testing.T) {
    t.Setenv(EnvEnvironment, "e")
    c := NewConfig()
    assert.Equal(t, "e", c.GetEnv())
}

func TestProvider_GetAppName_Value(t *testing.T) {
    t.Setenv(EnvAppName, "a")
    c := NewConfig()
    assert.Equal(t, "a", c.GetAppName())
}

func TestProvider_GetNRLicenseKey_Value(t *testing.T) {
    t.Setenv(EnvNRLicenseKey, "k")
    c := NewConfig()
    assert.Equal(t, "k", c.GetNRLicenseKey())
}

