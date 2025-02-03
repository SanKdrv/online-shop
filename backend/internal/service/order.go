package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type OrdersService struct {
	repo repository.Orders
}

func NewOrdersService(repo repository.Orders) *OrdersService {
	return &OrdersService{repo: repo}
}

func (s *OrdersService) CreateOrder(order domain.Order) (int64, error) {
	return 0, nil
}

func (s *OrdersService) GetOrderByID(orderID int64) (domain.Order, error) {
	return domain.Order{}, nil
}

func (s *OrdersService) GetOrdersByUserID(userID int64) ([]domain.Order, error) {
	return []domain.Order{}, nil
}

func (s *OrdersService) UpdateOrder(order domain.Order) error {
	return nil
}

func (s *OrdersService) DeleteOrder(orderID int64) error {
	return nil
}
