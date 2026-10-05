package models

import (
	"time"
	"gorm.io/gorm"
)

type PengajuanIzin struct {
	ID                   uint           `gorm:"primaryKey"`
	SiswaID              uint           `gorm:"not null;index"`
	Kategori             string         `gorm:"type:VARCHAR(20);not null"` // izin, sakit
	TanggalMulai         time.Time      `gorm:"not null;type:date"`
	TanggalSelesai       time.Time      `gorm:"not null;type:date"`
	AlasinDetail         string         `gorm:"type:text;not null"`
	AlasinPenolakanAdmin string         `gorm:"type:text"`
	ApprovedBy           *uint          `gorm:"index"`
	Status               string         `gorm:"type:VARCHAR(20);not null;default:'pending'"` // pending, approved, rejected
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

type PengajuanIzinResponse struct {
	ID                   uint   `json:"id"`
	SiswaID              uint   `json:"siswa_id"`
	Kategori             string `json:"kategori"`
	TanggalMulai         string `json:"tanggal_mulai"`
	TanggalSelesai       string `json:"tanggal_selesai"`
	AlasinDetail         string `json:"alasan_detail"`
	AlasinPenolakanAdmin string `json:"alasan_penolakan_admin"`
	ApprovedBy           *uint  `json:"approved_by"`
	Status               string `json:"status"`
}

type CreatePengajuanIzinRequest struct {
	SiswaID        uint   `json:"siswa_id" binding:"required"`
	Kategori       string `json:"kategori" binding:"required,oneof=izin sakit"`
	TanggalMulai   string `json:"tanggal_mulai" binding:"required"`
	TanggalSelesai string `json:"tanggal_selesai" binding:"required"`
	AlasinDetail   string `json:"alasan_detail" binding:"required"`
}

type ApprovePengajuanIzinRequest struct {
	Status                string `json:"status" binding:"required,oneof=approved rejected"`
	AlasinPenolakanAdmin  string `json:"alasan_penolakan_admin"`
}
