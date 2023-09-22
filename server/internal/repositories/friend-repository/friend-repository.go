package friendrepository

import (
	"context"
	"log"

	"github.com/2O77/chat-app/internal/domain/friend"
	userrepository "github.com/2O77/chat-app/internal/repositories/user-repository"
	"go.mongodb.org/mongo-driver/bson"

	MongoDB "go.mongodb.org/mongo-driver/mongo"
)

type FriendRepository struct {
	MongoDB        MongoDB.Database
	userRepository *userrepository.UserRepository
}

func NewFriendRepository(
	mongoDB MongoDB.Database,
	userRepository *userrepository.UserRepository,

) *FriendRepository {
	return &FriendRepository{
		MongoDB:        mongoDB,
		userRepository: userRepository,
	}
}

func (fr *FriendRepository) GetUserFriendsByUserID(userID string) ([]friend.Friend, error) {
	collection := fr.MongoDB.Collection("friends")
	filter := bson.M{"userid": userID}

	var result []friend.Friend

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var friendType friend.FriendRequest
		err := cur.Decode(&friendType)
		if err != nil {
			return nil, err
		}

		user, err := fr.userRepository.GetUser(friendType.FriendID)
		if err != nil {
			continue
		}

		friendType.FriendID = user.ID
		result = append(result, friend.Friend{
			UserID:   user.ID,
			Username: user.Username,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (fr *FriendRepository) GetUserFriendsByFriendID(userID string) ([]friend.Friend, error) {
	collection := fr.MongoDB.Collection("friends")
	filter := bson.M{"friendid": userID}

	var result []friend.Friend

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var friendType friend.FriendRequest
		err := cur.Decode(&friendType)
		if err != nil {
			return nil, err
		}

		user, err := fr.userRepository.GetUser(friendType.UserID)
		if err != nil {
			return nil, err
		}

		result = append(result, friend.Friend{
			UserID:   friendType.FriendID,
			Username: user.Username,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (fr *FriendRepository) GetIncomingFriendRequests(userID string) ([]friend.FriendResponse, error) {
	collection := fr.MongoDB.Collection("friend-requests")
	filter := bson.M{"friendid": userID}

	var result []friend.FriendResponse

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var friendRequest friend.FriendRequest
		err := cur.Decode(&friendRequest)
		if err != nil {
			return nil, err
		}

		user, err := fr.userRepository.GetUser(friendRequest.UserID)
		if err != nil {
			return nil, err
		}

		result = append(result, friend.FriendResponse{
			UserID:   user.ID,
			Username: user.Username,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (fr *FriendRepository) GetSentFriendRequests(userID string) ([]friend.FriendResponse, error) {
	collection := fr.MongoDB.Collection("friend-requests")
	filter := bson.M{"userid": userID}

	var result []friend.FriendResponse

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var friendRequest friend.FriendRequest
		err := cur.Decode(&friendRequest)
		if err != nil {
			return nil, err
		}

		log.Println(friendRequest.FriendID)

		user, err := fr.userRepository.GetUser(friendRequest.FriendID)
		if err != nil {
			return nil, err
		}

		result = append(result, friend.FriendResponse{
			UserID:   user.ID,
			Username: user.Username,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (fr *FriendRepository) CreateFriendRequest(userID string, friendID string) error {
	collection := fr.MongoDB.Collection("friend-requests")

	friendRequest := friend.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	}

	_, err := collection.InsertOne(context.Background(), friendRequest)

	if err != nil {
		return err
	}

	return nil
}

func (fr *FriendRepository) IsFriendExists(userID string, friendID string) (bool, error) {
	collection := fr.MongoDB.Collection("friends")

	filter := bson.M{
		"userid":   userID,
		"friendid": friendID,
	}

	var result friend.FriendRequest

	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == MongoDB.ErrNoDocuments {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (fr *FriendRepository) IsFriendRequestExists(userID string, friendID string) (bool, error) {
	collection := fr.MongoDB.Collection("friend-requests")

	filter := bson.M{
		"userid":   userID,
		"friendid": friendID,
	}

	var result friend.FriendRequest

	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == MongoDB.ErrNoDocuments {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (fr *FriendRepository) DeleteFriendRequest(userID string, friendID string) error {
	collection := fr.MongoDB.Collection("friend-requests")

	filter := bson.M{
		"userid":   userID,
		"friendid": friendID,
	}

	_, err := collection.DeleteOne(context.Background(), filter)

	return err
}

func (fr *FriendRepository) AddUserFriend(userID string, friendID string) error {
	collection := fr.MongoDB.Collection("friends")

	friend := friend.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	}

	_, err := collection.InsertOne(context.Background(), friend)

	return err
}

func (fr *FriendRepository) DeleteUserFriend(userID string, friendID string) error {
	collection := fr.MongoDB.Collection("friends")

	filter := bson.M{
		"userid":   userID,
		"friendid": friendID,
	}

	_, err := collection.DeleteOne(context.Background(), filter)

	return err
}
