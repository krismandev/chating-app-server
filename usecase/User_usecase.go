package usecase

import (
	"context"
	"server-chat/domain/request"
	"server-chat/domain/response"
)

type UserUseCase interface {
	GetUser(ctx context.Context, request request.UserRequest) (response.UserResponse, error)
	GetUsers(ctx context.Context, request request.UserRequest) ([]response.UserResponse, error)
	CreateUser(ctx context.Context, request request.CreateUserRequest) (response.UserResponse, error)
}
