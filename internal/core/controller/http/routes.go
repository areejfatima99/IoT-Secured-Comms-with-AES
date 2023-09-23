package http

import (
	"custom-blockchain-iot-protocol/internal/core/ports"
	"github.com/gin-gonic/gin"
)

func Map(group *gin.RouterGroup, h ports.MainHandler) error {
	group.POST("/send-message", h.SendMessage())
	group.GET("/get-messages", h.GetMessages())
	return nil
}
