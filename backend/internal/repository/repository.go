package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	GetUserByUsername(username, password string) (domain.User, error)
	GetUserByEmail(username, password string) (domain.User, error)
	CreateUser(user domain.User) (int64, error)
}

type RefreshSession interface {
	CreateRefreshSession(session domain.RefreshSession) (domain.RefreshSession, error)
	DeleteRefreshSession(token domain.RefreshToken) (domain.RefreshSession, error)
}

type WebSocket interface {
	CreateMessage(message domain.Message) error
}

type Repositories struct {
	Users          Users
	RefreshSession RefreshSession
	WebSocket      WebSocket
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:          NewUsersRepo(db),
		RefreshSession: NewRefreshSessionRepo(db),
		WebSocket:      NewWebSocketRepo(db),
	}
}
