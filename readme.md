Mantap banget tutorialnya! Ini sudah lengkap dari awal sampai akhir untuk deploy aplikasi Golang ke VPS IDCloudHost dengan reverse proxy + SSL. Sesuai permintaan kamu, berikut versi **README.md** yang siap kamu pakai langsung di repositori GitHub kamu:

---

```markdown
# 🚀 Deploy Aplikasi Golang ke VPS IDCloudHost (Full Setup Guide)

Tutorial ini menjelaskan cara **full setup** deploy aplikasi Golang ke VPS IDCloudHost lengkap dengan domain, reverse proxy, dan HTTPS.

---

## 🧱 1. Siapkan VPS dari IDCloudHost
1. Login ke [IDCloudHost Panel](https://portal.idcloudhost.com/)
2. Buat VPS (disarankan: Ubuntu 22.04 LTS)
3. Catat:
    - IP VPS
    - Username (biasanya `root`)
    - Password

---

## 💻 2. Login ke VPS via Terminal

```bash
ssh root@IP_VPS_KAMU
# Contoh:
ssh root@116.193.190.50
```

---

## 🔧 3. Update VPS & Install Tools

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install curl git ufw nginx unzip -y
```

---

## 🛡️ 4. Aktifkan Firewall & Izinkan Port Penting

```bash
sudo ufw allow OpenSSH
sudo ufw allow 'Nginx Full'
sudo ufw enable
```

---

## 🟨 5. Install Golang

```bash
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

---

## 🧑‍💻 6. Clone Aplikasi Golang Kamu

```bash
cd /home
git clone https://github.com/username/project-golang.git
cd project-golang
```

---

## 🗄️ 7. Setup Database (SQLite / PostgreSQL)

### Jika SQLite:

```bash
sudo apt install gcc libsqlite3-dev -y
```

### Jika PostgreSQL:

```bash
sudo apt install postgresql postgresql-contrib -y
sudo -u postgres createuser myuser --interactive
sudo -u postgres createdb mydb
```

---

## 🔨 8. Build Aplikasi Golang

```bash
CGO_ENABLED=1 go build -o myapp main.go
```

---

## ⚙️ 9. Install PM2

```bash
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs
sudo npm install pm2 -g
```

---

## 🚦 10. Jalankan Aplikasi dengan PM2

```bash
pm2 start ./myapp --name go-server
pm2 startup
pm2 save
```

---

## 🌐 11. Pasang Domain ke VPS

1. Masuk ke panel domain kamu (Cloudflare / Niagahoster / Rumahweb / dll.)
2. Edit A Record:

```
Name: @      | Type: A | Value: IP VPS kamu | TTL: Auto
Name: www    | Type: A | Value: IP VPS kamu | TTL: Auto
```

---

## 🔧 12. Setup Nginx sebagai Reverse Proxy

```bash
sudo nano /etc/nginx/sites-available/goserver
```

Isi file:

```nginx
server {
    listen 80;
    server_name infodesahaurkolot.my.id www.infodesahaurkolot.my.id;

    location / {
        proxy_pass http://localhost:3000; # sesuaikan port aplikasi kamu
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection keep-alive;
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

Aktifkan konfigurasi:

```bash
sudo ln -s /etc/nginx/sites-available/goserver /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

---

## 🔒 13. Pasang SSL dengan Let’s Encrypt

```bash
sudo apt install certbot python3-certbot-nginx -y
sudo certbot --nginx -d infodesahaurkolot.my.id -d www.infodesahaurkolot.my.id
```

Cek auto-renew:

```bash
sudo certbot renew --dry-run
```

---

## ✅ 14. Selesai!

Buka di browser:

```
https://infodesahaurkolot.my.id
```

Kalau muncul aplikasi kamu → SELAMAT BERHASIL 🎉

---

## 🎁 Bonus: Monitoring Log PM2

```bash
pm2 logs go-server
```

---

## 🛠️ Butuh Bantuan Tambahan?

Jika kamu ingin setup:

- `.env` dengan PM2
- `ecosystem.config.js`
- PostgreSQL lebih advance (role, permission, backup)
- Auto-deploy via **GitHub Actions CI/CD**

Tinggal buka issue atau kontak saya langsung ✨
```

---

Kalau kamu mau, aku juga bisa bantu bikinkan file `ecosystem.config.js` untuk PM2 dan contoh `.env` Golang, tinggal bilang aja ya.