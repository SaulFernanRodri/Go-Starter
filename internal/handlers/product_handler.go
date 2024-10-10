package handlers

import (
	"ejemplo1/pkg/models"
	"encoding/json"
	"net/http"
)

var productos []models.Product

// ProductHandler maneja las solicitudes relacionadas con productos
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listarProductos(w)
	case "POST":
		crearProducto(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido"})
	}
}

// listarProductos devuelve los productos en formato JSON
func listarProductos(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if len(productos) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay productos registrados"})
		return
	}

	// Devolver la lista de productos como JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productos)
}

// crearProducto crea un nuevo producto a partir de los datos enviados en la solicitud
func crearProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parsear los datos enviados en el cuerpo de la solicitud (en JSON)
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"})
		return
	}

	// Validar que los datos sean correctos
	if product.Nombre == "" || product.Precio <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Nombre o precio inválido"})
		return
	}

	// Asignar un ID al nuevo producto
	product.ID = len(productos) + 1
	productos = append(productos, product)

	// Devolver el producto creado en la respuesta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
