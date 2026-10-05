package controllers

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePengaturanSistem creates a new system setting
func CreatePengaturanSistem(c *gin.Context) {
	var req models.CreatePengaturanSistemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse times
	jamMasukStart, err := time.Parse("15:04:05", req.JamMasukStart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jam_masuk_start format (use HH:MM:SS)"})
		return
	}

	jamMasukEnd, err := time.Parse("15:04:05", req.JamMasukEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jam_masuk_end format (use HH:MM:SS)"})
		return
	}

	jamPulangStart, err := time.Parse("15:04:05", req.JamPulangStart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jam_pulang_start format (use HH:MM:SS)"})
		return
	}

	jamPulangEnd, err := time.Parse("15:04:05", req.JamPulangEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jam_pulang_end format (use HH:MM:SS)"})
		return
	}

	setting := &models.PengaturanSistem{
		Hari:          req.Hari,
		JamMasukStart: jamMasukStart,
		JamMasukEnd:   jamMasukEnd,
		JamPulangStart: jamPulangStart,
		JamPulangEnd:  jamPulangEnd,
	}

	if err := repository.CreatePengaturanSistem(setting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create system setting"})
		return
	}

	c.JSON(http.StatusCreated, setting)
}

// GetPengaturanSistemByHari retrieves system setting by day
func GetPengaturanSistemByHari(c *gin.Context) {
	hari, err := strconv.Atoi(c.Param("hari"))
	if err != nil || hari < 1 || hari > 7 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hari (must be 1-7)"})
		return
	}

	setting, err := repository.GetPengaturanSistemByHari(hari)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "system setting not found"})
		return
	}

	c.JSON(http.StatusOK, setting)
}

// UpdatePengaturanSistem updates a system setting
func UpdatePengaturanSistem(c *gin.Context) {
	hari, err := strconv.Atoi(c.Param("hari"))
	if err != nil || hari < 1 || hari > 7 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hari (must be 1-7)"})
		return
	}

	setting, err := repository.GetPengaturanSistemByHari(hari)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "system setting not found"})
		return
	}

	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdatePengaturanSistem(setting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update system setting"})
		return
	}

	c.JSON(http.StatusOK, setting)
}

// GetAllPengaturanSistem retrieves all system settings
func GetAllPengaturanSistem(c *gin.Context) {
	settings, err := repository.GetAllPengaturanSistem()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve system settings"})
		return
	}

	c.JSON(http.StatusOK, settings)
}
