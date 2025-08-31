# API Tests

This directory contains HTTP test files for each protocol supported by the VPN Account Management API.

## Test Files

- [vmess.http](vmess.http) - Tests for VMESS account management
- [ssh.http](ssh.http) - Tests for SSH account management
- [vless.http](vless.http) - Tests for VLESS account management
- [trojan.http](trojan.http) - Tests for TROJAN account management
- [shadowsocks.http](shadowsocks.http) - Tests for SHADOWSOCKS account management

## How to Use

These test files can be used with VS Code REST Client extension or any other HTTP client that supports .http files.

1. Make sure the API server is running on `localhost:3000` (or update the `@baseUrl` variable)
2. Update the `@apiKey` variable with your actual API key
3. Run each request individually or use the "Send Request" links in VS Code

## Test Endpoints

Each test file includes tests for:
- Health check endpoint
- Protocol-specific endpoints:
  - Create account
  - Check account
  - Renew account
  - Delete account
- General account management endpoints (specifying account_type in the request body)

## Environment Variables

- `@baseUrl` - The base URL of the API server (default: http://localhost:3000)
- `@apiKey` - The API key for authentication (default: your_api_key_here)

## Example Usage

1. Start the API server:
   ```bash
   cd ../
   go run main.go
   ```

2. In VS Code, open any of the .http files
3. Click on the "Send Request" link above any HTTP request
4. View the response in the output panel

## Notes

- The tests use placeholder data that should be updated with real values
- Make sure to use unique usernames for create operations to avoid conflicts
- The server_id should correspond to an actual server in your infrastructure