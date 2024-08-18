package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var rc *redis.Client
var ctx context.Context

type RedisClient struct {
}

func NewRedisClient() *RedisClient {
	redisClient := RedisClient{}

	if rc == nil {
		rc = redis.NewClient(&redis.Options{Addr: "localhost:6379",
			Password: "",
			DB:       0})
		ctx = context.Background()
	}

	err := rc.Set(ctx, "health", "GOOD", 0).Err()
	if err != nil {
		panic(err)
	}

	return &redisClient
}

func (r *RedisClient) AddKey() {
	err := rc.Set(ctx, "health3", "ok", 0).Err()
	if err != nil {
		panic(err)
	}

}

func (r *RedisClient) GetKey() string {
	val, err := rc.Get(ctx, "health").Result()
	if err != nil {
		panic(err)
	}
	return val

}
