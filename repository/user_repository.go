package repository

import (
	"absensi-sekolah-backend/config"
	"absensi-sekolah-backend/models"
)

// CreateUser creates a new user
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// GetUserByID retrieves a user by ID
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// UpdateUser updates a user
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// GetAllUsers retrieves all users with pagination
func GetAllUsers(page, pageSize int) ([]models.User, error) {
	var users []models.User
	offset := (page - 1) * pageSize
	err := config.DB.Offset(offset).Limit(pageSize).Find(&users).Error
	return users, err
}

// GetUsersByRole retrieves users by role
func GetUsersByRole(role string) ([]models.User, error) {
	var users []models.User
	err := config.DB.Where("role = ?", role).Find(&users).Error
	return users, err
}
