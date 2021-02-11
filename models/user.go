package models

type Usuario struct {
	Nombre   string `json:"nombre" db:"nombre"`
	Apellido string `json:"apellido" db:"apellido"`
	Correo   string `json:"correo" db:"correo"`
	Password string `json:"password" db:"password"`
}
