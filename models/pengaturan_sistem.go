package models

import (
	"time"
	"gorm.io/gorm"
)

type PengaturanSistem struct {
	ID          uint           `gorm:"primaryKey"`
	Hari        int            `gorm:"not null"` // 1-7 (Monday to Sunday)
	JamMasukStart time.Time    `gorm:"not null;type:time"`
	JamMasukEnd   time.Time    `gorm:"not null;type:time"`
	JamPulangStart time.Time   `gorm:"not null;type:time"`
	JamPulangEnd  time.Time    `gorm:"not null;type:time"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type PengaturanSistemResponse struct {
	ID             uint   `json:"id"`
	Hari           int    `json:"hari"`
	JamMasukStart  string `json:"jam_masuk_start"`
	JamMasukEnd    string `json:"jam_masuk_end"`
	JamPulangStart string `json:"jam_pulang_start"`
	JamPulangEnd   string `json:"jam_pulang_end"`
}

type CreatePengaturanSistemRequest struct {
	Hari           int    `json:"hari" binding:"required,min=1,max=7"`
	JamMasukStart  string `json:"jam_masuk_start" binding:"required"`
	JamMasukEnd    string `json:"jam_masuk_end" binding:"required"`
	JamPulangStart string `json:"jam_pulang_start" binding:"required"`
	JamPulangEnd   string `json:"jam_pulang_end" binding:"required"`
}
