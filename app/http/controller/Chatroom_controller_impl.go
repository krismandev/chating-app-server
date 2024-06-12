package controller

import (
	"context"
	"server-chat/domain/request"
	"server-chat/domain/response"
	"server-chat/usecase"
	"server-chat/utility"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ChatroomControllerImpl struct {
	ChatroomUseCase usecase.ChatroomUseCase
}

func NewChatroomController(chatroomUseCase usecase.ChatroomUseCase) ChatroomController {
	return &ChatroomControllerImpl{
		ChatroomUseCase: chatroomUseCase,
	}
}

func (controller *ChatroomControllerImpl) GetChatrooms(c *gin.Context) {
	var dttb response.GlobalListDataTableResponse
	var err error
	ctx := c.Request.Context()

	chatroomRequest := request.ChatroomRequest{}
	err = utility.ParseRequestBody(c, &chatroomRequest)
	if err != nil {
		utility.WriteResponseListJSON(c, dttb, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return
	}
	data, err := controller.ChatroomUseCase.GetChatrooms(ctx, chatroomRequest)
	if err != nil {
		utility.WriteResponseListJSON(c, dttb, err)
		return
	}

	for _, each := range data {
		dttb.List = append(dttb.List, each)
	}

	utility.WriteResponseListJSON(c, dttb, err)
}

func (controller *ChatroomControllerImpl) CreateChatroom(c *gin.Context) {
	var err error
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()
	var responseData interface{}

	request := request.CreateChatroomRequest{}
	err = utility.ParseRequestBody(c, &request)
	if err != nil {
		logrus.Info("Error in Controller : Parsing Error %v", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return
	}

	responseData, err = controller.ChatroomUseCase.CreateChatroom(ctx, request)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

}

func (controller *ChatroomControllerImpl) JoinChatroom(c *gin.Context) {
	var err error
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()
	var responseData interface{}

	request := request.CreateChatroomMemberRequest{}
	err = utility.ParseRequestBody(c, &request)
	if err != nil {
		logrus.Info("Error in Controller : Parsing Error %v", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return
	}

	responseData, err = controller.ChatroomUseCase.JoinChatroom(ctx, request)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return
	}
	utility.WriteResponseSingleJSON(c, responseData, err)

}

func (controller *ChatroomControllerImpl) LeaveChatroom(c *gin.Context) {
	var err error
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()
	var responseData interface{}

	request := request.LeaveChatroomRequest{}
	err = utility.ParseRequestBody(c, &request)
	if err != nil {
		logrus.Info("Error in Controller : Parsing Error %v", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return
	}

	err = controller.ChatroomUseCase.LeaveChatroom(ctx, request)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return
	}
	utility.WriteResponseSingleJSON(c, responseData, err)
}
