package types

import "backend/internal/domain"

type GetIdByCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type GetIdByCategoryResponse struct {
	CategoryId int64 `json:"category_id"`
}

type GetCategoryByIdRequest struct {
	CategoryId int64 `json:"category_id"`
}

type GetCategoryByIdResponse struct {
	CategoryName string `json:"category_name"`
}

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type CreateCategoryResponse struct {
	CategoryId int64 `json:"category_id"`
}

type DeleteCategoryRequest struct {
	CategoryId int64 `json:"category_id"`
}

type DeleteCategoryResponse struct {
	Status string `json:"status"`
}

type UpdateCategoryRequest struct {
	OldCategoryName string `json:"new_category_name"`
	NewCategoryName string `json:"old_category_name"`
}

type UpdateCategoryResponse struct {
	Status string `json:"status"`
}

type GetAllCategoriesResponse struct {
	Categories []domain.Category `json:"categories"`
	Total      int64             `json:"total"`
	Offset     int64             `json:"offset"`
	Limit      int64             `json:"limit"`
}
