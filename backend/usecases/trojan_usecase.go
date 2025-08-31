package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"os/exec"
	"strings"
)

type TrojanUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type trojanUsecase struct {
	validator *utils.Validator
}

func NewTrojanUsecase() TrojanUsecase {
	return &trojanUsecase{
		validator: utils.NewValidator(),
	}
}

func (uc *trojanUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/add-trojan"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		req.Username, req.Password, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot create account Trojan ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	parsedData := utils.OutputParseTrojan(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "TROJAN account created successfully",
		Data:    parsedData,
	}, nil
}

func (uc *trojanUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/check-trojan"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot check account Trojan ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	// Parse the output using the new parser
	parsedData := utils.ParseCheckAccountOutput(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "Check account Trojan successfully",
		Data:    parsedData,
	}, nil
}

func (uc *trojanUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/del-trojan"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot delete account Trojan ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "TROJAN account deleted successfully",
	}, nil
}

func (uc *trojanUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/renew-trojan"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		req.Username, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot renew account Trojan ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "TROJAN account renewed successfully",
	}, nil
}
