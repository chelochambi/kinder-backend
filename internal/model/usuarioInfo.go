package model

type UsuarioInfo struct {
	ID              int        `json:"id"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Nombres         string     `json:"nombres"`
	PrimerApellido  string     `json:"primer_apellido"`
	SegundoApellido string     `json:"segundo_apellido"`
	Telefono        string     `json:"telefono"`
	FotoURL         string     `json:"foto_url"`
	Roles           []string   `json:"roles"`
	Permisos        []string   `json:"permisos"`
	Menus           []*Menu    `json:"menus"`
	Sucursales      []Sucursal `json:"sucursales"`
}
