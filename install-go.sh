#!/usr/bin/env bash

set -euo pipefail

# Auto installer Golang terbaru untuk VPS Ubuntu/Debian
# - Mengambil versi Go terbaru dari go.dev
# - Mengunduh dan memasang ke /usr/local/go
# - Menambahkan PATH via /etc/profile.d/go.sh
# - Mendukung arsitektur amd64, arm64, 386

require_root() {
  if [[ $(id -u) -ne 0 ]]; then
    echo "[INFO] Memerlukan hak akses root. Coba dengan sudo..."
    exec sudo -E bash "$0" "$@"
  fi
}

require_cmd() {
  local cmd="$1"
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "[INFO] Menginstal dependensi: $cmd"
    if command -v apt-get >/dev/null 2>&1; then
      apt-get update -y >/dev/null 2>&1 || true
      apt-get install -y "$cmd"
    else
      echo "[ERROR] Tidak menemukan apt-get untuk memasang $cmd. Pasang manual lalu jalankan ulang." >&2
      exit 1
    fi
  fi
}

detect_arch() {
  local machine
  machine=$(uname -m)
  case "$machine" in
    x86_64|amd64) echo "amd64" ;;
    aarch64|arm64) echo "arm64" ;;
    i386|i686) echo "386" ;;
    *)
      echo "[ERROR] Arsitektur tidak didukung: $machine" >&2
      exit 1
      ;;
  esac
}

main() {
  require_root "$@"
  require_cmd curl
  require_cmd tar

  echo "[INFO] Mengambil versi Go terbaru..."
  local version
  version=$(curl -fsSL https://go.dev/VERSION?m=text)
  if [[ -z "$version" ]]; then
    echo "[ERROR] Gagal memperoleh versi Go terbaru." >&2
    exit 1
  fi
  echo "[INFO] Versi terbaru: $version"

  local arch
  arch=$(detect_arch)
  local url
  url="https://go.dev/dl/${version}.linux-${arch}.tar.gz"
  local tmpfile
  tmpfile="/tmp/${version}.linux-${arch}.tar.gz"

  echo "[INFO] Mengunduh paket: $url"
  curl -fL "$url" -o "$tmpfile"

  echo "[INFO] Menghapus instalasi Go lama jika ada..."
  rm -rf /usr/local/go

  echo "[INFO] Mengekstrak ke /usr/local"
  tar -C /usr/local -xzf "$tmpfile"

  echo "[INFO] Menulis PATH ke /etc/profile.d/go.sh"
  cat >/etc/profile.d/go.sh <<'EOF'
export GOROOT=/usr/local/go
export GOPATH="$HOME/go"
export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
EOF
  chmod 0644 /etc/profile.d/go.sh

  # Muat env untuk sesi saat ini jika memungkinkan
  # shellcheck disable=SC1091
  source /etc/profile.d/go.sh || true

  echo "[INFO] Selesai memasang Go. Versi terpasang:"
  /usr/local/go/bin/go version || (echo "[WARN] Pastikan sesi shell baru untuk memuat PATH" && true)

  echo "[INFO] GOPATH akan berada di: $HOME/go (per user)"
  echo "[INFO] Selesai. Buka sesi shell baru atau jalankan: source /etc/profile.d/go.sh"
}

main "$@"