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
	return s.repo.CreateOrderContent(orderContent)
}

func (s *OrdersContentService) UpdateOrderContent(orderContent domain.OrderContent) error {
	return s.repo.UpdateOrderContent(orderContent)
}

func (s *OrdersContentService) DeleteOrderContent(orderContentID int64) error {
	return s.repo.DeleteOrderContent(orderContentID)
}
