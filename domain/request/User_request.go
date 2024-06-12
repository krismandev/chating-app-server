package request

type UserRequest struct {
	ID        int    `json:"id"`
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

type CreateUserRequest struct {
	ID        int    `json:"id"`
	FullName  string `json:"full_name" validate:"required,min=3,max=100"`
	Username  string `json:"username" validate:"required,min=3,max=15"`
	Password  string `json:"password" validate:"required,min=6"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
