package model

type Permiso struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Codigo string `json:"codigo"`
	Accion string `json:"accion"`
}
