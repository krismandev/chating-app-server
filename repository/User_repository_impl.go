package repository

import (
	domain "server-chat/domain/entity"
	"server-chat/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository UserRepositoryImpl) GetUser(user *domain.User) error {
	var err error
	result := repository.DB.Where("username = ?", &user.Username).First(&user)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository UserRepositoryImpl) CreateUser(user *domain.User) error {
	var err error
	result := repository.DB.Create(&user)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository UserRepositoryImpl) GetUsers(users *[]domain.User) error {
	var err error
	result := repository.DB.Find(&users)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}
