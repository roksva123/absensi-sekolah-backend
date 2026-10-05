package repository

import (
	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
	"time"
)

// CreateAbsensi creates a new absensi record
func CreateAbsensi(absensi *models.Absensi) error {
	return config.DB.Create(absensi).Error
}

// GetAbsensiByID retrieves an absensi by ID
func GetAbsensiByID(id uint) (*models.Absensi, error) {
	var absensi models.Absensi
	err := config.DB.First(&absensi, id).Error
	return &absensi, err
}

// GetAbsensiBySiswaAndTanggal retrieves absensi for a student on a specific date
func GetAbsensiBySiswaAndTanggal(siswaID uint, tanggal time.Time) (*models.Absensi, error) {
	var absensi models.Absensi
	err := config.DB.Where("siswa_id = ? AND tanggal = ?", siswaID, tanggal.Format("2006-01-02")).First(&absensi).Error
	return &absensi, err
}

// UpdateAbsensi updates an absensi record
func UpdateAbsensi(absensi *models.Absensi) error {
	return config.DB.Save(absensi).Error
}

// GetAbsensiReport retrieves absensi report with filters
func GetAbsensiReport(siswaID *uint, kelasID *uint, tanggalStart, tanggalEnd time.Time, statusKehadiran *string) ([]models.Absensi, error) {
	var absensi []models.Absensi
	query := config.DB

	if siswaID != nil {
		query = query.Where("siswa_id = ?", *siswaID)
	}

	if tanggalStart.Year() > 1 {
		query = query.Where("tanggal >= ?", tanggalStart.Format("2006-01-02"))
	}

	if tanggalEnd.Year() > 1 {
		query = query.Where("tanggal <= ?", tanggalEnd.Format("2006-01-02"))
	}

	if statusKehadiran != nil && *statusKehadiran != "" {
		query = query.Where("status_kehadiran = ?", *statusKehadiran)
	}

	err := query.Order("tanggal DESC").Find(&absensi).Error
	return absensi, err
}

// GetAbsensiByKelasAndPeriod retrieves absensi for a class during a period
func GetAbsensiByKelasAndPeriod(kelasID uint, tanggalStart, tanggalEnd time.Time) ([]models.Absensi, error) {
	var absensi []models.Absensi
	err := config.DB.
		Joins("JOIN siswa ON absensi.siswa_id = siswa.id").
		Joins("JOIN riwayat_kelas_siswa ON siswa.id = riwayat_kelas_siswa.siswa_id").
		Where("riwayat_kelas_siswa.kelas_id = ? AND absensi.tanggal >= ? AND absensi.tanggal <= ?",
			kelasID, tanggalStart.Format("2006-01-02"), tanggalEnd.Format("2006-01-02")).
		Find(&absensi).Error
	return absensi, err
}
