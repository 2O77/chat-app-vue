package friendcontroller

import (
	"github.com/2O77/chat-app/internal/domain/friend"
	"github.com/gofiber/fiber/v2"
)

type FriendController struct {
	friendService friend.FriendService
}

func NewFriendController(friendService friend.FriendService) *FriendController {
	return &FriendController{
		friendService: friendService,
	}
}

func (fc *FriendController) GetUserFriends(c *fiber.Ctx) error {
	token := c.Get("Token")

	friends, err := fc.friendService.GetUserFriends(token)
	if err != nil {
		return err
	}

	return c.JSON(friends)

}

func (fc *FriendController) SendFriendRequest(c *fiber.Ctx) error {
	var requestBody struct {
		FriendID string `json:"friendID"`
	}

	token := c.Get("Token")

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = fc.friendService.SendFriendRequest(token, requestBody.FriendID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (fc *FriendController) AcceptFriendRequest(c *fiber.Ctx) error {
	var requestBody struct {
		FriendID string `json:"friendID"`
	}

	token := c.Get("Token")

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = fc.friendService.AcceptFriendRequest(token, requestBody.FriendID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (fc *FriendController) RejectFriendRequest(c *fiber.Ctx) error {
	var requestBody struct {
		FriendID string `json:"friendID"`
	}

	token := c.Get("Token")

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = fc.friendService.RejectFriendRequest(token, requestBody.FriendID)

	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (fc *FriendController) CancelFriendRequest(c *fiber.Ctx) error {
	var requestBody struct {
		FriendID string `json:"friendID"`
	}

	token := c.Get("Token")

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = fc.friendService.CancelFriendRequest(token, requestBody.FriendID)

	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (fc *FriendController) RemoveFriend(c *fiber.Ctx) error {
	var requestBody struct {
		FriendID string `json:"friendID"`
	}

	token := c.Get("Token")

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	err = fc.friendService.RemoveUserFriend(token, requestBody.FriendID)

	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (fc *FriendController) GetIncomingFriendRequests(c *fiber.Ctx) error {
	token := c.Get("Token")

	requests, err := fc.friendService.GetIncomingFriendRequests(token)
	if err != nil {
		return err
	}

	return c.JSON(requests)
}

func (fc *FriendController) GetSentFriendRequests(c *fiber.Ctx) error {
	token := c.Get("Token")

	requests, err := fc.friendService.GetSentFriendRequests(token)

	if err != nil {
		return err
	}

	return c.JSON(requests)
}
