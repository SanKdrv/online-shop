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

func (h *Handler) getImageHashByProductId(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.getImageHashByProductId"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.GetImageHashByProductIdRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		hash, err := h.services.ProductsImages.GetImageHashByProductId(req.ProductId)
		if err != nil {
			log.Error("failed to get image hash", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetImageHashByProductIdResponse{
			Hash: hash,
		})
	}
}

func (h *Handler) createProductImage(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.createProductImage"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.CreateProductImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		recordId, err := h.services.ProductsImages.CreateProductImage(req.ProductImage)
		if err != nil {
			log.Error("failed create record", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.CreateProductImageResponse{
			RecordId: recordId,
		})
	}
}

func (h *Handler) updateProductImage(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.updateProductImage"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.UpdateProductImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.ProductsImages.UpdateProductImage(req.OldHashName, req.ProductImage)
		if err != nil {
			log.Error("failed update record", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateProductImageResponse{
			Status: "OK",
		})
	}
}

func (h *Handler) deleteProductImageByName(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.deleteProductImageByName"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteProductImageByNameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.ProductsImages.DeleteProductImageByName(req.OldHashName)
		if err != nil {
			log.Error("failed delete record by hash name", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteProductImageByNameResponse{
			Status: "OK",
		})
	}
}

func (h *Handler) deleteProductImageById(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.deleteProductImageById"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteProductImageByIdRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.ProductsImages.DeleteProductImageById(req.OldHashName)
		if err != nil {
			log.Error("failed delete record by record id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteProductImageByIdResponse{
			Status: "OK",
		})
	}
}
