package controllers

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSiswa creates a new student
func CreateSiswa(c *gin.Context) {
	var req struct {
		NIS      string `json:"nis" binding:"required"`
		NISN     string `json:"nisn" binding:"required"`
		Nama     string `json:"nama" binding:"required"`
		Gender   string `json:"gender" binding:"required,oneof=L P"`
		NamaOrtu string `json:"nama_ortu"`
		UIDKartu string `json:"uid_kartu" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	siswa := &models.Siswa{
		NIS:      req.NIS,
		NISN:     req.NISN,
		Nama:     req.Nama,
		Gender:   req.Gender,
		NamaOrtu: req.NamaOrtu,
		UIDKartu: req.UIDKartu,
		StatusAktif: true,
	}

	if err := repository.CreateSiswa(siswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, siswa)
}

// GetSiswa retrieves a student by ID
func GetSiswa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	siswa, err := repository.GetSiswaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, siswa)
}

// UpdateSiswa updates a student
func UpdateSiswa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	siswa, err := repository.GetSiswaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdateSiswa(siswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update student"})
		return
	}

	c.JSON(http.StatusOK, siswa)
}

// DeleteSiswa soft-deletes a student
func DeleteSiswa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := repository.DeleteSiswa(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}
