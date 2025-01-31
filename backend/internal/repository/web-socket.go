package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type WebSocketRepo struct {
	db *gorm.DB
}

func NewWebSocketRepo(db *gorm.DB) *WebSocketRepo {
	return &WebSocketRepo{
		db: db,
	}
}

func (r *WebSocketRepo) CreateMessage(message domain.Message) error {
	//slog.Info(message.Message)
	//if err := r.db.Create(&message).Error; err != nil {
	//	return 0, err
	//}
	//return user.ID, nil // Assuming domain.User has an ID field
	if err := r.db.Create(&message).Error; err != nil {
		return err
	}
	return nil
}
