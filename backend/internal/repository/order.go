package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type OrdersRepo struct {
	db *gorm.DB
}

func NewOrdersRepo(db *gorm.DB) *OrdersRepo {
	return &OrdersRepo{
		db: db,
	}
}

func (r *OrdersRepo) CreateOrder(order domain.Order) (int64, error) {
	return 0, nil
}

func (r *OrdersRepo) GetOrderByID(orderID int64) (domain.Order, error) {
	return domain.Order{}, nil
}

func (r *OrdersRepo) GetOrdersByUserID(userID int64) ([]domain.Order, error) {
	return []domain.Order{}, nil
}

func (r *OrdersRepo) UpdateOrder(order domain.Order) error {
	return nil
}

func (r *OrdersRepo) DeleteOrder(orderID int64) error {
	return nil
}
