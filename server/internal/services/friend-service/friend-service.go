package friendservice

import (
	"errors"

	userauthenticator "github.com/2O77/chat-app/internal/authenticators"
	"github.com/2O77/chat-app/internal/domain/friend"
)

type FriendService struct {
	friendRepository  friend.FriendRepository
	userauthenticator *userauthenticator.UserAuthenticator
}

func NewFriendService(friendRepository friend.FriendRepository) *FriendService {
	return &FriendService{
		friendRepository:  friendRepository,
		userauthenticator: userauthenticator.NewUserAuthenticator(),
	}
}

func (fs *FriendService) GetUserFriends(token string) ([]friend.Friend, error) {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return nil, err
	}

	userFriends, err := fs.friendRepository.GetUserFriendsByUserID(userID)
	if err != nil {
		return nil, err
	}

	friendFriends, err := fs.friendRepository.GetUserFriendsByFriendID(userID)
	if err != nil {
		return nil, err
	}

	allFriends := append(userFriends, friendFriends...)

	return allFriends, nil
}

func (fs *FriendService) SendFriendRequest(token string, friendID string) error {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return err
	}

	if userID == friendID {
		return errors.New("you can not be your friend")
	}

	isFriendExists, err := fs.friendRepository.IsFriendExists(userID, friendID)
	if err != nil {
		return err
	}

	if isFriendExists {
		return errors.New("user is already your friend")
	}

	isFriendExists, err = fs.friendRepository.IsFriendExists(friendID, userID)
	if err != nil {
		return err
	}

	if isFriendExists {
		return errors.New("user is already your friend")
	}

	isFriendRequestExists, err := fs.friendRepository.IsFriendRequestExists(userID, friendID)
	if err != nil {
		return err
	}

	if isFriendRequestExists {
		return errors.New("friend request already exists")
	}

	isFriendRequestExists, err = fs.friendRepository.IsFriendRequestExists(friendID, userID)
	if err != nil {
		return err
	}

	if isFriendRequestExists {
		return errors.New("friend request already exists")
	}

	return fs.friendRepository.CreateFriendRequest(userID, friendID)
}

func (fs *FriendService) AcceptFriendRequest(token string, friendID string) error {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return err
	}

	exists, err := fs.friendRepository.IsFriendRequestExists(friendID, userID)
	if !exists {
		return errors.New("friend request does not exists")
	}

	if err != nil {
		return err
	}

	if exists {
		err = fs.friendRepository.DeleteFriendRequest(friendID, userID)
		if err != nil {
			return err
		}

		return fs.friendRepository.AddUserFriend(userID, friendID)
	}

	return nil
}

func (fs *FriendService) RejectFriendRequest(token string, friendID string) error {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return err
	}

	exists, err := fs.friendRepository.IsFriendRequestExists(friendID, userID)
	if !exists {
		return errors.New("friend request does not exists")
	}

	if err != nil {
		return err
	}

	if exists {
		return fs.friendRepository.DeleteFriendRequest(friendID, userID)
	}

	return errors.New("friend request does not exist")
}

func (fs *FriendService) CancelFriendRequest(token string, friendID string) error {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return errors.New("invalid token")
	}

	exists, err := fs.friendRepository.IsFriendRequestExists(userID, friendID)
	if !exists {
		return errors.New("friend request does not exists")
	}

	if err != nil {
		return err
	}

	if exists {
		return fs.friendRepository.DeleteFriendRequest(userID, friendID)
	}

	return errors.New("friend request does not exist")
}

func (fs *FriendService) GetIncomingFriendRequests(token string) ([]friend.FriendResponse, error) {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return nil, err
	}

	return fs.friendRepository.GetIncomingFriendRequests(userID)
}

func (fs *FriendService) GetSentFriendRequests(token string) ([]friend.FriendResponse, error) {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return nil, err
	}

	return fs.friendRepository.GetSentFriendRequests(userID)
}

func (fs *FriendService) RemoveUserFriend(token string, friendID string) error {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return err
	}

	err = fs.friendRepository.DeleteUserFriend(userID, friendID)
	if err != nil {
		return err
	}

	err = fs.friendRepository.DeleteUserFriend(friendID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FriendService) IsUserFriend(token string, friendID string) (bool, error) {
	userID, err := fs.userauthenticator.DecodeUser(token)
	if err != nil {
		return false, err
	}

	isFriend, err := fs.friendRepository.IsFriendExists(userID, friendID)
	if err != nil {
		return false, err
	}

	if isFriend {
		return true, nil
	}

	isFriend, err = fs.friendRepository.IsFriendExists(friendID, userID)
	if err != nil {
		return false, err
	}

	if isFriend {
		return true, nil
	}

	return false, nil
}
