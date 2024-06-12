package usecase

import (
	"context"
	"server-chat/domain/request"
	"server-chat/domain/response"
)

type ChatroomUseCase interface {
	GetChatrooms(ctx context.Context, request request.ChatroomRequest) ([]response.ChatroomResponse, error)
	CreateChatroom(ctx context.Context, request request.CreateChatroomRequest) (response.ChatroomResponse, error)
	JoinChatroom(ctx context.Context, request request.CreateChatroomMemberRequest) (response.ChatroomMemberResponse, error)
	LeaveChatroom(ctx context.Context, request request.LeaveChatroomRequest) error
}
