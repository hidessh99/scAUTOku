package controllers

import (
	"backend/models"
	"backend/usecases"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	accountUsecase        usecases.AccountUsecase
	vmessController       *VmessController
	sshController         *SshController
	vlessController       *VlessController
	trojanController      *TrojanController
	shadowsocksController *ShadowsocksController
}

func NewAccountController(
	accountUsecase usecases.AccountUsecase,
	vmessController *VmessController,
	sshController *SshController,
	vlessController *VlessController,
	trojanController *TrojanController,
	shadowsocksController *ShadowsocksController,
) *AccountController {
	return &AccountController{
		accountUsecase:        accountUsecase,
		vmessController:       vmessController,
		sshController:         sshController,
		vlessController:       vlessController,
		trojanController:      trojanController,
		shadowsocksController: shadowsocksController,
	}
}

// CreateAccount creates a new account
func (ac *AccountController) CreateAccount(c *fiber.Ctx) error {
	var req models.CreateAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" || req.Exp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username and expiration are required",
		})
	}

	response, err := ac.accountUsecase.CreateAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// CheckAccount checks an existing account
func (ac *AccountController) CheckAccount(c *fiber.Ctx) error {
	var req models.CheckAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username is required",
		})
	}

	response, err := ac.accountUsecase.CheckAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteAccount deletes an existing account
func (ac *AccountController) DeleteAccount(c *fiber.Ctx) error {
	var req models.DeleteAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username is required",
		})
	}

	response, err := ac.accountUsecase.DeleteAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// RenewAccount renews an existing account
func (ac *AccountController) RenewAccount(c *fiber.Ctx) error {
	var req models.RenewAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" || req.Exp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username and expiration are required",
		})
	}

	response, err := ac.accountUsecase.RenewAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Health check endpoint
func (ac *AccountController) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Account management API is running",
	})
}
