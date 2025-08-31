package models

type AccountType string

const (
	VMESS       AccountType = "vmess"
	SSH         AccountType = "ssh"
	TROJAN      AccountType = "trojan"
	VLESS       AccountType = "vless"
	SHADOWSOCKS AccountType = "shadowsocks"
)

type CreateAccountRequest struct {
	Username    string      `json:"username" validate:"required" xml:"username" form:"username"`
	Password    string      `json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
	Exp         string      `json:"exp" validate:"required" xml:"exp" form:"exp"`
	Quota       string      `json:"quota,omitempty" xml:"quota,omitempty" form:"quota,omitempty"`
	IPQuota     string      `json:"ip_quota,omitempty" xml:"ip_quota,omitempty" form:"ip_quota,omitempty"`
	AccountType AccountType `json:"account_type" xml:"account_type" form:"account_type"`
}

type AccountResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type VmessAccountData struct {
	Username        string `json:"username"`
	Domain          string `json:"domain"`
	Quota           string `json:"quota"`
	IPQuota         string `json:"ip_limit"`
	Expired         string `json:"expired"`
	UUID            string `json:"uuid"`
	Pubkey          string `json:"pubkey"`
	VmessTLSLink    string `json:"vmess_tls_link"`
	VmessNonTLSLink string `json:"vmess_nontls_link"`
	VmessGRPCLink   string `json:"vmess_grpc_link"`
}

type SSHAccountData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
	Expired  string `json:"expired"`
	IPQuota  string `json:"ip_limit"`
	Pubkey   string `json:"pubkey"`
}

type TrojanAccountData struct {
	Username      string `json:"username"`
	Domain        string `json:"domain"`
	Expired       string `json:"expired"`
	IPQuota       string `json:"ip_limit"`
	Password      string `json:"password"`
	TrojanTLSLink string `json:"trojan_tls_link"`
	TrojanGRPC    string `json:"trojan_grpc_link"`
}

type VlessAccountData struct {
	Username      string `json:"username"`
	Domain        string `json:"domain"`
	Expired       string `json:"expired"`
	IPQuota       string `json:"ip_limit"`
	UUID          string `json:"uuid"`
	VlessTLSLink  string `json:"vless_tls_link"`
	VlessGRPCLink string `json:"vless_grpc_link"`
}

type ShadowsocksAccountData struct {
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Expired  string `json:"expired"`
	IPQuota  string `json:"ip_limit"`
	Password string `json:"password"`
	Method   string `json:"method"`
	SSLink   string `json:"ss_link"`
}

type CheckAccountRequest struct {
	Username    string      `json:"username" validate:"required" xml:"username" form:"username"`
	AccountType AccountType `json:"account_type" xml:"account_type" form:"account_type"`
}

type DeleteAccountRequest struct {
	Username    string      `json:"username" validate:"required" xml:"username" form:"username"`
	AccountType AccountType `json:"account_type"  xml:"account_type" form:"account_type"`
}

// Renewal request
type RenewAccountRequest struct {
	Username    string      `json:"username" validate:"required" xml:"username" form:"username"`
	AccountType AccountType `json:"account_type"  xml:"account_type" form:"account_type"`
	Exp         string      `json:"exp" validate:"required" xml:"exp" form:"exp"`
	Quota       string      `json:"quota" validate:"required" xml:"quota" form:"quota"`
	IPQuota     string      `json:"ip_quota" validate:"required" xml:"ip_quota" form:"ip_quota"`
}
