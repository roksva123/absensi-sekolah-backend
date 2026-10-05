package models

import (
	"time"
	"gorm.io/gorm"
)

type LogAktivitasUser struct {
	ID             uint           `gorm:"primaryKey"`
	UserID         uint           `gorm:"not null;index"`
	Role           string         `gorm:"type:VARCHAR(20);not null"`
	DetailPerubahan string        `gorm:"type:text;not null"`
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type CreateLogAktivitasUserRequest struct {
	UserID          uint   `json:"user_id" binding:"required"`
	DetailPerubahan string `json:"detail_perubahan" binding:"required"`
}
