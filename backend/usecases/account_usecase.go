package usecases

import (
	"backend/models"
	"backend/utils"
	"fmt"
)

type AccountUsecase interface {
	CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error)
	CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error)
	DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error)
}

type accountUsecase struct{}

func NewAccountUsecase() AccountUsecase {
	return &accountUsecase{}
}

func (uc *accountUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	switch req.AccountType {
	case models.VMESS:
		return uc.createVmessAccount(req)
	case models.SSH:
		return uc.createSSHAccount(req)
	case models.TROJAN:
		return uc.createTrojanAccount(req)
	case models.VLESS:
		return uc.createVlessAccount(req)
	case models.SHADOWSOCKS:
		return uc.createShadowsocksAccount(req)
	default:
		return &models.AccountResponse{
			Status:  "error",
			Message: "Unsupported account type",
		}, fmt.Errorf("unsupported account type: %s", req.AccountType)
	}
}

func (uc *accountUsecase) CheckAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	switch req.AccountType {
	case models.VMESS:
		return uc.checkVmessAccount(req)
	case models.SSH:
		return uc.checkSSHAccount(req)
	case models.TROJAN:
		return uc.checkTrojanAccount(req)
	case models.VLESS:
		return uc.checkVlessAccount(req)
	case models.SHADOWSOCKS:
		return uc.checkShadowsocksAccount(req)
	default:
		return &models.AccountResponse{
			Status:  "error",
			Message: "Unsupported account type",
		}, fmt.Errorf("unsupported account type: %s", req.AccountType)
	}
}

func (uc *accountUsecase) DeleteAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	switch req.AccountType {
	case models.VMESS:
		return uc.deleteVmessAccount(req)
	case models.SSH:
		return uc.deleteSSHAccount(req)
	case models.TROJAN:
		return uc.deleteTrojanAccount(req)
	case models.VLESS:
		return uc.deleteVlessAccount(req)
	case models.SHADOWSOCKS:
		return uc.deleteShadowsocksAccount(req)
	default:
		return &models.AccountResponse{
			Status:  "error",
			Message: "Unsupported account type",
		}, fmt.Errorf("unsupported account type: %s", req.AccountType)
	}
}

// VMESS implementations
func (uc *accountUsecase) createVmessAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-vmess-user script with appropriate parameters
	scriptArgs := []string{
		req.Username,
		req.Password
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

func (uc *accountUsecase) checkVmessAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the checkvmess script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/checkuservmess.sh", scriptArgs...)
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

func (uc *accountUsecase) deleteVmessAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the dellvmess script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/dellaccvmess.sh", scriptArgs...)
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

// SSH implementations
func (uc *accountUsecase) createSSHAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
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

func (uc *accountUsecase) checkSSHAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the checkssh script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/checkuserssh.sh", scriptArgs...)
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

func (uc *accountUsecase) deleteSSHAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the dellssh script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/dellaccssh.sh", scriptArgs...)
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

// TROJAN implementations
func (uc *accountUsecase) createTrojanAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
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

func (uc *accountUsecase) checkTrojanAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the checktrojan script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/checkusertrojan.sh", scriptArgs...)
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

func (uc *accountUsecase) deleteTrojanAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the delltrojan script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/dellacctrojan.sh", scriptArgs...)
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

// VLESS implementations
func (uc *accountUsecase) createVlessAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-vless-user script
	scriptArgs := []string{
		req.Username,
		req.Password
		req.Exp,
		req.Quota,
		req.IPQuota,
		fmt.Sprintf("%d", req.ServerID),
	}

	_, err := utils.ExecuteShellCommand("/usr/local/bin/add-vless-user", scriptArgs...)
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

func (uc *accountUsecase) checkVlessAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the checkvless script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/checkuservless.sh", scriptArgs...)
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

func (uc *accountUsecase) deleteVlessAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the dellvless script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/dellaccvless.sh", scriptArgs...)
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

// SHADOWSOCKS implementations
func (uc *accountUsecase) createShadowsocksAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
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

func (uc *accountUsecase) checkShadowsocksAccount(req models.CheckAccountRequest) (*models.AccountResponse, error) {
	// Execute the checkshadowsocks script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/checkusershadowsocks.sh", scriptArgs...)
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

func (uc *accountUsecase) deleteShadowsocksAccount(req models.DeleteAccountRequest) (*models.AccountResponse, error) {
	// Execute the dellshadowsocks script
	scriptArgs := []string{req.Username}
	_, err := utils.ExecuteShellCommand("./project/dellaccshadowsocks.sh", scriptArgs...)
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
