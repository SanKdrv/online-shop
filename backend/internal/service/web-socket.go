package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type WebSocketService struct {
	repo repository.WebSocket
}

func NewWebSocketService(repo repository.WebSocket) *WebSocketService {
	return &WebSocketService{repo: repo}
}

func (s *WebSocketService) SendMessage(message domain.Message) error {
	return s.repo.CreateMessage(message)
}
