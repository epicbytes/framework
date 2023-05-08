package redis

import (
	"github.com/epicbytes/framework/config"
	"github.com/go-redis/redis/v8"
)

func New(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.URI,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})
}
