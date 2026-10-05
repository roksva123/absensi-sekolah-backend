package services

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"errors"
	"time"
)

// ScanTapService handles RFID card tap logic
func ScanTapService(uidKartu string) (*models.AbsensiResponse, error) {
	// Find student by RFID card UID
	siswa, err := repository.GetSiswaByUID(uidKartu)
	if err != nil {
		return nil, errors.New("student not found")
	}

	// Check if today is a holiday
	hariLibur, _ := repository.IsHariLibur(time.Now())
	if hariLibur {
		return nil, errors.New("today is a holiday")
	}

	// Get today's absensi record
	tanggalHari := time.Now()
	absensi, err := repository.GetAbsensiBySiswaAndTanggal(siswa.ID, tanggalHari)
	if err != nil {
		// No absensi record yet, create new entry (tap-in)
		return CreateAbsensiTapIn(siswa)
	}

	// Check if already has jam_masuk (tap-in exists)
	if absensi.JamMasuk != nil && absensi.JamPulang == nil {
		// Check if currently in jam_pulang time window
		isJamPulang, err := IsCurrentTimeInJamPulang(tanggalHari)
		if err == nil && isJamPulang {
			// Update jam_pulang
			return UpdateAbsensiTapOut(absensi)
		}
		return nil, errors.New("student already tapped in, not in pulang time")
	}

	if absensi.JamPulang != nil {
		return nil, errors.New("student already completed attendance for today")
	}

	return nil, errors.New("invalid attendance state")
}

// CreateAbsensiTapIn creates a new absensi record for tap-in
func CreateAbsensiTapIn(siswa *models.Siswa) (*models.AbsensiResponse, error) {
	// Get system settings for today
	hariIni := int(time.Now().Weekday())
	if hariIni == 0 {
		hariIni = 7 // Convert Sunday from 0 to 7
	}

	setting, err := repository.GetPengaturanSistemByHari(hariIni)
	if err != nil {
		return nil, errors.New("system setting not found")
	}

	// Determine status: hadir or terlambat
	jamMasukSekarang := time.Now()
	statusKehadiran := "hadir"
	if jamMasukSekarang.After(setting.JamMasukEnd) {
		statusKehadiran = "terlambat"
	}

	// Create absensi record
	jamMasuk := jamMasukSekarang
	absensi := &models.Absensi{
		SiswaID:         siswa.ID,
		Tanggal:         jamMasukSekarang,
		JamMasuk:        &jamMasuk,
		StatusKehadiran: statusKehadiran,
		EntryMode:       "rfid",
	}

	if err := repository.CreateAbsensi(absensi); err != nil {
		return nil, errors.New("failed to create absensi record")
	}

	return &models.AbsensiResponse{
		ID:              absensi.ID,
		SiswaID:         absensi.SiswaID,
		Tanggal:         absensi.Tanggal.Format("2006-01-02"),
		StatusKehadiran: absensi.StatusKehadiran,
		EntryMode:       absensi.EntryMode,
	}, nil
}

// UpdateAbsensiTapOut updates absensi record with tap-out time
func UpdateAbsensiTapOut(absensi *models.Absensi) (*models.AbsensiResponse, error) {
	jamPulang := time.Now()
	absensi.JamPulang = &jamPulang

	if err := repository.UpdateAbsensi(absensi); err != nil {
		return nil, errors.New("failed to update absensi record")
	}

	var jamMasukStr, jamPulangStr string
	if absensi.JamMasuk != nil {
		jamMasukStr = absensi.JamMasuk.Format("15:04:05")
	}
	if absensi.JamPulang != nil {
		jamPulangStr = absensi.JamPulang.Format("15:04:05")
	}

	return &models.AbsensiResponse{
		ID:              absensi.ID,
		SiswaID:         absensi.SiswaID,
		Tanggal:         absensi.Tanggal.Format("2006-01-02"),
		JamMasuk:        &jamMasukStr,
		JamPulang:       &jamPulangStr,
		StatusKehadiran: absensi.StatusKehadiran,
		EntryMode:       absensi.EntryMode,
	}, nil
}

// IsCurrentTimeInJamPulang checks if current time is within pulang time window
func IsCurrentTimeInJamPulang(tanggal time.Time) (bool, error) {
	hariIni := int(tanggal.Weekday())
	if hariIni == 0 {
		hariIni = 7
	}

	setting, err := repository.GetPengaturanSistemByHari(hariIni)
	if err != nil {
		return false, err
	}

	sekarang := time.Now()
	return sekarang.After(setting.JamPulangStart) && sekarang.Before(setting.JamPulangEnd), nil
}
