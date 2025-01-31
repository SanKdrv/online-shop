package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type RefreshSessionRepo struct {
	db *gorm.DB
}

func NewRefreshSessionRepo(db *gorm.DB) *RefreshSessionRepo {
	return &RefreshSessionRepo{
		db: db,
	}
}

func (r *RefreshSessionRepo) CreateRefreshSession(session domain.RefreshSession) (domain.RefreshSession, error) {
	if err := r.db.Create(&session).Error; err != nil {
		return domain.RefreshSession{}, err
	}
	return session, nil
}

func (r *RefreshSessionRepo) DeleteRefreshSession(token domain.RefreshToken) (domain.RefreshSession, error) {
	var session domain.RefreshSession

	// Ищем запись в базе данных
	if err := r.db.Where("refresh_token = ?", token.RefreshToken).First(&session).Error; err != nil {
		return domain.RefreshSession{}, err
	}

	// Удаляем найденный кортеж
	if err := r.db.Delete(&session).Error; err != nil {
		return domain.RefreshSession{}, err
	}

	// Возвращаем найденную запись
	return session, nil
}
