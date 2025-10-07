## Installer Go Terbaru untuk Ubuntu/Debian

Script `install-go.sh` ini memasang Golang versi terbaru secara otomatis di VPS berbasis Ubuntu/Debian, mengonfigurasi PATH, serta mendukung arsitektur `amd64`, `arm64`, dan `386`.

### Prasyarat
- Akses root atau sudo.
- Koneksi internet.
- Paket dasar: `curl` dan `tar` (script akan memasang jika belum ada).

### Langkah Instalasi
1. Salin file `install-go.sh` ke server Anda (atau buat file dengan isi dari repository ini).
2. Jadikan executable:
   ```bash
   chmod +x install-go.sh
   ```
3. Jalankan installer:
   ```bash
   sudo ./install-go.sh
   ```
4. Setelah selesai, mulai sesi shell baru atau muat PATH:
   ```bash
   source /etc/profile.d/go.sh
   ```
5. Verifikasi:
   ```bash
   go version
   ```

### Apa yang Dilakukan Script
- Mengambil versi Go terbaru dari `https://go.dev/VERSION?m=text`.
- Mengunduh tarball sesuai arsitektur: `https://go.dev/dl/<version>.linux-<arch>.tar.gz`.
- Mengekstrak ke `/usr/local/go`.
- Membuat `/etc/profile.d/go.sh` untuk mengatur `GOROOT`, `GOPATH` (per user: `$HOME/go`), dan menambahkan Go ke `PATH`.

### Uninstall / Upgrade Manual
- Uninstall:
  ```bash
  sudo rm -rf /usr/local/go
  sudo rm -f /etc/profile.d/go.sh
  ```
- Upgrade: cukup jalankan ulang `install-go.sh`; script akan menghapus instalasi lama dan memasang versi terbaru.

### Catatan
- `GOPATH` default diset ke `$HOME/go` untuk setiap user.
- Jika `go version` tidak langsung tersedia, jalankan `source /etc/profile.d/go.sh` atau buka sesi shell baru.