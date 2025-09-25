package config

import "time"

type Provider interface {
	GetServerHost() string
	GetServerPort() string
	GetServerReadTimeout() time.Duration
	GetServerWriteTimeout() time.Duration
	GetRedisHost() string
	GetEnv() string
	GetAppName() string
	GetOTLPEndpoint() string
}

var _ Provider = (*Config)(nil)

func (c *Config) GetServerHost() string {
	return c.Server.Host
}

func (c *Config) GetServerPort() string {
	return c.Server.Port
}

func (c *Config) GetServerReadTimeout() time.Duration {
	return c.Server.ReadTimeout
}

func (c *Config) GetServerWriteTimeout() time.Duration {
	return c.Server.WriteTimeout
}

func (c *Config) GetRedisHost() string {
	return c.Redis.Host
}

func (c *Config) GetEnv() string {
	return c.Env
}

func (c *Config) GetAppName() string {
	return c.Server.AppName
}

func (c *Config) GetOTLPEndpoint() string {
	return c.Tracer.OTLPEndpoint
}
