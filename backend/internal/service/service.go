package service

import (
	"backend/internal/config"
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/internal/types"
)

type RefreshTokensRequest = types.RefreshTokensRequest

type Users interface {
	CreateUser(user domain.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	VerifyUser(identifier, password, verifyBy string) (domain.User, error)
	GetUsernameByID(userID int64) (string, error)
}

type RefreshSession interface {
	CreateRefreshSession(domain.User, domain.RefreshSession, *config.Config) (string, string, error)
	GetRefreshSession(accessToken string) (domain.RefreshSession, error)
	DeleteRefreshSession(refreshToken string) error
	UpdateRefreshSession(token domain.RefreshToken, req RefreshTokensRequest, cfg *config.Config) (domain.AccessToken, domain.RefreshToken, error)
}

type WebSocket interface {
	SendMessage(message domain.Message) error
}

type Service struct {
	Users
	RefreshSession
	WebSocket
}

func NewService(repos *repository.Repositories) *Service {
	return &Service{
		Users:          NewUsersService(repos.Users),
		RefreshSession: NewRefreshSessionService(repos.RefreshSession),
		WebSocket:      NewWebSocketService(repos.WebSocket),
	}
}
