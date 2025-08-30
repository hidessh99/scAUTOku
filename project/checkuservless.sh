#!/bin/bash
# Valid Script
ipsaya=$(curl -sS ipv4.icanhazip.com)
data_server=$(curl -v --insecure --silent https://google.com/ 2>&1 | grep Date | sed -e 's/< Date: //')
date_list=$(date +"%Y-%m-%d" -d "$data_server")

# =====================================================
# VLESS SPECIFIC USER CHECKER
# =====================================================

# Colors
green="\e[38;5;82m"
red="\e[38;5;196m"
yellow="\e[38;5;226m"
orange="\e[38;5;99m"
reset="\e[0m"
bold="\e[1m"

spinner=("⠋" "⠙" "⠹" "⠸" "⠼" "⠴" "⠦" "⠧" "⠇" "⠏")

convert_size() {
    local -i bytes=$1
    if [[ $bytes -lt 1024 ]]; then echo "${bytes}B"
    elif [[ $bytes -lt 1048576 ]]; then echo "$(((bytes + 1023)/1024))KB"
    elif [[ $bytes -lt 1073741824 ]]; then echo "$(((bytes + 1048575)/1048576))MB"
    else echo "$(((bytes + 1073741823)/1073741824))GB"
    fi
}

get_login_count() {
    local user="$1"
    local log_file="$2"
    local since=$(date -d "-1 minutes" '+%Y/%m/%d %H:%M:%S')
    awk -v u="$user" -v t="$since" '$0 ~ u && $0 > t && $0 !~ "127.0.0.1" {
        for(i=1;i<=NF;i++) if($i ~ /^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+:/) {split($i,a,":"); print a[1]}
    }' "$log_file" | sort -u | wc -l
}

# Prompt for username
read -p "User: " -e user

# Check if user exists
if ! grep -q "^### $user " "/etc/xray/vless/.vless.db"; then
  echo -e "${red}User '$user' not found!${reset}"
  exit 1
fi

LOG_FILE="/var/log/xray/access.log"

# Spinner loading
echo -ne "${yellow}Checking User "
for ((i=0;i<20;i++)); do
  echo -ne "${spinner[i % ${#spinner[@]}]}"
  sleep 0.05
  echo -ne "\b"
done
echo -e "${reset}\n"

# Table header
echo -e "${orange}──────────────────────────────────────────${reset}"
echo -e "${green}${bold}       VLESS USER STATUS${reset}"
echo -e "${orange}──────────────────────────────────────────${reset}"
echo -e "Username     | Usage  | Quota  | Log | Lim | Status"
echo -e "──────────────────────────────────────────"

ip_count=$(get_login_count "$user" "$LOG_FILE")
usage=$(cat /etc/xray/vless/usage/${user} 2>/dev/null || echo 0)
quota=$(cat /etc/xray/vless/${user} 2>/dev/null || echo 0)
ip_limit=$(cat /etc/xray/vless/${user}IP 2>/dev/null || echo "0")
ip_limit=$(echo "$ip_limit" | tr -d '[:space:]')
[[ -z "$ip_limit" || "$ip_limit" == "0" ]] && ip_limit="Unlimited"

readable_usage=$(convert_size $usage)
readable_quota=$(convert_size $quota)

if [[ "$ip_limit" == "Unlimited" ]]; then
  status="${green}OK${reset}"
else
  ip_limit_num=$(echo "$ip_limit" | tr -dc '0-9')
  if [[ "$ip_count" -gt "$ip_limit_num" ]]; then
    status="${red}OVER${reset}"
  else
    status="${green}OK${reset}"
  fi
fi

printf "%-12s | %-6s | %-6s | %-3s | %-3s | %b\n" \
  "$user" "$readable_usage" "$readable_quota" "$ip_count" "$ip_limit" "$status"

echo -e "${orange}──────────────────────────────────────────${reset}"

if [[ "$ip_count" -gt 0 ]]; then
  echo -e "${green}User is currently online${reset}"
else
  echo -e "${yellow}User is currently offline${reset}"
fi

# Display account information
clear 
echo -e "——————————————————————————————————————"
echo -e  "    Check Xray/Vless Account    "
echo -e "——————————————————————————————————————"
echo -e "User       : ${user}"
echo -e "Status     : ${status}"
echo -e "IP Connect : ${ip_count}"
echo -e "Usage      : ${readable_usage}"