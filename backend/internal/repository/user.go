package repository

import (
	"gorm.io/gorm"

	//"time"
	//"backend/internal/domain"
	"backend/internal/domain"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) CreateUser(user domain.User) (int64, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil // Assuming domain.User has an ID field
}

func (r *UsersRepo) GetUserByUsername(username, password string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ? AND password_hash = ?", username, password).First(&user).Error
	return user, err
}

func (r *UsersRepo) GetUserByEmail(email, password string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ? AND password_hash = ?", email, password).First(&user).Error
	return user, err
}

func (r *UsersRepo) GetUsernameByID(id int64) (string, error) {
	var username string
	err := r.db.Where("id = ?", id).First(&username).Error
	return username, err
}
