package usecase

import (
	"context"
	domain "server-chat/domain/entity"
	"server-chat/domain/request"
	"server-chat/domain/response"
	"server-chat/repository"
	"server-chat/utility"

	"github.com/go-playground/validator/v10"
)

type UserUseCaseImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserUseCase(userRepository repository.UserRepository, validate *validator.Validate) UserUseCase {
	return &UserUseCaseImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (usecase UserUseCaseImpl) GetUser(ctx context.Context, request request.UserRequest) (response.UserResponse, error) {
	var err error
	var resp response.UserResponse

	user := new(domain.User)
	user.Username = request.Username
	user.ID = request.ID
	user.FullName = request.FullName

	usecase.UserRepository.GetUser(user)

	resp = response.ToUserResponse(*user)

	return resp, err
}

func (usecase UserUseCaseImpl) GetUsers(ctx context.Context, request request.UserRequest) ([]response.UserResponse, error) {
	var err error
	var resp []response.UserResponse

	users := new([]domain.User)

	usecase.UserRepository.GetUsers(users)

	for _, each := range *users {
		usr := response.ToUserResponse(each)
		resp = append(resp, usr)
	}
	// resp = response.ToUserResponse(*user)

	return resp, err
}

func (usecase UserUseCaseImpl) CreateUser(ctx context.Context, request request.CreateUserRequest) (response.UserResponse, error) {
	var err error
	var resp response.UserResponse

	err = usecase.Validate.Struct(request)
	if err != nil {
		return resp, err
	}

	user := new(domain.User)
	user.FullName = request.FullName
	user.Username = request.Username
	user.Password = utility.HashPassword(request.Password)

	err = usecase.UserRepository.CreateUser(user)
	if err != nil {
		return resp, err
	}

	resp = response.ToUserResponse(*user)

	return resp, err
}
