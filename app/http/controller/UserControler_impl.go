package controller

import (
	"server-chat/domain/request"
	"server-chat/usecase"
	"server-chat/utility"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserControllerImpl struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return &UserControllerImpl{
		UserUseCase: useCase,
	}
}

func (controller *UserControllerImpl) CreateUser(c *gin.Context) {
	var err error
	ctx := c.Request.Context()

	var responseData interface{}
	createUserRequest := request.CreateUserRequest{}
	err = utility.ParseRequestBody(c, &createUserRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return
	}

	responseData, err = controller.UserUseCase.CreateUser(ctx, createUserRequest)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return
	}

	utility.WriteResponseSingleJSON(c, responseData, err)
}

func (controller *UserControllerImpl) GetUsers(c *gin.Context) {
	var err error
	ctx := c.Request.Context()

	var responseData interface{}
	userReq := request.UserRequest{}
	err = utility.ParseRequestBody(c, &userReq)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return
	}

	responseData, err = controller.UserUseCase.GetUsers(ctx, userReq)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return
	}

	utility.WriteResponseSingleJSON(c, responseData, err)
}
