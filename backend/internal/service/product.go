package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"time"
)

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{repo: repo}
}

func (s *ProductsService) CreateProduct(description string, name string, price float64, categoryId int64, brandId int64) (int64, error) {
	product := domain.Product{
		BrandId:     brandId,
		CategoryId:  categoryId,
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
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

func (s *ProductsService) UpdateProductById(productId int64, description string, name string, price float64, categoryId int64, brandId int64) error {
	product := domain.Product{
		Id:          productId,
		BrandId:     brandId,
		CategoryId:  categoryId,
		Name:        name,
		Description: description,
		Price:       price,
		UpdatedAt:   time.Now(),
	}
	return s.repo.UpdateProduct(product)
}

func (s *ProductsService) DeleteProduct(productId int64) error {
	return s.repo.DeleteProduct(productId)
}
