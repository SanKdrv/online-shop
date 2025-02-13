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

func (r *ProductsImagesRepo) GetSequenceByProductId(productId int64) (int64, error) {
	var maxOrder int64
	if err := r.db.Model(&domain.ProductImage{}).
		Where("product_id = ?", productId).
		Select("COALESCE(MAX(image_order), 0)").
		Scan(&maxOrder).Error; err != nil {
		return -1, err
	}
	return maxOrder, nil
}

func (r *ProductsImagesRepo) GetImageIdByHash(imageHash string) (int64, error) {
	var imgId int64
	if err := r.db.Model(&domain.ProductImage{}).Where("image_hash = ?", imageHash).Pluck("id", imgId).Error; err != nil {
		return 0, err
	}
	return imgId, nil
}

func (r *ProductsImagesRepo) GetImageHashByImageId(imageId int64) (string, error) {
	var hash string
	if err := r.db.Model(&domain.ProductImage{}).Where("id = ?", imageId).Pluck("image_hash", &hash).Error; err != nil {
		return "", err
	}
	return hash, nil
}

func (r *ProductsImagesRepo) GetImageHashesByProductId(productId int64) ([]string, error) {
	var hashes []string
	if err := r.db.Model(&domain.ProductImage{}).Where("product_id = ?", productId).Order("image_order ASC").Pluck("image_hash", &hashes).Error; err != nil {
		return []string{}, err
	}
	return hashes, nil
}

func (r *ProductsImagesRepo) CreateProductImage(productImage domain.ProductImage) (int64, error) {
	if err := r.db.Create(&productImage).Error; err != nil {
		return 0, err
	}
	return productImage.Id, nil
}

func (r *ProductsImagesRepo) UpdateProductImage(oldName string, newName string) error {
	var image domain.ProductImage
	if err := r.db.First(&image, "image_hash = ?", oldName).Error; err != nil {
		return err
	}

	return r.db.Model(&image).Update("image_hash", newName).Error
}

func (r *ProductsImagesRepo) DeleteProductImageByName(name string) error {
	return r.db.Delete(&domain.ProductImage{}, "image_hash = ?", name).Error
}

func (r *ProductsImagesRepo) DeleteProductImageById(imageId int64) error {
	return r.db.Delete(&domain.ProductImage{}, "id = ?", imageId).Error
}
