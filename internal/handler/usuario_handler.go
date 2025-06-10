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

	// Primero decodificamos el cuerpo del JSON
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Ahora validamos los campos obligatorios
	if u.Username == "" || u.Email == "" {
		http.Error(w, "Faltan campos obligatorios", http.StatusBadRequest)
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
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var payload struct {
		EstadoID       int `json:"estado_id"`
		ActualizadoPor int `json:"actualizado_por"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.CambiarEstado(r.Context(), id, payload.EstadoID, payload.ActualizadoPor)
	if err != nil {
		if err.Error() == "estado no válido o inactivo" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Estado del usuario actualizado exitosamente",
	})
}
