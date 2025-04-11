package types

import "backend/internal/domain"

type CreateProductRequest struct {
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	BrandId     int64   `json:"brand_id"`
	CategoryId  int64   `json:"category_id"`
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
	Page     int              `json:"page,omitempty"`
	Limit    int              `json:"limit,omitempty"`
}

type GetAllByNameRequest struct {
	ProductName string `json:"product_name"`
}

type GetAllByNameResponse struct {
	Products []domain.Product `json:"products"`
	Page     int              `json:"page,omitempty"`
	Limit    int              `json:"limit,omitempty"`
}

type GetAllByBrandRequest struct {
	BrandId int64 `json:"brand_id"`
}

type GetAllByBrandResponse struct {
	Products []domain.Product `json:"products"`
}

type UpdateProductRequest struct {
	ProductId   int64   `json:"product_id"`
	Description string  `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	BrandId     int64   `json:"brand_id,omitempty"`
	CategoryId  int64   `json:"category_id,omitempty"`
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

type GetAllProductsResponse struct {
	Products []domain.Product `json:"products"`
	Total    int64            `json:"total"`
	Offset   int64            `json:"offset"`
	Limit    int64            `json:"limit"`
}
