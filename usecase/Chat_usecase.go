package usecase

import (
	"context"
	"server-chat/domain/request"
	"server-chat/domain/response"
)

type ChatUseCase interface {
	GetChats(ctx context.Context, request request.GetChatRequest) ([]response.ChatResponse, error)
	CreateChat(ctx context.Context, request request.CreateChatRequest) (response.ChatResponse, error)
}
