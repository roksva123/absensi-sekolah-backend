package repository

import (
	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
)

// CreateKelas creates a new class
func CreateKelas(kelas *models.Kelas) error {
	return config.DB.Create(kelas).Error
}

// GetKelasbyID retrieves a class by ID
func GetKelasbyID(id uint) (*models.Kelas, error) {
	var kelas models.Kelas
	err := config.DB.First(&kelas, id).Error
	return &kelas, err
}

// UpdateKelas updates a class
func UpdateKelas(kelas *models.Kelas) error {
	return config.DB.Save(kelas).Error
}

// DeleteKelas soft-deletes a class
func DeleteKelas(id uint) error {
	return config.DB.Delete(&models.Kelas{}, id).Error
}

// GetAllKelas retrieves all classes
func GetAllKelas() ([]models.Kelas, error) {
	var kelas []models.Kelas
	err := config.DB.Find(&kelas).Error
	return kelas, err
}
