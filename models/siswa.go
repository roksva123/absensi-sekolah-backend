package models

import (
    "time"
    "gorm.io/gorm"
)

type Siswa struct {
    ID          uint           `gorm:"primaryKey"`
    NIS         string         `gorm:"size:20;unique;not null"`
    NISN        string         `gorm:"size:20;unique;not null"`
    Nama        string         `gorm:"size:100;not null"`
    Gender      string         `gorm:"type:VARCHAR(1);not null"` // L or P
    NamaOrtu    string         `gorm:"size:100"`
    UIDKartu    string         `gorm:"size:50;unique;not null"`
    StatusAktif bool           `gorm:"default:true"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
