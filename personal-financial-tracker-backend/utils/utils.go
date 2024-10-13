package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ValidateRequestBody(body interface{}, ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(body)

	return err
}

// HashPassword hashes the provided password using bcrypt
func HashPassword(password string) (string, error) {
	// The second argument is the cost factor, 14 is a good balance between security and speed
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plaintext password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // Returns true if passwords match, false otherwise
}
