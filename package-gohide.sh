#!/bin/bash
red='\e[1;31m'
green='\e[0;32m'
purple='\e[0;35m'
orange='\e[0;33m'
NC='\e[0m'
clear

echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green}   INSTALLASI Add package HIdeSSH       ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"


cd
wget -q -O /usr/local/bin/add-vmess "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/addvmess" && chmod +x /usr/local/bin/add-vmess
wget -q -O /usr/local/bin/add-vless "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/addvless" && chmod +x /usr/local/bin/add-vless
wget -q -O /usr/local/bin/add-trojan "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/addtrojan" && chmod +x /usr/local/bin/add-trojan
wget -q -O /usr/local/bin/add-shadowsocks "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/addshadowsocks" && chmod +x /usr/local/bin/add-shadowsocks
wget -q -O /usr/local/bin/add-ssh "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/addssh" && chmod +x /usr/local/bin/add-ssh
clear

echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green} INSTALLASI delete package HIdeSSH      ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"
cd
wget -q -O /usr/local/bin/del-vmess "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/dellaccvmess.sh" && chmod +x /usr/local/bin/del-vmess
wget -q -O /usr/local/bin/del-trojan "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/dellacctrojan.sh" && chmod +x /usr/local/bin/del-trojan
wget -q -O /usr/local/bin/del-vless "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/dellaccvless.sh" && chmod +x /usr/local/bin/del-vless
wget -q -O /usr/local/bin/del-shadowsocks "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/dellaccshadowsocks.sh" && chmod +x /usr/local/bin/del-shadowsocks
wget -q -O /usr/local/bin/del-ssh "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/dellaccssh.sh" && chmod +x /usr/local/bin/del-ssh
clear



echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green} INSTALLASI Check  package HIdeSSH      ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"
cd

wget -q -O /usr/local/bin/check-vless "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/checkuservless.sh" && chmod +x /usr/local/bin/check-vless
wget -q -O /usr/local/bin/check-trojan "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/checkusertrojan.sh" && chmod +x /usr/local/bin/check-trojan
wget -q -O /usr/local/bin/check-shadowsocks "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/checkusershadowsocks.sh" && chmod +x /usr/local/bin/check-shadowsocks    
wget -q -O /usr/local/bin/check-ssh "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/checkuserssh.sh" && chmod +x /usr/local/bin/check-ssh
wget -q -O /usr/local/bin/check-vmess "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/checkuservmess.sh" && chmod +x /usr/local/bin/check-vmess      



echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green} INSTALLASI renew package HIdeSSH      ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"
cd
 
wget -q -O /usr/local/bin/renew-vmess "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/renewaccvmess.sh" && chmod +x /usr/local/bin/renew-vmess
wget -q -O /usr/local/bin/renew-ssh "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/renewaccssh.sh" && chmod +x /usr/local/bin/renew-ssh
wget -q -O /usr/local/bin/renew-vless "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/renewaccvless.sh" && chmod +x /usr/local/bin/renew-vless
wget -q -O /usr/local/bin/renew-trojan "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/renewacctrojan.sh" && chmod +x /usr/local/bin/renew-trojan
wget -q -O /usr/local/bin/renew-shadowsocks "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/renewshadowsocks.sh" && chmod +x /usr/local/bin/renew-shadowsocks


echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green} INSTALLASI Trial package HIdeSSH      ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"
cd

wget -q -O /usr/local/bin/trial-vmess "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/trialvmess" && chmod +x /usr/local/bin/trial-vmess
wget -q -O /usr/local/bin/trial-ssh "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/trialssh" && chmod +x /usr/local/bin/trial-ssh
wget -q -O /usr/local/bin/trial-vless "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/trialvless" && chmod +x /usr/local/bin/trial-vless
wget -q -O /usr/local/bin/trial-trojan "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/trialTrojan" && chmod +x /usr/local/bin/trial-trojan
wget -q -O /usr/local/bin/trial-shadowsocks "https://raw.githubusercontent.com/hidessh99/scAUTOku/refs/heads/main/project/trialshadowsocks" && chmod +x /usr/local/bin/trial-shadowsocks

echo -e "${blue}─────────────────────────────────────────${neutral}"
echo -e "${green} Finish install package HIdeSSH      ${neutral}"
echo -e "${blue}─────────────────────────────────────────${neutral}"
cd