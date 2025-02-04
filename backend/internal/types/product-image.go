package types

import "backend/internal/domain"

type GetImageHashByProductIdRequest struct {
	ProductId int64 `json:"product_id"`
}

type GetImageHashByProductIdResponse struct {
	Hash string `json:"hash"`
}

type CreateProductImageRequest struct {
	ProductImage domain.ProductImage `json:"product_image"`
}

type CreateProductImageResponse struct {
	RecordId int64 `json:"record_id"`
}

type UpdateProductImageRequest struct {
	OldHashName  string              `json:"old_hash_name"`
	ProductImage domain.ProductImage `json:"product_image"`
}

type UpdateProductImageResponse struct {
	Status string `json:"status"`
}

type DeleteProductImageByNameRequest struct {
	OldHashName string `json:"old_hash_name"`
}

type DeleteProductImageByNameResponse struct {
	Status string `json:"status"`
}

type DeleteProductImageByIdRequest struct {
	ImageId int64 `json:"record_id"`
}

type DeleteProductImageByIdResponse struct {
	Status string `json:"status"`
}
