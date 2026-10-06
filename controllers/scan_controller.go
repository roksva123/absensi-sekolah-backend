package controllers

import (
	"absensi-sekolah-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ScanTapRequest represents RFID tap request
type ScanTapRequest struct {
	UIDKartu string `json:"uid_kartu" binding:"required"`
}

// ScanTap handles RFID card tap
func ScanTap(c *gin.Context) {
	var req ScanTapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	absensi, err := services.ScanTapService(req.UIDKartu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    absensi,
	})
}
