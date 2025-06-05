package router

import (
	"database/sql"
	"net/http"

	"github.com/chelochambi/kinder-backend/internal/handler"
	"github.com/chelochambi/kinder-backend/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	// Endpoint público (sin protección)
	r.HandleFunc("/api/tipo_estado", handler.ListarTipoEstado(db)).Methods("GET")

	// Endpoint protegido con JWT
	r.Handle("/api/usuarios", middleware.AuthMiddleware(http.HandlerFunc(handler.ListarUsuariosHandler(db)))).Methods("GET")

	// Endpoint protegido con JWT para obtener un usuario específico
	r.Handle("/auth/me", middleware.AuthMiddleware(http.HandlerFunc(handler.AuthMeHandler(db)))).Methods("GET")

	// Endpoint para login, abierto al público
	r.HandleFunc("/auth/login", handler.LoginHandler(db)).Methods("POST")

	return r
}
