package routes

import (
	"backend/internal/lib/api/response"
	"backend/internal/types"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

// @Summary createProduct
// @Tags Product
// @Description creating product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param input body types.CreateProductRequest true "Создаёт товар"
// @Success 200 {object} types.CreateProductResponse
// @Failure 400,404 {object} types.CreateProductResponse
// @Failure 500 {object} types.CreateProductResponse
// @Failure default {object} types.CreateProductResponse
// @Router /api/product/create-product [post]
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

		productId, err := h.services.Products.CreateProduct(req.Description, req.Name, req.Price, req.CategoryId, req.BrandId)
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

// @Summary getProduct
// @Tags Product
// @Description getting product
// @ID get-product
// @Accept  json
// @Produce  json
// @Param product_name query string true "Название товара"
// @Param brand_id query int64 true "ID бренда"
// @Param category_id query int64 true "ID категории"
// @Success 200 {object} types.GetProductResponse
// @Failure 400,404 {object} types.GetProductResponse
// @Failure 500 {object} types.GetProductResponse
// @Failure default {object} types.GetProductResponse
// @Router /api/product/get-product [get]
func (h *Handler) getProduct(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getProduct"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		productName := r.URL.Query().Get("product_name")
		if productName == "" {
			log.Error("missed product_name", slog.String("error", "missed product_name"))
			render.JSON(w, r, response.Error("Missed product_name"))
			return
		}

		brandIdStr := r.URL.Query().Get("brand_id")
		if brandIdStr == "" {
			log.Error("missed brand_id", slog.String("error", "missed brand_id"))
			render.JSON(w, r, response.Error("Missed brand_id"))
			return
		}

		brandId, err := strconv.ParseInt(brandIdStr, 10, 64)
		if err != nil {
			log.Error("invalid brand_id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("invalid brand_id"))
			return
		}

		categoryIdStr := r.URL.Query().Get("category_id")
		if brandIdStr == "" {
			log.Error("missed category_id", slog.String("error", "missed category_id"))
			render.JSON(w, r, response.Error("Missed category_id"))
			return
		}

		categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
		if err != nil {
			log.Error("invalid category_id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("invalid category_id"))
			return
		}

		product, err := h.services.Products.Get(productName, brandId, categoryId)
		if err != nil {
			log.Error("failed to get product", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetProductResponse{
			Product: product,
		})
	}
}

// @Summary getAllByCategory
// @Tags Product
// @Description get all products by category
// @ID get-all-by-category
// @Accept  json
// @Produce  json
// @Param category_id query int64 true "ID категории"
// @Success 200 {object} types.GetAllByCategoryResponse
// @Failure 400,404 {object} types.GetAllByCategoryResponse
// @Failure 500 {object} types.GetAllByCategoryResponse
// @Failure default {object} types.GetAllByCategoryResponse
// @Router /api/product/get-all-by-category [get]
func (h *Handler) getAllByCategory(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getAllByCategory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		categoryIdStr := r.URL.Query().Get("category_id")
		if categoryIdStr == "" {
			log.Error("missed category_id", slog.String("error", "missing category_id"))
			render.JSON(w, r, response.Error("missed category_id"))
			return
		}

		categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
		if err != nil {
			log.Error("invalid category_id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("invalid category_id"))
			return
		}

		products, err := h.services.Products.GetAllByCategory(categoryId)
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

// @Summary getAllByName
// @Tags Product
// @Description getting all products by name
// @ID get-all-by-name
// @Accept  json
// @Produce  json
// @Param product_name query string true "Название товара"
// @Success 200 {object} types.GetAllByNameResponse
// @Failure 400,404 {object} types.GetAllByNameResponse
// @Failure 500 {object} types.GetAllByNameResponse
// @Failure default {object} types.GetAllByNameResponse
// @Router /api/product/get-all-by-name [get]
func (h *Handler) getAllByName(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getAllByName"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		productName := r.URL.Query().Get("product_name")
		if productName == "" {
			log.Error("missing product_name", slog.String("error", "missing product_name"))
			render.JSON(w, r, response.Error("missing product_name"))
			return
		}

		products, err := h.services.Products.GetAllByName(productName)
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

// @Summary getAllByBrand
// @Tags Product
// @Description getting all products by brand
// @ID get-all-by-brand
// @Accept  json
// @Produce  json
// @Param brand_id formData int64 true "ID бренда"
// @Success 200 {object} types.GetAllByBrandResponse
// @Failure 400,404 {object} types.GetAllByBrandResponse
// @Failure 500 {object} types.GetAllByBrandResponse
// @Failure default {object} types.GetAllByBrandResponse
// @Router /api/product/get-all-by-brand [get]
func (h *Handler) getAllByBrand(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product.getAllByBrand"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		userIdStr := r.FormValue("brand_id")
		if userIdStr == "" {
			log.Error("missing brand_id", slog.String("error", "missing brand_id"))
			render.JSON(w, r, response.Error("Missing brand_id"))
			return
		}

		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			log.Error("wrong brand_id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid brand_id"))
			return
		}

		products, err := h.services.Products.GetAllByBrand(userId)
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

// @Summary updateProduct
// @Tags Product
// @Description updating product by id
// @ID update-product
// @Accept  json
// @Produce  json
// @Param input body types.UpdateProductRequest true "Обновляет товар по айди"
// @Success 200 {object} types.UpdateProductResponse
// @Failure 400,404 {object} types.UpdateProductResponse
// @Failure 500 {object} types.UpdateProductResponse
// @Failure default {object} types.UpdateProductResponse
// @Router /api/product/update-product [put]
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

		err := h.services.Products.UpdateProductById(req.ProductId, req.Description, req.Name, req.Price, req.CategoryId, req.BrandId)
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

// @Summary deleteProduct
// @Tags Product
// @Description deleting product by id
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param input body types.DeleteProductRequest true "Удаляет товар по айди"
// @Success 200 {object} types.DeleteProductResponse
// @Failure 400,404 {object} types.DeleteProductResponse
// @Failure 500 {object} types.DeleteProductResponse
// @Failure default {object} types.DeleteProductResponse
// @Router /api/product/delete-product [delete]
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
