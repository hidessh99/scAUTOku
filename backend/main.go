package main

import (
	"log"

	"backend/config"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "VPN Account Management API",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization,X-API-Key",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Setup routes
	routes.AccountRoutes(app)

	// Get port from configuration
	port := config.AppConfig.ServerPort

	// Start server
	log.Printf("🚀 Starting server on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
