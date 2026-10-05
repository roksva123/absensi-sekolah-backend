package models

import (
	"time"
	"gorm.io/gorm"
)

type Absensi struct {
	ID                  uint           `gorm:"primaryKey"`
	SiswaID             uint           `gorm:"not null;index:,composite:idx_siswa_tanggal"`
	Tanggal             time.Time      `gorm:"not null;index:,composite:idx_siswa_tanggal;type:date"`
	JamMasuk            *time.Time     `gorm:"type:time"`
	JamPulang           *time.Time     `gorm:"type:time"`
	StatusKehadiran     string         `gorm:"type:VARCHAR(20);not null;default:'hadir'"` // hadir, terlambat, izin, sakit, alpa
	EntryMode           string         `gorm:"type:VARCHAR(20);not null;default:'rfid'"` // rfid, manual
	Keterangan          string         `gorm:"type:text"`
	PengajuanIzinID     *uint          `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type AbsensiResponse struct {
	ID              uint       `json:"id"`
	SiswaID         uint       `json:"siswa_id"`
	Tanggal         string     `json:"tanggal"`
	JamMasuk        *string    `json:"jam_masuk"`
	JamPulang       *string    `json:"jam_pulang"`
	StatusKehadiran string     `json:"status_kehadiran"`
	EntryMode       string     `json:"entry_mode"`
	Keterangan      string     `json:"keterangan"`
}

type CreateAbsensiRequest struct {
	SiswaID      uint   `json:"siswa_id" binding:"required"`
	Tanggal      string `json:"tanggal" binding:"required"`
	JamMasuk     string `json:"jam_masuk"`
	JamPulang    string `json:"jam_pulang"`
	Keterangan   string `json:"keterangan"`
}
