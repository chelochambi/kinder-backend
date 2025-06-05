package model

import "time"

type TipoEstado struct {
	ID             int       `json:"id"`
	Nombre         string    `json:"nombre"`
	Descripcion    string    `json:"descripcion"`
	CreadoEn       time.Time `json:"creado_en"`
	ActualizadoEn  time.Time `json:"actualizado_en"`
	Codigo         string    `json:"codigo"`
	CreadoPor      int       `json:"creado_por"`
	ActualizadoPor int       `json:"actualizado_por"`
}
