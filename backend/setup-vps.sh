#!/bin/bash

# VPS Setup Script for VPN Account Management API

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Setting up VPN Account Management API on Ubuntu VPS...${NC}"

# Update system
echo -e "${YELLOW}Updating system packages...${NC}"
sudo apt update && sudo apt upgrade -y

# Install Go if not already installed
if ! command -v go &> /dev/null
then
    echo -e "${YELLOW}Installing Go...${NC}"
    sudo apt install -y golang
else
    echo -e "${GREEN}Go is already installed${NC}"
fi

# Create directory structure
echo -e "${YELLOW}Creating directory structure...${NC}"
sudo mkdir -p /root/scAUTO/backend

# Copy service file to systemd
echo -e "${YELLOW}Setting up systemd service...${NC}"
sudo cp vpn-api.service /etc/systemd/system/

# Reload systemd daemon
echo -e "${YELLOW}Reloading systemd daemon...${NC}"
sudo systemctl daemon-reload

# Enable the service
echo -e "${YELLOW}Enabling VPN API service...${NC}"
sudo systemctl enable vpn-api

echo -e "${GREEN}VPS setup completed!${NC}"
echo -e "${YELLOW}Next steps:${NC}"
echo -e "${GREEN}1. Copy your vpn-api binary and .env file to /root/scAUTO/backend/${NC}"
echo -e "${GREEN}2. Start the service: sudo systemctl start vpn-api${NC}"
echo -e "${GREEN}3. Check status: sudo systemctl status vpn-api${NC}"