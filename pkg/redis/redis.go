package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Init() {
	url := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal("Redis parse error:", err)
	}

	Client = redis.NewClient(opt)

	// 測試連線
	_, err = Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Redis connect error:", err)
	}

	log.Println("Redis connected")
}
