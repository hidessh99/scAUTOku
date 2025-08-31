package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
)

type ShadowsocksUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type shadowsocksUsecase struct{}

func NewShadowsocksUsecase() ShadowsocksUsecase {
	return &shadowsocksUsecase{}
}

func (uc *shadowsocksUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-shadowsocks-user script
	scriptArgs := []string{
		req.Username,
		req.Password,
		req.Exp,
		req.IPQuota,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/add-shadowsocks-user", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to create SHADOWSOCKS account: %v", err),
		}, err
	}

	data := models.ShadowsocksAccountData{
		Username: req.Username,
		Password: req.Password,
		Domain:   "example.com", // Would be extracted from output
		Expired:  req.Exp,
		IPQuota:  req.IPQuota,
		Method:   "aes-256-gcm", // Would be extracted from output
		SSLink:   "ss://shadowsocks-link",
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "SHADOWSOCKS account created successfully",
		Data:    data,
	}, nil
}

func (uc *shadowsocksUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the check-shadowsocks script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/check-shadowsocks", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to check SHADOWSOCKS account: %v", err),
		}, err
	}

	// Parse output (simplified)

	return &models.AccountResponse{
		Status:  "success",
		Message: "SHADOWSOCKS account details retrieved",
		Data: map[string]interface{}{
			"username": req.Username,
			"status":   "active", // Would be parsed from output
			"usage":    "1GB",    // Would be parsed from output
		},
	}, nil
}

func (uc *shadowsocksUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the del-addshadowsocks script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/del-addshadowsocks", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to delete SHADOWSOCKS account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "SHADOWSOCKS account deleted successfully",
	}, nil
}

func (uc *shadowsocksUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-shadowsocks script
	scriptArgs := []string{
		req.Username,
		req.Exp,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/renew-shadowsocks", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to renew SHADOWSOCKS account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "SHADOWSOCKS account renewed successfully",
	}, nil
}