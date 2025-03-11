package types

import (
	"backend/internal/domain"
)

type GetCartContentByIdRequest struct {
}

type GetCartContentByIdResponse struct {
	CartContent domain.CartContent `json:"cart-content"`
}

type GetCartContentByUserIdRequest struct {
}

type GetCartContentByUserIdResponse struct {
	CartContent []domain.CartContent `json:"cart-content"`
}

type CreateCartContentRequest struct {
	Id        int64 `json:"id,omitempty"`
	UserId    int64 `json:"user_id"`
	ProductId int64 `json:"product_id"`
	Count     int   `json:"count"`
}

type CreateCartContentResponse struct {
	CartContentId int64 `json:"cart_content_id"`
}

type UpdateCartContentRequest struct {
	Id    int64 `json:"id"`
	Count int   `json:"count"`
}

type UpdateCartContentResponse struct {
	Status string `json:"status"`
}

type DeleteCartContentRequest struct {
	Id int64 `json:"id"`
}

type DeleteCartContentResponse struct {
	Status string `json:"status"`
}
