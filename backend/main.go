package main

import (
	"jsdc-api/config"
	"jsdc-api/middleware"
	"jsdc-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to database
	config.ConnectDB()
	defer config.CloseDB()

	// Set Gin mode
	gin.SetMode(config.AppConfig.Server.Mode)

	// Create Gin engine
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	port := config.AppConfig.Server.Port
	log.Printf("🚀 ISDC API server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
