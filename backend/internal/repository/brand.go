package repository

import "gorm.io/gorm"

type BrandsRepo struct {
	db *gorm.DB
}

func NewBrandsRepo(db *gorm.DB) *BrandsRepo {
	return &BrandsRepo{
		db: db,
	}
}

func (r *BrandsRepo) GetIDByBrand(name string) (int64, error) {
	return 0, nil
}

func (r *BrandsRepo) GetBrandByID(categoryID int64) (string, error) {
	return "", nil
}

func (r *BrandsRepo) CreateBrand(name string) (int64, error) {
	return 0, nil
}

func (r *BrandsRepo) DeleteBrand(categoryID int64) error {
	return nil
}

func (r *BrandsRepo) UpdateBrand(categoryID int64, name string) error {
	return nil
}
