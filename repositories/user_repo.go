package repositories

import (
	"context"
	"time"

	"github.com/akshayverma91/rydeapi/config"
	"github.com/akshayverma91/rydeapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Context for MongoDB operations
var cxt = context.Background()

// CreateUser inserts a new user into the database.
// It returns the ID of the created user or an error if the operation fails.
func CreateUser(user models.User) (primitive.ObjectID, error) {
	user.CreatedAt = time.Now()

	result, err := config.UserCollection.InsertOne(cxt, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// GetAllUsers retrieves all users from the database.
// It returns a slice of User models or an error if the operation fails.
func GetAllUsers() ([]models.User, error) {
	cursor, err := config.UserCollection.Find(cxt, primitive.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(cxt)

	var users []models.User
	for cursor.Next(cxt) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserById retrieves a user by their ID.
// It takes a primitive.ObjectID and returns the User model or an error if the operation fails.
func GetUserById(id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := config.UserCollection.FindOne(cxt, primitive.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database.
// It takes the user ID and the updated User model, returning an error if the operation fails
func UpdateUser(id primitive.ObjectID, updatedUser models.User) error {
	updatedUser.ID = id
	updatedUser.CreatedAt = time.Now()

	_, err := config.UserCollection.UpdateOne(cxt, primitive.M{"_id": id}, primitive.M{"$set": updatedUser})
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser removes a user from the database by their ID.
// It takes a primitive.ObjectID and returns an error if the operation fails.
func DeleteUser(id primitive.ObjectID) error {
	_, err := config.UserCollection.DeleteOne(cxt, primitive.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
