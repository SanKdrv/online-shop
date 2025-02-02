package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type ProductsImagesRepo struct {
	db *gorm.DB
}

func NewProductsImagesRepo(db *gorm.DB) *ProductsImagesRepo {
	return &ProductsImagesRepo{
		db: db,
	}
}

func (r *ProductsImagesRepo) GetImageHashByProductID(productID int64) (string, error) {
	return "", nil
}

func (r *ProductsImagesRepo) CreateProductImage(productImage domain.ProductImage) (int64, error) {
	return 0, nil
}

func (r *ProductsImagesRepo) UpdateProductImage(oldName string, productImage domain.ProductImage) error {
	return nil
}

func (r *ProductsImagesRepo) DeleteProductImageByName(name string) error {
	return nil
}

func (r *ProductsImagesRepo) DeleteProductImageByID(imageID int64) error {
	return nil
}
