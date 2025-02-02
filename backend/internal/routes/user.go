package routes

import (
	"backend/internal/config"
	"backend/internal/domain"
	"backend/internal/lib/api/response"
	"backend/internal/types"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

type Request = types.Request
type RefreshTokensRequest = types.RefreshTokensRequest
type SignInRequest = types.SignInRequest
type SignUpRequest = types.SignUpRequest
type SignOutRequest = types.SignOutRequest
type Response = types.Response

type GetUsernameByIDRequest = types.GetUsernameByIDRequest
type GetUsernameByIDResponse = types.GetUsernameByIDResponse

// TODO: добавить разлогинирование при появлении более 5 различных подключений одного пользователя
// @Summary Sign In
// @Tags Auth
// @Description Авторизация пользователя с возвратом access и refresh токенов
// @ID sign-in-user
// @Accept  json
// @Produce  json
// @Param input body SignInRequest true "Информация для авторизации"
// @Success 200 {object} Response "Успешная авторизация"
// @Failure 400 {object} Response "Ошибка запроса"
// @Failure 401 {object} Response "Неверные учетные данные"
// @Failure 500 {object} Response "Внутренняя ошибка сервера"
// @Router /api/auth/sign-in [post]
func (h *Handler) signIn(log *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.user.signIn"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// Декодируем входящий JSON
		var req SignInRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		username := req.Username
		email := req.Email
		password := req.Password

		log.Info("request body decoded; ", username, email, password)

		if password == "" {
			log.Error("failed to login user", slog.String("error", "Password is empty"))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		if username == "" && email == "" {
			log.Error("failed to login user", slog.String("error", "Username or Email is empty"))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}
		user := domain.User{}

		if username != "" {
			localUserContext, err := h.services.Users.VerifyUser(username, password, "username")
			if err != nil {
				log.Error("failed to login user", slog.String("error", err.Error()))
				render.JSON(w, r, response.Error("Wrong username or password"))
				return
			}
			user = localUserContext
		} else if email != "" {
			localUserContext, err := h.services.Users.VerifyUser(email, password, "email")
			if err != nil {
				log.Error("failed to login user", slog.String("error", err.Error()))
				render.JSON(w, r, response.Error("Wrong username or password"))
				return
			}
			user = localUserContext
		}

		if (user == domain.User{}) {
			log.Error("failed to login user", slog.String("error", "User not found"))
			render.JSON(w, r, response.Error("Something went wrong"))
			return
		}

		//sess, err := h.services.RefreshSession.CreateRefreshSession(user)
		refreshSession := domain.RefreshSession{
			//ID:
			UserID:       user.ID,
			RefreshToken: uuid.New().String(),
			UserAgent:    req.UserAgent,
			ExpiresIn:    time.Now().Add(cfg.Auth.RefreshTokenTTL).Unix(),
			CreatedAt:    time.Now(),
		}

		if accessToken, refreshToken, err := h.services.RefreshSession.CreateRefreshSession(user, refreshSession, cfg); err != nil {
			log.Error("failed to create refresh session", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		} else {
			render.JSON(w, r, map[string]interface{}{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			})
		}
		return
	}
}

// @Summary SignUp
// @Tags Auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body SignUpRequest true "account info"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/auth/sign-up [post]
func (h *Handler) signUp(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.user.signUp"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// Декодируем входящий JSON
		var req SignUpRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		// Преобразуем Request в domain.User
		user := domain.User{
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: req.Password,
		}

		// Создание пользователя
		_, err := h.services.Users.CreateUser(user)
		if err != nil {
			log.Error("failed to create user", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		// Успешный ответ
		render.JSON(w, r, map[string]interface{}{
			"status": response.StatusOK,
		})
	}
}

// @Summary RefreshTokens
// @Tags Auth
// @Description refresh access & refresh tokens
// @ID refresh-tokens
// @Accept  json
// @Produce  json
// @Param input body RefreshTokensRequest true "refresh"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/auth/refresh-tokens [post]
func (h *Handler) refreshTokens(log *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.user.refreshTokens"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// Декодируем входящий JSON
		var req types.RefreshTokensRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		// Преобразуем Request в refreshToken
		oldRefreshToken := domain.RefreshToken{
			RefreshToken: req.RefreshToken,
		}
		if oldRefreshToken.RefreshToken == "" {
			log.Error("refresh token is empty")
			render.JSON(w, r, response.Error("Refresh token is empty"))
			return
		}

		if accessToken, refreshToken, err := h.services.RefreshSession.UpdateRefreshSession(oldRefreshToken, req, cfg); err != nil {
			log.Error("failed to update session", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		} else {
			render.JSON(w, r, map[string]interface{}{
				"access_token":  accessToken.AccessToken,
				"refresh_token": refreshToken.RefreshToken,
			})
		}
		return
	}
}

// TODO: Просто получаем рефреш токен и удаляем его из таблицы
// @Summary SignOut
// @Tags Auth
// @Description delete refresh token
// @ID delete-refresh-tokens
// @Accept  json
// @Produce  json
// @Param input body SignOutRequest true "Удаляет рефреш токен из таблицы refresh_tokens"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/auth/sign-out [post]
func (h *Handler) signOut(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.user.signUp"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req SignOutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		if err := h.services.RefreshSession.DeleteRefreshSession(req.RefreshToken); err != nil {
			log.Error("failed to delete session", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"status": response.StatusOK,
		})
	}
}

// @Summary GetUsernameByID
// @Tags User
// @Description get username by user id
// @ID get-username-by-id
// @Accept  json
// @Produce  json
// @Param input body GetUsernameByIDRequest true "Возвращает имя пользователя по user id"
// @Success 200 {object} GetUsernameByIDResponse
// @Failure 400,404 {object} GetUsernameByIDResponse
// @Failure 500 {object} GetUsernameByIDResponse
// @Failure default {object} GetUsernameByIDResponse
// @Router /api/user/get-username-by-id [post]
func (h *Handler) getUsernameByID(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.user.getUsernameByID"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req GetUsernameByIDRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		username, err := h.services.Users.GetUsernameByID(req.UserID)
		if err != nil {
			log.Error("failed to get username by id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, GetUsernameByIDResponse{
			Username: username,
		})
	}
}
