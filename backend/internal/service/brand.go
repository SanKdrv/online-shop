package service

import "backend/internal/repository"

type BrandsService struct {
	repo repository.Brands
}

func NewBrandsService(repo repository.Brands) *BrandsService {
	return &BrandsService{repo: repo}
}

// GetIDByBrand retrieves the ID of a brand by its name.
//
// Returns the brand ID and an error if the operation fails or if the brand does not exist.
func (s *BrandsService) GetIDByBrand(name string) (int64, error) {
	return s.repo.GetIDByBrand(name)
}

// GetBrandByID returns brand name by brand id.
//
// Returns empty string if brand does not exist in DB.
func (s *BrandsService) GetBrandByID(brandId int64) (string, error) {
	return s.repo.GetBrandByID(brandId)
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
