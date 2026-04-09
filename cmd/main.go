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
	db.InitTables()

	r := gin.Default()

	r.GET("/ws", websocket.HandleWebSocket)
	r.POST("/update-location", handler.UpdateLocation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

	ws := &websocket.WSNotifier{}
	service.SetNotifier(ws)
}
