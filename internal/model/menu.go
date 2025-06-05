package model

type Menu struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Icono    string  `json:"icono"`
	Ruta     string  `json:"ruta"`
	Tipo     string  `json:"tipo"`
	Mostrar  bool    `json:"mostrar"`
	PadreID  *int    `json:"padre_id,omitempty"`
	Submenus []*Menu `json:"submenus,omitempty"`
}
