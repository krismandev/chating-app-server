package repository

import (
	domain "server-chat/domain/entity"

	"gorm.io/gorm"
)

type ChatRepositoryImpl struct {
	DB *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &ChatRepositoryImpl{
		DB: db,
	}
}

func (repository ChatRepositoryImpl) GetChatByUsername(chats *[]domain.Chat, dataParams map[string]string) {
	base := repository.DB
	if len(dataParams["id"]) > 0 {
		base = base.Where("id = ?", dataParams["id"])
	}
	if len(dataParams["username"]) > 0 {
		base = base.Where("from = ?", dataParams["username"]).Or("to = ?", dataParams["username"])
	}
	base.Find(&chats)
}

func (repository ChatRepositoryImpl) CreateChat(chat *domain.Chat) error {
	var err error
	result := repository.DB.Create(&chat)
	err = result.Error
	return err
}
