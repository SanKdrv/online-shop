package routes

import (
	"backend/internal/domain"
	"backend/internal/lib/api/response"
	"backend/internal/types"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

// @Summary getCartContentById
// @Tags Cart Content
// @Description getting single content from cart-content table
// @ID get-cart-content-by-id
// @Accept  json
// @Produce  json
// @Param id query int64 true "ID кортежа в корзине"
// @Success 200 {object} types.GetCartContentByIdResponse
// @Failure 400,404 {object} types.GetCartContentByIdResponse
// @Failure 500 {object} types.GetCartContentByIdResponse
// @Failure default {object} types.GetCartContentByIdResponse
// @Router /api/cart-content/get-cart-content-by-id [get]
func (h *Handler) getCartContentById(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.cart-content.getCartContentById"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			log.Error("missing id", slog.String("error", "missing id"))
			render.JSON(w, r, response.Error("missing id"))
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			render.JSON(w, r, response.Error("Invalid id"))
			return
		}

		cartContent, err := h.services.CartsContent.GetCartContentById(id)
		if err != nil {
			log.Error("failed get cart content", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetCartContentByIdResponse{
			CartContent: cartContent,
		})
	}
}

// @Summary getCartContentByUserId
// @Tags Cart Content
// @Description getting cart contents
// @ID get-cart-content-by-user-id
// @Accept  json
// @Produce  json
// @Param user_id query int64 true "ID пользователя"
// @Success 200 {object} types.GetCartContentByUserIdResponse
// @Failure 400,404 {object} types.GetCartContentByUserIdResponse
// @Failure 500 {object} types.GetCartContentByUserIdResponse
// @Failure default {object} types.GetCartContentByUserIdResponse
// @Router /api/cart-content/get-cart-content-by-user-id [get]
func (h *Handler) getCartContentByUserId(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.cart-content.getCartContentByUserId"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		userIdStr := r.URL.Query().Get("user_id")
		if userIdStr == "" {
			log.Error("missing user id", slog.String("error", "missing user id"))
			render.JSON(w, r, response.Error("missing user id"))
			return
		}

		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			render.JSON(w, r, response.Error("Invalid user id"))
			return
		}

		cartContent, err := h.services.CartsContent.GetCartContentByUserId(userId)
		if err != nil {
			log.Error("failed get cart content", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.GetCartContentByUserIdResponse{
			CartContent: cartContent,
		})
	}
}

// @Summary createCartContent
// @Tags Cart Content
// @Description creating cart content
// @ID create-cart-content
// @Accept  json
// @Produce  json
// @Param input body types.CreateCartContentRequest true "Создаёт запись в корзине пользователя"
// @Success 200 {object} types.CreateCartContentResponse
// @Failure 400,404 {object} types.CreateCartContentResponse
// @Failure 500 {object} types.CreateCartContentResponse
// @Failure default {object} types.CreateCartContentResponse
// @Router /api/cart-content/create-cart-content [post]
func (h *Handler) createCartContent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.cart-content.createCartContent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		// Декодируем входящий JSON
		var req types.CreateCartContentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		// Преобразуем Request в domain.User
		cartContent := domain.CartContent{
			Id:        req.Id,
			UserId:    req.UserId,
			ProductId: req.ProductId,
			Count:     req.Count,
		}

		// Создание пользователя
		cartContentId, err := h.services.CartsContent.CreateCartContent(cartContent)
		if err != nil {
			log.Error("failed to create cart content", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.CreateCartContentResponse{
			CartContentId: cartContentId,
		})
	}
}

// @Summary updateCartContent
// @Tags Cart Content
// @Description updating cart content
// @ID update-cart-content
// @Accept  json
// @Produce  json
// @Param input body types.UpdateCartContentRequest true "Обновляет запись в корзине пользователя"
// @Success 200 {object} types.UpdateCartContentResponse
// @Failure 400,404 {object} types.UpdateCartContentResponse
// @Failure 500 {object} types.UpdateCartContentResponse
// @Failure default {object} types.UpdateCartContentResponse
// @Router /api/cart-content/update-cart-content [put]
func (h *Handler) updateCartContent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.cart-content.createCartContent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		// Декодируем входящий JSON
		var req types.UpdateCartContentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		// Преобразуем Request в domain.User
		cartContent := domain.CartContent{
			Id:    req.Id,
			Count: req.Count,
		}

		// Создание пользователя
		err := h.services.CartsContent.UpdateCartContent(cartContent)
		if err != nil {
			log.Error("failed to update cart content", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateCartContentResponse{
			Status: "OK",
		})
	}
}

// @Summary deleteCartContent
// @Tags Cart Content
// @Description deleting cart content
// @ID delete-cart-content
// @Accept  json
// @Produce  json
// @Param input body types.DeleteCartContentRequest true "Удаляет запись в корзине пользователя"
// @Success 200 {object} types.DeleteCartContentResponse
// @Failure 400,404 {object} types.DeleteCartContentResponse
// @Failure 500 {object} types.DeleteCartContentResponse
// @Failure default {object} types.DeleteCartContentResponse
// @Router /api/cart-content/delete-cart-content [delete]
func (h *Handler) deleteCartContent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.cart-content.deleteCartContent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		// Декодируем входящий JSON
		var req types.DeleteCartContentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		// Преобразуем Request в domain.User
		cartContentId := req.Id

		// Создание пользователя
		err := h.services.CartsContent.DeleteCartContent(cartContentId)
		if err != nil {
			log.Error("failed to delete cart content", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteCartContentResponse{
			Status: "OK",
		})
	}
}
