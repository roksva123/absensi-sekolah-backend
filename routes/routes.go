package routes

import (
    "github.com/gin-gonic/gin"

    "absensi-sekolah-backend/controllers"
    "absensi-sekolah-backend/middleware"
)

// SetupRoutes registers all API routes and middleware.
func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")

    // Public auth routes
    auth := api.Group("/auth")
    {
        auth.POST("/login", controllers.Login)
    }

    // Protected routes (require JWT)
    protected := api.Group("/", middleware.JWTAuth())
    {
        // RFID scan endpoint
        protected.POST("/scan/tap", controllers.ScanTap)

        // Siswa CRUD
        siswa := protected.Group("/siswa")
        {
            siswa.POST("", controllers.CreateSiswa)
            siswa.GET("/:id", controllers.GetSiswa)
            siswa.PUT("/:id", controllers.UpdateSiswa)
            siswa.DELETE("/:id", controllers.DeleteSiswa)
        }

        // Kelas CRUD
        kelas := protected.Group("/kelas")
        {
            kelas.POST("", controllers.CreateKelas)
            kelas.GET("/:id", controllers.GetKelas)
            kelas.PUT("/:id", controllers.UpdateKelas)
            kelas.DELETE("/:id", controllers.DeleteKelas)
        }

        // Wali Kelas CRUD
        wali := protected.Group("/wali_kelas")
        {
            wali.POST("", controllers.CreateWaliKelas)
            wali.GET("/:id", controllers.GetWaliKelas)
            wali.PUT("/:id", controllers.UpdateWaliKelas)
            wali.DELETE("/:id", controllers.DeleteWaliKelas)
        }

        // Pengaturan Sistem CRUD
        pengaturan := protected.Group("/pengaturan-sistem")
        {
            pengaturan.POST("", controllers.CreatePengaturanSistem)
            pengaturan.GET("/:hari", controllers.GetPengaturanSistemByHari)
            pengaturan.PUT("/:hari", controllers.UpdatePengaturanSistem)
            pengaturan.GET("", controllers.GetAllPengaturanSistem)
        }

        // Kalender Akademik CRUD
        kalender := protected.Group("/kalender-akademik")
        {
            kalender.POST("", controllers.CreateKalenderAkademik)
            kalender.GET("/:tanggal", controllers.GetKalenderAkademikByTanggal)
            kalender.PUT("/:tanggal", controllers.UpdateKalenderAkademik)
        }

        // Pengajuan Izin endpoints
        izin := protected.Group("/pengajuan-izin")
        {
            izin.POST("", controllers.CreatePengajuanIzin)
            izin.PUT("/:id/approve", controllers.ApprovePengajuanIzin)
        }

        // Reports
        reports := protected.Group("/reports")
        {
            reports.GET("/absensi", controllers.GetAbsensiReport)
        }
    }
}
