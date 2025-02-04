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
	return order.Id, nil
}

func (r *OrdersRepo) GetOrderById(orderId int64) (domain.Order, error) {
	var order domain.Order
	if err := r.db.Model(&domain.Order{}).First(&order, "id = ?", orderId).Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (r *OrdersRepo) GetOrdersByUserId(userId int64) ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.db.Model(&domain.Order{}).Find(&orders, "user_id = ?", userId).Error; err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}

func (r *OrdersRepo) UpdateOrder(order domain.Order) error {
	return r.db.Find(&domain.Order{}, "id = ?", order.Id).Updates(&order).Error
}

func (r *OrdersRepo) DeleteOrder(orderId int64) error {
	return r.db.Delete(&domain.Order{Id: orderId}).Error
}
