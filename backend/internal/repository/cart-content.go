package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type CartsContentRepo struct {
	db *gorm.DB
}

func NewCartsContentRepo(db *gorm.DB) *CartsContentRepo {
	return &CartsContentRepo{
		db: db,
	}
}

func (r *CartsContentRepo) GetCartContentById(id int64) (domain.CartContent, error) {
	var cartContent domain.CartContent

	if err := r.db.Model(&domain.CartContent{}).Where("id = ?", id).First(&cartContent).Error; err != nil {
		return domain.CartContent{}, err
	}
	return cartContent, nil
}

func (r *CartsContentRepo) GetCartContentByUserId(userId int64) ([]domain.CartContent, error) {
	var cartContents []domain.CartContent

	if err := r.db.Model(&domain.CartContent{}).Where("user_id = ?", userId).Find(&cartContents).Error; err != nil {
		return []domain.CartContent{}, err
	}
	return cartContents, nil
}

func (r *CartsContentRepo) CreateCartContent(cartContent domain.CartContent) (int64, error) {
	if err := r.db.Model(&domain.CartContent{}).Create(&cartContent).Error; err != nil {
		return 0, err
	}
	return cartContent.Id, nil
}

func (r *CartsContentRepo) UpdateCartContent(cartContent domain.CartContent) error {
	return r.db.Model(&domain.CartContent{}).Where("id = ?", &cartContent.Id).Updates(&cartContent).Error
}

func (r *CartsContentRepo) DeleteCartContent(cartContentId int64) error {
	return r.db.Delete(&domain.CartContent{}, "id = ?", &cartContentId).Error
}
