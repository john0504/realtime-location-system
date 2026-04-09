# Realtime Location System

## 啟動方式

1. 啟動 Redis
   docker-compose up -d

2. 啟動 Server
   go run cmd/main.go

## API

POST /update-location
{
  "player_id": "A",
  "lat": 24.1,
  "lng": 120.6
}

## WebSocket
ws://localhost:8080/ws?player_id=A

## Client → Gin API → Redis(GEO) → WebSocket 推播

- Redis：即時位置與附近玩家查詢
- WebSocket：即時同步玩家位置
- PostgreSQL：儲存地標與歷史資料