package services

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"errors"
	"time"
)

// CreatePengajuanIzinService creates a new leave request
func CreatePengajuanIzinService(req *models.CreatePengajuanIzinRequest) (*models.PengajuanIzinResponse, error) {
	// Parse dates
	tglMulai, err := time.Parse("2006-01-02", req.TanggalMulai)
	if err != nil {
		return nil, errors.New("invalid start date format")
	}

	tglSelesai, err := time.Parse("2006-01-02", req.TanggalSelesai)
	if err != nil {
		return nil, errors.New("invalid end date format")
	}

	if tglSelesai.Before(tglMulai) {
		return nil, errors.New("end date must be after or equal to start date")
	}

	// Verify student exists
	_, err = repository.GetSiswaByID(req.SiswaID)
	if err != nil {
		return nil, errors.New("student not found")
	}

	pengajuan := &models.PengajuanIzin{
		SiswaID:        req.SiswaID,
		Kategori:       req.Kategori,
		TanggalMulai:   tglMulai,
		TanggalSelesai: tglSelesai,
		AlasinDetail:   req.AlasinDetail,
		Status:         "pending",
	}

	if err := repository.CreatePengajuanIzin(pengajuan); err != nil {
		return nil, errors.New("failed to create leave request")
	}

	return &models.PengajuanIzinResponse{
		ID:             pengajuan.ID,
		SiswaID:        pengajuan.SiswaID,
		Kategori:       pengajuan.Kategori,
		TanggalMulai:   pengajuan.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: pengajuan.TanggalSelesai.Format("2006-01-02"),
		AlasinDetail:   pengajuan.AlasinDetail,
		Status:         pengajuan.Status,
	}, nil
}

// ApprovePengajuanIzinService approves or rejects a leave request
func ApprovePengajuanIzinService(id uint, approvedBy uint, req *models.ApprovePengajuanIzinRequest) (*models.PengajuanIzinResponse, error) {
	pengajuan, err := repository.GetPengajuanIzinByID(id)
	if err != nil {
		return nil, errors.New("leave request not found")
	}

	if pengajuan.Status != "pending" {
		return nil, errors.New("leave request is not pending")
	}

	pengajuan.Status = req.Status
	pengajuan.ApprovedBy = &approvedBy

	if req.Status == "rejected" {
		pengajuan.AlasinPenolakanAdmin = req.AlasinPenolakanAdmin
	} else if req.Status == "approved" {
		// Automatically create absensi records for approved dates
		if err := CreateAbsensiForApprovedIzin(pengajuan); err != nil {
			return nil, err
		}
	}

	if err := repository.UpdatePengajuanIzin(pengajuan); err != nil {
		return nil, errors.New("failed to update leave request")
	}

	return &models.PengajuanIzinResponse{
		ID:                   pengajuan.ID,
		SiswaID:              pengajuan.SiswaID,
		Kategori:             pengajuan.Kategori,
		TanggalMulai:         pengajuan.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai:       pengajuan.TanggalSelesai.Format("2006-01-02"),
		AlasinDetail:         pengajuan.AlasinDetail,
		AlasinPenolakanAdmin: pengajuan.AlasinPenolakanAdmin,
		ApprovedBy:           pengajuan.ApprovedBy,
		Status:               pengajuan.Status,
	}, nil
}

// CreateAbsensiForApprovedIzin creates absensi records for all dates covered by approved leave request
func CreateAbsensiForApprovedIzin(pengajuan *models.PengajuanIzin) error {
	currentDate := pengajuan.TanggalMulai
	endDate := pengajuan.TanggalSelesai

	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		// Check if this date is already recorded
		_, err := repository.GetAbsensiBySiswaAndTanggal(pengajuan.SiswaID, currentDate)
		if err != nil {
			// Create new absensi record
			absensi := &models.Absensi{
				SiswaID:         pengajuan.SiswaID,
				Tanggal:         currentDate,
				StatusKehadiran: pengajuan.Kategori, // "izin" or "sakit"
				EntryMode:       "manual",
				PengajuanIzinID: &pengajuan.ID,
			}

			if err := repository.CreateAbsensi(absensi); err != nil {
				return err
			}
		}

		currentDate = currentDate.AddDate(0, 0, 1) // Add 1 day
	}

	return nil
}
