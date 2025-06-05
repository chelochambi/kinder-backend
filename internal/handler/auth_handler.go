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
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		var u model.UsuarioInfo
		var passwordHash string
		var telefono, fotoURL sql.NullString

		// Consulta principal del usuario
		query := `
			SELECT u.id, u.username, u.email, u.nombres, u.primer_apellido, u.segundo_apellido,
			       u.telefono, u.foto_url, u.password_hash
			FROM usuarios u
			WHERE u.username = $1 AND u.estado_id = 3
		`

		err := db.QueryRowContext(context.Background(), query, req.Username).Scan(
			&u.ID, &u.Username, &u.Email, &u.Nombres, &u.PrimerApellido, &u.SegundoApellido,
			&telefono, &fotoURL, &passwordHash,
		)
		if err != nil {
			http.Error(w, "Usuario no encontrado o inactivo", http.StatusUnauthorized)
			return
		}

		if err := utils.CompararPasswordHash(req.Password, passwordHash); err != nil {
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}

		if telefono.Valid {
			u.Telefono = telefono.String
		}
		if fotoURL.Valid {
			u.FotoURL = fotoURL.String
		}

		// TOKEN
		token, err := security.GenerarJWT(u.ID, u.Username)
		if err != nil {
			http.Error(w, "Error generando token", http.StatusInternalServerError)
			return
		}

		// ROLES
		rows, _ := db.Query(`SELECT r.codigo FROM usuario_rol ur JOIN roles r ON ur.rol_id = r.id WHERE ur.usuario_id = $1 AND ur.estado_id = 1`, u.ID)
		for rows.Next() {
			var codigo string
			rows.Scan(&codigo)
			u.Roles = append(u.Roles, codigo)
		}
		rows.Close()

		// PERMISOS
		rows, _ = db.Query(`
			SELECT DISTINCT p.codigo
			FROM usuario_rol ur
			JOIN rol_permiso rp ON ur.rol_id = rp.rol_id
			JOIN permisos p ON rp.permiso_id = p.id
			WHERE ur.usuario_id = $1 AND ur.estado_id = 1 AND rp.estado_id = 1
		`, u.ID)
		for rows.Next() {
			var codigo string
			rows.Scan(&codigo)
			u.Permisos = append(u.Permisos, codigo)
		}
		rows.Close()

		// MENUS
		type MenuItem struct {
			ID      int
			Nombre  string
			Icono   string
			Ruta    string
			Tipo    string
			Mostrar bool
			PadreID sql.NullInt64
		}
		menuMap := map[int]*model.Menu{}
		var flatMenus []MenuItem

		rows, _ = db.Query(`
			SELECT DISTINCT m.id, m.nombre, m.icono, m.ruta, m.tipo, m.mostrar, m.padre_id
			FROM usuario_rol ur
			JOIN rol_permiso rp ON ur.rol_id = rp.rol_id
			JOIN permisos p ON rp.permiso_id = p.id
			JOIN menus m ON m.id = p.menu_id
			WHERE ur.usuario_id = $1 AND m.estado_id = 1
		`, u.ID)

		for rows.Next() {
			var item MenuItem
			rows.Scan(&item.ID, &item.Nombre, &item.Icono, &item.Ruta, &item.Tipo, &item.Mostrar, &item.PadreID)
			flatMenus = append(flatMenus, item)
		}
		rows.Close()

		// Armar árbol de menús
		for _, item := range flatMenus {
			menu := &model.Menu{
				ID:      item.ID,
				Nombre:  item.Nombre,
				Icono:   item.Icono,
				Ruta:    item.Ruta,
				Tipo:    item.Tipo,
				Mostrar: item.Mostrar,
			}
			menuMap[menu.ID] = menu
		}
		for _, item := range flatMenus {
			if item.PadreID.Valid {
				padre := menuMap[int(item.PadreID.Int64)]
				if padre != nil {
					padre.Submenus = append(padre.Submenus, menuMap[item.ID])
				}
			} else {
				u.Menus = append(u.Menus, menuMap[item.ID])
			}
		}

		// SUCURSALES
		rows, _ = db.Query(`
			SELECT s.id, s.nombre
			FROM usuario_sucursal us
			JOIN sucursales s ON s.id = us.sucursal_id
			WHERE us.usuario_id = $1 AND us.estado_id = 1
		`, u.ID)

		for rows.Next() {
			var s model.Sucursal
			rows.Scan(&s.ID, &s.Nombre)
			u.Sucursales = append(u.Sucursales, s)
		}
		rows.Close()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token":   token,
			"usuario": u,
		})
	}
}
