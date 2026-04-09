package service

import (
	"context"
	"realtime-location/pkg/redis"

	redis2 "github.com/redis/go-redis/v9"
)

var notifier Notifier

func SetNotifier(n Notifier) {
	notifier = n
}

var ctx = context.Background()

func UpdatePlayerLocation(playerID string, lat, lng float64) {

	// 1. 存入 Redis GEO
	redis.Client.GeoAdd(ctx, "players", &redis2.GeoLocation{
		Name:      playerID,
		Latitude:  lat,
		Longitude: lng,
	})

	// 2. 查詢附近玩家 (300m)
	players, _ := redis.Client.GeoRadius(ctx, "players", lng, lat, &redis2.GeoRadiusQuery{
		Radius: 300,
		Unit:   "m",
	}).Result()

	// 3. 轉換成 targets
	var targets []string

	for _, p := range players {
		if p.Name == playerID {
			continue
		}
		targets = append(targets, p.Name)
	}

	// 4. 廣播
	notifier.Broadcast(playerID, lat, lng, targets)

	CheckLandmarks(playerID, lat, lng)
}
