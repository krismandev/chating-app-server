package usecase

import (
	"context"
	domain "server-chat/domain/entity"
	"server-chat/domain/request"
	"server-chat/domain/response"
	"server-chat/repository"
	"server-chat/utility"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ChatUseCaseImpl struct {
	Validate           *validator.Validate
	ChatRepository     repository.ChatRepository
	ChatroomRepository repository.ChatroomRepository
	UserRepository     repository.UserRepository
}

func NewChatUseCase(validate *validator.Validate, repository repository.ChatRepository, chatroomRepository repository.ChatroomRepository, userRepository repository.UserRepository) ChatUseCase {
	return &ChatUseCaseImpl{
		Validate:           validate,
		ChatRepository:     repository,
		ChatroomRepository: chatroomRepository,
		UserRepository:     userRepository,
	}
}

func (usecase *ChatUseCaseImpl) GetChats(ctx context.Context, request request.GetChatRequest) ([]response.ChatResponse, error) {
	var responses []response.ChatResponse
	var err error

	return responses, err
}

func (usecase *ChatUseCaseImpl) CreateChat(ctx context.Context, request request.CreateChatRequest) (response.ChatResponse, error) {
	var responses response.ChatResponse
	var err error

	chat := new(domain.Chat)
	err = usecase.Validate.Struct(request)
	if err != nil {
		return responses, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	isAllowed := usecase.checkMemberOrUser(request)
	if !isAllowed {
		logrus.Errorf("Error in UseCase: User or chatroom is not valid %v", err)
		return responses, err
	}

	chat.From = request.From
	chat.To = request.To
	chat.IsChatroom = request.IsChatroom
	chat.Message = request.Message

	now := time.Now().Format("2006-01-02 15:04:05")
	chat.CreatedAt = now

	err = usecase.ChatRepository.CreateChat(chat)
	if err != nil {
		return responses, err
	}

	responses = response.ToChatResponse(*chat)

	return responses, err
}

func (usecase *ChatUseCaseImpl) checkMemberOrUser(request request.CreateChatRequest) bool {
	// var err error
	//true means this request is allowed to be inserted

	userPass := new(domain.User)
	userPass.Username = request.From
	err1 := usecase.UserRepository.GetUser(userPass)

	checkMember := new(domain.ChatroomMember)
	checkMember.Username = request.From
	checkMember.ChatroomID = request.To
	err2 := usecase.ChatroomRepository.CheckChatroomMember(checkMember)

	if err1 == nil || err2 == nil {
		return true
	} else {
		return false
	}
}
