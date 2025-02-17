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

func (s *ProductsService) GetAllByCategoryPaginated(categoryId int64, page int, limit int) ([]domain.Product, error) {
	return s.repo.GetAllByCategoryPaginated(categoryId, page, limit)
}

func (s *ProductsService) GetAllByNamePaginated(name string, page int, limit int) ([]domain.Product, error) {
	return s.repo.GetAllByNamePaginated(name, page, limit)
}

func (s *ProductsService) GetAllByBrandPaginated(brandId int64, page int, limit int) ([]domain.Product, error) {
	return s.repo.GetAllByBrandPaginated(brandId, page, limit)
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
