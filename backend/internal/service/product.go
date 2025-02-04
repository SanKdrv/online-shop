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
	return s.repo.CreateProduct(product)
}

func (s *ProductsService) Get(name string, brandID int64, categoryID int64) (domain.Product, error) {
	return s.repo.Get(name, brandID, categoryID)
}

func (s *ProductsService) GetAllByCategory(categoryID int64) ([]domain.Product, error) {
	return s.repo.GetAllByCategory(categoryID)
}

func (s *ProductsService) GetAllByName(name string) ([]domain.Product, error) {
	return s.repo.GetAllByName(name)
}

func (s *ProductsService) GetAllByBrand(brandID int64) ([]domain.Product, error) {
	return s.repo.GetAllByBrand(brandID)
}

func (s *ProductsService) UpdateProduct(product domain.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *ProductsService) DeleteProduct(productID int64) error {
	return s.repo.DeleteProduct(productID)
}
