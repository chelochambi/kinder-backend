package router

import (
	"database/sql"
	"net/http"

	"github.com/chelochambi/kinder-backend/internal/handler"
	"github.com/chelochambi/kinder-backend/internal/middleware"
	"github.com/chelochambi/kinder-backend/internal/repository"
	"github.com/chelochambi/kinder-backend/internal/service"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	// Público
	r.HandleFunc("/auth/login", handler.LoginHandler(db)).Methods("POST")
	r.HandleFunc("/api/tipo_estado", handler.ListarTipoEstado(db)).Methods("GET")

	// Protegido - Auth Middleware
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	// ---------- Usuarios ----------
	// Inicializamos repositorio, servicio y handler
	usuarioRepo := repository.NewUsuarioRepository(db)
	usuarioService := service.NewUsuarioService(usuarioRepo)
	usuarioHandler := handler.NewUsuarioHandler(usuarioService)

	api.HandleFunc("/usuarios", usuarioHandler.Listar).Methods("GET")
	api.HandleFunc("/usuarios", usuarioHandler.Crear).Methods("POST")
	api.HandleFunc("/usuarios/{id:[0-9]+}", usuarioHandler.Actualizar).Methods("PUT")
	api.HandleFunc("/usuarios/{id:[0-9]+}/estado", usuarioHandler.CambiarEstado).Methods("PUT")

	return r
}
