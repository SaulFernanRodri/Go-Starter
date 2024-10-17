package models

type Product struct {
	ID        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Precio    float64   `json:"precio"`
	Imagen    string    `json:"imagen"`
	Milsymbol Milsymbol `json:"milsymbol"`
}
