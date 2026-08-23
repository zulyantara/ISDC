# 🚗 ISDC - Indonesia Safety Driving Center

Sistem manajemen pelatihan Safety Riding & Driving, mencakup pendaftaran peserta, ujian, penilaian, dan sertifikasi.

## 📋 Daftar Isi

- [Arsitektur](#arsitektur)
- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [API Documentation](#api-documentation)
- [Struktur Database](#struktur-database)
- [Troubleshooting](#troubleshooting)

---

## Arsitektur

```
┌──────────────────┐     ┌──────────────────┐     ┌──────────────────┐
│   React Frontend │────▶│   Go API Backend  │────▶│   MySQL Database │
│   (Port 3000)    │     │   (Port 8080)     │     │   (Port 3306)    │
│   Ant Design     │     │   Gin Framework   │     │   Docker         │
└──────────────────┘     └──────────────────┘     └──────────────────┘
```

| Komponen | Tech Stack | Port |
|----------|-----------|------|
| Frontend | React 18 + Ant Design 5 + Vite | 3000 |
| Backend | Go 1.21 + Gin + JWT | 8080 |
| Database | MySQL 8.0 (Docker) | 3306 |

---

## Prasyarat

Pastikan tools berikut terinstall di komputer Anda:

| Tool | Versi Minimum | Cek Instalasi |
|------|--------------|---------------|
| **Docker** | 20.10+ | `docker --version` |
| **Docker Compose** | 2.0+ | `docker compose version` |
| **Go** | 1.21+ | `go version` |
| **Node.js** | 18+ | `node --version` |
| **npm** | 9+ | `npm --version` |

---

## Instalasi

### 1. Clone Repository

```bash
git clone <url-repo>
cd isdc
```

### 2. Siapkan Database (Docker)

```bash
# Jalankan MySQL container
docker compose up -d

# Cek status container (tunggu sampai healthy)
docker compose ps
```

> **Catatan:** Database akan diinisialisasi dari file SQL di folder `database/` saat pertama kali dijalankan.

### 3. Setup Backend (Go API)

```bash
cd backend

# Buat file .env dari template
cp .env.example .env

# Edit .env sesuai kebutuhan (lihat tabel konfigurasi di bawah)
# ...

# Install dependencies
go mod tidy

# Jalankan server
go run main.go
```

**Konfigurasi `.env`:**

| Variable | Default | Keterangan |
|----------|---------|------------|
| `DB_HOST` | `localhost` | Host database |
| `DB_PORT` | `3306` | Port database |
| `DB_USER.*isdc` | Username database |
| `DB_PASSWORD.*isdc_pass` | Password database |
| `DB_NAME` | `isdc_db` | Nama database |
| `JWT_SECRET` | *(generate sendiri)* | Secret key JWT |
| `JWT_EXPIRY` | `24` | Token expired dalam jam |
| `SERVER_PORT` | `8080` | Port Go API |
| `SERVER_MODE` | `debug` | `debug` atau `release` |

### 4. Setup Frontend (React)

```bash
cd frontend

# Install dependencies
npm install

# Jalankan development server
npm run dev
```

---

## Menjalankan Aplikasi

### Semua Sekaligus (3 Terminal)

Buka **3 terminal** dan jalankan perintah berikut:

**Terminal 1 — Database:**
```bash
docker compose up -d
```

**Terminal 2 — Backend API:**
```bash
cd backend
go run main.go
```

**Terminal 3 — Frontend:**
```bash
cd frontend
npm run dev
```

### Akses Aplikasi

| Service | URL |
|---------|-----|
| Frontend | http://localhost:3000 |
| Backend API | http://localhost:8080 |
| Health Check | http://localhost:8080/health |

### Login Default

| User ID | Password | Level | Keterangan |
|---------|----------|-------|------------|
| `ADMIN` | *(lihat di mt_user)* | 1 (Admin) | Full akses |
| `JIMMY` | *(lihat di mt_user)* | 2 (Super Admin) | Full akses |
| `KASIR1` | *(lihat di mt_user)* | 3 (Kasir) | Pendaftaran & cetak |

> **Catatan:** Password di database dalam format MD5. Untuk reset password, jalankan:
> ```sql
> UPDATE mt_user SET user_pwd = MD5('password_baru') WHERE user_id = 'ADMIN';
> ```

### Menghentikan Aplikasi

```bash
# Hentikan database
docker compose down

# Hentikan Go API: Ctrl+C di terminal
# Hentikan React: Ctrl+C di terminal
```

### Hapus Database (Fresh Start)

```bash
docker compose down -v
docker compose up -d
```

---

## API Documentation

### Auth

| Method | Endpoint | Body | Keterangan |
|--------|----------|------|------------|
| `POST` | `/api/auth/login` | `{ user_id, user_pwd }` | Login, return JWT |
| `POST` | `/api/auth/logout` | - | Logout |
| `GET` | `/api/auth/me` | - | Info user login |

### Pendaftaran

| Method | Endpoint | Body | Keterangan |
|--------|----------|------|------------|
| `GET` | `/api/daftar` | - | List semua pendaftaran |
| `GET` | `/api/daftar/:id` | - | Detail pendaftaran |
| `POST` | `/api/daftar` | `{ nama, kelas_id, ... }` | Tambah pendaftaran (auto ID) |
| `PUT` | `/api/daftar/:id` | `{ nama, kelas_id, ... }` | Edit pendaftaran |
| `DELETE` | `/api/daftar/:id` | - | Hapus pendaftaran |

### Peserta & Ujian

| Method | Endpoint | Body | Keterangan |
|--------|----------|------|------------|
| `GET` | `/api/peserta` | - | List semua peserta |
| `GET` | `/api/peserta/search?q=keyword` | - | Cari peserta |
| `GET` | `/api/peserta/:id` | - | Detail peserta |
| `GET` | `/api/peserta/:id/soal` | - | Ambil soal untuk peserta |
| `POST` | `/api/peserta/:id/praktek` | `{ results: [{ soal_id, hasil }] }` | Submit nilai ujian |
| `POST` | `/api/peserta/:id/comment` | `{ pengetahuan, teknik, perilaku }` | Submit komentar |
| `GET` | `/api/peserta/:id/comment` | - | Lihat komentar |

### Master Data

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| `GET/POST/PUT/DELETE` | `/api/kelas` | CRUD Kelas |
| `GET/POST/PUT/DELETE` | `/api/area` | CRUD Area |
| `GET/POST/PUT/DELETE` | `/api/users` | CRUD User |
| `GET/POST/PUT/DELETE` | `/api/soal` | CRUD Soal |
| `GET/POST/PUT/DELETE` | `/api/nilai-lulus` | CRUD Nilai Lulus |
| `GET/POST/PUT/DELETE` | `/api/jenis-dokumen` | CRUD Jenis Dokumen |

> **Keterangan:** Semua endpoint (kecuali login) memerlukan header `Authorization: Bearer <token>`.

### Contoh Request dengan cURL

**Login:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"user_id":"ADMIN","user_pwd":"password"}'
```

**Ambil Data Pendaftaran:**
```bash
curl -X GET http://localhost:8080/api/daftar \
  -H "Authorization: Bearer <token-dari-login>"
```

**Submit Nilai Ujian:**
```bash
curl -X POST http://localhost:8080/api/peserta/2013.09.00252/praktek \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"results":[{"soal_id":61,"hasil":80},{"soal_id":66,"hasil":75}]}'
```

---

## Struktur Database

### Tabel Utama

| Tabel | Keterangan |
|-------|------------|
| `mt_user` | User/petugas (password MD5) |
| `mt_kelas` | Kelas/pelatihan (BASIC, ADVANCE, dll) |
| `mt_area` | Area/lokasi (JAKARTA, MEDAN, dll) |
| `mt_nilai_lulus` | Pengaturan nilai minimum lulus |
| `tb_daftar` | Pendaftaran peserta |
| `tb_peserta` | Data peserta pelatihan |
| `tb_soal` | Soal ujian per kategori |
| `tb_uji_praktek` | Hasil ujian praktek |
| `tb_comments` | Komentar penilaian |
| `tb_jenis_dokumen` | Jenis dokumen |
| `tb_daftar_dokumen` | Entri dokumen |

### ERD Ringkas

```
mt_user ──────┐
              ├──▶ tb_daftar ──▶ tb_peserta
mt_kelas ─────┘                  │
mt_area ──────┘                  ├──▶ tb_uji_praktek
                                 ├──▶ tb_comments
                                 └──▶ tb_soal (via teori_id)
```

---

## Troubleshooting

### Database tidak bisa connect

```bash
# Cek container MySQL berjalan
docker compose ps

# Cek log MySQL
docker compose logs

# Restart container
docker compose restart
```

### MySQL "Access denied"

```bash
# Login ke MySQL dan cek user
docker exec -it isdc-mysql mysql -u root -pisdc_root_2024 -e "SELECT user, host FROM mysql.user;"
```

### Port sudah digunakan

```bash
# Cari proses yang menggunakan port
lsof -i :3306   # untuk MySQL
lsof -i :8080   # untuk Go API
lsof -i :3000   # untuk React

# Ubah port di .env atau docker-compose.yml
```

### Go API tidak bisa start

```bash
# Pastikan dependency terinstall
cd backend
go mod tidy

# Cek apakah ada error compile
go build ./...
```

### React npm install error

```bash
# Hapus node_modules dan install ulang
cd frontend
rm -rf node_modules package-lock.json
npm install
```

### Reset Database ke Awal

```bash
docker compose down -v
docker compose up -d
```

---

## 📁 Struktur Project

```
isdc/
├── docker-compose.yml         # Docker MySQL
├── .gitignore
├── README.md
├── frontend/                  # React Frontend
│   ├── src/
│   │   ├── api/               # Axios config
│   │   ├── context/           # Auth context
│   │   ├── components/        # Layout
│   │   ├── pages/             # Halaman
│   │   └── utils/             # Helpers
│   ├── package.json
│   └── vite.config.js
├── backend/                   # Go Backend
│   ├── main.go                # Entry point
│   ├── config/                # DB, JWT, Env
│   ├── middleware/             # Auth, CORS
│   ├── models/                # Database models
│   ├── handlers/              # API handlers
│   ├── routes/                # Routing
│   ├── utils/                 # Helpers
│   ├── go.mod
│   └── go.sum
└── database/                  # SQL dump files
```

---

## Lisensi

Private - ISDC System
