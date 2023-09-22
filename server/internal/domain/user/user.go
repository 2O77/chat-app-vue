package user

type UserService interface {
	LoginUser(username string, password string) (string, error)
	DeleteUser(userID string) error
	RegisterUser(user User) error
	GetUser(token string) (UserResponse, error)
	GetAllUsers() ([]UserResponse, error)
}

type UserRepository interface {
	GetUser(userID string) (UserResponse, error)
	GetUserByUsername(username string) (User, error)
	IsUserExists(userID string) error
	IsUsernameExists(username string) error
	DeleteUser(userID string) error
	CreateUser(user User) error
	GetAllUsers() ([]UserResponse, error)
}

type Friend struct {
	UserID   string
	Username string
}

type User struct {
	ID       string
	Username string
	Password string
}

type UserResponse struct {
	ID       string
	Username string
}
