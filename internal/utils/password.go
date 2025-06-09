package utils

import "golang.org/x/crypto/bcrypt"

func CompararPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
