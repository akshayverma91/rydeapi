package controllers

import (
	"net/http"

	"github.com/akshayverma91/rydeapi/models"
	"github.com/akshayverma91/rydeapi/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserHandler handles the creation of a new user.
// It expects a JSON payload with user details and returns the created user's ID or an error.
// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User Body"
// @Success 201 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := repositories.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetAllUsersHandler retrieves all users from the database.
// It returns a list of users or an error if the operation fails.
// @Summary Get all users
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetAllUsersHandler(c *gin.Context) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByIdHandler retrieves a user by their ID.
// It expects the user ID as a URL parameter and returns the user details or an error if
// @Summary Get user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [get]
func GetUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := repositories.GetUserById(objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserHandler updates an existing user in the database.
// It expects the user ID as a URL parameter and the updated user details in the request body
// @Summary Update user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Param user body models.User true "Updated User Body"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = repositories.UpdateUser(objectID, updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUserHandler removes a user from the database by their ID.
// It expects the user ID as a URL parameter and returns a success message or an error if
// @Summary Delete user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = repositories.DeleteUser(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
