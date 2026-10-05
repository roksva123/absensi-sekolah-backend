package controllers

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateKelas creates a new class
func CreateKelas(c *gin.Context) {
	var req struct {
		NamaKelas string `json:"nama_kelas" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kelas := &models.Kelas{
		NamaKelas: req.NamaKelas,
	}

	if err := repository.CreateKelas(kelas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create class"})
		return
	}

	c.JSON(http.StatusCreated, kelas)
}

// GetKelas retrieves a class by ID
func GetKelas(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	kelas, err := repository.GetKelasbyID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "class not found"})
		return
	}

	c.JSON(http.StatusOK, kelas)
}

// UpdateKelas updates a class
func UpdateKelas(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	kelas, err := repository.GetKelasbyID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "class not found"})
		return
	}

	if err := c.ShouldBindJSON(&kelas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdateKelas(kelas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update class"})
		return
	}

	c.JSON(http.StatusOK, kelas)
}

// DeleteKelas soft-deletes a class
func DeleteKelas(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := repository.DeleteKelas(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "class deleted successfully"})
}
