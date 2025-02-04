package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type ProductsImagesService struct {
	repo repository.ProductsImages
}

func NewProductsImagesService(repo repository.ProductsImages) *ProductsImagesService {
	return &ProductsImagesService{repo: repo}
}

func (s *ProductsImagesService) GetImageHashByProductId(productId int64) (string, error) {
	return s.repo.GetImageHashByProductId(productId)
}

func (s *ProductsImagesService) CreateProductImage(productImage domain.ProductImage) (int64, error) {
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
