package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
)

type TrojanUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type trojanUsecase struct{}

func NewTrojanUsecase() TrojanUsecase {
	return &trojanUsecase{}
}

func (uc *trojanUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-trojan-user script
	scriptArgs := []string{
		req.Username,
		req.Password,
		req.Exp,
		req.IPQuota,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/add-trojan-user", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to create TROJAN account: %v", err),
		}, err
	}

	data := models.TrojanAccountData{
		Username:      req.Username,
		Password:      req.Password,
		Domain:        "example.com", // Would be extracted from output
		Expired:       req.Exp,
		IPQuota:       req.IPQuota,
		TrojanTLSLink: "trojan://tls-link",
		TrojanGRPC:    "trojan://grpc-link",
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "TROJAN account created successfully",
		Data:    data,
	}, nil
}

func (uc *trojanUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the check-trojan script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/check-trojan", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to check TROJAN account: %v", err),
		}, err
	}

	// Parse output (simplified)

	return &models.AccountResponse{
		Status:  "success",
		Message: "TROJAN account details retrieved",
		Data: map[string]interface{}{
			"username": req.Username,
			"status":   "active", // Would be parsed from output
			"usage":    "1GB",    // Would be parsed from output
		},
	}, nil
}

func (uc *trojanUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the del-trojan script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/del-trojan", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to delete TROJAN account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "TROJAN account deleted successfully",
	}, nil
}

func (uc *trojanUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-trojan script
	scriptArgs := []string{
		req.Username,
		req.Exp,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/renew-trojan", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to renew TROJAN account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "TROJAN account renewed successfully",
	}, nil
}