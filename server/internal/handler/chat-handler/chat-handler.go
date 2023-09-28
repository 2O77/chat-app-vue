package chathandler

import (
	"encoding/json"
	"time"
    
	"github.com/2O77/chat-app/internal/domain/chat"
	"github.com/gofiber/websocket/v2"
)

type MessageJson struct {
	Username string `json:"username"`
	ChatID   string `json:"chatID"`
	Text     string `json:"text"`
    Timestamp time.Time `json:"timestamp"`
}

type WebSocketHandler struct {
	clients map[*websocket.Conn]bool
    chatservice chat.ChatService
}

func ConvertMessageJSONToMessage(messageJSON MessageJson) chat.Message {
	return chat.Message{
		UserID:    messageJSON.Username,
		ChatID:    messageJSON.ChatID,
		Text:      messageJSON.Text,
		Timestamp: messageJSON.Timestamp.String(),
	}
}

func NewWebSocketHandler(chatservice chat.ChatService) *WebSocketHandler {
	return &WebSocketHandler{
		clients: make(map[*websocket.Conn]bool),
        chatservice: chatservice,
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
	msgJSON := MessageJson{
		Username: username,
		ChatID:   chatID,
		Text:     text,
        Timestamp: time.Now(),
	}

    msg := ConvertMessageJSONToMessage(msgJSON)

    wh.chatservice.SaveMessage(msg)

	for client := range wh.clients {
		err := client.WriteJSON(msg)
		if err != nil {
			delete(wh.clients, client)
			client.Close()
		}
	}

}








