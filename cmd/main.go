package main

import (
	"realtime-location/internal/handler"
	"realtime-location/internal/service"
	"realtime-location/internal/websocket"
	"realtime-location/pkg/db"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
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
