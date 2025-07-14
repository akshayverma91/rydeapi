package routes

import (
	"github.com/akshayverma91/rydeapi/controllers"
	"github.com/gin-gonic/gin"
)

// registers the user-related routes with the provided Gin router.
func RegisterRoutes(routes *gin.Engine) {
	routes.POST("/users", controllers.CreateUserHandler)
	routes.GET("/users", controllers.GetAllUsersHandler)
	routes.GET("/user/:id", controllers.GetUserByIdHandler)
	routes.PUT("/user/:id", controllers.UpdateUserHandler)
	routes.DELETE("/user/:id", controllers.DeleteUserHandler)
}
