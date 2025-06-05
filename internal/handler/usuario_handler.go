package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chelochambi/kinder-backend/internal/model"
)

func ListarUsuariosHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := `
		SELECT u.id, u.username, u.email, u.nombres, u.primer_apellido, u.segundo_apellido, 
		       u.telefono, u.foto_url,
		       e.id, e.nombre, e.codigo, e.descripcion
		FROM usuarios u
		INNER JOIN estados e ON u.estado_id = e.id
		ORDER BY u.id;
		`

		rows, err := db.QueryContext(context.Background(), query)
		if err != nil {
			http.Error(w, "Error al consultar usuarios", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var usuarios []model.Usuario

		for rows.Next() {
			var u model.Usuario
			var telefono sql.NullString
			var fotoURL sql.NullString

			err := rows.Scan(
				&u.ID, &u.Username, &u.Email, &u.Nombres, &u.PrimerApellido, &u.SegundoApellido,
				&telefono, &fotoURL,
				&u.Estado.ID, &u.Estado.Nombre, &u.Estado.Codigo, &u.Estado.Descripcion,
			)
			if err != nil {
				http.Error(w, "Error al leer datos", http.StatusInternalServerError)
				return
			}

			// Asignar valores seguros para posibles NULL
			if telefono.Valid {
				u.Telefono = telefono.String
			} else {
				u.Telefono = ""
			}

			if fotoURL.Valid {
				u.FotoURL = fotoURL.String
			} else {
				u.FotoURL = ""
			}

			usuarios = append(usuarios, u)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(usuarios)
	}
}
