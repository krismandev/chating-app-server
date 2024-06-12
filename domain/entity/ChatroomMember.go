package domain

type ChatroomMember struct {
	ChatroomID string `json:"chatroom_id" gorm:"column:chatroom_id;primaryKey"`
	Username   string `json:"username" gorm:"column:username"`
	JoinDate   string `json:"join_date" gorm:"column:join_date"`
	Status     int    `json:"status" gorm:"column:status"`
}
