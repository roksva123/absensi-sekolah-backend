# 🚀 Quick Start Guide

## ⚡ Fastest Way to Get Running

### Option 1: Docker (Recommended)

```bash
# Clone and navigate to project
cd absensi-sekolah-backend

# Windows
.\start.bat

# macOS/Linux
bash start.sh
```

Server akan berjalan di `http://localhost:8080`

### Option 2: Local Development

**Prasyarat:**
- Go 1.21+
- PostgreSQL 12+

**Setup:**

```bash
# 1. Setup environment
cp .env.example .env
# Edit .env dengan konfigurasi database Anda

# 2. Download dependencies
go mod download

# 3. Run server
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## 📝 First Steps

### 1. Login dengan Default Credentials

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'
```

Response:
```json
{
  "user": {
    "id": 1,
    "name": "Admin User",
    "email": "admin@example.com",
    "role": "admin"
  },
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### 2. Copy Token untuk Penggunaan Berikutnya

```bash
# Set sebagai environment variable
export TOKEN="eyJhbGciOiJIUzI1NiIs..."
```

### 3. Test RFID Tap Endpoint

```bash
curl -X POST http://localhost:8080/api/v1/scan/tap \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "uid_kartu": "12345ABC"
  }'
```

## 📊 Testing dengan Postman

1. Download [Postman Collection](./postman-collection.json) (jika ada)
2. Import ke Postman
3. Set Authorization tab dengan Bearer Token dari login
4. Test endpoints

## 🛠️ Development Commands

```bash
# Build binary
make build

# Run server
make run

# Run with hot reload (requires air)
make dev

# View Docker logs
make docker-logs

# Stop Docker
make docker-down
```

## 📚 Dokumentasi Lengkap

- [README.md](./README.md) - Dokumentasi lengkap API
- [ARCHITECTURE.md](./ARCHITECTURE.md) - Desain sistem
- [Makefile](./Makefile) - Daftar command development

## 🔗 API Endpoints Utama

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/v1/auth/login` | Login user |
| POST | `/api/v1/scan/tap` | RFID scan tap |
| POST | `/api/v1/siswa` | Buat siswa |
| GET | `/api/v1/siswa/:id` | Get siswa |
| POST | `/api/v1/pengajuan-izin` | Buat pengajuan izin |
| PUT | `/api/v1/pengajuan-izin/:id/approve` | Approve izin |
| GET | `/api/v1/reports/absensi` | Laporan absensi |

## 🐛 Troubleshooting

### "Connection refused"
- Pastikan PostgreSQL berjalan
- Verifikasi konfigurasi `.env`

### "Invalid token"
- Token mungkin sudah expired
- Login kembali untuk mendapatkan token baru

### Docker not working?
- Install Docker Desktop
- Pastikan Docker daemon berjalan

## 📞 Perlu Bantuan?

Cek dokumentasi lengkap di:
- [README.md](./README.md) - Penjelasan detail
- [ARCHITECTURE.md](./ARCHITECTURE.md) - Desain sistem

---

**Siap! 🎉 Backend Anda sudah berjalan.**
