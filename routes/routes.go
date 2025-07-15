package routes

import (
	"github.com/akshayverma91/rydeapi/controllers"
	"github.com/akshayverma91/rydeapi/middleware"
	"github.com/gin-gonic/gin"
)

// registers the user-related routes with the provided Gin router.
func RegisterRoutes(routes *gin.Engine) {
	api := routes.Group("/api")

	// public routes
	api.POST("/auth/register", controllers.Register)
	api.POST("/auth/login", controllers.Login)

	// auth routes
	authRoutes := api.Group("/v1/users")
	// Apply JWT middleware to user routes
	authRoutes.Use(middleware.JwtAuthMiddleware())
	{
		authRoutes.POST("/", controllers.CreateUserHandler)
		authRoutes.GET("/", controllers.GetAllUsersHandler)
		authRoutes.GET("/:id", controllers.GetUserByIdHandler)
		authRoutes.PUT("/:id", controllers.UpdateUserHandler)
		authRoutes.DELETE("/:id", controllers.DeleteUserHandler)
	}

}
