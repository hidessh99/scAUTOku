package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"os/exec"
	"strings"
)

type ShadowsocksUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type shadowsocksUsecase struct {
	validator *utils.Validator
}

func NewShadowsocksUsecase() ShadowsocksUsecase {
	return &shadowsocksUsecase{
		validator: utils.NewValidator(),
	}
}

func (uc *shadowsocksUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/add-addshadowsocks"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		req.Username, req.Password, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot create account SHADOWSOCKS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	parsedData := utils.OutputParseShadowsocks(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "SHADOWSOCKS account created successfully",
		Data:    parsedData,
	}, nil
}

func (uc *shadowsocksUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/check-shadowsocks"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot check account SHADOWSOCKS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	// Parse the output using the new parser
	parsedData := utils.ParseCheckAccountOutput(string(output))

	return &models.AccountResponse{
		Status:  "success",
		Message: "Check account SHADOWSOCKS successfully",
		Data:    parsedData,
	}, nil
}

func (uc *shadowsocksUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/del-addshadowsocks"

	input := fmt.Sprintf("%s\n",
		req.Username)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot delete account SHADOWSOCKS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "SHADOWSOCKS account deleted successfully",
	}, nil
}

func (uc *shadowsocksUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Validate the request
	if validationErrors := uc.validator.ValidateStruct(req); len(validationErrors) > 0 {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    validationErrors,
		}, fmt.Errorf("validation failed: %v", validationErrors)
	}

	scriptPath := "/usr/local/bin/renew-shadowsocks"

	input := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		req.Username, req.Exp, req.Quota, req.IPQuota)

	// Execute the script with input
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdin = strings.NewReader(input)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: "Cannot renew account SHADOWSOCKS ",
		}, fmt.Errorf("validation failed: %v", err)
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "SHADOWSOCKS account renewed successfully",
	}, nil
}
