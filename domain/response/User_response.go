package response

import domain "server-chat/domain/entity"

type UserResponse struct {
	ID        int    `json:"id"`
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

func ToUserResponse(user domain.User) UserResponse {
	var dt UserResponse
	dt.ID = user.ID
	dt.FullName = user.FullName
	dt.Username = user.Username
	dt.Password = user.Password
	dt.CreatedAt = user.CreatedAt
	dt.UpdateAt = user.UpdateAt

	return dt
}
