package repository

import domain "server-chat/domain/entity"

type ChatRepository interface {
	GetChatByUsername(chat *[]domain.Chat, dataParams map[string]string)
	CreateChat(chat *domain.Chat) error
}
