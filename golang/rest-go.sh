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


echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green}   INSTALLASI golang restapi HIdeSSH       ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"

cd
wget -q -O /usr/local/bin/vpn-api "https://github.com/hidessh99/scAUTOku/raw/refs/heads/main/golang/vpn-rest" && chmod +x /usr/local/bin/vpn-api

clear
read -p "Port example 3005 : " port
uuid=$(cat /proc/sys/kernel/random/uuid)

# Add nginx HTTPS configuration for the API port
echo "Adding nginx HTTPS configuration for port ${port}..."

# Backup existing nginx xray.conf
cp /etc/nginx/conf.d/xray.conf /etc/nginx/conf.d/xray.conf.bak 2>/dev/null || true

# Add server block for API HTTPS access
cat >>/etc/nginx/conf.d/xray.conf <<NGINX_EOF

# VPN API HTTPS Configuration
server {
    listen ${port} ssl http2 reuseport;
    ssl_certificate /etc/xray/xray.crt;
    ssl_certificate_key /etc/xray/xray.key;
    ssl_ciphers EECDH+CHACHA20:EECDH+CHACHA20-draft:EECDH+ECDSA+AES128:EECDH+aRSA+AES128:RSA+AES128:EECDH+ECDSA+AES256:EECDH+aRSA+AES256:RSA+AES256:EECDH+ECDSA+3DES:EECDH+aRSA+3DES:RSA+3DES:!MD5;
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
    
    location / {
        proxy_pass http://127.0.0.1:${port};
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
        proxy_set_header X-Forwarded-Host \$host;
        proxy_set_header X-Forwarded-Port \$server_port;
        
        # API specific headers
        proxy_set_header Authorization \$http_authorization;
        proxy_set_header Content-Type \$content_type;
        
        proxy_buffering off;
        proxy_request_buffering off;
        proxy_read_timeout 300s;
        proxy_send_timeout 300s;
    }
}
NGINX_EOF

cat >/etc/systemd/system/vpn-api.service <<EOF
[Unit]
Description=VPN Account Management API
After=network.target

[Service]
Type=simple
User=root
ExecStart=/usr/local/bin/vpn-api
Restart=always
RestartSec=10
Environment=PORT=${port}
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

# Test nginx configuration and restart nginx
echo "Testing nginx configuration..."
nginx -t
if [ $? -eq 0 ]; then
    echo "Nginx configuration is valid. Restarting nginx..."
    systemctl restart nginx
    systemctl reload nginx
    echo "Nginx restarted successfully."
else
    echo "Nginx configuration error! Please check the configuration."
    echo "Restoring backup configuration..."
    cp /etc/nginx/conf.d/xray.conf.bak /etc/nginx/conf.d/xray.conf 2>/dev/null || true
    systemctl restart nginx
fi

echo "Installation completed"
echo "You can access key header Bearer: ${uuid}"
echo "Port Rest Api: ${port}"
echo "HTTPS API URL: https://your-domain:${port}/"