package usercontroller

import (
	user "github.com/2O77/chat-app/internal/domain/user"
	fiber "github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) LoginUser(c *fiber.Ctx) error {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	user, err := uc.userService.LoginUser(requestBody.Username, requestBody.Password)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {

	var requestBody struct {
		UserID string `json:"token"`
	}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	user := uc.userService.DeleteUser(requestBody.UserID)

	return c.JSON(user)
}

func (uc *UserController) RegisterUser(c *fiber.Ctx) error {
	var user user.User
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	err = uc.userService.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	token := c.Get("Token")

	user, err := uc.userService.GetUser(token)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(users)
}
