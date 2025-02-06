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

// @Summary getIdByBrand
// @Tags Brand
// @Description getting brand id by brand name
// @ID get-brand-id-by-brand-name
// @Accept  json
// @Produce  json
// @Param brand_name query string true "Название бренда"
// @Success 200 {object} types.GetIdByBrandResponse
// @Failure 400,404 {object} types.GetIdByBrandResponse
// @Failure 500 {object} types.GetIdByBrandResponse
// @Failure default {object} types.GetIdByBrandResponse
// @Router /api/brand/get-id-by-brand [get]
func (h *Handler) getIdByBrand(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.brand.getIdByBrand"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		brandName := r.URL.Query().Get("brand_name")
		if brandName == "" {
			log.Error("missing brand_name in query params")
			render.JSON(w, r, response.Error("Missing brand_name parameter"))
			return
		}

		brandId, err := h.services.Brands.GetIdByBrand(brandName)
		if err != nil {
			log.Error("failed to get brand name by id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetIdByBrandResponse{
			BrandId: brandId,
		})
	}
}

// @Summary getBrandById
// @Tags Brand
// @Description getting brand id by brand name
// @ID get-brand-name-by-brand-id
// @Accept  json
// @Produce  json
// @Param brand_id query int64 true "id бренда"
// @Success 200 {object} types.GetBrandByIdResponse
// @Failure 400,404 {object} types.GetBrandByIdResponse
// @Failure 500 {object} types.GetBrandByIdResponse
// @Failure default {object} types.GetBrandByIdResponse
// @Router /api/brand/get-brand-by-id [get]
func (h *Handler) getBrandById(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.brand.getBrandById"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		brandId, err := strconv.ParseInt(r.URL.Query().Get("brand_id"), 10, 64)
		if err != nil {
			log.Error("missing brand_id in query params", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Missing brand_id parameter"))
			return
		}

		brandName, err := h.services.Brands.GetBrandById(brandId)
		if err != nil {
			log.Error("failed to get brand id by name", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetBrandByIdResponse{
			BrandName: brandName,
		})
	}
}

// @Summary createBrand
// @Tags Brand
// @Description creating brand
// @ID create-brand
// @Accept  json
// @Produce  json
// @Param input body types.CreateBrandRequest true "Создаёт бренд"
// @Success 200 {object} types.CreateBrandResponse
// @Failure 400,404 {object} types.CreateBrandResponse
// @Failure 500 {object} types.CreateBrandResponse
// @Failure default {object} types.CreateBrandResponse
// @Router /api/brand/create-brand [post]
func (h *Handler) createBrand(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.brand.createBrand"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.CreateBrandRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		brandId, err := h.services.Brands.CreateBrand(req.BrandName)
		if err != nil {
			log.Error("failed to create brand", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.CreateBrandResponse{
			BrandId: brandId,
		})
	}
}

// @Summary deleteBrand
// @Tags Brand
// @Description deleting brand
// @ID delete-brand
// @Accept  json
// @Produce  json
// @Param input body types.DeleteBrandRequest true "Удаляет бренд"
// @Success 200 {object} types.DeleteBrandResponse
// @Failure 400,404 {object} types.DeleteBrandResponse
// @Failure 500 {object} types.DeleteBrandResponse
// @Failure default {object} types.DeleteBrandResponse
// @Router /api/brand/delete-brand [delete]
func (h *Handler) deleteBrand(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.brand.deleteBrand"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteBrandRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.Brands.DeleteBrand(req.BrandId)
		if err != nil {
			log.Error("failed to delete brand", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteBrandResponse{
			Status: "OK",
		})
	}
}

// @Summary updateBrand
// @Tags Brand
// @Description updating brand
// @ID update-brand
// @Accept  json
// @Produce  json
// @Param input body types.UpdateBrandRequest true "Обновляет бренд"
// @Success 200 {object} types.UpdateBrandResponse
// @Failure 400,404 {object} types.UpdateBrandResponse
// @Failure 500 {object} types.UpdateBrandResponse
// @Failure default {object} types.UpdateBrandResponse
// @Router /api/brand/update-brand [put]
func (h *Handler) updateBrand(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.brand.updateBrand"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.UpdateBrandRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		var brandId int64
		brandId, err := h.services.Brands.GetIdByBrand(req.OldBrandName)
		if err != nil {
			log.Error("failed to find brand", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		err = h.services.Brands.UpdateBrand(brandId, req.NewBrandName)
		if err != nil {
			log.Error("failed to update brand", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateBrandResponse{
			Status: "OK",
		})
	}
}
