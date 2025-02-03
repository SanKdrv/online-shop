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
	return "", nil
}

func (s *ProductsImagesService) CreateProductImage(productImage domain.ProductImage) (int64, error) {
	return 0, nil
}

func (s *ProductsImagesService) UpdateProductImage(oldName string, productImage domain.ProductImage) error {
	return nil
}

func (s *ProductsImagesService) DeleteProductImageByName(name string) error {
	return nil
}

func (s *ProductsImagesService) DeleteProductImageByID(imageID int64) error {
	return nil
}
