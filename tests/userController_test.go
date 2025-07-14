package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akshayverma91/rydeapi/controllers"
	"github.com/akshayverma91/rydeapi/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock repository functions
var (
	mockUserID = "507f1f77bcf86cd799439011"
)

func TestCreateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/users", controllers.CreateUserHandler)

	user := models.User{
		Name:        "Test User",
		Address:     "Test Address",
		Description: "Test Description",
	}
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// You may want to mock repositories.CreateUser here for true unit test

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "id")
}

func TestCreateUserHandler_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/users", controllers.CreateUserHandler)

	invalidBody := []byte(`{"name":123}`) // name should be string
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(invalidBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
