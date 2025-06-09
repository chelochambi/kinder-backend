package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("clave_secreta_segura") // usar variable de entorno real en producción

type ContextKey string

const UsuarioIDKey ContextKey = "usuarioID"

// Middleware para validar el JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token no proporcionado", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Validar método de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Error al leer claims", http.StatusUnauthorized)
			return
		}

		// Extraer ID del usuario desde el claim
		usuarioID := int(claims["sub"].(float64))
		ctx := context.WithValue(r.Context(), UsuarioIDKey, usuarioID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
