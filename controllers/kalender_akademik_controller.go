package controllers

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateKalenderAkademik creates a new academic calendar entry
func CreateKalenderAkademik(c *gin.Context) {
	var req models.CreateKalenderAkademikRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse date
	tanggal, err := time.Parse("2006-01-02", req.Tanggal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tanggal format (use YYYY-MM-DD)"})
		return
	}

	kalender := &models.KalenderAkademik{
		Tanggal:    tanggal,
		Keterangan: req.Keterangan,
		IsLibur:    req.IsLibur,
	}

	if err := repository.CreateKalenderAkademik(kalender); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create academic calendar entry"})
		return
	}

	c.JSON(http.StatusCreated, kalender)
}

// GetKalenderAkademikByTanggal retrieves academic calendar by date
func GetKalenderAkademikByTanggal(c *gin.Context) {
	tanggalStr := c.Param("tanggal")
	tanggal, err := time.Parse("2006-01-02", tanggalStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tanggal format (use YYYY-MM-DD)"})
		return
	}

	kalender, err := repository.GetKalenderAkademikByTanggal(tanggal)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "academic calendar entry not found"})
		return
	}

	c.JSON(http.StatusOK, kalender)
}

// UpdateKalenderAkademik updates an academic calendar entry
func UpdateKalenderAkademik(c *gin.Context) {
	tanggalStr := c.Param("tanggal")
	tanggal, err := time.Parse("2006-01-02", tanggalStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tanggal format (use YYYY-MM-DD)"})
		return
	}

	kalender, err := repository.GetKalenderAkademikByTanggal(tanggal)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "academic calendar entry not found"})
		return
	}

	if err := c.ShouldBindJSON(&kalender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdateKalenderAkademik(kalender); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update academic calendar entry"})
		return
	}

	c.JSON(http.StatusOK, kalender)
}
