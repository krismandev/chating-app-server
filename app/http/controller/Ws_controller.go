package controller

import "github.com/gin-gonic/gin"

type WsController interface {
	Connect(c *gin.Context)
	HandleIncommingMessage()
}
