package chatrepository

import (
	"context"
	"errors"

	chat "github.com/2O77/chat-app/internal/domain/chat"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	MongoDB "go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	mongoDB MongoDB.Database
}

func NewChatRepository(mongoDB MongoDB.Database) *ChatRepository {
	return &ChatRepository{
		mongoDB: mongoDB,
	}
}

func (cr *ChatRepository) CreateChat(chat chat.Chat) error {
	collection := cr.mongoDB.Collection("chats")

	chat.ID = uuid.New().String()

	_, err := collection.InsertOne(context.Background(), chat)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ChatRepository) DeleteChat(chatID string) error {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ChatRepository) GetUserChats(userID string) ([]chat.Chat, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"users.id": userID}

	var result []chat.Chat
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return []chat.Chat{}, err
	}
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return []chat.Chat{}, err
	}
	return result, nil
}

func (cr *ChatRepository) GetUserChat(userID string, chatID string) (chat.Chat, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID, "users.id": userID}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return chat.Chat{}, err
	}
	return result, nil
}

func (cr *ChatRepository) GetChatByID(chatID string) (chat.Chat, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return chat.Chat{}, err
	}
	return result, nil
}

func (cr *ChatRepository) GetChatUsers(chatID string) ([]chat.ChatUser, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return []chat.ChatUser{}, err
	}
	return result.Users, nil
}

func (cr *ChatRepository) GetChatUser(userID string, chatID string) (chat.ChatUser, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID, "users.id": userID}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return chat.ChatUser{}, err
	}
	return result.Users[0], nil
}

func (cr *ChatRepository) AddUserToChat(chatID string, user chat.ChatUser) error {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID}

	update := bson.M{"$push": bson.M{"users": user}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ChatRepository) DeleteUserFromChat(chatID string, userID string) error {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID}

	update := bson.M{"$pull": bson.M{"users": bson.M{"id": userID}}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ChatRepository) IsUserInChat(chatID string, userID string) (bool, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID, "users.id": userID}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, MongoDB.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (cr *ChatRepository) IsChatExists(chatID string) (bool, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, MongoDB.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (cr *ChatRepository) IsUserAdmin(chatID string, userID string) (bool, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{
		"id": chatID,
		"users": bson.M{
			"$elemMatch": bson.M{
				"id":       userID,
				"is_admin": true,
			},
		},
	}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, MongoDB.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (cr *ChatRepository) IsUserCreator(chatID string, userID string) (bool, error) {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{
		"id": chatID,
		"users": bson.M{
			"$elemMatch": bson.M{
				"id":         userID,
				"is_creator": true,
			},
		},
	}

	var result chat.Chat
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, MongoDB.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (cr *ChatRepository) MakeUserAdmin(chatID string, userID string) error {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID, "users.id": userID}

	update := bson.M{"$set": bson.M{"users.$.is_admin": true}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ChatRepository) MakeUserNotAdmin(chatID string, userID string) error {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID, "users.id": userID}

	update := bson.M{"$set": bson.M{"users.$.is_admin": false}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ChatRepository) DeleteUserFromAdmins(chatID string, userID string) error {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{
		"id": chatID,
		"users": bson.M{
			"$elemMatch": bson.M{
				"id":       userID,
				"is_admin": true,
			},
		}}

	update := bson.M{"$pull": bson.M{"users.$.is_admin": false}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ChatRepository) RenameChat(chatID string, newName string) error {
	collection := cr.mongoDB.Collection("chats")

	filter := bson.M{"id": chatID}

	update := bson.M{"$set": bson.M{"name": newName}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ChatRepository) LoadMessage(message chat.Message) error {
    collection := cr.mongoDB.Collection("messages")

    _, err := collection.InsertOne(context.Background(),message)
    if err != nil {
        return err
    }

    return nil
}


















