package service

import "backend/internal/repository"

type BrandsService struct {
	repo repository.Brands
}

func NewBrandsService(repo repository.Brands) *BrandsService {
	return &BrandsService{repo: repo}
}

func (s *BrandsService) GetIDByBrand(name string) (int64, error) {
	return 0, nil
}

func (s *BrandsService) GetBrandByID(categoryID int64) (string, error) {
	return "", nil
}

func (s *BrandsService) CreateBrand(name string) (int64, error) {
	return 0, nil
}

func (s *BrandsService) DeleteBrand(categoryID int64) error {
	return nil
}

func (s *BrandsService) UpdateBrand(categoryID int64, name string) error {
	return nil
}
