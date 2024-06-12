package domain

type Chat struct {
	ID         int    `json:"id" gorm:"column:id;primaryKey"`
	Message    string `json:"message" gorm:"column:message"`
	From       string `json:"from" gorm:"column:from"`
	To         string `json:"to" gorm:"column:to"`
	IsChatroom int    `json:"is_chatroom" gorm:"column:is_chatroom"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`
}
