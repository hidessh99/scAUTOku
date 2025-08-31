package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
)

type VmessUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type vmessUsecase struct{}

func NewVmessUsecase() VmessUsecase {
	return &vmessUsecase{}
}

func (uc *vmessUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-vmess-user script with appropriate parameters
	scriptArgs := []string{
		req.Username,
		req.Exp,
		req.Quota,
		req.IPQuota,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/add-vmess-user", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to create VMESS account: %v", err),
		}, err
	}

	// Parse the output to extract account details
	// This would need to be adapted based on actual script output format
	data := models.VmessAccountData{
		Username:        req.Username,
		Domain:          "example.com", // Would be extracted from output
		Quota:           req.Quota,
		IPQuota:         req.IPQuota,
		Expired:         req.Exp,
		UUID:            "generated-uuid", // Would be extracted from output
		Pubkey:          "public-key",     // Would be extracted from output
		VmessTLSLink:    "vmess://tls-link",
		VmessNonTLSLink: "vmess://non-tls-link",
		VmessGRPCLink:   "vmess://grpc-link",
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "VMESS account created successfully",
		Data:    data,
	}, nil
}

func (uc *vmessUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the check-vmess script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/check-vmess", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to check VMESS account: %v", err),
		}, err
	}

	// Parse output (simplified)

	return &models.AccountResponse{
		Status:  "success",
		Message: "VMESS account details retrieved",
		Data: map[string]interface{}{
			"username": req.Username,
			"status":   "active", // Would be parsed from output
			"usage":    "1GB",    // Would be parsed from output
		},
	}, nil
}

func (uc *vmessUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the del-vmess script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/del-vmess", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to delete VMESS account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "VMESS account deleted successfully",
	}, nil
}

func (uc *vmessUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-vmess script
	scriptArgs := []string{
		req.Username,
		req.Exp,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/renew-vmess", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to renew VMESS account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "VMESS account renewed successfully",
	}, nil
}