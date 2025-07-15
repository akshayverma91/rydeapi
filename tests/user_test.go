package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/akshayverma91/rydeapi/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserJSONMarshaling(t *testing.T) {
	// Arrange
	dob := time.Date(1991, 10, 15, 0, 0, 0, 0, time.UTC)
	createdAt := time.Now()
	user := models.User{
		ID:          primitive.NewObjectID(),
		Name:        "John Doe",
		DOB:         dob,
		Address:     "123 Street",
		Description: "A Go developer",
		CreatedAt:   createdAt,
	}

	// Act
	data, err := json.Marshal(user)
	assert.NoError(t, err)

	// Unmarshal again
	var decoded models.User
	err = json.Unmarshal(data, &decoded)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, user.Name, decoded.Name)
	assert.Equal(t, user.DOB.Format("2006-01-02"), decoded.DOB.Format("2006-01-02"))
	assert.Equal(t, user.Address, decoded.Address)
	assert.Equal(t, user.Description, decoded.Description)
	assert.Equal(t, user.ID, decoded.ID)
}

func TestUserDefaultValues(t *testing.T) {
	user := models.User{}
	assert.Empty(t, user.Name)
	assert.True(t, user.ID.IsZero())
	assert.True(t, user.DOB.IsZero())
	assert.True(t, user.CreatedAt.IsZero())
}
