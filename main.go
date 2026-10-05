package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/routes"
)

func main() {
	// Load environment variables from .env (ignore error if file missing)
	_ = godotenv.Load()

	// Initialize DB connection
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run auto-migrations
	if err := runMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create Gin engine
	r := gin.Default()

	// Register API routes
	routes.SetupRoutes(r)

	// Server configuration
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// runMigrations runs all database migrations
func runMigrations() error {
	return config.DB.AutoMigrate(
		&models.User{},
		&models.Kelas{},
		&models.WaliKelas{},
		&models.Siswa{},
		&models.RiwayatKelasSiswa{},
		&models.PengaturanSistem{},
		&models.KalenderAkademik{},
		&models.PengajuanIzin{},
		&models.Absensi{},
		&models.LogAktivitasUser{},
	)
}
