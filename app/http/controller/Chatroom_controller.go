package controller

import "github.com/gin-gonic/gin"

type ChatroomController interface {
	GetChatrooms(c *gin.Context)
	CreateChatroom(c *gin.Context)
	JoinChatroom(c *gin.Context)
	LeaveChatroom(c *gin.Context)
}
