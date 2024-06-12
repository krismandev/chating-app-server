package response

import domain "server-chat/domain/entity"

type ChatResponse struct {
	ID         int    `json:"id"`
	Message    string `json:"message"`
	From       string `json:"from"`
	To         string `json:"to"`
	IsChatroom string `json:"is_chatroom"`
	CreatedAt  string `json:"created_at"`
}

func ToChatResponse(chat domain.Chat) ChatResponse {
	var dt ChatResponse
	dt.ID = chat.ID
	dt.From = chat.From
	dt.To = chat.To
	dt.Message = chat.Message
	dt.CreatedAt = chat.CreatedAt

	return dt
}
