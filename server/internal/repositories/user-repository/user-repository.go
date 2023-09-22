package userrepository

import (
	"context"

	"github.com/2O77/chat-app/internal/domain/user"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	MongoDB "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongoDB MongoDB.Database
}

func NewUserRepository(mongoDB MongoDB.Database) *UserRepository {
	return &UserRepository{
		mongoDB: mongoDB,
	}
}

func (ur *UserRepository) GetUser(userID string) (user.UserResponse, error) {
	collection := ur.mongoDB.Collection("users")

	filter := bson.M{"id": userID}

	var result user.UserResponse
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return user.UserResponse{}, err
	}
	return result, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (user.User, error) {
	collection := ur.mongoDB.Collection("users")

	filter := bson.M{"username": username}

	var result user.User
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (ur *UserRepository) GetAllUsers() ([]user.UserResponse, error) {
	collection := ur.mongoDB.Collection("users")

	var result []user.UserResponse
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []user.UserResponse{}, err
	}
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return []user.UserResponse{}, err
	}

	return result, nil
}

func (ur *UserRepository) IsUserExists(userID string) error {
	collection := ur.mongoDB.Collection("users")
	filter := bson.M{"id": userID}

	var result user.User
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) IsUsernameExists(username string) error {
	collection := ur.mongoDB.Collection("users")
	filter := bson.M{"username": username}

	var result user.User
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(userID string) error {
	collection := ur.mongoDB.Collection("users")
	filter := bson.M{"id": userID}

	_, err := collection.DeleteOne(context.Background(), filter)

	return err
}

func (ur *UserRepository) CreateUser(user user.User) error {
	collection := ur.mongoDB.Collection("users")

	user.ID = uuid.New().String()

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
