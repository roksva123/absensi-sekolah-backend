package repository

import (
	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
)

// CreateWaliKelas creates a new wali kelas
func CreateWaliKelas(wali *models.WaliKelas) error {
	return config.DB.Create(wali).Error
}

// GetWaliKelasByID retrieves a wali kelas by ID
func GetWaliKelasByID(id uint) (*models.WaliKelas, error) {
	var wali models.WaliKelas
	err := config.DB.First(&wali, id).Error
	return &wali, err
}

// UpdateWaliKelas updates a wali kelas
func UpdateWaliKelas(wali *models.WaliKelas) error {
	return config.DB.Save(wali).Error
}

// DeleteWaliKelas soft-deletes a wali kelas
func DeleteWaliKelas(id uint) error {
	return config.DB.Delete(&models.WaliKelas{}, id).Error
}

// GetWaliKelasByKelas retrieves wali kelas for a class
func GetWaliKelasByKelas(kelasID uint) (*models.WaliKelas, error) {
	var wali models.WaliKelas
	err := config.DB.Where("kelas_id = ?", kelasID).First(&wali).Error
	return &wali, err
}

// GetWaliKelasByUser retrieves all wali kelas for a user
func GetWaliKelasByUser(userID uint) ([]models.WaliKelas, error) {
	var wali []models.WaliKelas
	err := config.DB.Where("user_id = ?", userID).Find(&wali).Error
	return wali, err
}
