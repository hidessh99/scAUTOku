package routes

import (
	"backend/controllers"
	"backend/usecases"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

// AccountRoutes sets up the account management routes
func AccountRoutes(app *fiber.App) {
	// Initialize usecases and controllers
	accountUsecase := usecases.NewAccountUsecase()
	accountController := controllers.NewAccountController(accountUsecase)

	// Public routes
	// Health check endpoint
	app.Get("/health", accountController.Health)

	// Protected routes group
	api := app.Group("/api/v1")

	// Apply API key authentication middleware to all protected routes
	api.Use(utils.APIKeyAuth())

	// Account management endpoints
	api.Post("/accounts", accountController.CreateAccount)
	api.Post("/accounts/check", accountController.CheckAccount)
	api.Post("/accounts/delete", accountController.DeleteAccount)

	// Alternative routes for each account type
	accounts := api.Group("/accounts")

	// VMESS routes
	vmess := accounts.Group("/vmess")
	vmess.Post("/", accountController.CreateAccount)
	vmess.Post("/check", accountController.CheckAccount)
	vmess.Post("/delete", accountController.DeleteAccount)

	// SSH routes
	ssh := accounts.Group("/ssh")
	ssh.Post("/", accountController.CreateAccount)
	ssh.Post("/check", accountController.CheckAccount)
	ssh.Post("/delete", accountController.DeleteAccount)

	// TROJAN routes
	trojan := accounts.Group("/trojan")
	trojan.Post("/", accountController.CreateAccount)
	trojan.Post("/check", accountController.CheckAccount)
	trojan.Post("/delete", accountController.DeleteAccount)

	// VLESS routes
	vless := accounts.Group("/vless")
	vless.Post("/", accountController.CreateAccount)
	vless.Post("/check", accountController.CheckAccount)
	vless.Post("/delete", accountController.DeleteAccount)

	// SHADOWSOCKS routes
	shadowsocks := accounts.Group("/shadowsocks")
	shadowsocks.Post("/", accountController.CreateAccount)
	shadowsocks.Post("/check", accountController.CheckAccount)
	shadowsocks.Post("/delete", accountController.DeleteAccount)
}
