package response

import domain "server-chat/domain/entity"

type ChatroomMemberResponse struct {
	ChatroomID string `json:"chatroom_id"`
	Username   string `json:"username"`
	CreatedAt  string `json:"created_at"`
	JoinDate   string `json:"join_date"`
	Status     int    `json:"status"`
}

func ToChatroomMemberResponse(data domain.ChatroomMember) ChatroomMemberResponse {
	var resp ChatroomMemberResponse
	resp.ChatroomID = data.ChatroomID
	resp.Status = data.Status
	resp.Username = data.Username
	resp.JoinDate = data.JoinDate
	return resp
}
