package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chelochambi/kinder-backend/internal/model"
	"github.com/chelochambi/kinder-backend/internal/security"
	"github.com/chelochambi/kinder-backend/internal/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		var u model.Usuario
		query := `
			SELECT u.id, u.username, u.email, u.nombres, u.primer_apellido, u.segundo_apellido,
			       u.telefono, u.foto_url, u.password_hash,
			       e.id, e.nombre, e.codigo
			FROM usuarios u
			INNER JOIN estados e ON u.estado_id = e.id
			WHERE u.username = $1;
		`

		var telefono, fotoURL sql.NullString

		err = db.QueryRowContext(context.Background(), query, req.Username).Scan(
			&u.ID, &u.Username, &u.Email, &u.Nombres, &u.PrimerApellido, &u.SegundoApellido,
			&telefono, &fotoURL, &u.PasswordHash,
			&u.Estado.ID, &u.Estado.Nombre, &u.Estado.Codigo,
		)
		if err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
			return
		}

		if err := utils.CompararPasswordHash(req.Password, u.PasswordHash); err != nil {
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}

		token, err := security.GenerarJWT(u.ID, u.Username)
		if err != nil {
			http.Error(w, "Error generando token", http.StatusInternalServerError)
			return
		}

		// Asegura valores nulos
		u.Telefono = ""
		u.FotoURL = ""
		if telefono.Valid {
			u.Telefono = telefono.String
		}
		if fotoURL.Valid {
			u.FotoURL = fotoURL.String
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token":   token,
			"usuario": u,
		})
	}
}
