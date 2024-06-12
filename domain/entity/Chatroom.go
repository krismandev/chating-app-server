package domain

type Chatroom struct {
	ID           string `json:"id" gorm:"column:id;primaryKey"`
	ChatroomName string `json:"chatroom_name" gorm:"column:chatroom_name"`
	Username     string `json:"username" gorm:"column:username"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at"`
	UpdateAt     string `json:"update_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP()"`
}
