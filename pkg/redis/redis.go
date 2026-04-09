package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	Client.Ping(context.Background())
}
