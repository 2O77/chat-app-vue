package main

import (
	"context"
	"log"

	chatcontroller "github.com/2O77/chat-app/internal/controllers/chat-controller"
	friendcontroller "github.com/2O77/chat-app/internal/controllers/friend-controller"
	usercontroller "github.com/2O77/chat-app/internal/controllers/user-controller"
	chathandler "github.com/2O77/chat-app/internal/handler/chat-handler"
	chatrepository "github.com/2O77/chat-app/internal/repositories/chat-repository"
	friendrepository "github.com/2O77/chat-app/internal/repositories/friend-repository"
	userrepository "github.com/2O77/chat-app/internal/repositories/user-repository"
	chatservice "github.com/2O77/chat-app/internal/services/chat-service"
	friendservice "github.com/2O77/chat-app/internal/services/friend-service"
	userservice "github.com/2O77/chat-app/internal/services/user-service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
     
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
    
	err := godotenv.Load("/home/hyvinhiljaa/chat-app-server/.env")
	if err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	mongoURI := os.Getenv("MONGO_URL")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is missing in .env file")
	}

	dbName := "chat-app"

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(dbName)

	userRepository := userrepository.NewUserRepository(*database)

	friendRepository := friendrepository.NewFriendRepository(*database, userRepository)
	friendService := friendservice.NewFriendService(friendRepository)
	friendController := friendcontroller.NewFriendController(friendService)

	userService := userservice.NewUserService(userRepository)
	userController := usercontroller.NewUserController(userService)

	chatRepository := chatrepository.NewChatRepository(*database)
	chatService := chatservice.NewChatService(chatRepository, userRepository, friendService)
	chatController := chatcontroller.NewChatController(chatService)

	websockethandler := chathandler.NewWebSocketHandler()

	app.Post("/user/login", userController.LoginUser)
	app.Delete("/user", userController.DeleteUser)
	app.Post("/user/register", userController.RegisterUser)
	app.Get("/user", userController.GetUser)
	app.Get("/users", userController.GetAllUsers)

	app.Get("/friends", friendController.GetUserFriends)
	app.Post("/friends", friendController.AcceptFriendRequest)
	app.Post("/friend-requests", friendController.SendFriendRequest)
	app.Delete("/friend-requests", friendController.RejectFriendRequest)
	app.Delete("/friend-requests/sent", friendController.CancelFriendRequest)
	app.Delete("/friends", friendController.RemoveFriend)
	app.Get("/friend-requests", friendController.GetIncomingFriendRequests)
	app.Get("/friend-requests/sent", friendController.GetSentFriendRequests)

	app.Get("/chats", chatController.GetUserChats)
	app.Get("/chat/:chatID", chatController.GetUserChat)
	app.Get("/chat/users/:chatID", chatController.GetChatUsers)
	app.Post("/chats", chatController.CreateChat)
	app.Delete("/chats", chatController.DeleteChat)
	app.Post("/chat/adduser", chatController.AddUserToChat)
	app.Put("/chat/makeadmin", chatController.MakeUserAdmin)
	app.Put("/chat/notadmin", chatController.MakeUserNotAdmin)
	app.Delete("/chat/removeuser", chatController.RemoveUserFromChat)
	app.Put("/chat/changename", chatController.ChangeChatName)
	app.Delete("/chat/leave", chatController.LeftChat)

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(websockethandler.HandleWebSocketConnection))

	app.Listen(":2000")

	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MongoDB lost connection.")
}
