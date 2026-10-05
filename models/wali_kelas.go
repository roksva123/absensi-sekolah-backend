package models

import (
    "time"
    "gorm.io/gorm"
)

type WaliKelas struct {
    ID          uint           `gorm:"primaryKey"`
    UserID      uint           `gorm:"not null"`
    KelasID     uint           `gorm:"not null"`
    TahunAjaran string         `gorm:"size:20;not null"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
