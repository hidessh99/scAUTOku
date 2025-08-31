#!/bin/bash

# Go Installation Script for Ubuntu/Debian VPS
# This script installs the latest version of Go

set -e  # Exit on any error

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print status messages
print_status() {
    echo -e "${BLUE}[*]${NC} $1"
}

# Function to print success messages
print_success() {
    echo -e "${GREEN}[+]${NC} $1"
}

# Function to print warning messages
print_warning() {
    echo -e "${YELLOW}[!]${NC} $1"
}

# Function to print error messages
print_error() {
    echo -e "${RED}[-]${NC} $1"
}

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to get latest Go version
get_latest_go_version() {
    local version=$(curl -s https://go.dev/VERSION?m=text | head -n 1)
    if [ -z "$version" ]; then
        print_error "Failed to fetch latest Go version"
        exit 1
    fi
    echo "$version"
}

# Check if running as root
if [ "$EUID" -eq 0 ]; then
    print_warning "Running as root. It's recommended to run this script as a regular user."
    print_warning "The script will use sudo when necessary."
    echo
fi

# Check OS
print_status "Checking OS..."
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$NAME
    VER=$VERSION_ID
else
    print_error "Cannot detect OS. This script is designed for Ubuntu/Debian."
    exit 1
fi

print_status "Detected OS: $OS $VER"

# Check if OS is Ubuntu or Debian
if [[ "$OS" != *"Ubuntu"* && "$OS" != *"Debian"* ]]; then
    print_warning "This script is designed for Ubuntu/Debian. You are running: $OS"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Update package list
print_status "Updating package list..."
sudo apt update

# Install required packages
print_status "Installing required packages..."
sudo apt install -y curl wget tar gzip

# Get latest Go version
print_status "Getting latest Go version..."
GO_VERSION=$(get_latest_go_version)
print_success "Latest Go version: $GO_VERSION"

# Set architecture
ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    GO_ARCH="amd64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
    GO_ARCH="arm64"
elif [ "$ARCH" = "armv6l" ]; then
    GO_ARCH="armv6l"
else
    print_error "Unsupported architecture: $ARCH"
    exit 1
fi

print_status "Detected architecture: $ARCH ($GO_ARCH)"

# Download URL
DOWNLOAD_URL="https://go.dev/dl/${GO_VERSION}.linux-${GO_ARCH}.tar.gz"
print_status "Download URL: $DOWNLOAD_URL"

# Check if Go is already installed
if command_exists go; then
    INSTALLED_VERSION=$(go version | awk '{print $3}')
    print_warning "Go is already installed: $INSTALLED_VERSION"
    read -p "Do you want to reinstall/upgrade? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_status "Exiting..."
        exit 0
    fi
fi

# Download Go
print_status "Downloading Go $GO_VERSION..."
cd /tmp
wget -O go.tar.gz "$DOWNLOAD_URL"

# Verify download
if [ ! -f go.tar.gz ]; then
    print_error "Failed to download Go"
    exit 1
fi

print_success "Download completed"

# Remove existing Go installation
print_status "Removing existing Go installation..."
sudo rm -rf /usr/local/go

# Extract Go
print_status "Extracting Go..."
sudo tar -C /usr/local -xzf go.tar.gz

# Clean up
rm go.tar.gz

# Set up environment variables
print_status "Setting up environment variables..."

# Add to profile files
GO_PROFILE_LINES="export PATH=\$PATH:/usr/local/go/bin
export GOPATH=\$HOME/go
export PATH=\$PATH:\$GOPATH/bin"

# Add to .profile if not already present
if ! grep -q "export PATH=\$PATH:/usr/local/go/bin" ~/.profile 2>/dev/null; then
    echo "" >> ~/.profile
    echo "# Go settings" >> ~/.profile
    echo "$GO_PROFILE_LINES" >> ~/.profile
    print_success "Added Go environment variables to ~/.profile"
fi

# Add to .bashrc if not already present
if ! grep -q "export PATH=\$PATH:/usr/local/go/bin" ~/.bashrc 2>/dev/null; then
    echo "" >> ~/.bashrc
    echo "# Go settings" >> ~/.bashrc
    echo "$GO_PROFILE_LINES" >> ~/.bashrc
    print_success "Added Go environment variables to ~/.bashrc"
fi

# Add to .zshrc if it exists and not already present
if [ -f ~/.zshrc ] && ! grep -q "export PATH=\$PATH:/usr/local/go/bin" ~/.zshrc 2>/dev/null; then
    echo "" >> ~/.zshrc
    echo "# Go settings" >> ~/.zshrc
    echo "$GO_PROFILE_LINES" >> ~/.zshrc
    print_success "Added Go environment variables to ~/.zshrc"
fi

# Reload environment variables
print_status "Reloading environment variables..."
source ~/.profile 2>/dev/null || true
source ~/.bashrc 2>/dev/null || true
source ~/.zshrc 2>/dev/null || true

# Verify installation
print_status "Verifying installation..."
if command_exists go; then
    INSTALLED_VERSION=$(go version)
    print_success "Go installed successfully: $INSTALLED_VERSION"
else
    # Try to add to PATH manually for this session
    export PATH=$PATH:/usr/local/go/bin
    if command_exists go; then
        INSTALLED_VERSION=$(go version)
        print_success "Go installed successfully: $INSTALLED_VERSION"
    else
        print_error "Failed to verify Go installation"
        exit 1
    fi
fi

# Create GOPATH directory
if [ ! -d "$HOME/go" ]; then
    print_status "Creating GOPATH directory..."
    mkdir -p "$HOME/go"
    print_success "Created GOPATH directory: $HOME/go"
fi

# Test Go installation
print_status "Testing Go installation..."
cd /tmp
GO_TEST_FILE="hello.go"
cat > "$GO_TEST_FILE" << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
EOF

if go run "$GO_TEST_FILE" 2>/dev/null; then
    print_success "Go is working correctly!"
    rm "$GO_TEST_FILE"
else
    print_error "Go installation test failed"
    rm "$GO_TEST_FILE"
    exit 1
fi

print_success "Go installation completed successfully!"

echo
echo "To use Go in your current session, run:"
echo "  source ~/.profile"
echo "  source ~/.bashrc"
echo
echo "Or simply log out and log back in."
echo
echo "You can verify the installation by running:"
echo "  go version"
echo "  go env GOPATH"