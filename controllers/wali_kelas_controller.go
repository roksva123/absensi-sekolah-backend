package controllers

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateWaliKelas creates a new wali kelas
func CreateWaliKelas(c *gin.Context) {
	var req struct {
		UserID      uint   `json:"user_id" binding:"required"`
		KelasID     uint   `json:"kelas_id" binding:"required"`
		TahunAjaran string `json:"tahun_ajaran" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wali := &models.WaliKelas{
		UserID:      req.UserID,
		KelasID:     req.KelasID,
		TahunAjaran: req.TahunAjaran,
	}

	if err := repository.CreateWaliKelas(wali); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create wali kelas"})
		return
	}

	c.JSON(http.StatusCreated, wali)
}

// GetWaliKelas retrieves a wali kelas by ID
func GetWaliKelas(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	wali, err := repository.GetWaliKelasByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "wali kelas not found"})
		return
	}

	c.JSON(http.StatusOK, wali)
}

// UpdateWaliKelas updates a wali kelas
func UpdateWaliKelas(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	wali, err := repository.GetWaliKelasByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "wali kelas not found"})
		return
	}

	if err := c.ShouldBindJSON(&wali); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdateWaliKelas(wali); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update wali kelas"})
		return
	}

	c.JSON(http.StatusOK, wali)
}

// DeleteWaliKelas soft-deletes a wali kelas
func DeleteWaliKelas(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := repository.DeleteWaliKelas(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete wali kelas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "wali kelas deleted successfully"})
}
