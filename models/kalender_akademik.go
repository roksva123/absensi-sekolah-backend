package models

import (
	"time"
	"gorm.io/gorm"
)

type KalenderAkademik struct {
	ID        uint           `gorm:"primaryKey"`
	Tanggal   time.Time      `gorm:"unique;not null;type:date"`
	Keterangan string        `gorm:"type:text"`
	IsLibur   bool           `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type KalenderAkademikResponse struct {
	ID         uint   `json:"id"`
	Tanggal    string `json:"tanggal"`
	Keterangan string `json:"keterangan"`
	IsLibur    bool   `json:"is_libur"`
}

type CreateKalenderAkademikRequest struct {
	Tanggal    string `json:"tanggal" binding:"required"`
	Keterangan string `json:"keterangan"`
	IsLibur    bool   `json:"is_libur"`
}
