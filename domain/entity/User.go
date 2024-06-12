package domain

type User struct {
	ID        int    `json:"id" gorm:"column:id;primaryKey"`
	FullName  string `json:"full_name" gorm:"column:full_name"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP()"`
	UpdateAt  string `json:"update_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP()"`
}
