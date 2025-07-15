package controllers

import (
	"net/http"
	"time"

	"github.com/akshayverma91/rydeapi/config"
	"github.com/akshayverma91/rydeapi/models"
	"github.com/akshayverma91/rydeapi/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body models.AuthUser true "User registration info"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var authUser models.AuthUser
	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// here we hash password before saving it
	hash, _ := bcrypt.GenerateFromPassword([]byte(authUser.Password), bcrypt.DefaultCost)
	// Compose new User with extra auth fields
	newUser := bson.D{
		{Key: "name", Value: authUser.User.Name},
		{Key: "dob", Value: authUser.User.DOB},
		{Key: "address", Value: authUser.User.Address},
		{Key: "description", Value: authUser.User.Description},
		{Key: "created_at", Value: time.Now()},
		{Key: "email", Value: authUser.Email},
		{Key: "password", Value: string(hash)},
	}
	authUser.Password = string(hash)

	// insert in db
	res, err := config.UserCollection.InsertOne(c, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"new user id": res.InsertedID})
}

// @Summary Login and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var loginUser models.LoginRequest
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find user by email
	var user models.LoginRequest
	err := config.UserCollection.FindOne(c, gin.H{"email": loginUser.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJwtToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
