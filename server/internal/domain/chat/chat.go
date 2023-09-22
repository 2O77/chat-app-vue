package chat

type ChatService interface {
	CreateChat(adminToken string, chatname string) error
	DeleteChat(adminToken string, chatID string) error
	GetUserChats(token string) ([]Chat, error)
	GetUserChat(token string, chatID string) (Chat, error)
	GetChatUsers(token string, chatID string) ([]ChatUser, error)
	GetChatUser(userID string, chatID string) (ChatUser, error)
	AddUserToChat(adminToken string, chatID string, userID string) error
	MakeUserAdmin(adminToken string, chatID string, userID string) error
	MakeUserNotAdmin(creatorToken string, chatID string, userID string) error
	RemoveUserFromChat(adminToken string, chatID string, userID string) error
	ChangeChatName(adminToken string, chatID string, newName string) error
	LeftChat(token string, chatID string) error
}

type ChatRepository interface {
	CreateChat(chat Chat) error
	DeleteChat(chatID string) error
	GetUserChats(userID string) ([]Chat, error)
	GetUserChat(userID string, chatID string) (Chat, error)
	GetChatUsers(chatID string) ([]ChatUser, error)
	GetChatUser(userID string, chatID string) (ChatUser, error)
	GetChatByID(chatID string) (Chat, error)
	AddUserToChat(chatID string, user ChatUser) error
	DeleteUserFromChat(chatID string, userID string) error
	IsUserInChat(chatID string, userID string) (bool, error)
	IsChatExists(chatID string) (bool, error)
	IsUserAdmin(chatID string, userID string) (bool, error)
	IsUserCreator(chatID string, userID string) (bool, error)
	MakeUserAdmin(chatID string, userID string) error
	MakeUserNotAdmin(chatID string, userID string) error
	RenameChat(chatID string, newName string) error
}

type Chat struct {
	ID    string     `bson:"id"`
	Name  string     `bson:"name"`
	Users []ChatUser `bson:"users"`
}

type ChatUser struct {
	ID        string `bson:"id"`
	Username  string `bson:"username"`
	IsAdmin   bool   `bson:"is_admin"`
	IsCreator bool   `bson:"is_creator"`
}

type Message struct {
	UserID string `bson:"id"`
	ChatID string `bson:"chat_id"`
	Text   string `bson:"text"`
}
