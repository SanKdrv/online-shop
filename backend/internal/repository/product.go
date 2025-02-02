package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type ProductsRepo struct {
	db *gorm.DB
}

func NewProductsRepo(db *gorm.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (r *ProductsRepo) CreateProduct(product domain.Product) (int64, error) {
	return 0, nil
}

func (r *ProductsRepo) Get(name string, brandID int64, categoryID int64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *ProductsRepo) GetAllByCategory(categoryID int64) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *ProductsRepo) GetAllByName(name string) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *ProductsRepo) GetAllByBrand(brandID int64) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *ProductsRepo) UpdateProduct(product domain.Product) error {
	return nil
}

func (r *ProductsRepo) DeleteProduct(productID int64) error {
	return nil
}
