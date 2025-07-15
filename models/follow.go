package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// swagger:model Follow
// Follow represents a follow relationship between users.
type Follow struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"` // read only: true
	FollowerID  primitive.ObjectID `bson:"followerId" json:"followerId"`
	FollowingID primitive.ObjectID `bson:"followingId" json:"followingId"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
