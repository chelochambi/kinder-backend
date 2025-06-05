package handlers

import (
	"encoding/json"
	"kinder-backend/config"
	"kinder-backend/internal/models"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decodificar el JSON enviado en la solicitud
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Guardar el usuario en la base de datos
	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el usuario creado
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
