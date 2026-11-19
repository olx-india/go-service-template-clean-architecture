package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
	Env    string
}

type ServerConfig struct {
	Host               string
	Port               string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	AppName            string
	OTLPEndpoint       string
	OTELEnabled        bool
	CORSAllowedOrigins []string
}

type RedisConfig struct {
	Host string
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:         getEnv(EnvHost, DefaultHost),
			Port:         getEnv(EnvPort, DefaultPort),
			ReadTimeout:  getEnvAsDuration(EnvReadTimeout, DefaultReadTimeout),
			WriteTimeout: getEnvAsDuration(EnvWriteTimeout, DefaultWriteTimeout),
			AppName:      getEnv(EnvAppName, DefaultAppName),
			OTLPEndpoint: getEnv(EnvOLTPEndpoint, DefaultOTLPEndpoint),
			OTELEnabled:  getEnvAsBool(EnvOTELEnabled, DefaultOTELEnabled),
			CORSAllowedOrigins: getEnvAsStringSlice(
				EnvCORSAllowedOrigins,
				DefaultCORSAllowedOrigins,
			),
		},
		Redis: RedisConfig{
			Host: getEnv(EnvRedisHost, DefaultRedisHost),
		},
		Env: getEnv(EnvEnvironment, DefaultEnv),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != EmptyString {
		return value
	}
	return fallback
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	if value := os.Getenv(key); value != EmptyString {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != EmptyString {
		return value == "true" || value == "1"
	}
	return fallback
}

func getEnvAsStringSlice(key, fallback string) []string {
	raw := getEnv(key, EmptyString)
	if raw == EmptyString {
		return []string{fallback}
	}

	parts := strings.Split(raw, ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		if value := strings.TrimSpace(part); value != EmptyString {
			values = append(values, value)
		}
	}
	if len(values) == 0 {
		log.Printf("warning: invalid %s value %s; using default %q", key, strconv.Quote(raw), fallback)
		return []string{fallback}
	}
	return values
}

const (
	EmptyString = ""
)

const (
	EnvHost               = "HOST"
	EnvPort               = "PORT"
	EnvReadTimeout        = "READ_TIMEOUT"
	EnvWriteTimeout       = "WRITE_TIMEOUT"
	EnvRedisHost          = "REDIS_HOST"
	EnvEnvironment        = "ENV"
	EnvAppName            = "APP_NAME"
	EnvOLTPEndpoint       = "OTLP_ENDPOINT"
	EnvOTELEnabled        = "OTEL_ENABLED"
	EnvCORSAllowedOrigins = "CORS_ALLOWED_ORIGINS"
)

const (
	DefaultHost               = "0.0.0.0"
	DefaultPort               = "8080"
	DefaultReadTimeout        = 60 * time.Second
	DefaultWriteTimeout       = 60 * time.Second
	DefaultRedisHost          = "localhost:6379"
	DefaultAppName            = "go-service-template"
	DefaultOTLPEndpoint       = "localhost:4317"
	DefaultOTELEnabled        = false
	DefaultEnv                = "local"
	DefaultCORSAllowedOrigins = "*"
)
