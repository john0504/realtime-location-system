package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	redis2 "github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[string]*websocket.Conn)

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	playerID := c.Query("player_id")
	clients[playerID] = conn

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(clients, playerID)
			conn.Close()
			break
		}
	}
}

func Broadcast(players []redis2.GeoLocation, playerID string, lat, lng float64) {
	for _, p := range players {
		if p.Name == playerID {
			continue
		}

		if conn, ok := clients[p.Name]; ok && conn != nil {
			err := conn.WriteJSON(map[string]interface{}{
				"player_id": playerID,
				"lat":       lat,
				"lng":       lng,
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}

type WSNotifier struct{}

func (w *WSNotifier) Broadcast(playerID string, lat, lng float64, targets []string) {
	for _, id := range targets {
		if conn, ok := clients[id]; ok {
			conn.WriteJSON(map[string]interface{}{
				"player_id": playerID,
				"lat":       lat,
				"lng":       lng,
			})
		}
	}
}

func (w *WSNotifier) NotifyLandmark(playerID string, name string) {
	if conn, ok := clients[playerID]; ok {
		conn.WriteJSON(map[string]interface{}{
			"type":    "landmark",
			"name":    name,
			"message": "你已進入地標範圍",
		})
	}
}
