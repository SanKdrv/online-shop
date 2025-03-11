package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type CartsContentService struct {
	repo repository.CartsContent
}

func NewCartsContentService(repo repository.CartsContent) *CartsContentService {
	return &CartsContentService{repo: repo}
}

func (s *CartsContentService) GetCartContentById(id int64) (domain.CartContent, error) {
	return s.repo.GetCartContentById(id)
}

func (s *CartsContentService) GetCartContentByUserId(userId int64) ([]domain.CartContent, error) {
	return s.repo.GetCartContentByUserId(userId)
}

func (s *CartsContentService) CreateCartContent(cartContent domain.CartContent) (int64, error) {
	return s.repo.CreateCartContent(cartContent)
}

func (s *CartsContentService) UpdateCartContent(cartContent domain.CartContent) error {
	return s.repo.UpdateCartContent(cartContent)
}

func (s *CartsContentService) DeleteCartContent(cartContentId int64) error {
	return s.repo.DeleteCartContent(cartContentId)
}
