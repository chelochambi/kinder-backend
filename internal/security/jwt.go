package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("clave_secreta_segura") // reemplazar con variable de entorno en producción

func GenerarJWT(usuarioID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  usuarioID,
		"name": username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
