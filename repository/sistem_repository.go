package repository

import (
	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
	"time"
)

// CreatePengaturanSistem creates a new system setting
func CreatePengaturanSistem(setting *models.PengaturanSistem) error {
	return config.DB.Create(setting).Error
}

// GetPengaturanSistemByHari retrieves system setting by day
func GetPengaturanSistemByHari(hari int) (*models.PengaturanSistem, error) {
	var setting models.PengaturanSistem
	err := config.DB.Where("hari = ?", hari).First(&setting).Error
	return &setting, err
}

// UpdatePengaturanSistem updates a system setting
func UpdatePengaturanSistem(setting *models.PengaturanSistem) error {
	return config.DB.Save(setting).Error
}

// GetAllPengaturanSistem retrieves all system settings
func GetAllPengaturanSistem() ([]models.PengaturanSistem, error) {
	var settings []models.PengaturanSistem
	err := config.DB.Order("hari ASC").Find(&settings).Error
	return settings, err
}

// CreateKalenderAkademik creates a new academic calendar entry
func CreateKalenderAkademik(kalender *models.KalenderAkademik) error {
	return config.DB.Create(kalender).Error
}

// GetKalenderAkademikByTanggal retrieves academic calendar by date
func GetKalenderAkademikByTanggal(tanggal time.Time) (*models.KalenderAkademik, error) {
	var kalender models.KalenderAkademik
	err := config.DB.Where("tanggal = ?", tanggal.Format("2006-01-02")).First(&kalender).Error
	return &kalender, err
}

// UpdateKalenderAkademik updates an academic calendar entry
func UpdateKalenderAkademik(kalender *models.KalenderAkademik) error {
	return config.DB.Save(kalender).Error
}

// IsHariLibur checks if a date is a holiday
func IsHariLibur(tanggal time.Time) (bool, error) {
	kalender, err := GetKalenderAkademikByTanggal(tanggal)
	if err != nil {
		return false, nil // Return false if date not found (normal day)
	}
	return kalender.IsLibur, nil
}
