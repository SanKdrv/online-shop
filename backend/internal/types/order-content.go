package types

import "backend/internal/domain"

type CreateOrderContentRequest struct {
	OrderContent domain.OrderContent `json:"order_content"`
}

type CreateOrderContentResponse struct {
	ContentId int64 `json:"content_id"`
}

type UpdateOrderContentRequest struct {
	OrderContent domain.OrderContent `json:"order_content"`
}

type UpdateOrderContentResponse struct {
	Status string `json:"status"`
}

type DeleteOrderContentRequest struct {
	ContentId int64 `json:"content_id"`
}

type DeleteOrderContentResponse struct {
	Status string `json:"status"`
}
