package chathandler

import (
	"encoding/json"

	"github.com/gofiber/websocket/v2"
)

type MessageJson struct {
	Username string `json:"username"`
	ChatID   string `json:"chatID"`
	Text     string `json:"text"`
}

type WebSocketHandler struct {
	clients map[*websocket.Conn]bool
}

func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (wh *WebSocketHandler) HandleWebSocketConnection(c *websocket.Conn) {
	wh.clients[c] = true

	defer func() {
		delete(wh.clients, c)
		c.Close()
	}()

	for {
		_, rawMsg, err := c.ReadMessage()
		if err != nil {
			break
		}

		msg := MessageJson{}
		err = json.Unmarshal(rawMsg, &msg)
		if err != nil {
			continue
		}

		wh.BroadcastMessage(msg.Username, msg.ChatID, msg.Text)
	}
}

func (wh *WebSocketHandler) BroadcastMessage(username string, chatID string, text string) {
	msg := MessageJson{
		Username: username,
		ChatID:   chatID,
		Text:     text,
	}

	for client := range wh.clients {
		err := client.WriteJSON(msg)
		if err != nil {
			delete(wh.clients, client)
			client.Close()
		}
	}

}
