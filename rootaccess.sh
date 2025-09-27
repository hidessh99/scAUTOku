#!/bin/bash

red='\e[1;31m'
green='\e[0;32m'
purple='\e[0;35m'
orange='\e[0;33m'
NC='\e[0m'
clear

echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green}           Root Access           ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"

sshd_conf_url="https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/fodder/examples/sshd"
banner_url="https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/fodder/examples/banner"
common_password_url="https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/fodder/examples/common-password"

echo "Memulai proses instalasi, mohon tunggu..."


echo "Proses instalasi selesai."


# Download SSHD config
if [ -n "$sshd_conf_url" ]; then
    [ -f /etc/ssh/sshd_config ] && rm /etc/ssh/sshd_config
    wget -q -O /etc/ssh/sshd_config "$sshd_conf_url" >/dev/null 2>&1 || echo -e "${red}Failed to download sshd_config${neutral}"
else
    echo -e "${yellow}sshd_conf_url is not set, skipping download of sshd_config${neutral}"
fi

# Download Banner
if [ -n "$banner_url" ]; then
    wget -q -O /etc/gerhanatunnel.txt "$banner_url" && chmod +x /etc/gerhanatunnel.txt >/dev/null 2>&1 || echo -e "${red}Failed to download gerhanatunnel.txt${neutral}"
else
    echo -e "${yellow}banner_url is not set, skipping download of gerhanatunnel.txt${neutral}"
fi

# Download Common Password
if [ -n "$common_password_url" ]; then
    [ -f /etc/pam.d/common-password ] && rm /etc/pam.d/common-password
    wget -O /etc/pam.d/common-password "$common_password_url" >/dev/null 2>&1 || echo -e "${red}Failed to download common-password${neutral}"
else
    echo -e "${yellow}common_password_url is not set, skipping download of common-password${neutral}"
fi


# Permission for common-password
if [ -f /etc/pam.d/common-password ]; then
    chmod +x /etc/pam.d/common-password || echo -e "${red}Failed to give execute permission to common-password${neutral}"
else
    echo -e "${yellow}/etc/pam.d/common-password not found, skipping permission change${neutral}"
fi


systemctl daemon-reload || echo -e "${red}Failed to reload systemd daemon${neutral}"
systemctl restart sshd || echo -e "${red}Failed to restart sshd${neutral}"
systemctl restart ssh || echo -e "${red}Failed to restart ssh${neutral}"
sleep 2
systemctl enable sshd || echo -e "${red}Failed to enable sshd${neutral}"
systemctl status sshd || echo -e "${red}Failed to check sshd status${neutral}"
