package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
)

type SshUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type sshUsecase struct{}

func NewSshUsecase() SshUsecase {
	return &sshUsecase{}
}

func (uc *sshUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-ssh-user script
	scriptArgs := []string{
		req.Username,
		req.Password,
		req.Exp,
		req.IPQuota,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/add-ssh-user", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to create SSH account: %v", err),
		}, err
	}

	data := models.SSHAccountData{
		Username: req.Username,
		Password: req.Password,
		Domain:   "example.com", // Would be extracted from output
		Expired:  req.Exp,
		IPQuota:  req.IPQuota,
		Pubkey:   "public-key", // Would be extracted from output
	}

	return &models.AccountResponse{
		Status:  "success",
		Message: "SSH account created successfully",
		Data:    data,
	}, nil
}

func (uc *sshUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the check-ssh script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/check-ssh", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to check SSH account: %v", err),
		}, err
	}

	// Parse output (simplified)

	return &models.AccountResponse{
		Status:  "success",
		Message: "SSH account details retrieved",
		Data: map[string]interface{}{
			"username": req.Username,
			"status":   "active", // Would be parsed from output
			"usage":    "1GB",    // Would be parsed from output
		},
	}, nil
}

func (uc *sshUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the del-ssh script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("/usr/local/bin/del-ssh", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to delete SSH account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "SSH account deleted successfully",
	}, nil
}

func (uc *sshUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-ssh script
	scriptArgs := []string{
		req.Username,
		req.Exp,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/renew-ssh", scriptArgs...)
	if err != nil {
		return &models.AccountResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to renew SSH account: %v", err),
		}, err
	}

	// Parse output (simplified)
	return &models.AccountResponse{
		Status:  "success",
		Message: "SSH account renewed successfully",
	}, nil
}