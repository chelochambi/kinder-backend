// internal/model/menu.go
package model

type Menu struct {
	ID       int      `json:"id"`
	Nombre   string   `json:"nombre"`
	Icono    string   `json:"icono"`
	Ruta     string   `json:"ruta"`
	PadreID  *int     `json:"padre_id"`
	Tipo     string   `json:"tipo"`
	Mostrar  bool     `json:"mostrar"`
	Permisos []string `json:"permisos"`
	Submenus []*Menu  `json:"submenus,omitempty"`
}
