package controllers

import (
	"backend/models"
	"backend/usecases"

	"github.com/gofiber/fiber/v2"
)

type TrojanController struct {
	trojanUsecase usecases.TrojanUsecase
}

func NewTrojanController(trojanUsecase usecases.TrojanUsecase) *TrojanController {
	return &TrojanController{
		trojanUsecase: trojanUsecase,
	}
}

// CreateAccount creates a new TROJAN account
func (ac *TrojanController) CreateAccount(c *fiber.Ctx) error {
	var req models.CreateAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if req.Username == "" || req.Password == "" || req.Exp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.AccountResponse{
			Status:  "error",
			Message: "Username, password, and expiration are required",
		})
	}

	// Set account type
	req.AccountType = models.TROJAN

	response, err := ac.trojanUsecase.CreateAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// CheckAccount checks an existing TROJAN account
func (ac *TrojanController) CheckAccount(c *fiber.Ctx) error {
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
	req.AccountType = models.TROJAN

	response, err := ac.trojanUsecase.CheckAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteAccount deletes an existing TROJAN account
func (ac *TrojanController) DeleteAccount(c *fiber.Ctx) error {
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
	req.AccountType = models.TROJAN

	response, err := ac.trojanUsecase.DeleteAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// RenewAccount renews an existing TROJAN account
func (ac *TrojanController) RenewAccount(c *fiber.Ctx) error {
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
	req.AccountType = models.TROJAN

	response, err := ac.trojanUsecase.RenewAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
