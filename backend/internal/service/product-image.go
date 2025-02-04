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

func (s *ProductsImagesService) GetImageHashByProductID(productID int64) (string, error) {
	return s.repo.GetImageHashByProductID(productID)
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

func (s *ProductsImagesService) DeleteProductImageByID(imageID int64) error {
	return s.repo.DeleteProductImageByID(imageID)
}
