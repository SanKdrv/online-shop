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

// @Summary сreateOrder
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
func (h *Handler) сreateOrder(log *slog.Logger) http.HandlerFunc {
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

		order, err := h.services.Orders.GetOrderById(req.CategoryName)
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
			Order: orders,
		})
	}
}

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
