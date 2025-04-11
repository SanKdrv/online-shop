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

// @Summary getAllCategories
// @Tags Category
// @Description getting all rows
// @ID get-all-categories
// @Accept  json
// @Produce  json
// @Param offset query int64 true "offset"
// @Param limit query int64 true "limit"
// @Success 200 {object} types.GetAllCategoriesResponse
// @Failure 400,404 {object} types.GetAllCategoriesResponse
// @Failure 500 {object} types.GetAllCategoriesResponse
// @Failure default {object} types.GetAllCategoriesResponse
// @Router /api/category/get-all-categories [get]
func (h *Handler) getAllCategories(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.brand.getAllCategories"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
		if err != nil || offset < 0 {
			offset = 0
		}

		limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
		if err != nil || limit <= 0 {
			limit = 10
		}

		categories, total, err := h.services.Categories.GetAll(offset, limit)
		if err != nil {
			log.Error("failed to get categories", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetAllCategoriesResponse{
			Categories: categories,
			Total:      total,
			Offset:     offset,
			Limit:      limit,
		})
	}
}

// @Summary getIdByCategory
// @Tags Category
// @Description getting category id by category name
// @ID get-category-id-by-category-name
// @Accept  json
// @Produce  json
// @Param category_name query string true "Название категории"
// @Success 200 {object} types.GetIdByCategoryResponse
// @Failure 400,404 {object} types.GetIdByCategoryResponse
// @Failure 500 {object} types.GetIdByCategoryResponse
// @Failure default {object} types.GetIdByCategoryResponse
// @Router /api/category/get-id-by-category [get]
func (h *Handler) getIdByCategory(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.category.getIdByCategory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		categoryName := r.URL.Query().Get("category_name")
		if categoryName == "" {
			log.Error("category name is empty", slog.String("error", "category name is empty"))
			render.JSON(w, r, response.Error("Category name is empty"))
			return
		}

		categoryId, err := h.services.Categories.GetIdByCategory(categoryName)
		if err != nil {
			log.Error("failed to get category name by id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetIdByCategoryResponse{
			CategoryId: categoryId,
		})
	}
}

// @Summary getCategoryById
// @Tags Category
// @Description getting category id by category name
// @ID get-category-name-by-category-id
// @Accept  json
// @Produce  json
// @Param category_id query int64 true "ID категории"
// @Success 200 {object} types.GetCategoryByIdResponse
// @Failure 400,404 {object} types.GetCategoryByIdResponse
// @Failure 500 {object} types.GetCategoryByIdResponse
// @Failure default {object} types.GetCategoryByIdResponse
// @Router /api/category/get-category-by-id [get]
func (h *Handler) getCategoryById(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.category.getCategoryById"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		categoryIdStr := r.URL.Query().Get("category_id")
		if categoryIdStr == "" {
			log.Error("category id is empty", slog.String("error", "category id is empty"))
			render.JSON(w, r, response.Error("Category id is empty"))
			return
		}

		categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
		if err != nil {
			log.Error("failed to parse category id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid category id"))
			return
		}

		categoryName, err := h.services.Categories.GetCategoryById(categoryId)
		if err != nil {
			log.Error("failed to get category id by name", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetCategoryByIdResponse{
			CategoryName: categoryName,
		})
	}
}

// @Summary createCategory
// @Tags Category
// @Description creating category
// @ID create-category
// @Accept  json
// @Produce  json
// @Param input body types.CreateCategoryRequest true "Создаёт категорию"
// @Success 200 {object} types.CreateCategoryResponse
// @Failure 400,404 {object} types.CreateCategoryResponse
// @Failure 500 {object} types.CreateCategoryResponse
// @Failure default {object} types.CreateCategoryResponse
// @Router /api/category/create-category [post]
func (h *Handler) createCategory(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.category.createCategory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.CreateCategoryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		categoryId, err := h.services.Categories.CreateCategory(req.CategoryName)
		if err != nil {
			log.Error("failed to create category", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.CreateCategoryResponse{
			CategoryId: categoryId,
		})
	}
}

// @Summary deleteCategory
// @Tags Category
// @Description deleting category
// @ID delete-category
// @Accept  json
// @Produce  json
// @Param input body types.DeleteCategoryRequest true "Удаляет категорию"
// @Success 200 {object} types.DeleteCategoryResponse
// @Failure 400,404 {object} types.DeleteCategoryResponse
// @Failure 500 {object} types.DeleteCategoryResponse
// @Failure default {object} types.DeleteCategoryResponse
// @Router /api/category/delete-category [delete]
func (h *Handler) deleteCategory(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.category.deleteCategory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteCategoryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := h.services.Categories.DeleteCategory(req.CategoryId)
		if err != nil {
			log.Error("failed to delete category", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteCategoryResponse{
			Status: "OK",
		})
	}
}

// @Summary updateCategory
// @Tags Category
// @Description updating category
// @ID update-category
// @Accept  json
// @Produce  json
// @Param input body types.UpdateCategoryRequest true "Обновляет категорию"
// @Success 200 {object} types.UpdateCategoryResponse
// @Failure 400,404 {object} types.UpdateCategoryResponse
// @Failure 500 {object} types.UpdateCategoryResponse
// @Failure default {object} types.UpdateCategoryResponse
// @Router /api/category/update-category [put]
func (h *Handler) updateCategory(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.category.updateCategory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.UpdateCategoryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		var categoryId int64
		categoryId, err := h.services.Categories.GetIdByCategory(req.OldCategoryName)
		if err != nil {
			log.Error("failed to find category", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		err = h.services.Categories.UpdateCategory(categoryId, req.NewCategoryName)
		if err != nil {
			log.Error("failed to update category", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateCategoryResponse{
			Status: "OK",
		})
	}
}
