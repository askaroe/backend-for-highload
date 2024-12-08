package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func ConnectToRedis(ctx context.Context) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis address
		DB:   0,                // Default DB
	})
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic("Could not connect to Redis")
	}
}
