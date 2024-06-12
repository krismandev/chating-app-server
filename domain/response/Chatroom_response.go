package response

import domain "server-chat/domain/entity"

type ChatroomResponse struct {
	ID           string `json:"id"`
	ChatroomName string `json:"chatroom_name"`
	Username     string `json:"username"`
	CreatedAt    string `json:"created_at"`
	UpdateAt     string `json:"update_at"`
}

func ToChatroomResponse(chatroom domain.Chatroom) ChatroomResponse {
	var dt ChatroomResponse

	dt.ID = chatroom.ID
	dt.ChatroomName = chatroom.ChatroomName
	dt.CreatedAt = chatroom.CreatedAt
	dt.UpdateAt = chatroom.UpdateAt
	dt.Username = chatroom.Username
	return dt
}
