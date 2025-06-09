package model

type Estado struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
}
