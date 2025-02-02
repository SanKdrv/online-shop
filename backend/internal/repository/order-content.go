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
	return 0, nil
}

func (r *OrdersContentRepo) UpdateOrderContent(orderContent domain.OrderContent) error {
	return nil
}

func (r *OrdersContentRepo) DeleteOrderContent(orderContentID int64) error {
	return nil
}
