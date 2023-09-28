package chatservice

import (
	"errors"

	userauthenticator "github.com/2O77/chat-app/internal/authenticators"
	"github.com/2O77/chat-app/internal/domain/chat"
	"github.com/2O77/chat-app/internal/domain/friend"
	user "github.com/2O77/chat-app/internal/domain/user"
	"github.com/google/uuid"
)

type ChatService struct {
	chatrepository    chat.ChatRepository
	userauthenticator userauthenticator.UserAuthenticator
	userrepository    user.UserRepository
	friendservice     friend.FriendService
}

func NewChatService(chatrepository chat.ChatRepository, userrepository user.UserRepository, friendservice friend.FriendService) *ChatService {
	return &ChatService{
		userrepository:    userrepository,
		chatrepository:    chatrepository,
		friendservice:     friendservice,
		userauthenticator: *userauthenticator.NewUserAuthenticator(),
	}
}

func (cs *ChatService) CreateChat(adminToken string, chatName string) error {
	decodeUser, err := cs.userauthenticator.DecodeUser(adminToken)
	if err != nil {
		return err
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return errors.New("user does not exists")
	}

	user, err := cs.userrepository.GetUser(decodeUser)
	if err != nil {
		return err
	}

	chat := chat.Chat{
		ID:   uuid.New().String(),
		Name: chatName,
		Users: []chat.ChatUser{
			{
				ID:        user.ID,
				Username:  user.Username,
				IsAdmin:   true,
				IsCreator: true,
			},
		},
	}

	err = cs.chatrepository.CreateChat(chat)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChatService) DeleteChat(adminToken string, chatID string) error {
	decodeUser, err := cs.userauthenticator.DecodeUser(adminToken)
	if err != nil {
		return errors.New("not authorized")
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return errors.New("user does not exists")
	}

	isAdmin, err := cs.chatrepository.IsUserAdmin(chatID, decodeUser)
	if err != nil {
		return errors.New("errors while checking if user is admin")
	}

	if !isAdmin {
		return errors.New("user is not admin")
	}

	err = cs.chatrepository.DeleteChat(chatID)
	if err != nil {
		return errors.New("chat does not exists")
	}
	return nil
}

func (cs *ChatService) GetUserChats(userToken string) ([]chat.Chat, error) {
	decodeUser, err := cs.userauthenticator.DecodeUser(userToken)
	if err != nil {
		return []chat.Chat{}, err
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return []chat.Chat{}, errors.New("user does not exists")
	}

	user, err := cs.userrepository.GetUser(decodeUser)
	if err != nil {
		return []chat.Chat{}, err
	}

	chats, err := cs.chatrepository.GetUserChats(user.ID)
	if err != nil {
		return []chat.Chat{}, err
	}
	return chats, nil
}

func (cs *ChatService) GetUserChat(chatID string, userToken string) (chat.Chat, error) {
	decodeUser, err := cs.userauthenticator.DecodeUser(userToken)
	if err != nil {
		return chat.Chat{}, err
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return chat.Chat{}, errors.New("user does not exists")
	}

	isUserInChat, err := cs.chatrepository.IsUserInChat(decodeUser, chatID)
	if err != nil {
		return chat.Chat{}, err
	}

	if !isUserInChat {
		return chat.Chat{}, errors.New("user is not in chat")
	}

	user, err := cs.userrepository.GetUser(decodeUser)
	if err != nil {
		return chat.Chat{}, err
	}

	userChat, err := cs.chatrepository.GetUserChat(chatID, user.ID)
	if err != nil {
		return chat.Chat{}, err
	}
	return userChat, nil
}

func (cs *ChatService) GetChatUsers(userToken string, chatID string) ([]chat.ChatUser, error) {
	decodeUser, err := cs.userauthenticator.DecodeUser(userToken)
	if err != nil {
		return []chat.ChatUser{}, errors.New("not authorized")
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return []chat.ChatUser{}, errors.New("user does not exists")
	}

	isUserInChat, err := cs.chatrepository.IsUserInChat(chatID, decodeUser)
	if err != nil {
		return []chat.ChatUser{}, err
	}

	if !isUserInChat {
		return []chat.ChatUser{}, errors.New("user is not in chat")
	}

	users, err := cs.chatrepository.GetChatUsers(chatID)
	if err != nil {
		return []chat.ChatUser{}, err
	}

	return users, nil
}

func (cs *ChatService) GetChatUser(userID string, chatID string) (chat.ChatUser, error) {
	chatUser, err := cs.chatrepository.GetChatUser(userID, chatID)
	if err != nil {
		return chat.ChatUser{}, err
	}

	return chatUser, nil
}

func (cs *ChatService) AddUserToChat(adminToken string, chatID string, userID string) error {
	decodedAdmin, err := cs.userauthenticator.DecodeUser(adminToken)
	if err != nil {
		return err
	}

	isChatExists, err := cs.chatrepository.IsChatExists(chatID)
	if err != nil {
		return errors.New("chat does not exists")
	}

	if !isChatExists {
		return errors.New("chat does not exists")
	}

	err = cs.userrepository.IsUserExists(decodedAdmin)
	if err != nil {
		return errors.New("user does not exists")
	}

	isAdmin, err := cs.chatrepository.IsUserAdmin(chatID, decodedAdmin)
	if err != nil {
		return errors.New("errors while checking if user is admin")
	}

	if !isAdmin {
		return errors.New("user is not admin")
	}

	err = cs.userrepository.IsUserExists(userID)
	if err != nil {
		return errors.New("user does not exists")
	}

	isUserInChat, err := cs.chatrepository.IsUserInChat(chatID, userID)
	if err != nil {
		return errors.New("errors while checking if user is in chat")
	}

	if isUserInChat {
		return errors.New("user is already in chat")
	}

	isUserFriend, err := cs.friendservice.IsUserFriend(decodedAdmin, userID)
	if err != nil {
		return errors.New("errors while checking if user is friend")
	}

	if !isUserFriend {
		return errors.New("user is not friend")
	}

	user, err := cs.userrepository.GetUser(userID)
	if err != nil {
		return err
	}

	chat := chat.ChatUser{
		ID:        user.ID,
		Username:  user.Username,
		IsAdmin:   false,
		IsCreator: false,
	}

	err = cs.chatrepository.AddUserToChat(chatID, chat)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChatService) RemoveUserFromChat(adminToken string, chatID string, userID string) error {
	decodeUser, err := cs.userauthenticator.DecodeUser(adminToken)
	if err != nil {
		return err
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return errors.New("user does not exists")
	}

	isAdmin, err := cs.chatrepository.IsUserAdmin(chatID, decodeUser)
	if err != nil {
		return err
	}

	if !isAdmin {
		return errors.New("user is not admin")
	}

	isAdmin, err = cs.chatrepository.IsUserAdmin(chatID, userID)
	if err != nil {
		return err
	}

	if isAdmin {
		return errors.New("user is admin")
	}

	err = cs.userrepository.IsUserExists(userID)
	if err != nil {
		return err
	}

	isUserInChat, err := cs.chatrepository.IsUserInChat(chatID, userID)
	if err != nil {
		return err
	}

	if !isUserInChat {
		return errors.New("user is not in chat")
	}

	user, err := cs.userrepository.GetUser(userID)
	if err != nil {
		return err
	}

	err = cs.chatrepository.DeleteUserFromChat(chatID, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChatService) LeftChat(userToken string, chatID string) error {
	decodeUser, err := cs.userauthenticator.DecodeUser(userToken)
	if err != nil {
		return err
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return errors.New("user does not exists")
	}

	isUserInChat, err := cs.chatrepository.IsUserInChat(chatID, decodeUser)
	if err != nil {
		return err
	}

	if !isUserInChat {
		return errors.New("user is not in chat")
	}

	isCreator, err := cs.chatrepository.IsUserCreator(chatID, decodeUser)
	if err != nil {
		return err
	}

	if isCreator {
		return errors.New("user is creator")
	}

	err = cs.chatrepository.DeleteUserFromChat(chatID, decodeUser)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChatService) MakeUserAdmin(adminToken string, chatID string, userID string) error {
	decodeUser, err := cs.userauthenticator.DecodeUser(adminToken)
	if err != nil {
		return err
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return errors.New("user does not exists")
	}

	isAdmin, err := cs.chatrepository.IsUserAdmin(chatID, decodeUser)
	if err != nil {
		return err
	}

	if !isAdmin {
		return errors.New("user is not admin")
	}

	err = cs.userrepository.IsUserExists(userID)
	if err != nil {
		return errors.New("user does not exists")
	}

	err = cs.chatrepository.MakeUserAdmin(chatID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChatService) MakeUserNotAdmin(creatorToken string, chatID string, adminID string) error {
	decodeCreator, err := cs.userauthenticator.DecodeUser(creatorToken)
	if err != nil {
		return err
	}

	err = cs.userrepository.IsUserExists(decodeCreator)
	if err != nil {
		return errors.New("user does not exists")
	}

	isCreator, err := cs.chatrepository.IsUserCreator(chatID, decodeCreator)
	if err != nil {
		return err
	}

	if !isCreator {
		return errors.New("creator is not creator")
	}

	isAdmin, err := cs.chatrepository.IsUserAdmin(chatID, adminID)
	if err != nil {
		return err
	}

	if !isAdmin {
		return errors.New("admin is not admin")
	}

	isCreator, err = cs.chatrepository.IsUserCreator(chatID, adminID)
	if err != nil {
		return err
	}

	if isCreator {
		return errors.New("admin is creator")
	}

	err = cs.chatrepository.MakeUserNotAdmin(chatID, adminID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChatService) ChangeChatName(adminToken string, chatID string, newName string) error {
	decodeUser, err := cs.userauthenticator.DecodeUser(adminToken)
	if err != nil {
		return err
	}

	err = cs.userrepository.IsUserExists(decodeUser)
	if err != nil {
		return errors.New("user does not exists")
	}

	isUserInChat, err := cs.chatrepository.IsUserInChat(chatID, decodeUser)
	if err != nil {
		return err
	}

	if !isUserInChat {
		return errors.New("user is not in chat")
	}

	isAdmin, err := cs.chatrepository.IsUserAdmin(chatID, decodeUser)
	if err != nil {
		return err
	}

	if !isAdmin {
		return errors.New("user is not admin")
	}

	err = cs.chatrepository.RenameChat(chatID, newName)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChatService) SaveMessage (message chat.Message) error {

    err := cs.chatrepository.LoadMessage(message)
    if err != nil {
        return err 
    }

    return nil
}


