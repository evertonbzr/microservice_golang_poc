package recache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type service struct {
	c *redis.Client
}

func New(options *redis.Options) *service {
	redisClient := redis.NewClient(options)
	return &service{
		c: redisClient,
	}
}

func (s *service) Get(ctx context.Context, key string) (string, error) {
	status := s.c.Get(ctx, key)
	return status.Result()
}

func (s *service) Put(ctx context.Context, key string, value interface{}) error {
	status := s.c.Set(ctx, key, value, time.Minute*10)
	return status.Err()
}
