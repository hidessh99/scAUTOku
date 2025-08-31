package routes

import (
	"backend/controllers"
	"backend/usecases"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

// AccountRoutes sets up the account management routes
func AccountRoutes(app *fiber.App) {
	// Initialize usecases with validators
	vmessUsecase := usecases.NewVmessUsecase()
	sshUsecase := usecases.NewSshUsecase()
	vlessUsecase := usecases.NewVlessUsecase()
	trojanUsecase := usecases.NewTrojanUsecase()
	shadowsocksUsecase := usecases.NewShadowsocksUsecase()

	// Initialize main account usecase
	accountUsecase := usecases.NewAccountUsecase(
		vmessUsecase,
		sshUsecase,
		vlessUsecase,
		trojanUsecase,
		shadowsocksUsecase,
	)

	// Initialize controllers
	vmessController := controllers.NewVmessController(vmessUsecase)
	sshController := controllers.NewSshController(sshUsecase)
	vlessController := controllers.NewVlessController(vlessUsecase)
	trojanController := controllers.NewTrojanController(trojanUsecase)
	shadowsocksController := controllers.NewShadowsocksController(shadowsocksUsecase)

	// Initialize main account controller
	accountController := controllers.NewAccountController(
		accountUsecase,
		vmessController,
		sshController,
		vlessController,
		trojanController,
		shadowsocksController,
	)

	// Public routes
	// Health check endpoint
	app.Get("/health", accountController.Health)

	// Protected routes group
	api := app.Group("/api/v1")

	// Apply API key authentication middleware to all protected routes
	api.Use(utils.APIKeyAuth())

	// General account management endpoints
	api.Post("/accounts", accountController.CreateAccount)
	api.Post("/accounts/check", accountController.CheckAccount)
	api.Post("/accounts/delete", accountController.DeleteAccount)
	api.Post("/accounts/renew", accountController.RenewAccount)

	// Protocol specific routes
	// VMESS routes
	vmess := api.Group("/vmess")
	vmess.Post("/", vmessController.CreateAccount)
	vmess.Post("/check", vmessController.CheckAccount)
	vmess.Post("/delete", vmessController.DeleteAccount)
	vmess.Post("/renew", vmessController.RenewAccount)

	// SSH routes
	ssh := api.Group("/ssh")
	ssh.Post("/", sshController.CreateAccount)
	ssh.Post("/check", sshController.CheckAccount)
	ssh.Post("/delete", sshController.DeleteAccount)
	ssh.Post("/renew", sshController.RenewAccount)

	// VLESS routes
	vless := api.Group("/vless")
	vless.Post("/", vlessController.CreateAccount)
	vless.Post("/check", vlessController.CheckAccount)
	vless.Post("/delete", vlessController.DeleteAccount)
	vless.Post("/renew", vlessController.RenewAccount)

	// TROJAN routes
	trojan := api.Group("/trojan")
	trojan.Post("/", trojanController.CreateAccount)
	trojan.Post("/check", trojanController.CheckAccount)
	trojan.Post("/delete", trojanController.DeleteAccount)
	trojan.Post("/renew", trojanController.RenewAccount)

	// SHADOWSOCKS routes
	shadowsocks := api.Group("/shadowsocks")
	shadowsocks.Post("/", shadowsocksController.CreateAccount)
	shadowsocks.Post("/check", shadowsocksController.CheckAccount)
	shadowsocks.Post("/delete", shadowsocksController.DeleteAccount)
	shadowsocks.Post("/renew", shadowsocksController.RenewAccount)
}
