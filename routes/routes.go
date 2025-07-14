package routes

import (
	"github.com/akshayverma91/rydeapi/controllers"
	"github.com/gin-gonic/gin"
)

// registers the user-related routes with the provided Gin router.
func RegisterRoutes(routes *gin.Engine) {
	api := routes.Group("/api")
	{
		api.POST("/users", controllers.CreateUserHandler)
		api.GET("/users", controllers.GetAllUsersHandler)
		api.GET("/users/:id", controllers.GetUserByIdHandler)
		api.PUT("/users/:id", controllers.UpdateUserHandler)
		api.DELETE("/users/:id", controllers.DeleteUserHandler)
	}

}
