package repository

import (
	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
)

// GetSiswaByID retrieves a student by ID
func GetSiswaByID(id uint) (*models.Siswa, error) {
	var siswa models.Siswa
	err := config.DB.First(&siswa, id).Error
	return &siswa, err
}

// GetSiswaByUID retrieves a student by RFID card UID
func GetSiswaByUID(uid string) (*models.Siswa, error) {
	var siswa models.Siswa
	err := config.DB.Where("uid_kartu = ? AND status_aktif = true", uid).First(&siswa).Error
	return &siswa, err
}

// GetSiswaByNIS retrieves a student by NIS
func GetSiswaByNIS(nis string) (*models.Siswa, error) {
	var siswa models.Siswa
	err := config.DB.Where("nis = ?", nis).First(&siswa).Error
	return &siswa, err
}

// CreateSiswa creates a new student record
func CreateSiswa(siswa *models.Siswa) error {
	return config.DB.Create(siswa).Error
}

// UpdateSiswa updates a student record
func UpdateSiswa(siswa *models.Siswa) error {
	return config.DB.Save(siswa).Error
}

// DeleteSiswa soft-deletes a student
func DeleteSiswa(id uint) error {
	return config.DB.Delete(&models.Siswa{}, id).Error
}

// GetAllSiswa retrieves all active students
func GetAllSiswa(page, pageSize int) ([]models.Siswa, error) {
	var siswa []models.Siswa
	offset := (page - 1) * pageSize
	err := config.DB.Where("status_aktif = true").Offset(offset).Limit(pageSize).Find(&siswa).Error
	return siswa, err
}

// GetSiswaByKelas retrieves students in a specific class
func GetSiswaByKelas(kelasID uint) ([]models.Siswa, error) {
	var siswa []models.Siswa
	err := config.DB.
		Joins("JOIN riwayat_kelas_siswa ON siswa.id = riwayat_kelas_siswa.siswa_id").
		Where("riwayat_kelas_siswa.kelas_id = ? AND siswa.status_aktif = true", kelasID).
		Find(&siswa).Error
	return siswa, err
}
