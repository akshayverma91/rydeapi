// @title Ryde API
// @version 1.0
// @description This is the API documentation for the Ryde api application.
// @host localhost:8080
// @BasePath /api
// @schemes http
package main

import (
	"log"

	"github.com/akshayverma91/rydeapi/config"
	_ "github.com/akshayverma91/rydeapi/docs"
	"github.com/akshayverma91/rydeapi/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	port := "8080"

	// Initialize the Gin router and MongoDB connection
	r := gin.Default()
	config.InitializeMongoDB()
	routes.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port) // Start the server on the specified port
	log.Println("Server running on port:", port)
}
