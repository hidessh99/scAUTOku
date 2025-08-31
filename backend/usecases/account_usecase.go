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
	RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error)
}

type accountUsecase struct {
	vmessUsecase       VmessUsecase
	sshUsecase         SshUsecase
	vlessUsecase       VlessUsecase
	trojanUsecase      TrojanUsecase
	shadowsocksUsecase ShadowsocksUsecase
}

func NewAccountUsecase(
	vmessUsecase VmessUsecase,
	sshUsecase SshUsecase,
	vlessUsecase VlessUsecase,
	trojanUsecase TrojanUsecase,
	shadowsocksUsecase ShadowsocksUsecase,
) AccountUsecase {
	return &accountUsecase{
		vmessUsecase:       vmessUsecase,
		sshUsecase:         sshUsecase,
		vlessUsecase:       vlessUsecase,
		trojanUsecase:      trojanUsecase,
		shadowsocksUsecase: shadowsocksUsecase,
	}
}

func (uc *accountUsecase) CreateAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	switch req.AccountType {
	case models.VMESS:
		return uc.vmessUsecase.CreateAccount(req)
	case models.SSH:
		return uc.sshUsecase.CreateAccount(req)
	case models.TROJAN:
		return uc.trojanUsecase.CreateAccount(req)
	case models.VLESS:
		return uc.vlessUsecase.CreateAccount(req)
	case models.SHADOWSOCKS:
		return uc.shadowsocksUsecase.CreateAccount(req)
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
		return uc.vmessUsecase.CheckAccount(req)
	case models.SSH:
		return uc.sshUsecase.CheckAccount(req)
	case models.TROJAN:
		return uc.trojanUsecase.CheckAccount(req)
	case models.VLESS:
		return uc.vlessUsecase.CheckAccount(req)
	case models.SHADOWSOCKS:
		return uc.shadowsocksUsecase.CheckAccount(req)
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
		return uc.vmessUsecase.DeleteAccount(req)
	case models.SSH:
		return uc.sshUsecase.DeleteAccount(req)
	case models.TROJAN:
		return uc.trojanUsecase.DeleteAccount(req)
	case models.VLESS:
		return uc.vlessUsecase.DeleteAccount(req)
	case models.SHADOWSOCKS:
		return uc.shadowsocksUsecase.DeleteAccount(req)
	default:
		return &models.AccountResponse{
			Status:  "error",
			Message: "Unsupported account type",
		}, fmt.Errorf("unsupported account type: %s", req.AccountType)
	}
}

func (uc *accountUsecase) RenewAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	switch req.AccountType {
	case models.VMESS:
		return uc.vmessUsecase.RenewAccount(req)
	case models.SSH:
		return uc.sshUsecase.RenewAccount(req)
	case models.TROJAN:
		return uc.trojanUsecase.RenewAccount(req)
	case models.VLESS:
		return uc.vlessUsecase.RenewAccount(req)
	case models.SHADOWSOCKS:
		return uc.shadowsocksUsecase.RenewAccount(req)
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
		req.Exp,
		req.Quota,
		req.IPQuota,
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

func (uc *accountUsecase) renewVmessAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-vmess script
	scriptArgs := []string{
		req.Username,
		req.Exp,
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

// SSH implementations
func (uc *accountUsecase) createSSHAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-ssh-user script
	scriptArgs := []string{
		req.Username,
		req.Password,
		req.Exp,
		req.IPQuota,
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

func (uc *accountUsecase) renewSSHAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-ssh script
	scriptArgs := []string{
		req.Username,
		req.Exp,
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

// TROJAN implementations
func (uc *accountUsecase) createTrojanAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-trojan-user script
	scriptArgs := []string{
		req.Username,
		req.Password,
		req.Exp,
		req.IPQuota,
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

func (uc *accountUsecase) renewTrojanAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-trojan script
	scriptArgs := []string{
		req.Username,
		req.Exp,
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

// VLESS implementations
func (uc *accountUsecase) createVlessAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-vless-user script
	scriptArgs := []string{
		req.Username,
		req.Exp,
		req.Quota,
		req.IPQuota,
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

func (uc *accountUsecase) renewVlessAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
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

// SHADOWSOCKS implementations
func (uc *accountUsecase) createShadowsocksAccount(req models.CreateAccountRequest) (*models.AccountResponse, error) {
	// Execute the add-shadowsocks-user script
	scriptArgs := []string{
		req.Username,
		req.Password,
		req.Exp,
		req.IPQuota,
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

func (uc *accountUsecase) renewShadowsocksAccount(req models.RenewAccountRequest) (*models.AccountResponse, error) {
	// Execute the renew-shadowsocks script
	scriptArgs := []string{
		req.Username,
		req.Exp,
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
