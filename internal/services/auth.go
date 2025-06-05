package services

import (
	"encoding/json"
	"fmt"
	"kinder-backend/config"
	"kinder-backend/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Definir la clave secreta para firmar los tokens
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY")) // Usualmente guardada en .env

// Estructura del JWT Claims
type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

// Función para generar el token JWT
func GenerateJWT(user models.User) (string, error) {
	// Crear los claims del JWT
	claims := &Claims{
		Username: user.Username,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Expira en 24 horas
			Issuer:    "kinder-backend",
		},
	}

	// Crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error al firmar el token:", err)
		return "", err
	}

	return tokenString, nil
}

// Función para verificar el token JWT
func ValidateToken(tokenStr string) (*Claims, error) {
	// Parsear el token
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el método de firma sea válido
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inválido")
		}
		return jwtKey, nil
	})

	// Validar el token
	if err != nil {
		return nil, err
	}

	// Verificar que los claims sean válidos
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return claims, nil
}

// Función para autenticar al usuario con email y contraseña
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decodificar el JSON recibido
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Buscar el usuario en la base de datos por su email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
		return
	}

	// Verificar que la contraseña coincida
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
		return
	}

	// Generar el JWT
	token, err := GenerateJWT(user)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	// Responder con el token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
