package model

type Usuario struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Nombres         string `json:"nombres"`
	PrimerApellido  string `json:"primer_apellido"`
	SegundoApellido string `json:"segundo_apellido"`
	Telefono        string `json:"telefono"`
	FotoURL         string `json:"foto_url"`

	Password     string `json:"password"` // <- se recibe desde el JSON
	PasswordHash string `json:"-"`        // <- se genera internamente

	Estado         Estado `json:"estado"`
	CreadoPor      int    `json:"creado_por,omitempty"`
	ActualizadoPor int    `json:"actualizado_por,omitempty"`
}
