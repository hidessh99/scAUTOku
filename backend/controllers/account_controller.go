package controllers

import (
	"backend/models"
	"backend/usecases"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	accountUsecase usecases.AccountUsecase
}

func NewAccountController(accountUsecase usecases.AccountUsecase) *AccountController {
	return &AccountController{
		accountUsecase: accountUsecase,
	}
}

// CreateAccount creates a new account
// @Summary Create a new account
// @Description Create a new VPN account (VMESS, SSH, TROJAN, VLESS, SHADOWSOCKS)
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body models.CreateAccountRequest true "Account creation request"
// @Success 200 {object} models.AccountResponse
// @Failure 400 {object} models.AccountResponse
// @Failure 500 {object} models.AccountResponse
// @Router /accounts [post]
func (ac *AccountController) CreateAccount(c *fiber.Ctx) error {
	var req models.CreateAccountRequest

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

	response, err := ac.accountUsecase.CreateAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// CheckAccount checks an existing account
// @Summary Check an account
// @Description Check the status and details of an existing VPN account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body models.CheckAccountRequest true "Account check request"
// @Success 200 {object} models.AccountResponse
// @Failure 400 {object} models.AccountResponse
// @Failure 500 {object} models.AccountResponse
// @Router /accounts/check [post]
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
// @Summary Delete an account
// @Description Delete an existing VPN account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body models.DeleteAccountRequest true "Account deletion request"
// @Success 200 {object} models.AccountResponse
// @Failure 400 {object} models.AccountResponse
// @Failure 500 {object} models.AccountResponse
// @Router /accounts/delete [post]
func (ac *AccountController) DeleteAccount(c *fiber.Ctx) error {
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

	response, err := ac.accountUsecase.DeleteAccount(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Health check endpoint
// @Summary Health check
// @Description Check if the API is running
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (ac *AccountController) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Account management API is running",
	})
}
