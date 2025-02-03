package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{repo: repo}
}

func (s *ProductsService) CreateProduct(product domain.Product) (int64, error) {
	return 0, nil
}

func (s *ProductsService) Get(name string, brandID int64, categoryID int64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (s *ProductsService) GetAllByCategory(categoryID int64) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (s *ProductsService) GetAllByName(name string) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (s *ProductsService) GetAllByBrand(brandID int64) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (s *ProductsService) UpdateProduct(product domain.Product) error {
	return nil
}

func (s *ProductsService) DeleteProduct(productID int64) error {
	return nil
}
