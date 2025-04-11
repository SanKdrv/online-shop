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
	return s.repo.CreateOrder(order)
}

func (s *OrdersService) GetOrderById(orderId int64) (domain.Order, error) {
	return s.repo.GetOrderById(orderId)
}

func (s *OrdersService) GetOrdersByUserId(userId int64) ([]domain.Order, error) {
	return s.repo.GetOrdersByUserId(userId)
}

func (s *OrdersService) UpdateOrder(order domain.Order) error {
	return s.repo.UpdateOrder(order)
}

func (s *OrdersService) DeleteOrder(orderId int64) error {
	return s.repo.DeleteOrder(orderId)
}

func (s *OrdersService) GetAll(offset int64, limit int64) ([]domain.Order, int64, error) {
	return s.repo.GetAll(offset, limit)
}
