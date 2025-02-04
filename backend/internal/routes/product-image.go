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

// @Summary getImageHashByProductId
// @Tags Product Image
// @Description getting image hash by product id
// @ID get-image-hash-by-product-id
// @Accept  json
// @Produce  json
// @Param input body types.GetImageHashByProductIdRequest true "Получает хэш-имя изображения по его id"
// @Success 200 {object} types.GetImageHashByProductIdResponse
// @Failure 400,404 {object} types.GetImageHashByProductIdResponse
// @Failure 500 {object} types.GetImageHashByProductIdResponse
// @Failure default {object} types.GetImageHashByProductIdResponse
// @Router /api/product-image/get-image-hash-by-product-id [get]
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

// @Summary createProductImage
// @Tags Product Image
// @Description creating database record
// @ID create-product-image
// @Accept  json
// @Produce  json
// @Param input body types.CreateProductImageRequest true "Создаёт запись об изображении в базе данных"
// @Success 200 {object} types.CreateProductImageResponse
// @Failure 400,404 {object} types.CreateProductImageResponse
// @Failure 500 {object} types.CreateProductImageResponse
// @Failure default {object} types.CreateProductImageResponse
// @Router /api/product-image/create-product-image [post]
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

// @Summary updateProductImage
// @Tags Product Image
// @Description updating database record
// @ID update-product-image
// @Accept  json
// @Produce  json
// @Param input body types.CreateProductImageRequest true "Обновляет запись об изображении в базе данных по id"
// @Success 200 {object} types.CreateProductImageResponse
// @Failure 400,404 {object} types.CreateProductImageResponse
// @Failure 500 {object} types.CreateProductImageResponse
// @Failure default {object} types.CreateProductImageResponse
// @Router /api/product-image/update-product-image [put]
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

// @Summary deleteProductImageByName
// @Tags Product Image
// @Description deleting database record by name
// @ID delete-product-image-by-name
// @Accept  json
// @Produce  json
// @Param input body types.DeleteProductImageByNameRequest true "Удаляет запись об изображении в базе данных по хэшу имени"
// @Success 200 {object} types.DeleteProductImageByNameResponse
// @Failure 400,404 {object} types.DeleteProductImageByNameResponse
// @Failure 500 {object} types.DeleteProductImageByNameResponse
// @Failure default {object} types.DeleteProductImageByNameResponse
// @Router /api/product-image/delete-product-image-by-name [delete]
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

// @Summary deleteProductImageById
// @Tags Product Image
// @Description deleting database record by id
// @ID delete-product-image-by-id
// @Accept  json
// @Produce  json
// @Param input body types.DeleteProductImageByIdRequest true "Удаляет запись об изображении в базе данных по id"
// @Success 200 {object} types.DeleteProductImageByIdResponse
// @Failure 400,404 {object} types.DeleteProductImageByIdResponse
// @Failure 500 {object} types.DeleteProductImageByIdResponse
// @Failure default {object} types.DeleteProductImageByIdResponse
// @Router /api/product-image/delete-product-image-by-id [delete]
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

		err := h.services.ProductsImages.DeleteProductImageById(req.ImageId)
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
