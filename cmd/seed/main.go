package main

import (
	"fmt"
	"log"

	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
	"golang.org/x/crypto/bcrypt"
)

// SeedDatabase creates initial data for the system
func SeedDatabase() error {
	// Check if users already exist
	var userCount int64
	if err := config.DB.Model(&models.User{}).Count(&userCount).Error; err != nil {
		return err
	}

	if userCount > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	log.Println("Seeding database...")

	// Create admin user
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := &models.User{
		Name:     "Admin User",
		Email:    "admin@example.com",
		Password: string(adminPassword),
		Role:     "admin",
	}

	if err := config.DB.Create(admin).Error; err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}
	log.Printf("✅ Created admin user (email: admin@example.com, password: admin123)")

	// Create teacher user
	teacherPassword, _ := bcrypt.GenerateFromPassword([]byte("teacher123"), bcrypt.DefaultCost)
	teacher := &models.User{
		Name:     "Guru Budi",
		Email:    "guru@example.com",
		Password: string(teacherPassword),
		Role:     "wali_kelas",
	}

	if err := config.DB.Create(teacher).Error; err != nil {
		return fmt.Errorf("failed to create teacher user: %w", err)
	}
	log.Printf("✅ Created teacher user (email: guru@example.com, password: teacher123)")

	// Create classes
	classes := []models.Kelas{
		{NamaKelas: "XI IPA 1"},
		{NamaKelas: "XI IPA 2"},
		{NamaKelas: "XI IPS 1"},
		{NamaKelas: "XII IPA 1"},
	}

	for _, kelas := range classes {
		if err := config.DB.Create(&kelas).Error; err != nil {
			return fmt.Errorf("failed to create class: %w", err)
		}
	}
	log.Printf("✅ Created %d classes", len(classes))

	// Create sample students
	students := []models.Siswa{
		{
			NIS:      "12345",
			NISN:     "0012345678901",
			Nama:     "Budi Santoso",
			Gender:   "L",
			NamaOrtu: "Santoso",
			UIDKartu: "12345ABC",
		},
		{
			NIS:      "12346",
			NISN:     "0012345678902",
			Nama:     "Siti Nurhaliza",
			Gender:   "P",
			NamaOrtu: "Halim",
			UIDKartu: "12346DEF",
		},
		{
			NIS:      "12347",
			NISN:     "0012345678903",
			Nama:     "Ahmad Hidayat",
			Gender:   "L",
			NamaOrtu: "Hidayat",
			UIDKartu: "12347GHI",
		},
	}

	for _, siswa := range students {
		if err := config.DB.Create(&siswa).Error; err != nil {
			return fmt.Errorf("failed to create student: %w", err)
		}
	}
	log.Printf("✅ Created %d students", len(students))

	// Create system settings (jam masuk/pulang for each day)
	settings := []models.PengaturanSistem{
		// Monday to Friday
		{
			Hari:            1,
			JamMasukStart:   parseTime("07:00:00"),
			JamMasukEnd:     parseTime("08:00:00"),
			JamPulangStart:  parseTime("14:00:00"),
			JamPulangEnd:    parseTime("15:00:00"),
		},
		{
			Hari:            2,
			JamMasukStart:   parseTime("07:00:00"),
			JamMasukEnd:     parseTime("08:00:00"),
			JamPulangStart:  parseTime("14:00:00"),
			JamPulangEnd:    parseTime("15:00:00"),
		},
		{
			Hari:            3,
			JamMasukStart:   parseTime("07:00:00"),
			JamMasukEnd:     parseTime("08:00:00"),
			JamPulangStart:  parseTime("14:00:00"),
			JamPulangEnd:    parseTime("15:00:00"),
		},
		{
			Hari:            4,
			JamMasukStart:   parseTime("07:00:00"),
			JamMasukEnd:     parseTime("08:00:00"),
			JamPulangStart:  parseTime("14:00:00"),
			JamPulangEnd:    parseTime("15:00:00"),
		},
		{
			Hari:            5,
			JamMasukStart:   parseTime("07:00:00"),
			JamMasukEnd:     parseTime("08:00:00"),
			JamPulangStart:  parseTime("14:00:00"),
			JamPulangEnd:    parseTime("15:00:00"),
		},
		{
			Hari:            6,
			JamMasukStart:   parseTime("08:00:00"),
			JamMasukEnd:     parseTime("09:00:00"),
			JamPulangStart:  parseTime("12:00:00"),
			JamPulangEnd:    parseTime("13:00:00"),
		},
		// Sunday (no class)
		{
			Hari:            7,
			JamMasukStart:   parseTime("00:00:00"),
			JamMasukEnd:     parseTime("00:00:00"),
			JamPulangStart:  parseTime("00:00:00"),
			JamPulangEnd:    parseTime("00:00:00"),
		},
	}

	for _, setting := range settings {
		if err := config.DB.Create(&setting).Error; err != nil {
			return fmt.Errorf("failed to create system setting: %w", err)
		}
	}
	log.Printf("✅ Created %d system settings", len(settings))

	log.Println("✅ Database seeding complete!")
	return nil
}

func parseTime(timeStr string) interface{} {
	// This will be parsed by database driver as TIME type
	return timeStr
}
