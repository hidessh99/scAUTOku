# VPN Account Management API

A REST API built with GoFiber implementing Clean Architecture for managing VPN accounts (VMESS, SSH, TROJAN, VLESS, SHADOWSOCKS).

## Architecture

This API follows Clean Architecture principles with the following layers:

1. **Models**: Data structures and entities
2. **Controllers**: Handle HTTP requests and responses
3. **Usecases**: Business logic implementation
4. **Utils**: Utility functions and helpers
5. **Routes**: API route definitions
6. **Config**: Configuration management

The architecture has been separated by protocol:
- Each protocol (VMESS, SSH, VLESS, TROJAN, SHADOWSOCKS) has its own usecase and controller
- A main account usecase and controller coordinate between the protocol-specific components

## Authentication

The API requires authentication for all protected endpoints. You can authenticate using:

1. **API Key in Authorization Header**:
   ```
   Authorization: Bearer YOUR_API_KEY
   ```

2. **API Key in X-API-Key Header**:
   ```
   X-API-Key: YOUR_API_KEY
   ```

The API key can be configured in the `.env` file.

## API Endpoints

### Health Check
- `GET /health` - Check if the API is running (No authentication required)

### General Account Management
- `POST /api/v1/accounts` - Create a new account (specify account_type in request)
- `POST /api/v1/accounts/check` - Check an existing account (specify account_type in request)
- `POST /api/v1/accounts/delete` - Delete an existing account (specify account_type in request)
- `POST /api/v1/accounts/renew` - Renew an existing account (specify account_type in request)

### Protocol Specific Routes

#### VMESS
- `POST /api/v1/vmess/` - Create VMESS account
- `POST /api/v1/vmess/check` - Check VMESS account
- `POST /api/v1/vmess/delete` - Delete VMESS account
- `POST /api/v1/vmess/renew` - Renew VMESS account

#### SSH
- `POST /api/v1/ssh/` - Create SSH account
- `POST /api/v1/ssh/check` - Check SSH account
- `POST /api/v1/ssh/delete` - Delete SSH account
- `POST /api/v1/ssh/renew` - Renew SSH account

#### VLESS
- `POST /api/v1/vless/` - Create VLESS account
- `POST /api/v1/vless/check` - Check VLESS account
- `POST /api/v1/vless/delete` - Delete VLESS account
- `POST /api/v1/vless/renew` - Renew VLESS account

#### TROJAN
- `POST /api/v1/trojan/` - Create TROJAN account
- `POST /api/v1/trojan/check` - Check TROJAN account
- `POST /api/v1/trojan/delete` - Delete TROJAN account
- `POST /api/v1/trojan/renew` - Renew TROJAN account

#### SHADOWSOCKS
- `POST /api/v1/shadowsocks/` - Create SHADOWSOCKS account
- `POST /api/v1/shadowsocks/check` - Check SHADOWSOCKS account
- `POST /api/v1/shadowsocks/delete` - Delete SHADOWSOCKS account
- `POST /api/v1/shadowsocks/renew` - Renew SHADOWSOCKS account

## Request Examples

### Create Account (General)
```json
{
  "username": "testuser",
  "password": "testpass", // Required for SSH, TROJAN, SHADOWSOCKS
  "exp": "30", // Expiration in days
  "quota": "10GB", // Optional quota
  "ip_quota": "5", // IP limit
  "server_id": 1,
  "account_type": "vmess" // vmess, ssh, trojan, vless, shadowsocks
}
```

### Create VMESS Account (Protocol Specific)
```json
{
  "username": "testuser",
  "exp": "30",
  "quota": "10GB",
  "ip_quota": "5",
  "server_id": 1
}
```

### Check Account
```json
{
  "username": "testuser",
  "account_type": "vmess"
}
```

### Delete Account
```json
{
  "username": "testuser",
  "server_id": 1,
  "account_type": "vmess"
}
```

### Renew Account
```json
{
  "username": "testuser",
  "exp": "30", // Additional days to extend expiration
  "server_id": 1,
  "account_type": "vmess"
}
```

## Response Format

All responses follow this format:
```json
{
  "status": "success|error",
  "message": "Description of the operation result",
  "data": {} // Optional data field for successful operations
}
```

## Configuration

The API can be configured using environment variables or a `.env` file:

```
# Authentication Configuration
API_KEY=your_api_key_here
JWT_SECRET=your_jwt_secret_here
ADMIN_USERNAME=admin
ADMIN_PASSWORD=admin123

# Server Configuration
PORT=3000
HOST=localhost

# Database Configuration (if needed)
DB_HOST=localhost
DB_PORT=5432
DB_USER=vpnuser
DB_PASSWORD=vpnpassword
DB_NAME=vpnaccounts
```

## Installation

1. Make sure Go is installed (version 1.21 or higher)
2. Navigate to the backend directory: `cd backend`
3. Install dependencies: `go mod tidy`
4. Run the application: `go run main.go`

## Testing

The [tests](tests/) directory contains HTTP test files for each protocol:
- [VMESS tests](tests/vmess.http)
- [SSH tests](tests/ssh.http)
- [VLESS tests](tests/vless.http)
- [TROJAN tests](tests/trojan.http)
- [SHADOWSOCKS tests](tests/shadowsocks.http)
- [All protocols tests](tests/all_protocols.http)

These can be run using VS Code REST Client extension or any HTTP client that supports .http files.

## Deployment to Ubuntu VPS

1. Build the application for Linux:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o vpn-api main.go
   ```

2. Copy the binary and configuration files to your VPS:
   ```bash
   scp vpn-api root@your-vps-ip:/root/scAUTO/backend/
   scp .env root@your-vps-ip:/root/scAUTO/backend/
   scp vpn-api.service root@your-vps-ip:/etc/systemd/system/
   ```

3. On your VPS, reload the systemd daemon:
   ```bash
   sudo systemctl daemon-reload
   ```

4. Enable the service:
   ```bash
   sudo systemctl enable vpn-api
   ```

5. Start the service:
   ```bash
   sudo systemctl start vpn-api
   ```

6. Check the service status:
   ```bash
   sudo systemctl status vpn-api
   ```

7. To view logs:
   ```bash
   sudo journalctl -u vpn-api -f
   ```

## Command Paths

The application expects the following shell commands to be available on the system:

### Account Creation:
- VMESS: `/usr/local/bin/add-vmess-user`
- SSH: `/usr/local/bin/add-ssh-user`
- TROJAN: `/usr/local/bin/add-trojan-user`
- VLESS: `/usr/local/bin/add-vless-user`
- SHADOWSOCKS: `/usr/local/bin/add-shadowsocks-user`

### Account Deletion:
- VMESS: `/usr/local/bin/del-vmess`
- SSH: `/usr/local/bin/del-ssh`
- TROJAN: `/usr/local/bin/del-trojan`
- VLESS: `/usr/local/bin/del-vless`
- SHADOWSOCKS: `/usr/local/bin/del-addshadowsocks`

### Account Renewal:
- VMESS: `/usr/local/bin/renew-vmess`
- SSH: `/usr/local/bin/renew-ssh`
- TROJAN: `/usr/local/bin/renew-trojan`
- VLESS: `/usr/local/bin/renew-vless`
- SHADOWSOCKS: `/usr/local/bin/renew-shadowsocks`

Make sure these commands are properly installed and accessible on your system.

## Dependencies

- [GoFiber](https://gofiber.io/) - Express-inspired web framework for Go
- [godotenv](https://github.com/joho/godotenv) - Loading environment variables from .env files