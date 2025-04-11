package types

import "backend/internal/domain"

type CreateOrderRequest struct {
	Order domain.Order `json:"order"`
}

type CreateOrderResponse struct {
	OrderId int64 `json:"order_id"`
}

type GetOrderByIdRequest struct {
	CategoryId int64 `json:"category_id"`
}

type GetOrderByIdResponse struct {
	Order domain.Order `json:"order"`
}

type GetOrdersByUserIdRequest struct {
	UserId int64 `json:"user_id"`
}

type GetOrdersByUserIdResponse struct {
	Orders []domain.Order `json:"orders"`
}

type UpdateOrderRequest struct {
	Order domain.Order `json:"order"`
}

type UpdateOrderResponse struct {
	Status string `json:"status"`
}

type DeleteOrderRequest struct {
	OrderId int64 `json:"order_id"`
}

type DeleteOrderResponse struct {
	Status string `json:"status"`
}

type GetAllOrdersResponse struct {
	Orders []domain.Order `json:"orders"`
	Total  int64          `json:"total"`
	Offset int64          `json:"offset"`
	Limit  int64          `json:"limit"`
}
