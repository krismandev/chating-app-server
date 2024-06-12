package route

import (
	"server-chat/app/http/controller"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Gin                *gin.Engine
	AuthMiddleware     gin.HandlerFunc
	WsController       controller.WsController
	UserController     controller.UserController
	ChatroomController controller.ChatroomController
}

func (c *RouteConfig) Setup() {
	c.SetUpRoute()
}

func (rc *RouteConfig) SetUpRoute() {
	publicRoute := rc.Gin.Group("")
	publicRoute.POST("/register", rc.UserController.CreateUser)
	publicRoute.GET("/ws", rc.WsController.Connect)
	publicRoute.GET("/chatroom", rc.ChatroomController.GetChatrooms)
	publicRoute.POST("/chatroom", rc.ChatroomController.CreateChatroom)
	publicRoute.POST("/join-chatroom", rc.ChatroomController.JoinChatroom)
	publicRoute.POST("/leave-chatroom", rc.ChatroomController.LeaveChatroom)
	publicRoute.GET("/user", rc.UserController.GetUsers)
	publicRoute.GET("/", func(ctx *gin.Context) {

	})
}
