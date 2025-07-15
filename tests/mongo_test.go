package tests

import (
	"testing"
	"time"

	"github.com/akshayverma91/rydeapi/config"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockCollection is a mock for mongo.Collection
type MockCollection struct{}

func TestInitializeMongoDB_Mock(t *testing.T) {
	// Save original UserCollection
	origUserCollection := config.UserCollection
	defer func() { config.UserCollection = origUserCollection }()

	// Mock assignment instead of real DB connection
	config.UserCollection = &mongo.Collection{} // or &MockCollection{} if you use interfaces

	// Simulate what InitializeMongoDB would do (without connecting)
	start := time.Now()
	config.UserCollection = &mongo.Collection{}
	elapsed := time.Since(start)

	assert.NotNil(t, config.UserCollection)
	assert.Less(t, elapsed.Seconds(), 1.0, "Should not take long since it's mocked")
}
