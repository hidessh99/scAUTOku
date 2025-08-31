package controllers

import (
	"backend/models"
	"backend/usecases"

	"github.com/gofiber/fiber/v2"
)

type ShadowsocksController struct {
	shadowsocksUsecase usecases.ShadowsocksUsecase
}

func NewShadowsocksController(shadowsocksUsecase usecases.ShadowsocksUsecase) *ShadowsocksController {
	return &ShadowsocksController{
		shadowsocksUsecase: shadowsocksUsecase,
	}
}

// CreateAccount creates a new SHADOWSOCKS account
func (ac *ShadowsocksController) CreateAccount(c *fiber.Ctx) error {
	var req models.CreateAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" || req.Password == "" || req.Exp == "" || req.ServerID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username, password, expiration, and server ID are required",
		})
	}

	// Set account type
	req.AccountType = models.SHADOWSOCKS

	response, err := ac.shadowsocksUsecase.CreateAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// CheckAccount checks an existing SHADOWSOCKS account
func (ac *ShadowsocksController) CheckAccount(c *fiber.Ctx) error {
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

	// Set account type
	req.AccountType = models.SHADOWSOCKS

	response, err := ac.shadowsocksUsecase.CheckAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteAccount deletes an existing SHADOWSOCKS account
func (ac *ShadowsocksController) DeleteAccount(c *fiber.Ctx) error {
	var req models.DeleteAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" || req.ServerID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username and server ID are required",
		})
	}

	// Set account type
	req.AccountType = models.SHADOWSOCKS

	response, err := ac.shadowsocksUsecase.DeleteAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// RenewAccount renews an existing SHADOWSOCKS account
func (ac *ShadowsocksController) RenewAccount(c *fiber.Ctx) error {
	var req models.RenewAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" || req.Exp == "" || req.ServerID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username, expiration, and server ID are required",
		})
	}

	// Set account type
	req.AccountType = models.SHADOWSOCKS

	response, err := ac.shadowsocksUsecase.RenewAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
