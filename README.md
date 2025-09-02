# Script SSH/VPN
## Installation


### Step 1

    apt-get update && \
    apt-get --reinstall --fix-missing install -y whois bzip2 gzip coreutils wget screen nscd && \
    wget --inet4-only --no-check-certificate -O setup.sh https://raw.githubusercontent.com/hidessh99/goautoscript/refs/heads/main/setup.sh && \
    chmod +x setup.sh && \
    screen -S setup ./setup.sh

### Informasi

- Jika dalam proses instalasi [Step 1](#Step-1), terjadi diskoneksi pada terminal. Jangan masukkan kembali perintah instalasi [Step 1](#Step-1). Silahkan masukkan perintah `screen -r setup` untuk melihat proses yang telah berjalan.
- Jika ingin melihat log instalasi dapat dilihat pada `/root/syslog.log`.
- Laporan bug bisa dilakukan pada akun [hidessh admin](https://t.me/hidessh).

## Tutorial

### Auto Reboot

Secara default script ini tidak diberikan sistem auto reboot karena tidak semua pengguna membutuhkannya. Jika kamu ingin memasang auto reboot pada VPS bisa gunakan perintah berikut ini

    crontab -l > /tmp/cron.txt
    sed -i "/reboot$/d" /tmp/cron.txt
    echo -e "\n"'0 4 * * * '"$(which reboot)" >> /tmp/cron.txt
    crontab /tmp/cron.txt
    rm -rf /tmp/cron.txt

Perintah di atas akan memasang auto reboot setiap jam 04.00.

Perintah untuk membatalkan.

    crontab -l > /tmp/cron.txt
    sed -i "/reboot$/d" /tmp/cron.txt
    crontab /tmp/cron.txt
    rm -rf /tmp/cron.txt