package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"go-service-template/internal/infrastructure/config"
	"go-service-template/internal/infrastructure/logger"
)

type Provider struct {
	client *redis.Client
}

func NewProvider(cfg config.Provider) (*Provider, error) {
	ctx := context.Background()
	logger.Info(ctx, "Connecting to Redis",
		logger.String("host", cfg.GetRedisHost()),
	)

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetRedisHost(),
		PoolSize: poolSize,
	})

	pingCtx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &Provider{
		client: client,
	}, nil
}

func (p *Provider) GetClient() *redis.Client {
	return p.client
}

func (p *Provider) Close() error {
	return p.client.Close()
}

const (
	contextTimeout = 5 * time.Second
	poolSize       = 10
)
