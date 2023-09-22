package chatcontroller

import (
	chat "github.com/2O77/chat-app/internal/domain/chat"
	"github.com/gofiber/fiber/v2"
)

type ChatController struct {
	chatservice chat.ChatService
}

func NewChatController(chatservice chat.ChatService) *ChatController {
	return &ChatController{
		chatservice: chatservice,
	}
}

func (cc *ChatController) CreateChat(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatName string `json:"chatName"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.CreateChat(token, requestBody.ChatName)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (cc *ChatController) DeleteChat(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatID string `json:"chatID"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.DeleteChat(token, requestBody.ChatID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (cc *ChatController) GetUserChats(c *fiber.Ctx) error {
	token := c.Get("Token")

	chats, err := cc.chatservice.GetUserChats(token)
	if err != nil {
		return err
	}

	return c.JSON(chats)
}

func (cc *ChatController) GetUserChat(c *fiber.Ctx) error {
	token := c.Get("Token")
	param := c.Params("chatID")

	chat, err := cc.chatservice.GetUserChat(token, param)
	if err != nil {
		return err
	}

	return c.JSON(chat)
}

func (cc *ChatController) GetChatUsers(c *fiber.Ctx) error {
	token := c.Get("Token")
	param := c.Params("chatID")

	users, err := cc.chatservice.GetChatUsers(token, param)
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (cc *ChatController) GetChatUser(c *fiber.Ctx) error {
	chatID := c.Params("chatID")
	userID := c.Params("userID")

	user, err := cc.chatservice.GetChatUser(chatID, userID)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (cc *ChatController) AddUserToChat(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatID string `json:"chatID"`
		UserID string `json:"userID"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.AddUserToChat(token, requestBody.ChatID, requestBody.UserID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (cc *ChatController) MakeUserAdmin(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatID string `json:"chatID"`
		UserID string `json:"userID"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.MakeUserAdmin(token, requestBody.ChatID, requestBody.UserID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (cc *ChatController) MakeUserNotAdmin(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatID  string `json:"chatID"`
		AdminID string `json:"userID"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.MakeUserNotAdmin(token, requestBody.ChatID, requestBody.AdminID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (cc *ChatController) RemoveUserFromChat(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatID string `json:"chatID"`
		UserID string `json:"userID"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.RemoveUserFromChat(token, requestBody.ChatID, requestBody.UserID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (cc *ChatController) LeftChat(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatID string `json:"chatID"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.LeftChat(token, requestBody.ChatID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (cc *ChatController) ChangeChatName(c *fiber.Ctx) error {
	token := c.Get("Token")

	var requestBody struct {
		ChatID   string `json:"chatID"`
		ChatName string `json:"chatName"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = cc.chatservice.ChangeChatName(token, requestBody.ChatID, requestBody.ChatName)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
