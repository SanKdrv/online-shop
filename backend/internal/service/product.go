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

func (s *ProductsService) Get(name string, brandId int64, categoryId int64) (domain.Product, error) {
	return s.repo.Get(name, brandId, categoryId)
}

func (s *ProductsService) GetAllByCategory(categoryId int64) ([]domain.Product, error) {
	return s.repo.GetAllByCategory(categoryId)
}

func (s *ProductsService) GetAllByName(name string) ([]domain.Product, error) {
	return s.repo.GetAllByName(name)
}

func (s *ProductsService) GetAllByBrand(brandId int64) ([]domain.Product, error) {
	return s.repo.GetAllByBrand(brandId)
}

func (s *ProductsService) UpdateProduct(product domain.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *ProductsService) DeleteProduct(productId int64) error {
	return s.repo.DeleteProduct(productId)
}
