package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chelochambi/kinder-backend/internal/model"
	"github.com/chelochambi/kinder-backend/internal/service"
	"github.com/gorilla/mux"

	"github.com/go-chi/chi/v5"
)

type UsuarioHandler struct {
	Service *service.UsuarioService
}

func NewUsuarioHandler(s *service.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{Service: s}
}

func (h *UsuarioHandler) RegistrarRutas(r chi.Router) {
	r.Get("/", h.Listar)
	r.Post("/", h.Crear)
	r.Put("/{id}", h.Actualizar)
	r.Put("/{id}/estado", h.CambiarEstado)
}

func (h *UsuarioHandler) Listar(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.Service.ListarUsuarios(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(usuarios)
}

func (h *UsuarioHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var u model.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if err := h.Service.CrearUsuario(r.Context(), &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": "Usuario creado exitosamente",
		"usuario": map[string]interface{}{
			"username": u.Username,
			"email":    u.Email,
		},
	})
}

func (h *UsuarioHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	var u model.Usuario

	// Cambiado para usar mux
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Falta el parámetro ID en la URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido: debe ser un número", http.StatusBadRequest)
		return
	}
	u.ID = id

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.ActualizarUsuario(r.Context(), &u)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Usuario actualizado exitosamente",
	})
}

func (h *UsuarioHandler) CambiarEstado(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var payload struct {
		EstadoID int `json:"estado_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if err := h.Service.CambiarEstado(r.Context(), id, payload.EstadoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
