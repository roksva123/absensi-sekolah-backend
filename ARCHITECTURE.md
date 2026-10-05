# Architecture Documentation

## System Design

### 1. Layered Architecture

```
┌─────────────────────────────────────┐
│      HTTP Layer (Gin Router)         │
│      - routes.go                     │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Controller Layer                │
│     - Parse HTTP requests           │
│     - Call services                 │
│     - Format responses              │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Service Layer                   │
│     - Business logic                │
│     - Validations                   │
│     - Orchestration                 │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Repository Layer                │
│     - Database queries              │
│     - Data retrieval/storage        │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Database Layer                  │
│     - PostgreSQL                    │
│     - GORM ORM                      │
└─────────────────────────────────────┘
```

### 2. Authentication Flow

```
Client Request
     │
     ▼
POST /api/v1/auth/login
     │
     ▼
auth_controller.Login()
     │
     ▼
auth_service.LoginService()
     │
     ├─ repository.GetUserByEmail()
     │
     ├─ bcrypt.CompareHashAndPassword()
     │
     └─ middleware.GenerateToken()
     │
     ▼
Response with JWT Token
     │
     ▼
Client stores token
     │
     ▼
Subsequent requests with "Authorization: Bearer <token>"
     │
     ▼
middleware.JWTAuth() validates token
```

### 3. RFID Tap Flow

```
Device sends RFID uid_kartu
     │
     ▼
POST /api/v1/scan/tap
     │
     ▼
scan_controller.ScanTap()
     │
     ▼
scan_service.ScanTapService()
     │
     ├─ repository.GetSiswaByUID() - Find student
     │
     ├─ repository.IsHariLibur() - Check if holiday
     │
     ├─ repository.GetAbsensiBySiswaAndTanggal()
     │
     ├─ Check if tap-in or tap-out
     │
     ├─ Compare with pengaturan_sistem (jam_masuk_start/end)
     │
     ├─ Determine status (hadir/terlambat)
     │
     └─ Create/Update absensi record
     │
     ▼
Response with attendance status
```

### 4. Leave Request Approval Flow

```
Student submits leave request
     │
     ▼
POST /api/v1/pengajuan-izin
     │
     ▼
izin_controller.CreatePengajuanIzin()
     │
     ▼
pengajuan_izin_service.CreatePengajuanIzinService()
     │
     ▼
pengajuan_izin record created (status: pending)
     │
     ▼
Admin/Teacher reviews
     │
     ▼
PUT /api/v1/pengajuan-izin/:id/approve
     │
     ▼
izin_controller.ApprovePengajuanIzin()
     │
     ▼
pengajuan_izin_service.ApprovePengajuanIzinService()
     │
     ├─ Update status (approved/rejected)
     │
     ├─ If approved:
     │  └─ pengajuan_izin_service.CreateAbsensiForApprovedIzin()
     │     └─ Create absensi records for all dates
     │
     └─ If rejected:
        └─ Add rejection reason
     │
     ▼
Response with updated leave request
```

## Data Models

### User Model
- Represents system users (admin, teachers, staff)
- Password is hashed with bcrypt
- Role-based access control

### Siswa Model
- Student information
- Linked to RFID card via uid_kartu (unique)
- Can be marked inactive without deletion (soft delete)

### Absensi Model
- Attendance records
- Unique constraint on (siswa_id, tanggal)
- Tracks both tap-in (jam_masuk) and tap-out (jam_pulang)
- Entry mode: RFID or manual
- Status: hadir, terlambat, izin, sakit, alpa

### PengajuanIzin Model
- Leave requests
- Date range based (tanggal_mulai to tanggal_selesai)
- Approval workflow (pending → approved/rejected)
- Approval creates corresponding absensi records

## Security Considerations

1. **Password Security**
   - Bcrypt hashing with salt
   - Never store plaintext passwords

2. **Authentication**
   - JWT tokens valid for 24 hours
   - Token validation on all protected endpoints
   - Bearer token in Authorization header

3. **Authorization**
   - Role-based access control in middleware
   - Can be extended with endpoint-specific role checks

4. **Data Validation**
   - Input validation at controller level
   - Database constraints for data integrity
   - Unique constraints on critical fields

5. **SQL Injection Prevention**
   - GORM parameterized queries
   - No string concatenation in SQL

## Performance Considerations

1. **Database Indexes**
   - Foreign key indexes
   - Composite index on (siswa_id, tanggal) for absensi

2. **Query Optimization**
   - Eager loading where needed
   - Pagination support in list endpoints

3. **Caching Opportunities**
   - System settings (pengaturan_sistem) - rarely changes
   - Holiday calendar (kalender_akademik) - static per year

## Deployment

### Development
- Local PostgreSQL or Docker Compose
- `go run main.go`
- Hot reload with `air` (optional)

### Production
- Docker containerization
- PostgreSQL on managed service
- Environment-specific .env files
- Reverse proxy (Nginx/Apache)
- SSL/TLS certificates
- Monitoring and logging

## Future Enhancements

1. **Role-based endpoint access**
   - Admin-only endpoints
   - Teacher-specific data access

2. **Audit logging**
   - Track all changes via log_aktivitas_user
   - Implement audit middleware

3. **Attendance reports**
   - Enhanced reporting with statistics
   - Export to Excel/PDF

4. **Notification system**
   - Email notifications for approvals
   - SMS alerts for irregularities

5. **API documentation**
   - Swagger/OpenAPI documentation
   - Interactive API explorer

6. **Testing**
   - Unit tests for services
   - Integration tests for API
   - End-to-end tests

7. **Caching layer**
   - Redis for frequently accessed data
   - Token blacklist for logout

8. **Rate limiting**
   - Prevent API abuse
   - Per-user rate limits
