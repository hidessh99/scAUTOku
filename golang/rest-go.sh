#!/bin/bash
red='\e[1;31m'
green='\e[0;32m'
purple='\e[0;35m'
orange='\e[0;33m'
NC='\e[0m'
clear

echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green}   INSTALLASI rest-go HIdeSSH       ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"
cd

clear

uuid=$(cat /proc/sys/kernel/random/uuid)



cat >/etc/systemd/system/vpn-api.service <<EOF
[Unit]
Description=VPN Account Management API
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/backend
ExecStart=/root/backend/vpn-api
Restart=always
RestartSec=10
Environment=PORT=3005
Environment=API_KEY=${uuid}
Environment=AllowOrigins=*

[Install]
WantedBy=multi-user.target
EOF

cd
chmod 644 /etc/systemd/system/vpn-api.service

sudo systemctl daemon-reload
sudo systemctl start vpn-api
sudo systemctl enable vpn-api

echo "Installation completed"
echo "You can access key header Bearer: ${uuid}"