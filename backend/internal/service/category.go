package service

import "backend/internal/repository"

type CategoriesService struct {
	repo repository.Categories
}

func NewCategoriesService(repo repository.Categories) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) GetIDByCategory(name string) (int64, error) {
	return 0, nil
}

func (s *CategoriesService) GetCategoryByID(categoryID int64) (string, error) {
	return "", nil
}

func (s *CategoriesService) CreateCategory(name string) (int64, error) {
	return 0, nil
}

func (s *CategoriesService) DeleteCategory(categoryID int64) error {
	return nil
}

func (s *CategoriesService) UpdateCategory(categoryID int64, name string) error {
	return nil
}
