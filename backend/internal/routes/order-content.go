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

// @Summary createOrderContent
// @Tags Order Content
// @Description creating order content
// @ID create-order-content
// @Accept  json
// @Produce  json
// @Param input body types.CreateOrderContentRequest true "Создаёт заказ"
// @Success 200 {object} types.CreateOrderContentResponse
// @Failure 400,404 {object} types.CreateOrderContentResponse
// @Failure 500 {object} types.CreateOrderContentResponse
// @Failure default {object} types.CreateOrderContentResponse
// @Router /api/order-content/create-order-content [post]
func (h *Handler) createOrderContent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order-content.createOrderContent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.CreateOrderContentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		contentId, err := h.services.OrdersContent.CreateOrderContent(req.OrderContent)
		if err != nil {
			log.Error("failed to create order content", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.CreateOrderContentResponse{
			ContentId: contentId,
		})
	}
}

// @Summary updateOrderContent
// @Tags Order Content
// @Description updating order content
// @ID update-order-content
// @Accept  json
// @Produce  json
// @Param input body types.UpdateOrderContentRequest true "Обновляет заказ"
// @Success 200 {object} types.UpdateOrderContentResponse
// @Failure 400,404 {object} types.UpdateOrderContentResponse
// @Failure 500 {object} types.UpdateOrderContentResponse
// @Failure default {object} types.UpdateOrderContentResponse
// @Router /api/order-content/update-order-content [put]
func (h *Handler) updateOrderContent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order-content.updateOrderContent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.UpdateOrderContentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.OrdersContent.UpdateOrderContent(req.OrderContent)
		if err != nil {
			log.Error("failed to update order content", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateOrderContentResponse{
			Status: "OK",
		})
	}
}

// @Summary deleteOrderContent
// @Tags Order Content
// @Description deleting order content
// @ID delete-order-content
// @Accept  json
// @Produce  json
// @Param input body types.DeleteOrderContentRequest true "Удаляет заказ"
// @Success 200 {object} types.DeleteOrderContentResponse
// @Failure 400,404 {object} types.DeleteOrderContentResponse
// @Failure 500 {object} types.DeleteOrderContentResponse
// @Failure default {object} types.DeleteOrderContentResponse
// @Router /api/order-content/delete-order-content [delete]
func (h *Handler) deleteOrderContent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order-content.deleteOrderContent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteOrderContentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.OrdersContent.DeleteOrderContent(req.ContentId)
		if err != nil {
			log.Error("failed to delete order content by id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteOrderContentResponse{
			Status: "OK",
		})
	}
}
