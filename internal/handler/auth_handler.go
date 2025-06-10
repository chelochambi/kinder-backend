package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"log"

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

		query := `
			SELECT u.id, u.username, u.email, u.nombres, u.primer_apellido, u.segundo_apellido,
			       u.telefono, u.foto_url, u.password_hash
			FROM usuarios u
			JOIN estados e ON u.estado_id = e.id
			WHERE u.username = $1 AND e.codigo = 'UACT'
		`

		err := db.QueryRowContext(context.Background(), query, req.Username).Scan(
			&u.ID, &u.Username, &u.Email, &u.Nombres, &u.PrimerApellido, &u.SegundoApellido,
			&telefono, &fotoURL, &passwordHash,
		)

		if err != nil {
			http.Error(w, "Usuario no encontrado o inactivo", http.StatusUnauthorized)
			return
		}

		// código de análisis
		errHash := utils.CompararPasswordHash(req.Password, passwordHash)
		if errHash != nil {
			log.Printf("Entró en errHash != nil - Contraseña incorrecta %s", u.Username)
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}
		log.Println("Contraseña válida, continuando con login...")
		// Si llegamos aquí, la contraseña es correcta

		if telefono.Valid {
			u.Telefono = telefono.String
		}
		if fotoURL.Valid {
			u.FotoURL = fotoURL.String
		}

		token, err := security.GenerarJWT(u.ID, u.Username)
		if err != nil {
			http.Error(w, "Error generando token", http.StatusInternalServerError)
			return
		}

		// ROLES
		rows, err := db.Query(`
			SELECT r.codigo
			FROM usuario_rol ur
			JOIN roles r ON ur.rol_id = r.id
			JOIN estados e ON ur.estado_id = e.id
			WHERE ur.usuario_id = $1 AND e.codigo = 'HAB'
		`, u.ID)
		if err != nil {
			log.Println("Error al obtener roles:", err)
			http.Error(w, "Error obteniendo roles", http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var codigo string
			rows.Scan(&codigo)
			u.Roles = append(u.Roles, codigo)
		}
		rows.Close()
		if len(u.Roles) == 0 {
			log.Println("Usuario sin roles habilitados:", u.Username)
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized) // Mensaje neutro
			return
		}
		// MENÚS + PERMISOS
		type MenuRow struct {
			MenuID  int
			Nombre  string
			Icono   string
			Ruta    string
			Tipo    string
			Mostrar bool
			PadreID sql.NullInt64
			Permiso string
		}

		rows, err = db.Query(`
			SELECT DISTINCT m.id, m.nombre, m.icono, m.ruta, m.tipo, m.mostrar, m.padre_id, p.codigo
			FROM usuario_rol ur
			JOIN rol_menu_permiso rmp ON ur.rol_id = rmp.rol_id
			JOIN permisos p ON rmp.permiso_id = p.id
			JOIN menus m ON rmp.menu_id = m.id
			JOIN estados e1 ON ur.estado_id = e1.id
			JOIN estados e2 ON rmp.estado_id = e2.id
			JOIN estados e3 ON m.estado_id = e3.id
			WHERE ur.usuario_id = $1 AND e1.codigo = 'HAB' AND e2.codigo = 'HAB' AND e3.codigo = 'HAB'
		`, u.ID)

		if err != nil {
			log.Println("Error al obtener menús:", err)
			http.Error(w, "Error obteniendo menús", http.StatusInternalServerError)
			return
		}

		menuMap := map[int]*model.Menu{}
		permMap := map[int][]string{}
		var flatMenuData []MenuRow

		for rows.Next() {
			var row MenuRow
			rows.Scan(&row.MenuID, &row.Nombre, &row.Icono, &row.Ruta, &row.Tipo, &row.Mostrar, &row.PadreID, &row.Permiso)
			flatMenuData = append(flatMenuData, row)
			permMap[row.MenuID] = append(permMap[row.MenuID], row.Permiso)
		}
		rows.Close()

		if len(flatMenuData) == 0 {
			log.Println("Usuario sin menús asignados:", u.Username)
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}
		// Construir los menús base
		for _, row := range flatMenuData {
			if _, exists := menuMap[row.MenuID]; !exists {
				menu := &model.Menu{
					ID:       row.MenuID,
					Nombre:   row.Nombre,
					Icono:    row.Icono,
					Ruta:     row.Ruta,
					Tipo:     row.Tipo,
					Mostrar:  row.Mostrar,
					Permisos: permMap[row.MenuID],
				}
				menuMap[row.MenuID] = menu
			}
		}

		// Construir jerarquía de menús sin duplicados
		submenuAdded := make(map[int]bool)
		for _, row := range flatMenuData {
			menu := menuMap[row.MenuID]
			if row.PadreID.Valid {
				padreID := int(row.PadreID.Int64)
				padre := menuMap[padreID]
				if padre != nil && !submenuAdded[menu.ID] {
					padre.Submenus = append(padre.Submenus, menu)
					submenuAdded[menu.ID] = true
				}
			} else {
				if !submenuAdded[menu.ID] {
					u.Menus = append(u.Menus, menu)
					submenuAdded[menu.ID] = true
				}
			}
		}

		// SUCURSALES
		rows, err = db.Query(`
			SELECT s.id, s.nombre
			FROM usuario_sucursal us
			JOIN sucursales s ON s.id = us.sucursal_id
			JOIN estados e ON us.estado_id = e.id
			WHERE us.usuario_id = $1 AND e.codigo = 'ACT'
		`, u.ID)
		if err != nil {
			log.Println("Error al obtener sucursales:", err)
			http.Error(w, "Error obteniendo sucursales", http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var s model.Sucursal
			rows.Scan(&s.ID, &s.Nombre)
			u.Sucursales = append(u.Sucursales, s)
		}
		rows.Close()
		if len(u.Sucursales) == 0 {
			log.Println("Usuario sin sucursales activas:", u.Username)
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}
		// RESPUESTA FINAL
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token":   token,
			"usuario": u,
		})
	}
}
