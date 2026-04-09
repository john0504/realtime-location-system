package handler

import (
	"context"
	"net/http"
	"realtime-location/internal/service"

	"github.com/gin-gonic/gin"
)

type LocationRequest struct {
	PlayerID string  `json:"player_id"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

var ctx = context.Background()

func UpdateLocation(c *gin.Context) {
	var req LocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.UpdatePlayerLocation(req.PlayerID, req.Lat, req.Lng)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
