package friend

type FriendService interface {
	GetUserFriends(token string) ([]Friend, error)
	SendFriendRequest(token string, friendID string) error
	AcceptFriendRequest(token string, friendID string) error
	RejectFriendRequest(token string, friendID string) error
	CancelFriendRequest(token string, friendID string) error
	GetIncomingFriendRequests(token string) ([]FriendResponse, error)
	GetSentFriendRequests(token string) ([]FriendResponse, error)
	RemoveUserFriend(token string, friendID string) error
	IsUserFriend(token string, friendID string) (bool, error)
}

type FriendRepository interface {
	GetUserFriendsByUserID(userID string) ([]Friend, error)
	GetUserFriendsByFriendID(userID string) ([]Friend, error)
	GetIncomingFriendRequests(userID string) ([]FriendResponse, error)
	GetSentFriendRequests(userID string) ([]FriendResponse, error)
	CreateFriendRequest(userID string, friendID string) error
	IsFriendExists(userID string, friendID string) (bool, error)
	IsFriendRequestExists(userID string, friendID string) (bool, error)
	DeleteFriendRequest(userID string, friendID string) error
	AddUserFriend(userID string, friendID string) error
	DeleteUserFriend(userID string, friendID string) error
}

type Friend struct {
	UserID   string
	Username string
}

type FriendRequest struct {
	UserID   string `bson:"userid"`
	FriendID string `bson:"friendid"`
}

type FriendResponse struct {
	UserID   string
	Username string
}
