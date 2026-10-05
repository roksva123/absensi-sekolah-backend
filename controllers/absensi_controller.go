package controllers

import (
	"absensi-sekolah-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAbsensiReport retrieves absensi report with filters
func GetAbsensiReport(c *gin.Context) {
	siswaIDStr := c.Query("siswa_id")
	kelasIDStr := c.Query("kelas_id")
	tanggalStart := c.Query("tanggal_start")
	tanggalEnd := c.Query("tanggal_end")

	var siswaID, kelasID *uint

	if siswaIDStr != "" {
		id, err := strconv.ParseUint(siswaIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid siswa_id"})
			return
		}
		uintID := uint(id)
		siswaID = &uintID
	}

	if kelasIDStr != "" {
		id, err := strconv.ParseUint(kelasIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid kelas_id"})
			return
		}
		uintID := uint(id)
		kelasID = &uintID
	}

	reports, err := services.GetAbsensiReportService(siswaID, kelasID, tanggalStart, tanggalEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_records": len(reports),
		"data":          reports,
	})
}
