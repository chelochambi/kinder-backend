package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chelochambi/kinder-backend/internal/middleware"
	"github.com/chelochambi/kinder-backend/internal/model"
)

func AuthMeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtenemos el usuarioID del contexto (seteado por el middleware)
		usuarioIDVal := r.Context().Value(middleware.UsuarioIDKey)
		if usuarioIDVal == nil {
			http.Error(w, "Usuario no autenticado", http.StatusUnauthorized)
			return
		}

		usuarioID, ok := usuarioIDVal.(int)
		if !ok {
			http.Error(w, "Error con el ID de usuario", http.StatusInternalServerError)
			return
		}

		var u model.Usuario
		query := `
			SELECT u.id, u.username, u.email, u.nombres, u.primer_apellido, u.segundo_apellido,
			       u.telefono, u.foto_url,
			       e.id, e.nombre, e.codigo
			FROM usuarios u
			INNER JOIN estados e ON u.estado_id = e.id
			WHERE u.id = $1;
		`

		var telefono, fotoURL sql.NullString

		err := db.QueryRowContext(context.Background(), query, usuarioID).Scan(
			&u.ID, &u.Username, &u.Email, &u.Nombres, &u.PrimerApellido, &u.SegundoApellido,
			&telefono, &fotoURL,
			&u.Estado.ID, &u.Estado.Nombre, &u.Estado.Codigo,
		)
		if err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}

		if telefono.Valid {
			u.Telefono = telefono.String
		}
		if fotoURL.Valid {
			u.FotoURL = fotoURL.String
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}
