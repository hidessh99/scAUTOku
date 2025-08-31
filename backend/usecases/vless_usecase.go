package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
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

	// Execute the add-vless script
	scriptArgs := []string{
		req.Username,
		req.Exp,
		req.Quota,
		req.IPQuota,
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/add-vless", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to create VLESS account: %v", err),
		}, err
	}

	data := models.VlessAccountData{
		Username:      req.Username,
		Domain:        "example.com", // Would be extracted from output
		Expired:       req.Exp,
		IPQuota:       req.IPQuota,
		UUID:          "generated-uuid", // Would be extracted from output
		VlessTLSLink:  "vless://tls-link",
		VlessGRPCLink: "vless://grpc-link",
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "VLESS account created successfully",
		Data:    data,
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

	// Execute the check-vless script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/check-vless", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to check VLESS account: %v", err),
		}, err
	}

	// Parse output (simplified)

	return &models.AccountResponse{
		Status:  "success",
		Message: "VLESS account details retrieved",
		Data: map[string]interface{}{
			"username": req.Username,
			"status":   "active", // Would be parsed from output
			"usage":    "1GB",    // Would be parsed from output
		},
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

	// Execute the del-vless script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/del-vless", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to delete VLESS account: %v", err),
		}, err
	}

	// Parse output (simplified)
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

	// Execute the renew-vless script
	scriptArgs := []string{
		req.Username,
		req.Exp,
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/renew-vless", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to renew VLESS account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "VLESS account renewed successfully",
	}, nil
}
