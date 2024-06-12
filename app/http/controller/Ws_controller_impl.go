package controller

import (
	"context"
	"net/http"
	"server-chat/config"
	"server-chat/domain/request"
	"server-chat/usecase"
	"server-chat/utility"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsControllerImpl struct {
	UserUseCase   usecase.UserUseCase
	WsConnections *[]*config.WebSocketConnection
	ChatUseCase   usecase.ChatUseCase
}

var broadcast = make(chan config.MessagePayload, 10)

func NewWsController(userUseCase usecase.UserUseCase, wsConnections *[]*config.WebSocketConnection, chatUseCase usecase.ChatUseCase) WsController {
	return &WsControllerImpl{
		UserUseCase:   userUseCase,
		WsConnections: wsConnections,
		ChatUseCase:   chatUseCase,
	}
}

func (controller WsControllerImpl) HandleIncommingMessage() {
	for {
		payload := <-broadcast
		// logrus.Info("Lihat All Connections ", *controller.WsConnections)
		logrus.Info("Lihat Payload ", payload)
		for _, conn := range *controller.WsConnections {
			logrus.Info("Each ", conn)
			// only those in the same chatroom will receive messages
			if conn.Chatroom == payload.To || conn.Chatroom == payload.From {
				err := conn.WriteJSON(payload)
				if err != nil {
					logrus.Errorf("Error :%v", err)
					conn.Conn.Close()
					removeConnections(*controller.WsConnections, conn)
				}
			}

			if conn.Username == payload.From {
				ctx := context.Background()
				var chatRequest request.CreateChatRequest
				chatRequest.From = payload.From
				chatRequest.To = payload.To
				chatRequest.Message = payload.Message
				resp, err := controller.ChatUseCase.CreateChat(ctx, chatRequest)
				if err != nil {
					logrus.Errorf("Error create chat : %v", err)
				}
				logrus.Info(resp)
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (controller WsControllerImpl) Connect(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	defer ws.Close()

	ctx := c.Request.Context()

	var credentials map[string]string
	err = ws.ReadJSON(&credentials)
	if err != nil {
		logrus.Printf("error reading credentials: %v", err)
		return
	}

	username, password := credentials["username"], credentials["password"]
	userRequest := request.UserRequest{}
	userRequest.Username = strings.TrimSpace(username)
	userRequest.Password = strings.TrimSpace(password)

	user, err := controller.UserUseCase.GetUser(ctx, userRequest)
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("Get User Error"))
		ws.Close()
	}

	checkPassword := utility.ComparePass([]byte(user.Password), []byte(userRequest.Password))
	if !checkPassword {
		ws.WriteMessage(websocket.TextMessage, []byte("Invalid credentials"))
		ws.Close()
		logrus.Info("Auth failed")
	} else {
		ws.WriteMessage(websocket.TextMessage, []byte("Authenticated Successfully"))
	}

	var chatroom map[string]string
	err = ws.ReadJSON(&chatroom)
	if err != nil {
		logrus.Printf("error reading chatroom: %v", err)
		return
	}

	chatroomOrDmUsername := chatroom["chatroom"]

	currentConnection := config.WebSocketConnection{Conn: ws, Username: user.Username, Chatroom: chatroomOrDmUsername}
	*controller.WsConnections = append(*controller.WsConnections, &currentConnection)
	logrus.Info("Adding new connection ", currentConnection)

	for {
		var msg config.MessagePayload
		err := currentConnection.Conn.ReadJSON(&msg)
		if err != nil {
			logrus.Printf("error read json: %v", err)

			*controller.WsConnections = removeConnections(*controller.WsConnections, &currentConnection)
			break
		}
		// msg.Chatroom = client.Chatroom
		broadcast <- msg
	}

}

func removeConnections(connections []*config.WebSocketConnection, elem *config.WebSocketConnection) []*config.WebSocketConnection {
	newConnections := []*config.WebSocketConnection{}
	for _, conn := range connections {
		if conn != elem {
			newConnections = append(newConnections, conn)
		}
	}

	return newConnections

}
