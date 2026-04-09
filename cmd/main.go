package main

import (
	"os"
	"realtime-location/internal/handler"
	"realtime-location/internal/service"
	"realtime-location/internal/websocket"
	"realtime-location/pkg/db"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	db.Init()

	db.DB.Exec(`
	CREATE TABLE IF NOT EXISTS landmarks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		latitude DOUBLE PRECISION NOT NULL,
		longitude DOUBLE PRECISION NOT NULL,
		radius INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)

	db.DB.Exec(`
	INSERT INTO landmarks (name, latitude, longitude, radius)
	VALUES 
	('台中火車站', 24.1367, 120.6850, 300)
	ON CONFLICT DO NOTHING;
	`)

	ws := &websocket.WSNotifier{}
	service.SetNotifier(ws)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.GET("/ws", websocket.HandleWebSocket)
	r.POST("/update-location", handler.UpdateLocation)

	r.Run(":" + port)
}
