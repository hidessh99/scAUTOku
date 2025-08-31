package controllers

import (
	"backend/models"
	"backend/usecases"

	"github.com/gofiber/fiber/v2"
)

type VlessController struct {
	vlessUsecase usecases.VlessUsecase
}

func NewVlessController(vlessUsecase usecases.VlessUsecase) *VlessController {
	return &VlessController{
		vlessUsecase: vlessUsecase,
	}
}

// CreateAccount creates a new VLESS account
func (ac *VlessController) CreateAccount(c *fiber.Ctx) error {
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

	// Set account type
	req.AccountType = models.VLESS

	response, err := ac.vlessUsecase.CreateAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// CheckAccount checks an existing VLESS account
func (ac *VlessController) CheckAccount(c *fiber.Ctx) error {
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
	req.AccountType = models.VLESS

	response, err := ac.vlessUsecase.CheckAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteAccount deletes an existing VLESS account
func (ac *VlessController) DeleteAccount(c *fiber.Ctx) error {
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

	// Set account type
	req.AccountType = models.VLESS

	response, err := ac.vlessUsecase.DeleteAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// RenewAccount renews an existing VLESS account
func (ac *VlessController) RenewAccount(c *fiber.Ctx) error {
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

	// Set account type
	req.AccountType = models.VLESS

	response, err := ac.vlessUsecase.RenewAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
