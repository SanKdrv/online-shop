package service

import "backend/internal/repository"

type CategoriesService struct {
	repo repository.Categories
}

func NewCategoriesService(repo repository.Categories) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) GetIdByCategory(name string) (int64, error) {
	return s.repo.GetIdByCategory(name)
}

func (s *CategoriesService) GetCategoryById(categoryID int64) (string, error) {
	return s.repo.GetCategoryById(categoryID)
}

func (s *CategoriesService) CreateCategory(name string) (int64, error) {
	return s.repo.CreateCategory(name)
}

func (s *CategoriesService) DeleteCategory(categoryID int64) error {
	return s.repo.DeleteCategory(categoryID)
}

func (s *CategoriesService) UpdateCategory(categoryID int64, name string) error {
	return s.repo.UpdateCategory(categoryID, name)
}
