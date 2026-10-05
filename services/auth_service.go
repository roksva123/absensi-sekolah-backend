package services

import (
	"absensi-sekolah-backend/middleware"
	"absensi-sekolah-backend/models"
	"absensi-sekolah-backend/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// LoginService handles user authentication
func LoginService(email, password string) (*models.User, string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

// CreateUserService creates a new user with hashed password
func CreateUserService(user *models.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	return repository.CreateUser(user)
}

// UpdateUserService updates user information
func UpdateUserService(user *models.User) error {
	return repository.UpdateUser(user)
}
