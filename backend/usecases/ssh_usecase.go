package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"os/exec"
	"strings"
)

type SshUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type sshUsecase struct {
	validator *utils.Validator
}

func NewSshUsecase() SshUsecase {
	return &sshUsecase{
		validator: utils.NewValidator(),
	}
}

func (uc *sshUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/add-ssh"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		req.Username, req.Password, req.Exp, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot create account SSH ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	parsedData := utils.OutputParseSsh(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "SSH account created successfully",
		Data:    parsedData,
	}, nil
}

func (uc *sshUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/check-ssh"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot check account SSH ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	// Parse the output using the new parser
	parsedData := utils.ParseCheckAccountOutput(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "Check account SSH successfully",
		Data:    parsedData,
	}, nil
}

func (uc *sshUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/del-ssh"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot delete account SSH ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "SSH account deleted successfully",
	}, nil
}

func (uc *sshUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/renew-ssh"

	input := fmt.Sprintf("%s\n%s\n",
		req.Username, req.Exp)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot renew account SSH ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "SSH account renewed successfully",
	}, nil
}
