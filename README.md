# Absensi Sekolah Backend API

Backend RESTful API untuk sistem absensi sekolah berbasis kartu RFID yang dibangun dengan Golang, Gin Framework, GORM, dan PostgreSQL.

## 📋 Daftar Isi

- [Fitur Utama](#fitur-utama)
- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Konfigurasi](#konfigurasi)
- [Menjalankan Server](#menjalankan-server)
- [API Endpoints](#api-endpoints)
- [Struktur Database](#struktur-database)
- [Struktur Kode](#struktur-kode)

## ✨ Fitur Utama

- **Autentikasi JWT** - Keamanan endpoint dengan JWT token
- **Manajemen User** - Login dan manajemen user dengan role-based access
- **RFID Scanning** - Endpoint untuk tap kartu RFID siswa
- **Manajemen Data Siswa** - CRUD untuk data siswa dengan RFID card tracking
- **Manajemen Kelas** - CRUD untuk data kelas dan wali kelas
- **Pengajuan Izin** - Sistem pengajuan izin/sakit dengan approval workflow
- **Absensi Otomatis** - Recording absensi berdasarkan tap RFID atau manual
- **Laporan Absensi** - Rekapitulasi absensi dengan filter tanggal dan kelas
- **Pengaturan Sistem** - Konfigurasi jam masuk/pulang per hari
- **Kalender Akademik** - Manajemen tanggal libur dan hari istimewa

## 🔧 Prasyarat

- Go 1.21+
- PostgreSQL 12+
- Git

## 📦 Instalasi

### 1. Clone Repository

```bash
git clone <repository-url>
cd absensi-sekolah-backend
```

### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

Atau langsung jalankan go run (Go akan otomatis download dependencies):

```bash
go run main.go
```

## ⚙️ Konfigurasi

### 1. Setup Database PostgreSQL

Buat database baru di PostgreSQL:

```sql
CREATE DATABASE absensi_sekolah;
```

### 2. Setup Environment Variables

Buat file `.env` berdasarkan `.env.example`:

```bash
cp .env.example .env
```

Edit file `.env` dengan konfigurasi Anda:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=absensi_sekolah

# Application Configuration
APP_PORT=8080

# JWT Secret (ubah dengan secret yang aman)
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
```

## 🚀 Menjalankan Server

### Development Mode

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

### Build untuk Production

```bash
go build -o absensi-sekolah-backend .
./absensi-sekolah-backend
```

### Menggunakan Docker (Optional)

```bash
docker build -t absensi-backend .
docker run -p 8080:8080 --env-file .env absensi-backend
```

## 📡 API Endpoints

### Authentication

#### Login
```
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "admin@example.com",
  "password": "password123"
}

Response: 200 OK
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

### RFID Scanning

#### Tap Kartu RFID
```
POST /api/v1/scan/tap
Authorization: Bearer <token>
Content-Type: application/json

{
  "uid_kartu": "12345ABC"
}

Response: 200 OK
{
  "success": true,
  "data": {
    "id": 1,
    "siswa_id": 5,
    "tanggal": "2024-01-15",
    "jam_masuk": "07:30:00",
    "status_kehadiran": "hadir",
    "entry_mode": "rfid"
  }
}
```

### Siswa Management

#### Buat Siswa
```
POST /api/v1/siswa
Authorization: Bearer <token>
Content-Type: application/json

{
  "nis": "12345",
  "nisn": "0012345678901",
  "nama": "Budi Santoso",
  "gender": "L",
  "nama_ortu": "Santoso",
  "uid_kartu": "12345ABC"
}

Response: 201 Created
```

#### Get Siswa
```
GET /api/v1/siswa/:id
Authorization: Bearer <token>

Response: 200 OK
```

#### Update Siswa
```
PUT /api/v1/siswa/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "nama": "Budi Santoso Updated",
  "nama_ortu": "Santoso Updated"
}

Response: 200 OK
```

#### Delete Siswa
```
DELETE /api/v1/siswa/:id
Authorization: Bearer <token>

Response: 200 OK
```

### Kelas Management

#### Buat Kelas
```
POST /api/v1/kelas
Authorization: Bearer <token>
Content-Type: application/json

{
  "nama_kelas": "XI IPA 1"
}

Response: 201 Created
```

#### Get Kelas
```
GET /api/v1/kelas/:id
Authorization: Bearer <token>

Response: 200 OK
```

#### Update Kelas
```
PUT /api/v1/kelas/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "nama_kelas": "XI IPA 1 Updated"
}

Response: 200 OK
```

#### Delete Kelas
```
DELETE /api/v1/kelas/:id
Authorization: Bearer <token>

Response: 200 OK
```

### Wali Kelas Management

#### Buat Wali Kelas
```
POST /api/v1/wali_kelas
Authorization: Bearer <token>
Content-Type: application/json

{
  "user_id": 2,
  "kelas_id": 1,
  "tahun_ajaran": "2024/2025"
}

Response: 201 Created
```

#### Get Wali Kelas
```
GET /api/v1/wali_kelas/:id
Authorization: Bearer <token>

Response: 200 OK
```

#### Update Wali Kelas
```
PUT /api/v1/wali_kelas/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "user_id": 3
}

Response: 200 OK
```

#### Delete Wali Kelas
```
DELETE /api/v1/wali_kelas/:id
Authorization: Bearer <token>

Response: 200 OK
```

### Pengaturan Sistem

#### Buat Pengaturan Sistem
```
POST /api/v1/pengaturan-sistem
Authorization: Bearer <token>
Content-Type: application/json

{
  "hari": 1,
  "jam_masuk_start": "07:00:00",
  "jam_masuk_end": "08:00:00",
  "jam_pulang_start": "14:00:00",
  "jam_pulang_end": "15:00:00"
}

Response: 201 Created
```

#### Get Pengaturan Sistem by Hari
```
GET /api/v1/pengaturan-sistem/:hari
Authorization: Bearer <token>

Response: 200 OK
```

#### Get All Pengaturan Sistem
```
GET /api/v1/pengaturan-sistem
Authorization: Bearer <token>

Response: 200 OK
```

#### Update Pengaturan Sistem
```
PUT /api/v1/pengaturan-sistem/:hari
Authorization: Bearer <token>
Content-Type: application/json

{
  "jam_masuk_start": "07:15:00"
}

Response: 200 OK
```

### Kalender Akademik

#### Buat Kalender Akademik
```
POST /api/v1/kalender-akademik
Authorization: Bearer <token>
Content-Type: application/json

{
  "tanggal": "2024-01-01",
  "keterangan": "Libur Tahun Baru",
  "is_libur": true
}

Response: 201 Created
```

#### Get Kalender Akademik by Tanggal
```
GET /api/v1/kalender-akademik/:tanggal
Authorization: Bearer <token>

Response: 200 OK
```

#### Update Kalender Akademik
```
PUT /api/v1/kalender-akademik/:tanggal
Authorization: Bearer <token>
Content-Type: application/json

{
  "keterangan": "Libur Nasional"
}

Response: 200 OK
```

### Pengajuan Izin

#### Buat Pengajuan Izin
```
POST /api/v1/pengajuan-izin
Authorization: Bearer <token>
Content-Type: application/json

{
  "siswa_id": 5,
  "kategori": "sakit",
  "tanggal_mulai": "2024-01-15",
  "tanggal_selesai": "2024-01-16",
  "alasan_detail": "Demam tinggi"
}

Response: 201 Created
```

#### Approve/Reject Pengajuan Izin
```
PUT /api/v1/pengajuan-izin/:id/approve
Authorization: Bearer <token>
Content-Type: application/json

{
  "status": "approved"
}

Response: 200 OK
```

Untuk reject:
```json
{
  "status": "rejected",
  "alasan_penolakan_admin": "Dokter tidak dilengkapi"
}
```

### Reports

#### Get Laporan Absensi
```
GET /api/v1/reports/absensi?siswa_id=5&tanggal_start=2024-01-01&tanggal_end=2024-01-31
Authorization: Bearer <token>

Response: 200 OK
{
  "total_records": 20,
  "data": [
    {
      "id": 1,
      "siswa_id": 5,
      "tanggal": "2024-01-15",
      "jam_masuk": "07:30:00",
      "jam_pulang": "14:30:00",
      "status_kehadiran": "hadir",
      "entry_mode": "rfid"
    }
  ]
}
```

Query parameters:
- `siswa_id` (optional) - Filter by student ID
- `kelas_id` (optional) - Filter by class ID
- `tanggal_start` (optional) - Start date (YYYY-MM-DD)
- `tanggal_end` (optional) - End date (YYYY-MM-DD)

## 🗄️ Struktur Database

### Tabel Users
```sql
- id (PK)
- name
- email (UNIQUE)
- password (hashed)
- role (admin, wali_kelas, guru_bk)
- created_at, updated_at
```

### Tabel Kelas
```sql
- id (PK)
- nama_kelas
- created_at, updated_at
```

### Tabel Wali Kelas
```sql
- id (PK)
- user_id (FK)
- kelas_id (FK)
- tahun_ajaran
- created_at, updated_at
```

### Tabel Siswa
```sql
- id (PK)
- nis (UNIQUE)
- nisn (UNIQUE)
- nama
- gender (L/P)
- nama_ortu
- uid_kartu (UNIQUE)
- status_aktif (boolean)
- created_at, updated_at
```

### Tabel Riwayat Kelas Siswa
```sql
- id (PK)
- siswa_id (FK)
- kelas_id (FK)
- tahun_ajaran
- created_at, updated_at
```

### Tabel Pengaturan Sistem
```sql
- id (PK)
- hari (1-7)
- jam_masuk_start (TIME)
- jam_masuk_end (TIME)
- jam_pulang_start (TIME)
- jam_pulang_end (TIME)
- created_at, updated_at
```

### Tabel Kalender Akademik
```sql
- id (PK)
- tanggal (DATE, UNIQUE)
- keterangan
- is_libur (boolean)
- created_at, updated_at
```

### Tabel Pengajuan Izin
```sql
- id (PK)
- siswa_id (FK)
- kategori (izin/sakit)
- tanggal_mulai (DATE)
- tanggal_selesai (DATE)
- alasan_detail
- alasan_penolakan_admin
- approved_by (FK users)
- status (pending/approved/rejected)
- created_at, updated_at
```

### Tabel Absensi
```sql
- id (PK)
- siswa_id (FK)
- tanggal (DATE)
- jam_masuk (TIME)
- jam_pulang (TIME)
- status_kehadiran (hadir/terlambat/izin/sakit/alpa)
- entry_mode (rfid/manual)
- keterangan
- pengajuan_izin_id (FK)
- created_at, updated_at
- UNIQUE(siswa_id, tanggal)
```

### Tabel Log Aktivitas User
```sql
- id (PK)
- user_id (FK)
- role
- detail_perubahan
- created_at
```

## 📁 Struktur Kode

```
absensi-sekolah-backend/
├── config/              # Database & environment configuration
│   ├── config.go
│   └── database.go
├── models/              # Data structures & DTOs
│   ├── user.go
│   ├── kelas.go
│   ├── wali_kelas.go
│   ├── siswa.go
│   ├── riwayat_kelas_siswa.go
│   ├── pengaturan_sistem.go
│   ├── kalender_akademik.go
│   ├── pengajuan_izin.go
│   ├── absensi.go
│   └── log_aktivitas_user.go
├── repository/          # Database queries
│   ├── user_repository.go
│   ├── siswa_repository.go
│   ├── kelas_repository.go
│   ├── wali_kelas_repository.go
│   ├── absensi_repository.go
│   ├── pengajuan_izin_repository.go
│   └── sistem_repository.go
├── services/            # Business logic
│   ├── auth_service.go
│   ├── scan_service.go
│   ├── absensi_service.go
│   └── pengajuan_izin_service.go
├── controllers/         # HTTP handlers
│   ├── auth_controller.go
│   ├── scan_controller.go
│   ├── siswa_controller.go
│   ├── kelas_controller.go
│   ├── wali_kelas_controller.go
│   ├── pengaturan_sistem_controller.go
│   ├── kalender_akademik_controller.go
│   ├── izin_controller.go
│   └── absensi_controller.go
├── middleware/          # JWT & role middleware
│   └── jwt_middleware.go
├── routes/              # Route definitions
│   └── routes.go
├── main.go              # Application entry point
├── go.mod               # Go module definition
├── go.sum               # Go dependencies checksums
├── .env                 # Environment variables (git ignored)
├── .env.example         # Example environment variables
└── README.md            # This file
```

## 🔐 Keamanan

- Password di-hash menggunakan bcrypt
- JWT token berlaku 24 jam
- Perlu role-based access control di middleware (bisa dikembangkan)
- Gunakan HTTPS di production
- Ubah JWT_SECRET di production dengan value yang kuat

## 🧪 Testing

### Menggunakan cURL

Login dan dapatkan token:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password123"}'
```

Gunakan token di request berikutnya:
```bash
curl -X GET http://localhost:8080/api/v1/siswa/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### Menggunakan Postman

1. Import API ke Postman
2. Set Authorization type menjadi Bearer Token
3. Paste token dari response login
4. Gunakan endpoints yang tersedia

## 📝 Contoh Data

### Inisialisasi User Admin

Anda perlu membuat user admin secara manual di database:

```sql
INSERT INTO users (name, email, password, role, created_at, updated_at)
VALUES ('Admin', 'admin@example.com', '$2a$10$...', 'admin', NOW(), NOW());
```

Password hash dapat dibuat menggunakan bcrypt generator online atau kode Go.

## 🐛 Troubleshooting

### Error: "database configuration missing in environment"
- Pastikan file `.env` sudah dibuat dan terisi dengan benar
- Verifikasi semua variabel database ada di `.env`

### Error: "failed to connect to database"
- Pastikan PostgreSQL server berjalan
- Verifikasi kredensial database di `.env`
- Pastikan database `absensi_sekolah` sudah dibuat

### Error: "invalid token"
- Token mungkin sudah expired (24 jam)
- Login kembali untuk mendapatkan token baru
- Verifikasi JWT_SECRET sudah diset di `.env`

## 🚀 Deploy ke Production

1. Set environment variables dengan nilai production
2. Build aplikasi: `go build -o absensi-backend .`
3. Gunakan systemd atau docker untuk menjalankan service
4. Setup reverse proxy (nginx) dengan SSL
5. Setup monitoring dan logging

## 📖 Dokumentasi Lanjutan

- [GORM Documentation](https://gorm.io/)
- [Gin Framework Documentation](https://gin-gonic.com/)
- [JWT Go Documentation](https://github.com/golang-jwt/jwt)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

## 📄 Lisensi

MIT License

## 👥 Kontribusi

Kontribusi, bug reports, dan feature requests welcome!

## 📞 Support

Untuk pertanyaan dan support, hubungi tim development.

---

**Happy Coding! 🚀**
