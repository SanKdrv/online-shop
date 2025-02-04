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

func (h *Handler) createProduct(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.createProduct"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.CreateProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		productId, err := h.services.Products.CreateProduct(req.Product)
		if err != nil {
			log.Error("failed to create product", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.CreateProductResponse{
			ProductId: productId,
		})
	}
}

func (h *Handler) getProduct(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getProduct"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.GetProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		product, err := h.services.Products.Get(req.Name, req.BrandId, req.CategoryId)
		if err != nil {
			log.Error("failed to get product", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.getProductResponse{
			Product: product,
		})
	}
}

func (h *Handler) getAllByCategory(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getAllByCategory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.GetAllByCategoryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		products, err := h.services.Products.GetAllByCategory(req.CategoryId)
		if err != nil {
			log.Error("failed to get products by category id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetAllByCategoryResponse{
			Products: products,
		})
	}
}

func (h *Handler) getAllByName(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getAllByName"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.GetAllByNameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		products, err := h.services.Products.GetAllByName(req.ProductName)
		if err != nil {
			log.Error("failed to find products by name", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetAllByNameResponse{
			Products: products,
		})
	}
}

func (h *Handler) getAllByBrand(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getAllByBrand"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.GetAllByBrandRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		products, err := h.services.Products.GetAllByBrand(req.BrandId)
		if err != nil {
			log.Error("failed to find products by brand", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetAllByBrandResponse{
			Products: products,
		})
	}
}

func (h *Handler) updateProduct(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.updateProduct"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.UpdateProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.Products.UpdateProduct(req.OrderContent)
		if err != nil {
			log.Error("failed to update product", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateProductResponse{
			Status: "OK",
		})
	}
}

func (h *Handler) deleteProduct(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.deleteProduct"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.Products.DeleteProduct(req.ProductId)
		if err != nil {
			log.Error("failed to delete product", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteProductResponse{
			Status: "OK",
		})
	}
}
