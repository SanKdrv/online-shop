package routes

import (
	"backend/internal/lib/api/response"
	"backend/internal/types"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

// @Summary createOrder
// @Tags Order
// @Description creating order
// @ID create-order
// @Accept  json
// @Produce  json
// @Param input body types.CreateOrderRequest true "Создаёт заказ"
// @Success 200 {object} types.CreateOrderResponse
// @Failure 400,404 {object} types.CreateOrderResponse
// @Failure 500 {object} types.CreateOrderResponse
// @Failure default {object} types.CreateOrderResponse
// @Router /api/order/create-order [post]
func (h *Handler) createOrder(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order.сreateOrder"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		orderId, err := h.services.Orders.CreateOrder(req.Order)
		if err != nil {
			log.Error("failed to create order", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.CreateOrderResponse{
			OrderId: orderId,
		})
	}
}

// TODO
// @Summary getOrderById
// @Tags Order
// @Description getting order by id
// @ID get-order-by-id
// @Accept  json
// @Produce  json
// @Param input body types.GetOrderByIdRequest true "Ищет заказ по id"
// @Success 200 {object} types.GetOrderByIdResponse
// @Failure 400,404 {object} types.GetOrderByIdResponse
// @Failure 500 {object} types.GetOrderByIdResponse
// @Failure default {object} types.GetOrderByIdResponse
// @Router /api/order/get-order-by-id [get]
func (h *Handler) getOrderById(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order.getOrderById"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.GetOrderByIdRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		order, err := h.services.Orders.GetOrderById(req.CategoryId)
		if err != nil {
			log.Error("failed to get order by order id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetOrderByIdResponse{
			Order: order,
		})
	}
}

// @Summary getOrdersByUserId
// @Tags Order
// @Description getting orders by user id
// @ID get-orders-by-user-id
// @Accept  json
// @Produce  json
// @Param input body types.GetOrdersByUserIdRequest true "Ищет все заказы по айди пользователя"
// @Success 200 {object} types.GetOrdersByUserIdResponse
// @Failure 400,404 {object} types.GetOrdersByUserIdResponse
// @Failure 500 {object} types.GetOrdersByUserIdResponse
// @Failure default {object} types.GetOrdersByUserIdResponse
// @Router /api/order/get-orders-by-user-id [get]
func (h *Handler) getOrdersByUserId(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order.getIdByCategory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.GetOrdersByUserIdRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		orders, err := h.services.Orders.GetOrdersByUserId(req.UserId)
		if err != nil {
			log.Error("failed to get orders by user id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetOrdersByUserIdResponse{
			Orders: orders,
		})
	}
}

// @Summary updateOrder
// @Tags Order
// @Description updating order
// @ID update-order
// @Accept  json
// @Produce  json
// @Param input body types.UpdateOrderRequest true "Обновляет заказ"
// @Success 200 {object} types.UpdateOrderResponse
// @Failure 400,404 {object} types.UpdateOrderResponse
// @Failure 500 {object} types.UpdateOrderResponse
// @Failure default {object} types.UpdateOrderResponse
// @Router /api/order/update-order [put]
func (h *Handler) updateOrder(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order.updateOrder"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.UpdateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.Orders.UpdateOrder(req.Order)
		if err != nil {
			log.Error("failed to update order", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateOrderResponse{
			Status: "OK",
		})
	}
}

// @Summary deleteOrder
// @Tags Order
// @Description deleting order
// @ID delete-order
// @Accept  json
// @Produce  json
// @Param input body types.DeleteOrderRequest true "Удаляет заказ"
// @Success 200 {object} types.DeleteOrderResponse
// @Failure 400,404 {object} types.DeleteOrderResponse
// @Failure 500 {object} types.DeleteOrderResponse
// @Failure default {object} types.DeleteOrderResponse
// @Router /api/order/delete-order [delete]
func (h *Handler) deleteOrder(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order.deleteOrder"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.Orders.DeleteOrder(req.OrderId)
		if err != nil {
			log.Error("failed to get category name by id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteOrderResponse{
			Status: "OK",
		})
	}
}
