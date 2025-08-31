package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"os/exec"
	"strings"
)

type VlessUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type vlessUsecase struct {
	validator *utils.Validator
}

func NewVlessUsecase() VlessUsecase {
	return &vlessUsecase{
		validator: utils.NewValidator(),
	}
}

func (uc *vlessUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/add-vless"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		req.Username, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot create account VLESS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	parsedData := utils.OutputParseVLess(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "VLESS account created successfully",
		Data:    parsedData,
	}, nil
}

func (uc *vlessUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/check-vless"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot check account VLESS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	// Parse the output using the new parser
	parsedData := utils.ParseCheckAccountOutput(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "Check account VLESS successfully",
		Data:    parsedData,
	}, nil
}

func (uc *vlessUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/del-vless"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot delete account VLESS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "VLESS account deleted successfully",
	}, nil
}

func (uc *vlessUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/renew-vless"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		req.Username, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot renew account VLESS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "VLESS account renewed successfully",
	}, nil
}
