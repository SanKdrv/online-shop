package types

import (
	"backend/internal/domain"
)

type SendMessageRequest struct {
	Message domain.Message `json:"message"`
}

type MessageRequest struct {
	Message string `json:"message"`
	UserID  int64  `json:"user_id"`
	ChatID  int64  `json:"chat_id,string"`
}

type SendMessageResponse struct {
	Response string `json:"response,omitempty"`
}
