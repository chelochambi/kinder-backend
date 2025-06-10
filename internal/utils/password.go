package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CompararPasswordHash(password string, hash string) error {
	password = strings.TrimSpace(password)
	hash = strings.TrimSpace(hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err // devuelve el error para que el handler lo maneje
	}
	return nil
}
