#!/bin/bash

# Deployment script for VPN Account Management API

# Exit on any error
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting deployment of VPN Account Management API...${NC}"

# Build the application for Linux
echo -e "${YELLOW}Building application for Linux...${NC}"
GOOS=linux GOARCH=amd64 go build -o vpn-api main.go

# Check if build was successful
if [ $? -ne 0 ]; then
    echo -e "${RED}Build failed!${NC}"
    exit 1
fi

echo -e "${GREEN}Build completed successfully!${NC}"

# Copy files to VPS (you'll need to replace with your VPS details)
echo -e "${YELLOW}Copying files to VPS...${NC}"
# scp vpn-api root@your-vps-ip:/root/scAUTO/backend/
# scp .env root@your-vps-ip:/root/scAUTO/backend/
# scp vpn-api.service root@your-vps-ip:/etc/systemd/system/

echo -e "${YELLOW}Files copied to VPS. Now run these commands on your VPS:${NC}"
echo -e "${GREEN}1. sudo systemctl daemon-reload${NC}"
echo -e "${GREEN}2. sudo systemctl enable vpn-api${NC}"
echo -e "${GREEN}3. sudo systemctl start vpn-api${NC}"
echo -e "${GREEN}4. sudo systemctl status vpn-api${NC}"

echo -e "${YELLOW}To check logs:${NC}"
echo -e "${GREEN}sudo journalctl -u vpn-api -f${NC}"

echo -e "${GREEN}Deployment script completed!${NC}"