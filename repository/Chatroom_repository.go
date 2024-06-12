package repository

import domain "server-chat/domain/entity"

type ChatroomRepository interface {
	GetChatrooms(chatrooms *[]domain.Chatroom) error
	GetChatroom(chatrooms *domain.Chatroom) error
	CreateChatroom(chatroom *domain.Chatroom) error
	JoinChatroom(member *domain.ChatroomMember) error
	LeaveChatroom(member *domain.ChatroomMember) error
	CheckChatroomMember(member *domain.ChatroomMember) error
}
