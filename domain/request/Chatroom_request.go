package request

type ChatroomRequest struct {
	ID           string `json:"id"`
	ChatroomName string `json:"chatroom_name"`
	Username     string `json:"username"`
	CreatedAt    string `json:"created_at"`
	UpdateAt     string `json:"update_at"`
}

type CreateChatroomRequest struct {
	ChatroomName string `json:"chatroom_name" validate:"required,min=3,max=50"`
	Username     string `json:"username" validate:"required"`
}
