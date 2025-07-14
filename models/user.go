package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `json:"name"`
	DOB         time.Time          `json:"dob"`
	Address     string             `json:"address"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
}
