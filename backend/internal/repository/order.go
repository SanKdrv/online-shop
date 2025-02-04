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
	if err := r.db.Create(&order).Error; err != nil {
		return -1, err
	}
	return order.ID, nil
}

func (r *OrdersRepo) GetOrderByID(orderID int64) (domain.Order, error) {
	var order domain.Order
	if err := r.db.Model(&domain.Order{}).First(&order, "id = ?", orderID).Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (r *OrdersRepo) GetOrdersByUserID(userID int64) ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.db.Model(&domain.Order{}).Find(&orders, "user_id = ?", userID).Error; err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}

func (r *OrdersRepo) UpdateOrder(order domain.Order) error {
	return r.db.Find(&domain.Order{}, "id = ?", order.ID).Updates(&order).Error
}

func (r *OrdersRepo) DeleteOrder(orderID int64) error {
	return r.db.Delete(&domain.Order{ID: orderID}).Error
}
