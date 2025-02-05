package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log/slog"
	"strconv"
	"sync"
)

// Client представляет клиента WebSocket
type Client struct {
	Conn   *websocket.Conn
	Send   chan []byte
	ChatId string
}

// Hub управляет всеми активными WebSocket-соединениями
type Hub struct {
	Clients    map[string]map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// Если чата ещё нет, создаём его
			if _, exists := h.Clients[client.ChatId]; !exists {
				h.Clients[client.ChatId] = make(map[*Client]bool)
			}
			// Добавляем клиента в соответствующий чат
			h.Clients[client.ChatId][client] = true
			slog.Info("Client connected to chat", slog.String("chatId", client.ChatId))

		case client := <-h.Unregister:
			if clients, exists := h.Clients[client.ChatId]; exists {
				delete(clients, client)
				close(client.Send) // Закрываем канал
				if len(clients) == 0 {
					delete(h.Clients, client.ChatId) // Удаляем чат, если в нём больше нет клиентов
				}
			}
			slog.Info("Client disconnected from chat", slog.String("chatId", client.ChatId))

		case message := <-h.Broadcast:
			var msgRequest MessageRequest
			if err := json.Unmarshal(message, &msgRequest); err != nil {
				slog.Error("Ошибка парсинга сообщения", slog.String("error", err.Error()))
				continue
			}

			chatID := msgRequest.ChatId

			if clients, exists := h.Clients[strconv.FormatInt(chatID, 10)]; exists {
				for client := range clients {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(clients, client)
					}
				}
			}
		}
	}
}
