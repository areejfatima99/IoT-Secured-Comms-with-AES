package ports

import "github.com/gin-gonic/gin"

type MainHandler interface {
	SendMessage() gin.HandlerFunc
	GetMessages() gin.HandlerFunc
}
