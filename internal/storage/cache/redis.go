package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func ConnectToRedis(ctx context.Context) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic("Could not connect to Redis")
	}
}
