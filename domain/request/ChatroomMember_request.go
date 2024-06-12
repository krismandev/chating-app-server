package request

type ChatroomMemberRequest struct {
	ChatroomID string `json:"chatroom_id"`
	Username   string `json:"username"`
	CreatedAt  string `json:"created_at"`
	JoinDate   string `json:"join_date"`
	Status     int    `json:"status"`
}

type CreateChatroomMemberRequest struct {
	ChatroomID string `json:"chatroom_id" validate:"required"`
	Username   string `json:"username" validate:"required"`
}

type LeaveChatroomRequest struct {
	ChatroomID string `json:"chatroom_id" validate:"required"`
	Username   string `json:"username" validate:"required"`
}
