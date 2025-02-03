package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type OrdersContentService struct {
	repo repository.OrdersContent
}

func NewOrdersContentService(repo repository.OrdersContent) *OrdersContentService {
	return &OrdersContentService{repo: repo}
}

func (s *OrdersContentService) CreateOrderContent(orderContent domain.OrderContent) (int64, error) {
	return 0, nil
}

func (s *OrdersContentService) UpdateOrderContent(orderContent domain.OrderContent) error {
	return nil
}

func (s *OrdersContentService) DeleteOrderContent(orderContentID int64) error {
	return nil
}
