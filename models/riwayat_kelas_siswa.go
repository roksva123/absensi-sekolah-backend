package models

import (
    "time"
    "gorm.io/gorm"
)

type RiwayatKelasSiswa struct {
    ID          uint           `gorm:"primaryKey"`
    SiswaID     uint           `gorm:"not null"`
    KelasID     uint           `gorm:"not null"`
    TahunAjaran string         `gorm:"size:20;not null"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
