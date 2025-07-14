package tests

import (
	"testing"
	"time"

	"github.com/akshayverma91/rydeapi/config"
	"github.com/akshayverma91/rydeapi/models"
	"github.com/akshayverma91/rydeapi/repositories"
)

func TestCreateUser(t *testing.T) {
	config.InitializeMongoDB()

	user := models.User{
		Name:        "Test User One",
		DOB:         time.Date(2010, time.June, 12, 01, 10, 0, 0, time.UTC),
		Address:     "Test Address One",
		Description: "Test user for testing.",
	}

	id, err := repositories.CreateUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if id.Hex() == "" || len(id.Hex()) != 24 {
		t.Errorf("Expected a valid ObjectID, got: %v", id.Hex())
	}
}
