package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"time"
)

type ProductsImagesService struct {
	repo repository.ProductsImages
}

func NewProductsImagesService(repo repository.ProductsImages) *ProductsImagesService {
	return &ProductsImagesService{repo: repo}
}

func (s *ProductsImagesService) GetImageIdByHash(imageHash string) (int64, error) {
	return s.repo.GetImageIdByHash(imageHash)
}
func (s *ProductsImagesService) GetImageHashByImageId(imageId int64) (string, error) {
	return s.repo.GetImageHashByImageId(imageId)
}

func (s *ProductsImagesService) GetImageHashesByProductId(productId int64) ([]string, error) {
	return s.repo.GetImageHashesByProductId(productId)
}

func (s *ProductsImagesService) CreateProductImage(productId int64, hashString string) (int64, error) {
	var productImage = domain.ProductImage{
		ProductId: productId,
		ImageHash: hashString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.repo.CreateProductImage(productImage)
}

func (s *ProductsImagesService) UpdateProductImage(oldName string, productImage domain.ProductImage) error {
	return s.repo.UpdateProductImage(oldName, productImage)
}

func (s *ProductsImagesService) DeleteProductImageByName(name string) error {
	return s.repo.DeleteProductImageByName(name)
}

func (s *ProductsImagesService) DeleteProductImageById(imageId int64) error {
	return s.repo.DeleteProductImageById(imageId)
}
