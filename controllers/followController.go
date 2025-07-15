package controllers

import (
	"net/http"

	"github.com/akshayverma91/rydeapi/config"
	"github.com/akshayverma91/rydeapi/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Follow a user
// @Tags Users v2
// @Security BearerAuth
// @Param id path string true "User to follow"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v2/users/{id}/follow [post]
func FollowUserHandler(c *gin.Context) {
	targetID := c.Param("id")
	authEmail := c.GetString("email")

	var user struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err := config.UserCollection.FindOne(c, bson.M{"email": authEmail}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth user not found"})
		return
	}

	if err := repositories.FollowUser(user.ID.Hex(), targetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "follow failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "followed successfully"})
}

// @Summary Unfollow a user
// @Tags Users v2
// @Security BearerAuth
// @Param id path string true "User to unfollow"
// @Success 200 {object} map[string]string
// @Router /api/v2/users/{id}/unfollow [post]
func UnfollowUserHandler(c *gin.Context) {
	targetID := c.Param("id")
	authEmail := c.GetString("email")

	var user struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err := config.UserCollection.FindOne(c, bson.M{"email": authEmail}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth user not found"})
		return
	}

	if err := repositories.UnfollowUser(user.ID.Hex(), targetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unfollow failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "unfollowed successfully"})
}

// @Summary Get followers of a user
// @Tags Users v2
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {array} models.User
// @Router /api/v2/users/{id}/followers [get]
func GetFollowersHandler(c *gin.Context) {
	id := c.Param("id")
	users, err := repositories.GetFollowerUsers(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get followers"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get following of a user
// @Tags Users v2
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {array} models.User
// @Router /api/v2/users/{id}/following [get]
func GetFollowingHandler(c *gin.Context) {
	id := c.Param("id")
	users, err := repositories.GetFollowingUsers(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get following"})
		return
	}

	c.JSON(http.StatusOK, users)
}
