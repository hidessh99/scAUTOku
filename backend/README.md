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

### Account Management
- `POST /api/v1/accounts` - Create a new account
- `POST /api/v1/accounts/check` - Check an existing account
- `POST /api/v1/accounts/delete` - Delete an existing account

### Protocol Specific Routes
- `POST /api/v1/accounts/vmess/` - Create VMESS account
- `POST /api/v1/accounts/vmess/check` - Check VMESS account
- `POST /api/v1/accounts/vmess/delete` - Delete VMESS account

Similar routes exist for:
- SSH: `/api/v1/accounts/ssh/`
- TROJAN: `/api/v1/accounts/trojan/`
- VLESS: `/api/v1/accounts/vless/`
- SHADOWSOCKS: `/api/v1/accounts/shadowsocks/`

## Request Examples

### Create Account
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

## Dependencies

- [GoFiber](https://gofiber.io/) - Express-inspired web framework for Go
- [godotenv](https://github.com/joho/godotenv) - Loading environment variables from .env files