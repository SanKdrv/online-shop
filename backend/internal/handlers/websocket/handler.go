package websocket

import (
	"backend/internal/domain"
	"backend/internal/service"
	"backend/internal/types"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
)

// MessageRequest представляет структуру входящего JSON сообщения
type MessageRequest = types.MessageRequest

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Настройте правила CORS при необходимости
	},
}

func ServeWebSocket(hub *Hub, w http.ResponseWriter, r *http.Request, services *service.Service) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Ошибка обновления до WebSocket:", err)
		return
	}

	// Получаем chatId из URL (например, ws://localhost:8082/stream/chat/ws?chatId=123)
	chatId := r.URL.Query().Get("chatId")
	if chatId == "" {
		slog.Error("chatId отсутствует в запросе")
		err := conn.Close()
		if err != nil {
			return
		}
		return
	}

	// Создаём клиента с привязкой к chatId
	client := &Client{
		Conn:   conn,
		Send:   make(chan []byte, 256),
		ChatId: chatId,
	}

	hub.Register <- client

	// Обработчик чтения
	go client.readPump(hub, services)
	// Обработчик записи
	go client.writePump()
}

func (c *Client) readPump(hub *Hub, services *service.Service) {
	defer func() {
		hub.Unregister <- c
		err := c.Conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		// Разбор JSON сообщения
		var msgRequest MessageRequest
		if err := json.Unmarshal(message, &msgRequest); err != nil {
			slog.Error("failed to parse message", slog.String("error", err.Error()))
			continue
		}

		// Создание нормализованного сообщения с данными из запроса
		normalizedMessage := domain.Message{
			Message:     msgRequest.Message,
			ChatOwnerId: msgRequest.ChatId,
			UserId:      msgRequest.UserId,
		}

		if err := services.WebSocket.SendMessage(normalizedMessage); err != nil {
			slog.Error("failed to save message", slog.String("error", err.Error()))
			break
		}

		hub.Broadcast <- message
	}
}

func (c *Client) writePump() {
	defer func() {
		err := c.Conn.Close()
		if err != nil {
			return
		}
	}()

	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}
