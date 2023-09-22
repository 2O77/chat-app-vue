package userservice

import (
	"errors"

	userauthenticator "github.com/2O77/chat-app/internal/authenticators"
	user "github.com/2O77/chat-app/internal/domain/user"
)

type UserService struct {
	userAuthenticator *userauthenticator.UserAuthenticator
	userRepository    user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *UserService {
	return &UserService{
		userRepository:    userRepository,
		userAuthenticator: userauthenticator.NewUserAuthenticator(),
	}
}

func (u *UserService) LoginUser(username string, password string) (string, error) {
	user, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if user.Password != password {
		return "", errors.New("wrong password")
	}

	token, err := u.userAuthenticator.EncodeUser(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) DeleteUser(token string) error {
	userID, err := u.userAuthenticator.DecodeUser(token)
	if err != nil {
		return err
	}

	err = u.userRepository.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) RegisterUser(user user.User) error {
	err := u.userRepository.IsUsernameExists(user.Username)
	if err == nil {
		return errors.New("user already exists")
	}

	return u.userRepository.CreateUser(user)
}

func (u *UserService) GetUser(token string) (user.UserResponse, error) {
	userID, err := u.userAuthenticator.DecodeUser(token)
	if err != nil {
		return user.UserResponse{}, err
	}

	getUser, err := u.userRepository.GetUser(userID)
	if err != nil {
		return user.UserResponse{}, err
	}
	return user.UserResponse{
		ID:       getUser.ID,
		Username: getUser.Username,
	}, nil

}

func (u *UserService) GetAllUsers() ([]user.UserResponse, error) {
	return u.userRepository.GetAllUsers()
}
