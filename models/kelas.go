package models

import (
    "time"
    "gorm.io/gorm"
)

type Kelas struct {
    ID        uint           `gorm:"primaryKey"`
    NamaKelas string         `gorm:"size:100;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
