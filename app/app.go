package app

import (
	"server-chat/app/http/controller"
	"server-chat/app/http/route"
	"server-chat/config"
	"server-chat/repository"
	"server-chat/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Application struct {
	DB            *gorm.DB
	Validate      *validator.Validate
	Gin           *gin.Engine
	Config        config.Configuration
	WsConnections []*config.WebSocketConnection
}

func InitApp(app *Application) {

	userRepository := repository.NewUserRepository(app.DB)
	userUseCase := usecase.NewUserUseCase(userRepository, app.Validate)
	userController := controller.NewUserController(userUseCase)

	chatroomRepository := repository.NewChatroomRepository(app.DB)
	chatroomUseCase := usecase.NewChatroomUseCase(app.Validate, chatroomRepository, userRepository)
	chatroomController := controller.NewChatroomController(chatroomUseCase)

	chatRepository := repository.NewChatRepository(app.DB)
	chatUseCase := usecase.NewChatUseCase(app.Validate, chatRepository, chatroomRepository, userRepository)

	wsController := controller.NewWsController(userUseCase, &app.WsConnections, chatUseCase)

	routeConfig := route.RouteConfig{
		Gin:                app.Gin,
		UserController:     userController,
		WsController:       wsController,
		ChatroomController: chatroomController,
	}

	routeConfig.Setup()

	go wsController.HandleIncommingMessage()
}
