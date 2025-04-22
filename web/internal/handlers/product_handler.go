package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-starter/web/pkg/models"
	"go-starter/web/pkg/utils"
)

var productos []models.Product

// ProductHandler maneja las solicitudes relacionadas con productos
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listarProductos(w)
	case "POST":
		crearProducto(w, r)
	case "PUT":
		actualizarProducto(w, r)
	case "DELETE":
		eliminarProducto(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}

// listarProductos devuelve los productos en formato JSON
func listarProductos(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if len(productos) == 0 {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "No hay productos registrados"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Devolver la lista de productos como JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(productos); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// crearProducto crea un nuevo producto a partir de los datos enviados en la solicitud
func crearProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parsear los datos enviados en el cuerpo de la solicitud (en JSON)
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Validar que los datos sean correctos
	if product.Nombre == "" || product.Precio <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Nombre o precio inválido"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Asignar un ID al nuevo producto
	product.ID = len(productos) + 1

	product.Imagen, err = utils.GenerateMilsymbol(product.Milsymbol)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Error al generar el símbolo"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	productos = append(productos, product)

	// Devolver el producto creado en la respuesta
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// actualizarProducto actualiza un producto existente
func actualizarProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parsear los datos enviados en el cuerpo de la solicitud (en JSON)
	var updatedProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Buscar el producto a actualizar
	for i, product := range productos {
		if product.ID == updatedProduct.ID {
			// Actualizar los datos del producto
			productos[i] = updatedProduct

			// Devolver el producto actualizado en la respuesta
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(updatedProduct); err != nil {
				http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
			}
			return
		}
	}

	// Si no se encuentra el producto, devolver un error
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": "Producto no encontrado"}); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// eliminarProducto elimina un producto existente
func eliminarProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obtener el ID del producto desde los parámetros de la URL
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "ID del producto no proporcionado"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Convertir el ID de string a int
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Buscar el producto a eliminar
	for i, product := range productos {
		if product.ID == id {
			// Eliminar el producto de la lista
			productos = append(productos[:i], productos[i+1:]...)

			// Devolver un mensaje de éxito en la respuesta
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(map[string]string{"message": "Producto eliminado"}); err != nil {
				http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
			}
			return
		}
	}

	// Si no se encuentra el producto, devolver un error
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": "Producto no encontrado"}); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}
