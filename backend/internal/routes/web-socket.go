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

type SendMessageRequest = types.SendMessageRequest
type SendMessageResponse = types.SendMessageResponse

// @Summary sendMessage
// @Tags Chat
// @Description Отправка сообщения в чат
// @ID send-message
// @Accept  json
// @Produce  json
// @Param input body SendMessageRequest true "Информация для авторизации"
// @Success 200 {object} SendMessageResponse "Успешная отправка сообщения"
// @Failure 400 {object} SendMessageResponse "Ошибка запроса"
// @Failure 401 {object} SendMessageResponse "Неверные учетные данные"
// @Failure 500 {object} SendMessageResponse "Внутренняя ошибка сервера"
// @Router /stream/chat/send-message [post]
func (h *Handler) sendMessage(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.web-socket.sendMessage"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// Декодируем входящий JSON
		var req SendMessageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		msg := req.Message

		if err := h.services.WebSocket.SendMessage(msg); err != nil {
			log.Error("failed to send message", slog.String("error", err.Error()))
			render.JSON(w, r, map[string]interface{}{
				"response": err.Error(),
			})
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"response": "OK",
		})
		return
	}
}
