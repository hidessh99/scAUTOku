package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"os/exec"
	"strings"
)

type VmessUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type vmessUsecase struct {
	validator *utils.Validator
}

func NewVmessUsecase() VmessUsecase {
	return &vmessUsecase{
		validator: utils.NewValidator(),
	}
}

func (uc *vmessUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/add-vmess"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		req.Username, req.Password, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot create account Vmess ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	parsedData := utils.OutputParseVMess(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "VMESS account created successfully",
		Data:    parsedData,
	}, nil
}

func (uc *vmessUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/check-vmess"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot check account Vmess ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	// Parse the output using the new parser
	parsedData := utils.ParseCheckAccountOutput(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "Check account Vmess successfully",
		Data:    parsedData,
	}, nil

}

func (uc *vmessUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/del-vmess"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot delete account Vmess ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "VMESS account deleted successfully",
	}, nil
}

func (uc *vmessUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/renew-vmess"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		req.Username, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot create renew account Vmess ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "VMESS account renewed successfully",
		data:    req.Exp,
	}, nil
}
