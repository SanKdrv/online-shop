package types

import "backend/internal/domain"

type CreateProductRequest struct {
	Product domain.Product `json:"product"`
}

type CreateProductResponse struct {
	ProductId int64 `json:"product_id"`
}

type GetProductRequest struct {
	Name       string `json:"name"`
	BrandId    int64  `json:"brand_id"`
	CategoryId int64  `json:"category_id"`
}

type GetProductResponse struct {
	Product domain.Product `json:"product"`
}

type GetAllByCategoryRequest struct {
	CategoryId int64 `json:"category_id"`
}

type GetAllByCategoryResponse struct {
	Products []domain.Product `json:"products"`
}

type GetAllByNameRequest struct {
	ProductName string `json:"product_name"`
}

type GetAllByNameResponse struct {
	Products []domain.Product `json:"products"`
}

type GetAllByBrandRequest struct {
	BrandId int64 `json:"brand_id"`
}

type GetAllByBrandResponse struct {
	Products []domain.Product `json:"products"`
}

type UpdateProductRequest struct {
	Product domain.Product `json:"product"`
}

type UpdateProductResponse struct {
	Status string `json:"status"`
}

type DeleteProductRequest struct {
	ProductId int64 `json:"product_id"`
}

type DeleteProductResponse struct {
	Status string `json:"status"`
}
