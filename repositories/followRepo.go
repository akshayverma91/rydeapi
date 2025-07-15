package repositories

import (
	"context"
	"time"

	"github.com/akshayverma91/rydeapi/config"
	"github.com/akshayverma91/rydeapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.TODO()

// FollowUser creates a follow relationship between two users.
// It takes the follower's user ID and the following user's ID as strings.
func FollowUser(followerID, followingID string) error {
	followerObj, _ := primitive.ObjectIDFromHex(followerID)
	followingObj, _ := primitive.ObjectIDFromHex(followingID)

	// Avoid duplicate follow
	filter := bson.M{
		"followerId":  followerObj,
		"followingId": followingObj,
	}
	count, _ := config.FollowCollection.CountDocuments(ctx, filter)
	if count > 0 {
		return nil // already followed
	}

	follow := models.Follow{
		FollowerID:  followerObj,
		FollowingID: followingObj,
		CreatedAt:   time.Now(),
	}
	_, err := config.FollowCollection.InsertOne(ctx, follow)
	return err
}

// UnfollowUser removes a follow relationship between two users.
// It takes the follower's user ID and the following user's ID as strings.
func UnfollowUser(followerID, followingID string) error {
	followerObj, _ := primitive.ObjectIDFromHex(followerID)
	followingObj, _ := primitive.ObjectIDFromHex(followingID)

	_, err := config.FollowCollection.DeleteOne(ctx, bson.M{
		"followerId":  followerObj,
		"followingId": followingObj,
	})
	return err
}

// GetFollowers retrieves a list of user IDs that follow a specific user.
// It takes the user's ID as a string and returns a slice of primitive.ObjectID.
func GetFollowerUsers(userID string) ([]models.User, error) {
	targetID, _ := primitive.ObjectIDFromHex(userID)

	// Find follower IDs
	cursor, err := config.FollowCollection.Find(ctx, bson.M{
		"followingId": targetID,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var followerIDs []primitive.ObjectID
	for cursor.Next(ctx) {
		var follow models.Follow
		cursor.Decode(&follow)
		followerIDs = append(followerIDs, follow.FollowerID)
	}

	// Lookup users
	userCursor, err := config.UserCollection.Find(ctx, bson.M{"_id": bson.M{"$in": followerIDs}})
	if err != nil {
		return nil, err
	}
	defer userCursor.Close(ctx)

	var users []models.User
	for userCursor.Next(ctx) {
		var u models.User
		userCursor.Decode(&u)
		users = append(users, u)
	}
	return users, nil
}

// GetFollowing retrieves a list of user IDs that the specified user is following.
// It takes the user's ID as a string and returns a slice of primitive.ObjectID.
func GetFollowingUsers(userID string) ([]models.User, error) {
	currID, _ := primitive.ObjectIDFromHex(userID)

	cursor, err := config.FollowCollection.Find(ctx, bson.M{"followerId": currID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var followingIDs []primitive.ObjectID
	for cursor.Next(ctx) {
		var follow models.Follow
		cursor.Decode(&follow)
		followingIDs = append(followingIDs, follow.FollowingID)
	}

	userCursor, err := config.UserCollection.Find(ctx, bson.M{"_id": bson.M{"$in": followingIDs}})
	if err != nil {
		return nil, err
	}
	defer userCursor.Close(ctx)

	var users []models.User
	for userCursor.Next(ctx) {
		var u models.User
		userCursor.Decode(&u)
		users = append(users, u)
	}
	return users, nil
}
