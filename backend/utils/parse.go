package utils

import (
	"encoding/base64"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

type ParseJsonSSH struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Remarks        string `json:"remarks"`
	Domain         string `json:"domain"`
	Host           string `json:"host"`
	OpenSSH        string `json:"openssh_port"`
	Dropbear       string `json:"dropbear_ports"`
	SSLTLS         string `json:"ssl_tls_ports"`
	PortSuid       string `json:"port_suid"`
	WebsocketHTTP  string `json:"websocket_http"`
	WebsocketHTTPS string `json:"websocket_https"`
	Badvpn         string `json:"badvpn_port_range"`
	MasaAktif      string `json:"masa_aktif"`
	Payload        string `json:"payload"`
	HostSSHSetting string `json:"host_ssh_setting"`
}

func OutputParseSsh(raw string) ParseJsonSSH {
	info := ParseJsonSSH{}
	lines := strings.Split(raw, "\n")

	for i, line := range lines {
		// Menangani header khusus yang nilainya ada di baris berikutnya
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "PAYLOAD" && i+1 < len(lines) {
			info.Payload = strings.TrimSpace(lines[i+1])
			continue
		}
		if trimmedLine == "SETING HOST SSH" && i+1 < len(lines) {
			info.HostSSHSetting = strings.TrimSpace(lines[i+1])
			continue
		}

		// Menangani baris dengan format "Key : Value"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "Username":
				info.Username = value
			case "Password":
				info.Password = value
			case "Remarks":
				info.Remarks = value
			case "Domain":
				// Mengambil domain dari baris pertama dan kedua jika ada
				if info.Domain == "" {
					info.Domain = value
				}
			case "Host":
				info.Host = value
			case "OpenSSH":
				info.OpenSSH = value
			case "Dropbear":
				info.Dropbear = value
			case "SSL/TLS":
				info.SSLTLS = value
			case "Port Suid":
				info.PortSuid = value
			case "Websocket HTTP":
				info.WebsocketHTTP = value
			case "Websocket HTTPS":
				info.WebsocketHTTPS = value
			case "badvpn":
				info.Badvpn = value
			case "Masa Aktif":
				info.MasaAktif = value
			}
		}
	}

	return info
}

// ParseJsonShadowsocks represents the structure of parsed Shadowsocks account data
type ParseJsonShadowsocks struct {
	Username       string                 `json:"username"`
	Description    string                 `json:"description"`
	ServerHost     string                 `json:"server_host"`
	Location       string                 `json:"location"`
	Port           string                 `json:"port"`
	Method         string                 `json:"method"`
	Password       string                 `json:"password"`
	WSPath         string                 `json:"ws_path"`
	ServiceName    string                 `json:"service_name"`
	PublicKey      string                 `json:"public_key"`
	SSWSLink       string                 `json:"ss_ws_link"`
	SSGRPCLink     string                 `json:"ss_grpc_link"`
	OpenClashURL   string                 `json:"openclash_url"`
	ExpiresOn      string                 `json:"expires_on"`
	AccountDetails map[string]interface{} `json:"account_details"`
}

// OutputParseShadowsocks parses the Shadowsocks response string and extracts all relevant data
func OutputParseShadowsocks(raw string) ParseJsonShadowsocks {
	info := ParseJsonShadowsocks{
		AccountDetails: make(map[string]interface{}),
	}

	lines := strings.Split(raw, "\n")

	// Parse basic information from key-value pairs
	for _, line := range lines {
		// Skip separator lines
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "————") || trimmedLine == "" {
			continue
		}

		// Handle special case for Username at the beginning
		if strings.HasPrefix(trimmedLine, "Username :") {
			parts := strings.SplitN(trimmedLine, ":", 2)
			if len(parts) == 2 {
				info.Username = strings.TrimSpace(parts[1])
			}
			continue
		}

		// Parse lines with format "Key : Value"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "Description":
				info.Description = value
			case "Server Host":
				info.ServerHost = value
			case "Location":
				info.Location = value
			case "Port":
				info.Port = value
			case "Method":
				info.Method = value
			case "Password":
				info.Password = value
			case "WS Path":
				info.WSPath = value
			case "ServiceName":
				info.ServiceName = value
			case "Public Key":
				info.PublicKey = value
			case "SS WS Link":
				info.SSWSLink = value
			case "SS gRPC Link":
				info.SSGRPCLink = value
			case "OpenClash Format":
				info.OpenClashURL = value
			case "Expires On":
				info.ExpiresOn = value
			}
		}
	}

	// Parse Shadowsocks links to extract account details
	parseSSLink := func(link, linkType string) map[string]interface{} {
		if link == "" || !strings.HasPrefix(link, "ss://") {
			return nil
		}

		// Parse the shadowsocks URL
		u, err := url.Parse(link)
		if err != nil {
			return nil
		}

		// Extract the base64 encoded user info
		encodedUserInfo := u.User.Username()
		decodedUserInfo, err := base64.StdEncoding.DecodeString(encodedUserInfo)
		if err != nil {
			// If base64 decoding fails, try to parse as method:password format
			userInfoParts := strings.Split(encodedUserInfo, ":")
			if len(userInfoParts) == 2 {
				decodedUserInfo = []byte(userInfoParts[0] + ":" + userInfoParts[1])
			}
		}

		// Parse method and password from decoded user info
		userInfoStr := string(decodedUserInfo)
		methodPassword := strings.SplitN(userInfoStr, ":", 2)
		var method, password string
		if len(methodPassword) == 2 {
			method = methodPassword[0]
			password = methodPassword[1]
		}

		// Extract query parameters
		queryParams := u.Query()

		// Extract plugin parameters
		pluginParams := strings.Split(queryParams.Get("plugin"), ";")
		pluginName := ""
		if len(pluginParams) > 0 {
			pluginName = pluginParams[0]
		}

		// Create account data map
		accountData := map[string]interface{}{
			"method":        method,
			"password":      password,
			"host":          u.Hostname(),
			"port":          u.Port(),
			"path":          queryParams.Get("path"),
			"serviceName":   queryParams.Get("serviceName"),
			"host_param":    queryParams.Get("host"),
			"network":       queryParams.Get("network"),
			"plugin":        pluginName,
			"plugin_params": pluginParams,
		}

		// Add fragment (the part after #)
		if u.Fragment != "" {
			accountData["fragment"] = u.Fragment
		}

		return accountData
	}

	// Parse each Shadowsocks link and add to account details
	if wsData := parseSSLink(info.SSWSLink, "ws"); wsData != nil {
		info.AccountDetails["ws"] = wsData
	}

	if grpcData := parseSSLink(info.SSGRPCLink, "grpc"); grpcData != nil {
		info.AccountDetails["grpc"] = grpcData
	}

	return info
}

// ParseJsonTrojan represents the structure of parsed Trojan account data
type ParseJsonTrojan struct {
	Description    string                 `json:"description"`
	HostServer     string                 `json:"host_server"`
	Location       string                 `json:"location"`
	TLSPort        string                 `json:"tls_port"`
	DNSPort        string                 `json:"dns_port"`
	GRPCPort       string                 `json:"grpc_port"`
	Security       string                 `json:"security"`
	Network        string                 `json:"network"`
	Path           string                 `json:"path"`
	ServiceName    string                 `json:"service_name"`
	Password       string                 `json:"password"`
	PublicKey      string                 `json:"public_key"`
	TLSLink        string                 `json:"tls_link"`
	GRPCLink       string                 `json:"grpc_link"`
	OpenClashURL   string                 `json:"openclash_url"`
	ExpiresOn      string                 `json:"expires_on"`
	AccountDetails map[string]interface{} `json:"account_details"`
}

// OutputParseTrojan parses the Trojan response string and extracts all relevant data
func OutputParseTrojan(raw string) ParseJsonTrojan {
	info := ParseJsonTrojan{
		AccountDetails: make(map[string]interface{}),
	}

	lines := strings.Split(raw, "\n")

	// Parse basic information from key-value pairs
	for _, line := range lines {
		// Skip separator lines
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "————") || trimmedLine == "" {
			continue
		}

		// Parse lines with format "Key : Value"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "Description":
				info.Description = value
			case "Host Server":
				info.HostServer = value
			case "Location":
				info.Location = value
			case "TLS Port":
				info.TLSPort = value
			case "DNS Port":
				info.DNSPort = value
			case "GRPC Port":
				info.GRPCPort = value
			case "Security":
				info.Security = value
			case "Network":
				info.Network = value
			case "Path":
				info.Path = value
			case "ServiceName":
				info.ServiceName = value
			case "Password":
				info.Password = value
			case "Public Key":
				info.PublicKey = value
			case "TLS Link":
				info.TLSLink = value
			case "GRPC Link":
				info.GRPCLink = value
			case "OpenClash Format":
				info.OpenClashURL = value
			case "Expires On":
				info.ExpiresOn = value
			}
		}
	}

	// Parse Trojan links to extract account details
	parseTrojanLink := func(link, linkType string) map[string]interface{} {
		if link == "" || !strings.HasPrefix(link, "trojan://") {
			return nil
		}

		// Parse the trojan URL
		u, err := url.Parse(link)
		if err != nil {
			return nil
		}

		// Extract query parameters
		queryParams := u.Query()

		// Create account data map
		accountData := map[string]interface{}{
			"password":    u.User.Username(),
			"host":        u.Hostname(),
			"port":        u.Port(),
			"path":        queryParams.Get("path"),
			"type":        queryParams.Get("type"),
			"security":    queryParams.Get("security"),
			"host_param":  queryParams.Get("host"),
			"serviceName": queryParams.Get("serviceName"),
			"sni":         queryParams.Get("sni"),
			"mode":        queryParams.Get("mode"),
		}

		// Add fragment (the part after #)
		if u.Fragment != "" {
			accountData["fragment"] = u.Fragment
		}

		return accountData
	}

	// Parse each Trojan link and add to account details
	if tlsData := parseTrojanLink(info.TLSLink, "tls"); tlsData != nil {
		info.AccountDetails["tls"] = tlsData
	}

	if grpcData := parseTrojanLink(info.GRPCLink, "grpc"); grpcData != nil {
		info.AccountDetails["grpc"] = grpcData
	}

	return info
}

// ParseJsonVLess represents the structure of parsed VLess account data
type ParseJsonVLess struct {
	Description    string                 `json:"description"`
	HostServer     string                 `json:"host_server"`
	Location       string                 `json:"location"`
	PortTLS        string                 `json:"port_tls"`
	PortNonTLS     string                 `json:"port_non_tls"`
	PortDNS        string                 `json:"port_dns"`
	PortGRPC       string                 `json:"port_grpc"`
	Security       string                 `json:"security"`
	Network        string                 `json:"network"`
	Path           string                 `json:"path"`
	ServiceName    string                 `json:"service_name"`
	UserID         string                 `json:"user_id"`
	PublicKey      string                 `json:"public_key"`
	TLSLink        string                 `json:"tls_link"`
	NTLSLink       string                 `json:"ntls_link"`
	GRPCLink       string                 `json:"grpc_link"`
	OpenClashURL   string                 `json:"openclash_url"`
	ExpiresOn      string                 `json:"expires_on"`
	AccountDetails map[string]interface{} `json:"account_details"`
}

// OutputParseVLess parses the VLess response string and extracts all relevant data
func OutputParseVLess(raw string) ParseJsonVLess {
	info := ParseJsonVLess{
		AccountDetails: make(map[string]interface{}),
	}

	lines := strings.Split(raw, "\n")

	// Parse basic information from key-value pairs
	for _, line := range lines {
		// Skip separator lines
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "————") || trimmedLine == "" {
			continue
		}

		// Parse lines with format "Key : Value"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "Description":
				info.Description = value
			case "Host Server":
				info.HostServer = value
			case "Location":
				info.Location = value
			case "Port TLS":
				info.PortTLS = value
			case "Port non TLS":
				info.PortNonTLS = value
			case "Port DNS":
				info.PortDNS = value
			case "Port GRPC":
				info.PortGRPC = value
			case "Security":
				info.Security = value
			case "Network":
				info.Network = value
			case "Path":
				info.Path = value
			case "ServiceName":
				info.ServiceName = value
			case "User ID":
				info.UserID = value
			case "Public Key":
				info.PublicKey = value
			case "TLS Link":
				info.TLSLink = value
			case "NTLS Link":
				info.NTLSLink = value
			case "GRPC Link":
				info.GRPCLink = value
			case "OpenClash Format":
				info.OpenClashURL = value
			case "Expires On":
				info.ExpiresOn = value
			}
		}
	}

	// Parse VLess links to extract account details
	parseVLessLink := func(link, linkType string) map[string]interface{} {
		if link == "" || !strings.HasPrefix(link, "vless://") {
			return nil
		}

		// Parse the vless URL
		u, err := url.Parse(link)
		if err != nil {
			return nil
		}

		// Extract query parameters
		queryParams := u.Query()

		// Create account data map
		accountData := map[string]interface{}{
			"uuid":        u.User.Username(),
			"host":        u.Hostname(),
			"port":        u.Port(),
			"path":        queryParams.Get("path"),
			"type":        queryParams.Get("type"),
			"security":    queryParams.Get("security"),
			"encryption":  queryParams.Get("encryption"),
			"host_param":  queryParams.Get("host"),
			"serviceName": queryParams.Get("serviceName"),
			"sni":         queryParams.Get("sni"),
			"mode":        queryParams.Get("mode"),
		}

		// Add fragment (the part after #)
		if u.Fragment != "" {
			accountData["fragment"] = u.Fragment
		}

		return accountData
	}

	// Parse each VLess link and add to account details
	if tlsData := parseVLessLink(info.TLSLink, "tls"); tlsData != nil {
		info.AccountDetails["tls"] = tlsData
	}

	if ntlsData := parseVLessLink(info.NTLSLink, "ntls"); ntlsData != nil {
		info.AccountDetails["ntls"] = ntlsData
	}

	if grpcData := parseVLessLink(info.GRPCLink, "grpc"); grpcData != nil {
		info.AccountDetails["grpc"] = grpcData
	}

	return info
}

// ParseJsonVMess represents the structure of parsed VMess account data
type ParseJsonVMess struct {
	Remarks        string                 `json:"remarks"`
	HostServer     string                 `json:"host_server"`
	Location       string                 `json:"location"`
	PortTLS        string                 `json:"port_tls"`
	PortNonTLS     string                 `json:"port_non_tls"`
	PortDNS        string                 `json:"port_dns"`
	PortGRPC       string                 `json:"port_grpc"`
	AlterId        string                 `json:"alter_id"`
	Security       string                 `json:"security"`
	Network        string                 `json:"network"`
	Path           string                 `json:"path"`
	ServiceName    string                 `json:"service_name"`
	UserID         string                 `json:"user_id"`
	PublicKey      string                 `json:"public_key"`
	TLSLink        string                 `json:"tls_link"`
	NTLSLink       string                 `json:"ntls_link"`
	GRPCLink       string                 `json:"grpc_link"`
	OpenClashURL   string                 `json:"openclash_url"`
	ExpiresOn      string                 `json:"expires_on"`
	AccountDetails map[string]interface{} `json:"account_details"`
}

// OutputParseVMess parses the VMess response string and extracts all relevant data
func OutputParseVMess(raw string) ParseJsonVMess {
	info := ParseJsonVMess{
		AccountDetails: make(map[string]interface{}),
	}

	lines := strings.Split(raw, "\n")

	// Parse basic information from key-value pairs
	for _, line := range lines {
		// Skip separator lines
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "————") || trimmedLine == "" {
			continue
		}

		// Handle special headers that might span multiple lines
		// For this implementation, we'll focus on the key-value pairs

		// Parse lines with format "Key : Value"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "Remarks":
				info.Remarks = value
			case "Host Server":
				info.HostServer = value
			case "Location":
				info.Location = value
			case "Port TLS":
				info.PortTLS = value
			case "Port non TLS":
				info.PortNonTLS = value
			case "Port DNS":
				info.PortDNS = value
			case "Port GRPC":
				info.PortGRPC = value
			case "AlterId":
				info.AlterId = value
			case "Security":
				info.Security = value
			case "Network":
				info.Network = value
			case "Path":
				info.Path = value
			case "ServiceName":
				info.ServiceName = value
			case "User ID":
				info.UserID = value
			case "Public Key":
				info.PublicKey = value
			case "TLS Link":
				info.TLSLink = value
			case "NTLS Link":
				info.NTLSLink = value
			case "GRPC Link":
				info.GRPCLink = value
			case "OpenClash Format":
				info.OpenClashURL = value
			case "Expires On":
				info.ExpiresOn = value
			}
		}
	}

	// Parse and decode VMess links to extract account details
	parseVMessLink := func(link, linkType string) map[string]interface{} {
		if link == "" || !strings.HasPrefix(link, "vmess://") {
			return nil
		}

		// Extract base64 part
		base64Data := strings.TrimPrefix(link, "vmess://")

		// Decode base64
		jsonData, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			return nil
		}

		// Parse JSON
		var accountData map[string]interface{}
		if err := json.Unmarshal(jsonData, &accountData); err != nil {
			return nil
		}

		return accountData
	}

	// Parse each VMess link and add to account details
	if tlsData := parseVMessLink(info.TLSLink, "tls"); tlsData != nil {
		info.AccountDetails["tls"] = tlsData
	}

	if ntlsData := parseVMessLink(info.NTLSLink, "ntls"); ntlsData != nil {
		info.AccountDetails["ntls"] = ntlsData
	}

	if grpcData := parseVMessLink(info.GRPCLink, "grpc"); grpcData != nil {
		info.AccountDetails["grpc"] = grpcData
	}

	return info
}

// VMessAccountInfo represents the structure of a VMess account
type VMessAccountInfo struct {
	Username  string `json:"username"`
	Usage     string `json:"usage"`
	Quota     string `json:"quota"`
	Log       string `json:"log"`
	Limit     string `json:"limit"`
	Status    string `json:"status"`
	User      string `json:"user"`
	IPConnect string `json:"ip_connect"`
	UsageMB   int    `json:"usage_mb,omitempty"`
	QuotaGB   int    `json:"quota_gb,omitempty"`
}

// ParseVMessAccountData parses the raw VMess output byte array and extracts account information
func ParseVMessAccountData(rawData []byte) (*VMessAccountInfo, error) {
	// Convert byte array to string
	rawString := string(rawData)

	// Remove ANSI escape sequences
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	cleanOutput := ansiRegex.ReplaceAllString(rawString, "")

	// Split into lines
	lines := strings.Split(cleanOutput, "\n")

	// Initialize account info
	accountInfo := &VMessAccountInfo{}

	// Parse the data
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines
		if trimmedLine == "" {
			continue
		}

		// Parse table data (the line with | separators)
		if strings.Contains(trimmedLine, "|") && !strings.Contains(trimmedLine, "————") {
			// This is the table data row
			fields := strings.Split(trimmedLine, "|")
			if len(fields) >= 6 {
				accountInfo.Username = strings.TrimSpace(fields[0])
				accountInfo.Usage = strings.TrimSpace(fields[1])
				accountInfo.Quota = strings.TrimSpace(fields[2])
				accountInfo.Log = strings.TrimSpace(fields[3])
				accountInfo.Limit = strings.TrimSpace(fields[4])
				accountInfo.Status = strings.TrimSpace(fields[5])

				// Try to convert usage and quota to numbers
				if usageNum, err := strconv.Atoi(strings.TrimSuffix(accountInfo.Usage, "MB")); err == nil {
					accountInfo.UsageMB = usageNum
				}
				if quotaNum, err := strconv.Atoi(strings.TrimSuffix(accountInfo.Quota, "GB")); err == nil {
					accountInfo.QuotaGB = quotaNum
				}
			}
		}

		// Parse key-value pairs
		if strings.Contains(trimmedLine, ":") {
			parts := strings.SplitN(trimmedLine, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				switch key {
				case "User":
					accountInfo.User = value
				case "Status":
					accountInfo.Status = value
				case "IP Connect":
					accountInfo.IPConnect = value
				case "Usage":
					accountInfo.Usage = value
					// Try to convert to number
					if usageNum, err := strconv.Atoi(strings.TrimSuffix(value, "MB")); err == nil {
						accountInfo.UsageMB = usageNum
					}
				case "Quota":
					accountInfo.Quota = value
					// Try to convert to number
					if quotaNum, err := strconv.Atoi(strings.TrimSuffix(value, "GB")); err == nil {
						accountInfo.QuotaGB = quotaNum
					}
				}
			}
		}
	}

	return accountInfo, nil
}

// ConvertVMessAccountToJSON converts VMess account data to JSON format
func CheckAccount(rawData []byte) ([]byte, error) {
	accountInfo, err := ParseVMessAccountData(rawData)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(accountInfo)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// CheckAccountInfo represents the structure of parsed check account data
type CheckAccountInfo struct {
	Username  string `json:"username"`
	Status    string `json:"status"`
	IPConnect string `json:"ip_connect"`
	Usage     string `json:"usage"`
}

// ParseCheckAccountOutput parses the check account response string and extracts all relevant data
func ParseCheckAccountOutput(raw string) CheckAccountInfo {
	info := CheckAccountInfo{}

	// Remove ANSI escape codes
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	cleanedRaw := ansiRegex.ReplaceAllString(raw, "")

	lines := strings.Split(cleanedRaw, "\n")

	// Parse basic information from key-value pairs
	for _, line := range lines {
		// Skip separator lines
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "————") || trimmedLine == "" {
			continue
		}

		// Parse lines with format "Key : Value"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "User":
				info.Username = value
			case "Username":
				info.Username = value
			case "Status":
				info.Status = value
			case "IP Connect":
				info.IPConnect = value
			case "Usage":
				info.Usage = value
			}
		}
	}

	return info
}
