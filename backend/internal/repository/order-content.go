package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type OrdersContentRepo struct {
	db *gorm.DB
}

func NewOrdersContentRepo(db *gorm.DB) *OrdersContentRepo {
	return &OrdersContentRepo{
		db: db,
	}
}

func (r *OrdersContentRepo) CreateOrderContent(orderContent domain.OrderContent) (int64, error) {
	if err := r.db.Create(&orderContent).Error; err != nil {
		return 0, err
	}
	return orderContent.Id, nil
}

func (r *OrdersContentRepo) UpdateOrderContent(orderContent domain.OrderContent) error {
	return r.db.Model(&domain.OrderContent{Id: orderContent.Id}).Updates(orderContent).Error
	//return r.db.Model(&domain.OrderContent{}).Where("id = ?", orderContent.Id).Updates(orderContent).Error
}

func (r *OrdersContentRepo) DeleteOrderContent(orderContentId int64) error {
	return r.db.Delete(&domain.OrderContent{Id: orderContentId}).Error
}
