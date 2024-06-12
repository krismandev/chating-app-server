package repository

import (
	domain "server-chat/domain/entity"
	"server-chat/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatroomRepositoryImpl struct {
	DB *gorm.DB
}

func NewChatroomRepository(db *gorm.DB) ChatroomRepository {
	return &ChatroomRepositoryImpl{
		DB: db,
	}
}

func (repository *ChatroomRepositoryImpl) GetChatrooms(chatrooms *[]domain.Chatroom) error {
	var err error
	result := repository.DB.Find(&chatrooms)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository *ChatroomRepositoryImpl) GetChatroom(chatrooms *domain.Chatroom) error {
	var err error
	result := repository.DB.Where("id = ?", chatrooms.ID).First(&chatrooms)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository *ChatroomRepositoryImpl) CreateChatroom(chatroom *domain.Chatroom) error {
	var err error

	result := repository.DB.Create(&chatroom)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository *ChatroomRepositoryImpl) JoinChatroom(member *domain.ChatroomMember) error {
	var err error

	result := repository.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chatroom_id"}, {Name: "username"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"status": 1}),
	}).Create(&member)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository *ChatroomRepositoryImpl) LeaveChatroom(member *domain.ChatroomMember) error {
	var err error

	result := repository.DB.Model(&member).Where("username = ? AND chatroom_id = ?", member.Username, member.ChatroomID).Update("status", 0)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository *ChatroomRepositoryImpl) CheckChatroomMember(member *domain.ChatroomMember) error {
	var err error

	result := repository.DB.Where("username = ? AND chatroom_id = ?", member.Username, member.ChatroomID).First(&member)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}
