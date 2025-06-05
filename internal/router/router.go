package router

import (
	"net/http"

	"database/sql"

	"github.com/chelochambi/kinder-backend/internal/handler"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	// Endpoints existentes
	r.HandleFunc("/api/tipo_estado", handler.ListarTipoEstado(db)).Methods("GET")

	// 🆕 Endpoint para usuarios
	r.HandleFunc("/api/usuarios", handler.ListarUsuariosHandler(db)).Methods("GET")

	return r
}

