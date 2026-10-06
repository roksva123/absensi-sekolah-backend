package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:100;not null"`
	Email     string         `gorm:"size:100;unique;not null"`
	Password  string         `gorm:"size:255;not null"`
	Role      string         `gorm:"type:VARCHAR(20);not null"` // admin, wali_kelas, guru_bk
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
