package main

import (
	"log"

	"github.com/akshayverma91/rydeapi/config"
	"github.com/akshayverma91/rydeapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	// Initialize the Gin router and MongoDB connection
	r := gin.Default()
	config.InitializeMongoDB()
	routes.RegisterRoutes(r)

	r.Run(":" + port) // Start the server on the specified port
	log.Println("Server running on port:", port)
}
