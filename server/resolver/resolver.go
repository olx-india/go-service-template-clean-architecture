package resolver

import (
	"context"

	"github.com/olx-india/go-service-template-clean-architecture/internal/api"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/config"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/logger"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/provider/redis"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/repo"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/repo/persistent"
	"github.com/olx-india/go-service-template-clean-architecture/internal/infrastructure/repo/webapi"
	"github.com/olx-india/go-service-template-clean-architecture/internal/usecase/limit"
	"github.com/olx-india/go-service-template-clean-architecture/internal/usecase/user"
)

type Resolver interface {
	ResolveServerContext() *ServerContext
}

type ServerContext struct {
	UserHandler    api.IUserHandler
	LimiterHandler api.ILimiterHandler
	RedisProvider  *redis.Provider
}

type Provider struct {
	redisProvider      *redis.Provider
	userRepo           repo.UserRepo
	userWebAPIProvider repo.UserWebAPI
}

type resolver struct {
	*Provider
	*ServerContext
	config config.Provider
}

func NewResolver(cfg config.Provider) Resolver {
	return &resolver{
		Provider:      &Provider{},
		ServerContext: &ServerContext{},
		config:        cfg,
	}
}

func (r *resolver) ResolveServerContext() *ServerContext {
	ctx := context.Background()
	resolver := r.resolveProviders()
	if resolver == nil {
		logger.Error(ctx, "Failed to resolve providers - continuing without Redis")
		resolver = r
	}

	resolver = resolver.
		repositories().
		provider().
		createServerContext()

	return resolver.ServerContext
}

func (r *resolver) createServerContext() *resolver {
	// Adopter: wire your handlers and use cases here.
	r.UserHandler = api.NewUserHandler(user.NewUserUseCase(r.userRepo, r.userWebAPIProvider))
	r.LimiterHandler = api.NewLimiterHandler(limit.NewLimitUseCase(r.redisProvider))
	r.RedisProvider = r.redisProvider
	return r
}

func (r *resolver) resolveProviders() *resolver {
	ctx := context.Background()
	redisProvider, err := redis.NewProvider(r.config)
	if err != nil {
		logger.Errorf(ctx, "Failed to create redis provider, Error %s: ", err.Error())
		logger.Error(ctx, "Please check your Redis configuration and ensure Redis is running")
		logger.Error(ctx, "For local development, you can:")
		logger.Error(ctx, "1. Start Redis locally: docker run -d -p 6379:6379 redis:alpine")
		logger.Error(ctx, "2. Or set REDIS_HOST=<host>:<port> in your environment")
		return nil
	}
	r.redisProvider = redisProvider
	logger.Info(ctx, "Redis provider initialized successfully")
	return r
}

func (r *resolver) repositories() *resolver {
	r.userRepo = persistent.NewUserRepo()
	return r
}

func (r *resolver) provider() *resolver {
	r.userWebAPIProvider = webapi.NewUserWebAPI()
	return r
}
