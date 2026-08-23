# 🚀 DEPLOY ISDC ke Server

Panduan lengkap deploy aplikasi ISDC (React + Go + MySQL) ke production server.

---

## 📋 Prasyarat Server

| Tool | Versi | Status |
|------|-------|--------|
| **Docker** | 20.10+ | ✅ Sudah ada |
| **Docker Compose** | 2.0+ | ✅ Sudah ada |
| **Node.js** | 24.18.0 | ✅ Sudah ada |
| **npm** | 12.0.2 | ✅ Sudah ada |
| **pm2** | Latest | ✅ Sudah ada |
| **Go** | 1.21+ | ❌ Perlu install |

---

## 🛠️ Step 1: Install Go

### Rekomendasi: Install Go Langsung (Bukan Docker)

**Alasan install langsung:**
- Go compile ke single binary → tidak butuh runtime
- Lebih ringan, tidak perlu Docker layer tambahan
- PM2 bisa manage langsung
- MySQL sudah pakai Docker, tidak perlu tambah complexity

```bash
# Download Go 1.22 (latest stable)
cd /tmp
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz

# Hapus Go lama jika ada
sudo rm -rf /usr/local/go

# Extract ke /usr/local
sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

# Tambahkan Go ke PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc

# Verifikasi
go version
# Output: go version go1.22.2 linux/amd64
```

---

## 📦 Step 2: Clone Repository

```bash
cd /home/$USER
git clone git@github.com:zulyantara/ISDC.git
cd ISDC
```

---

## 🗄️ Step 3: Setup Database

### 3.1 Buat file `.env` di root

```bash
cat > .env << 'EOF'
# MySQL Configuration
DB_ROOT_PASSWORD=isdc_root_2024
DB_NAME=isdc_db
DB_USER=isdc
DB_PASSWORD=isdc_pass
DB_PORT=3306
EOF
```

### 3.2 Jalankan MySQL

```bash
docker compose up -d
sleep 10  # Tunggu MySQL ready

# Cek status
docker compose ps
# Harus "healthy"
```

### 3.3 Import Database

```bash
# Import schema utama
docker exec -i isdc-mysql mysql -u root -pisdc_root_2024 isdc_db < database/db_jsdc.sql

# Verifikasi
docker exec isdc-mysql mysql -u root -pisdc_root_2024 isdc_db -e "SHOW TABLES;"
```

### 3.4 Resize kolom password (penting!)

```bash
# Kolom user_pwd harus cukup untuk bcrypt hash (60 chars)
docker exec isdc-mysql mysql -u root -pisdc_root_2024 isdc_db \
  -e "ALTER TABLE mt_user MODIFY COLUMN user_pwd VARCHAR(255) NOT NULL;"
```

### 3.5 Reset password default

```bash
docker exec isdc-mysql mysql -u root -pisdc_root_2024 isdc_db \
  -e "UPDATE mt_user SET user_pwd=MD5('admin123') WHERE user_id='ADMIN';"
```

---

## ⚙️ Step 4: Setup Backend (Go API)

### 4.1 Buat file `.env`

```bash
cd backend
cat > .env << 'EOF'
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=isdc
DB_PASSWORD=isdc_pass
DB_NAME=isdc_db

# JWT Configuration
JWT_SECRET=your-production-secret-change-this-$(openssl rand -hex 32)
JWT_EXPIRY=24

# Server Configuration
SERVER_PORT=8080
SERVER_MODE=release

# Default Password
DEFAULT_PASSWORD=password321!*
EOF
```

### 4.2 Generate JWT Secret

```bash
# Ganti JWT_SECRET di .env dengan secret yang aman
JWT_SECRET=$(openssl rand -hex 32)
sed -i "s|JWT_SECRET=.*|JWT_SECRET=$JWT_SECRET|" .env
```

### 4.3 Build & Test

```bash
cd backend

# Install dependencies
go mod tidy

# Build binary
go build -o /home/$USER/isdc-api main.go

# Test run
./isdc-api &
sleep 2
curl http://localhost:8080/health
# {"status":true,"message":"ISDC API is running"}
kill %1
```

### 4.4 Setup PM2

```bash
cd /home/$USER/ISDC

# Start with PM2
pm2 start backend/isdc-api \
  --name "isdc-api" \
  --cwd /home/$USER/ISDC/backend \
  --log /home/$USER/ISDC/logs/api.log \
  --time

# Save PM2 config
pm2 save

# Setup auto-start on boot
pm2 startup
# Ikuti instruksi yang muncul
```

### 4.5 Cek Status

```bash
pm2 status
pm2 logs isdc-api --lines 20
```

---

## 🎨 Step 5: Setup Frontend (React)

### 5.1 Build untuk Production

```bash
cd frontend

# Install dependencies
npm install

# Build
npm run build
# Output ke dist/
```

### 5.2 Serve dengan Nginx

```bash
# Install Nginx (jika belum ada)
sudo apt update && sudo apt install -y nginx

# Copy build ke Nginx
sudo cp -r dist/* /var/www/html/

# Atau buat location khusus
sudo mkdir -p /var/www/isdc
sudo cp -r dist/* /var/www/isdc/
```

### 5.3 Config Nginx

```bash
sudo cat > /etc/nginx/sites-available/isdc << 'EOF'
server {
    listen 80;
    server_name your-domain.com;  # Ganti dengan domain/IP server

    # Frontend React
    location / {
        root /var/www/isdc;
        index index.html;
        try_files $uri $uri/ /index.html;  # SPA fallback
    }

    # Backend API Proxy
    location /api/ {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Health check proxy
    location /health {
        proxy_pass http://127.0.0.1:8080;
    }

    # Gzip compression
    gzip on;
    gzip_types text/plain text/css application/json application/javascript text/xml;
}
EOF

# Aktifkan config
sudo ln -sf /etc/nginx/sites-available/isdc /etc/nginx/sites-enabled/
sudo rm -f /etc/nginx/sites-enabled/default

# Test & reload
sudo nginx -t
sudo systemctl reload nginx
```

---

## ✅ Step 6: Verifikasi

```bash
# 1. Cek semua service
docker compose ps          # MySQL
pm2 status                 # Go API
sudo systemctl status nginx  # Nginx

# 2. Test API
curl http://localhost:8080/health

# 3. Test Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"user_id":"ADMIN","user_pwd":"admin123"}'

# 4. Test Frontend
curl -I http://localhost/
# Harus return 200 OK
```

---

## 🔧 Step 7: Setup SSL (Opsional)

### Menggunakan Let's Encrypt

```bash
# Install Certbot
sudo apt install -y certbot python3-certbot-nginx

# Dapatkan SSL certificate
sudo certbot --nginx -d your-domain.com

# Auto-renew
sudo certbot renew --dry-run
```

---

## 📝 Step 8: Maintenance

### Restart Services

```bash
# Restart MySQL
docker compose restart

# Restart Go API
pm2 restart isdc-api

# Restart Nginx
sudo systemctl restart nginx
```

### Lihat Logs

```bash
# MySQL logs
docker compose logs -f

# Go API logs
pm2 logs isdc-api

# Nginx logs
sudo tail -f /var/log/nginx/access.log
sudo tail -f /var/log/nginx/error.log
```

### Update Application

```bash
cd /home/$USER/ISDC

# Pull latest code
git pull

# Rebuild backend
cd backend
go build -o isdc-api main.go
pm2 restart isdc-api

# Rebuild frontend
cd ../frontend
npm run build
sudo cp -r dist/* /var/www/isdc/
```

---

## 🐛 Troubleshooting

### MySQL tidak bisa start
```bash
docker compose down -v
docker compose up -d
```

### Go API tidak bisa connect ke database
```bash
# Cek apakah MySQL running
docker compose ps

# Cek koneksi
docker exec isdc-mysql mysql -u isdc -pisdc_pass isdc_db -e "SELECT 1;"
```

### PM2 Go API crash
```bash
# Lihat logs
pm2 logs isdc-api --err --lines 50

# Restart
pm2 restart isdc-api

# Jika binary corrupt, rebuild
cd backend && go build -o isdc-api main.go
pm2 restart isdc-api
```

### Port sudah digunakan
```bash
# Cari proses yang menggunakan port
sudo lsof -i :8080
sudo lsof -i :3306
sudo lsof -i :80

# Kill proses
sudo kill -9 <PID>
```

---

## 📊 Ringkasan Architecture

```
┌─────────────────────────────────────────────────────┐
│                    PRODUCTION SERVER                  │
│                                                      │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐       │
│  │  Nginx   │───▶│ Go API   │───▶│  MySQL   │       │
│  │  :80     │    │  :8080   │    │  :3306   │       │
│  │  (React) │    │  (PM2)   │    │ (Docker) │       │
│  └──────────┘    └──────────┘    └──────────┘       │
│       │                              │               │
│       ▼                              ▼               │
│  /var/www/isdc                  Docker Volume        │
│  (static files)                (persistent data)     │
│                                                      │
└─────────────────────────────────────────────────────┘
```

---

## 🔒 Security Notes

1. **Ganti JWT_SECRET** di `backend/.env` dengan secret yang unik
2. **Ganti password MySQL** di `docker-compose.yml` dan `backend/.env`
3. **Ganti password admin** setelah login pertama
4. **Enable SSL** untuk production
5. **Block port 3306** dari luar (hanya localhost yang bisa akses)
6. **Gunakan firewall** (UFW/iptables)

---

## 📞 Quick Command Reference

```bash
# === Status Check ===
docker compose ps
pm2 status
sudo systemctl status nginx

# === Restart All ===
docker compose restart
pm2 restart all
sudo systemctl restart nginx

# === Logs ===
docker compose logs -f
pm2 logs
sudo tail -f /var/log/nginx/error.log

# === Database ===
docker exec isdc-mysql mysql -u root -pisdc_root_2024 isdc_db

# === Update ===
git pull && cd backend && go build -o isdc-api main.go && pm2 restart isdc-api && cd ../frontend && npm run build && sudo cp -r dist/* /var/www/isdc/
```
