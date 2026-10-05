package controllers

import (
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePengajuanIzin creates a new leave request
func CreatePengajuanIzin(c *gin.Context) {
	var req models.CreatePengajuanIzinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.CreatePengajuanIzinService(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// ApprovePengajuanIzin approves or rejects a leave request
func ApprovePengajuanIzin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found in context"})
		return
	}

	var req models.ApprovePengajuanIzinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.ApprovePengajuanIzinService(uint(id), userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
