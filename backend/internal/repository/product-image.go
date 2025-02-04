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

func (r *ProductsImagesRepo) GetImageHashByProductId(productId int64) (string, error) {
	var hash string
	if err := r.db.First(&domain.ProductImage{}, "product_id = ?", productId).Pluck("image_hash", &hash).Error; err != nil {
		return "", err
	}
	return hash, nil
}

func (r *ProductsImagesRepo) CreateProductImage(productImage domain.ProductImage) (int64, error) {
	if err := r.db.Create(&productImage).Error; err != nil {
		return 0, err
	}
	return productImage.Id, nil
}

func (r *ProductsImagesRepo) UpdateProductImage(oldName string, productImage domain.ProductImage) error {
	return r.db.First(&domain.ProductImage{}, "image_hash = ?", oldName).Updates(&productImage).Error
}

func (r *ProductsImagesRepo) DeleteProductImageByName(name string) error {
	return r.db.Delete(&domain.ProductImage{}, "image_hash = ?", name).Error
}

func (r *ProductsImagesRepo) DeleteProductImageById(imageId int64) error {
	return r.db.Delete(&domain.ProductImage{}, "id = ?", imageId).Error
}
