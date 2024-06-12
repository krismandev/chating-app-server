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

type ChatroomUseCaseImpl struct {
	Validate           *validator.Validate
	ChatroomRepository repository.ChatroomRepository
	UserRepository     repository.UserRepository
}

func NewChatroomUseCase(validate *validator.Validate, chatroomRepository repository.ChatroomRepository, userRepository repository.UserRepository) ChatroomUseCase {
	return &ChatroomUseCaseImpl{
		Validate:           validate,
		ChatroomRepository: chatroomRepository,
		UserRepository:     userRepository,
	}
}

func (usecase *ChatroomUseCaseImpl) GetChatrooms(ctx context.Context, request request.ChatroomRequest) ([]response.ChatroomResponse, error) {
	var err error
	var responses []response.ChatroomResponse
	var chatrooms []domain.Chatroom
	err = usecase.ChatroomRepository.GetChatrooms(&chatrooms)
	if err != nil {
		return responses, err
	}

	for _, chatroom := range chatrooms {
		resp := response.ToChatroomResponse(chatroom)
		responses = append(responses, resp)
	}

	return responses, err
}

func (usecase *ChatroomUseCaseImpl) CreateChatroom(ctx context.Context, request request.CreateChatroomRequest) (response.ChatroomResponse, error) {
	var chatroomResponse response.ChatroomResponse
	chatroom := new(domain.Chatroom)
	var err error
	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase: Validation Error %v", err)
		return chatroomResponse, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	chatroom.ChatroomName = request.ChatroomName
	now := time.Now().Format("2006-01-02 15:04:05")
	chatroom.CreatedAt = now
	chatroom.Username = request.Username
	chatroom.ID = utility.RandString(8)

	err = usecase.ChatroomRepository.CreateChatroom(chatroom)
	if err != nil {
		return chatroomResponse, err
	}

	chatroomResponse = response.ToChatroomResponse(*chatroom)

	return chatroomResponse, err
}

func (usecase *ChatroomUseCaseImpl) JoinChatroom(ctx context.Context, request request.CreateChatroomMemberRequest) (response.ChatroomMemberResponse, error) {
	var resp response.ChatroomMemberResponse
	var err error

	userPass := new(domain.User)
	userPass.Username = request.Username
	err = usecase.UserRepository.GetUser(userPass)
	if err != nil {
		return resp, err
	}

	chatroomPass := new(domain.Chatroom)
	chatroomPass.ID = request.ChatroomID
	err = usecase.ChatroomRepository.GetChatroom(chatroomPass)
	if err != nil {
		return resp, err
	}

	member := new(domain.ChatroomMember)
	member.ChatroomID = request.ChatroomID
	member.Username = request.Username
	member.JoinDate = time.Now().Format("2006-01-02 15:04:05")
	member.Status = 1
	err = usecase.ChatroomRepository.JoinChatroom(member)
	if err != nil {
		return resp, err
	}

	resp = response.ToChatroomMemberResponse(*member)

	return resp, err
}

func (usecase *ChatroomUseCaseImpl) LeaveChatroom(ctx context.Context, request request.LeaveChatroomRequest) error {
	var err error

	userPass := new(domain.User)
	userPass.Username = request.Username
	err = usecase.UserRepository.GetUser(userPass)
	if err != nil {
		return err
	}

	chatroomPass := new(domain.Chatroom)
	chatroomPass.ID = request.ChatroomID
	err = usecase.ChatroomRepository.GetChatroom(chatroomPass)
	if err != nil {
		return err
	}

	member := new(domain.ChatroomMember)
	member.ChatroomID = request.ChatroomID
	member.Username = request.Username
	member.Status = 0
	err = usecase.ChatroomRepository.LeaveChatroom(member)
	if err != nil {
		return err
	}

	return err
}
