package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type BrandsService struct {
	repo repository.Brands
}

func NewBrandsService(repo repository.Brands) *BrandsService {
	return &BrandsService{repo: repo}
}

// GetIdByBrand retrieves the id of a brand by its name.
//
// Returns the brand ID and an error if the operation fails or if the brand does not exist.
func (s *BrandsService) GetIdByBrand(name string) (int64, error) {
	return s.repo.GetIdByBrand(name)
}

// GetBrandById returns brand name by brand id.
//
// Returns empty string if brand does not exist in DB.
func (s *BrandsService) GetBrandById(brandId int64) (string, error) {
	return s.repo.GetBrandById(brandId)
}

// CreateBrand creates a new brand with the given name.
//
// Returns the ID of the newly created brand and an error, if any occurred during creation.
func (s *BrandsService) CreateBrand(name string) (int64, error) {
	return s.repo.CreateBrand(name)
}

// DeleteBrand deletes a brand with the given ID.
//
// Returns an error if the deletion fails or if the brand does not exist.
func (s *BrandsService) DeleteBrand(brandId int64) error {
	return s.repo.DeleteBrand(brandId)
}

// UpdateBrand updates the name of the brand with the given ID.
//
// Returns an error if the update fails or if the brand does not exist.
func (s *BrandsService) UpdateBrand(brandId int64, name string) error {
	return s.repo.UpdateBrand(brandId, name)
}

// GetAll getting all rows from Brands Table.
//
// Returns an error if the error occurred.
func (s *BrandsService) GetAll(offset int64, limit int64) ([]domain.Brand, int64, error) {
	return s.repo.GetAll(offset, limit)
}
