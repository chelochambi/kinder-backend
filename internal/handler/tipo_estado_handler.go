package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chelochambi/kinder-backend/internal/model"
)

// ListarTipoEstado recibe la conexión a DB y retorna un handler para listar tipo_estado
func ListarTipoEstado(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.QueryContext(context.Background(), "SELECT id, nombre, descripcion, codigo FROM tipo_estado")
		if err != nil {
			http.Error(w, "Error al obtener tipo_estado", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tipos []model.TipoEstado
		for rows.Next() {
			var t model.TipoEstado
			if err := rows.Scan(&t.ID, &t.Nombre, &t.Descripcion, &t.Codigo); err != nil {
				http.Error(w, "Error al escanear tipo_estado", http.StatusInternalServerError)
				return
			}
			tipos = append(tipos, t)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tipos)
	}
}
