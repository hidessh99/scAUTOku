package utils

import (
	"backend/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// APIKeyAuth middleware for API key authentication
func APIKeyAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the API key from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format",
			})
		}

		// Extract the API key
		apiKey := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the API key
		if apiKey != config.AppConfig.APIKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}

		// Continue to the next handler
		return c.Next()
	}
}

// BasicAuth middleware for basic authentication
func BasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		// Check if the header starts with "Basic "
		if !strings.HasPrefix(authHeader, "Basic ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format",
			})
		}

		// For simplicity, we're just checking if the header exists
		// In a real application, you would decode and validate the credentials
		return c.Next()
	}
}

// APIKeyHeaderAuth middleware for X-API-Key header authentication
func APIKeyHeaderAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the API key from the X-API-Key header
		apiKey := c.Get("X-API-Key")
		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing X-API-Key header",
			})
		}

		// Validate the API key
		if apiKey != config.AppConfig.APIKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}

		// Continue to the next handler
		return c.Next()
	}
}
