package utils

import (
	"golang-template-api-service/app/storage"
	"fmt"
)

var cache *storage.RedisClient

type CacheClient struct {
}

func InitCacheRedis() {
	cache = storage.NewRedisClient()
}

func AddKey() {

	cache.AddKey()
}

func FetchCacheKey() string {

	fmt.Printf("Redis as %s", cache.GetKey())

	return cache.GetKey()
}
