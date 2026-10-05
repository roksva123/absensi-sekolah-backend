package repository

import (
	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
)

// CreatePengajuanIzin creates a new leave request
func CreatePengajuanIzin(pengajuan *models.PengajuanIzin) error {
	return config.DB.Create(pengajuan).Error
}

// GetPengajuanIzinByID retrieves a leave request by ID
func GetPengajuanIzinByID(id uint) (*models.PengajuanIzin, error) {
	var pengajuan models.PengajuanIzin
	err := config.DB.First(&pengajuan, id).Error
	return &pengajuan, err
}

// UpdatePengajuanIzin updates a leave request
func UpdatePengajuanIzin(pengajuan *models.PengajuanIzin) error {
	return config.DB.Save(pengajuan).Error
}

// GetPengajuanIzinBySiswa retrieves all leave requests for a student
func GetPengajuanIzinBySiswa(siswaID uint) ([]models.PengajuanIzin, error) {
	var pengajuan []models.PengajuanIzin
	err := config.DB.Where("siswa_id = ?", siswaID).Order("created_at DESC").Find(&pengajuan).Error
	return pengajuan, err
}

// GetPengajuanIzinPending retrieves all pending leave requests
func GetPengajuanIzinPending() ([]models.PengajuanIzin, error) {
	var pengajuan []models.PengajuanIzin
	err := config.DB.Where("status = ?", "pending").Order("created_at ASC").Find(&pengajuan).Error
	return pengajuan, err
}

// GetPengajuanIzinByStatus retrieves leave requests by status
func GetPengajuanIzinByStatus(status string) ([]models.PengajuanIzin, error) {
	var pengajuan []models.PengajuanIzin
	err := config.DB.Where("status = ?", status).Order("created_at DESC").Find(&pengajuan).Error
	return pengajuan, err
}
