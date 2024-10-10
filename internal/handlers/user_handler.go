package handlers

import (
	"ejemplo1/pkg/models"
	"encoding/json"
	"net/http"
)

var usuarios []models.User

// UserHandler maneja las solicitudes relacionadas con usuarios
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listarUsuarios(w)
	case "POST":
		crearUsuario(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido"})
	}
}

// listarUsuarios devuelve los usuarios en formato JSON
func listarUsuarios(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if len(usuarios) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay usuarios registrados"})
		return
	}

	// Devolver la lista de usuarios como JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

// crearUsuario crea un nuevo usuario a partir de los datos enviados en la solicitud
func crearUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parsear los datos enviados en el cuerpo de la solicitud (en JSON)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"})
		return
	}

	// Validar que los datos sean correctos
	if user.Nombre == "" || user.Correo == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Nombre o correo inválido"})
		return
	}

	// Asignar un ID al nuevo usuario
	user.ID = len(usuarios) + 1
	usuarios = append(usuarios, user)

	// Devolver el usuario creado en la respuesta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
