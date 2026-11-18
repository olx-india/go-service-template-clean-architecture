package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
	Env    string
}

type ServerConfig struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	AppName      string
	OTLPEndpoint string
	OTELEnabled  bool
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
			OTLPEndpoint: getEnv(EnvOTLPEndpoint, DefaultOTLPEndpoint),
			OTELEnabled:  getEnvAsBool(EnvOTELEnabled, DefaultOTELEnabled),
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
		log.Printf("warning: invalid %s value %s; using default %s", key, strconv.Quote(value), fallback)
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != EmptyString {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
		log.Printf("warning: invalid %s value %s; using default %t", key, strconv.Quote(value), fallback)
	}
	return fallback
}

const (
	EmptyString = ""
)

const (
	EnvHost         = "HOST"
	EnvPort         = "PORT"
	EnvReadTimeout  = "READ_TIMEOUT"
	EnvWriteTimeout = "WRITE_TIMEOUT"
	EnvRedisHost    = "REDIS_HOST"
	EnvEnvironment  = "ENV"
	EnvAppName      = "APP_NAME"
	EnvOTLPEndpoint = "OTLP_ENDPOINT"
	EnvOTELEnabled  = "OTEL_ENABLED"
)

const (
	DefaultHost         = "0.0.0.0"
	DefaultPort         = "8080"
	DefaultReadTimeout  = 60 * time.Second
	DefaultWriteTimeout = 60 * time.Second
	DefaultRedisHost    = "localhost:6379"
	DefaultAppName      = "go-service-template"
	DefaultOTLPEndpoint = "localhost:4317"
	DefaultOTELEnabled  = false
	DefaultEnv          = "local"
)
