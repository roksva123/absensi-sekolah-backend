package services

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"errors"
	"time"
)

// GetAbsensiReportService retrieves absensi report with various filters
func GetAbsensiReportService(siswaID *uint, kelasID *uint, tanggalStart, tanggalEnd string) ([]models.AbsensiResponse, error) {
	var tglStart, tglEnd time.Time
	var err error

	if tanggalStart != "" {
		tglStart, err = time.Parse("2006-01-02", tanggalStart)
		if err != nil {
			return nil, errors.New("invalid start date format")
		}
	}

	if tanggalEnd != "" {
		tglEnd, err = time.Parse("2006-01-02", tanggalEnd)
		if err != nil {
			return nil, errors.New("invalid end date format")
		}
	}

	var absensiList []models.Absensi
	var getErr error

	if siswaID != nil {
		absensiList, getErr = repository.GetAbsensiReport(siswaID, nil, tglStart, tglEnd, nil)
	} else if kelasID != nil {
		absensiList, getErr = repository.GetAbsensiByKelasAndPeriod(*kelasID, tglStart, tglEnd)
	} else {
		absensiList, getErr = repository.GetAbsensiReport(nil, nil, tglStart, tglEnd, nil)
	}

	if getErr != nil {
		return nil, errors.New("failed to retrieve absensi report")
	}

	// Convert to response format
	var responses []models.AbsensiResponse
	for _, abs := range absensiList {
		var jamMasukStr, jamPulangStr *string

		if abs.JamMasuk != nil {
			jm := abs.JamMasuk.Format("15:04:05")
			jamMasukStr = &jm
		}

		if abs.JamPulang != nil {
			jp := abs.JamPulang.Format("15:04:05")
			jamPulangStr = &jp
		}

		responses = append(responses, models.AbsensiResponse{
			ID:              abs.ID,
			SiswaID:         abs.SiswaID,
			Tanggal:         abs.Tanggal.Format("2006-01-02"),
			JamMasuk:        jamMasukStr,
			JamPulang:       jamPulangStr,
			StatusKehadiran: abs.StatusKehadiran,
			EntryMode:       abs.EntryMode,
			Keterangan:      abs.Keterangan,
		})
	}

	return responses, nil
}
