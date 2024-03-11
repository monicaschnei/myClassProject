package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hasshedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password %w", err)
	}
	return string(hasshedPassword), nil
}

// checkPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func IsValidFormatPassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password lentgh must be greater or equal than 8"
	}

	regexPassword := regexp.MustCompile("[A-Z]")
	if !regexPassword.MatchString(password) {
		return false, "Password must have at least one uppercase letter"
	}

	regexEspecialPassword := regexp.MustCompile(`[!@#~$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
	if !regexEspecialPassword.MatchString(password) {
		return false, "Password must have at least one special character"
	}

	return true, ""
}
