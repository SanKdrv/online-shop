package routes

import (
	"backend/internal/domain"
	"backend/internal/lib/api/response"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

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

func (h *Handler) deleteOrderContent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.order-content.deleteOrderContent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteOrderContent
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
