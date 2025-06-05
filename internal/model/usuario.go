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
	PasswordHash    string `json:"-"` // nunca se expone
	Estado          Estado `json:"estado"`
}
