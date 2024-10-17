package models

// User representa un usuario con ID, nombre y correo electrónico.
type User struct {
	ID        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Correo    string    `json:"correo"`
	Imagen    string    `json:"imagen"`
	Milsymbol Milsymbol `json:"milsymbol"`
}
