package repository

import domain "server-chat/domain/entity"

type UserRepository interface {
	GetUsers(user *[]domain.User) error
	GetUser(user *domain.User) error
	CreateUser(user *domain.User) error
}
