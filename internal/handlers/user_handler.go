package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SaulFernanRodri/go-starter/pkg/models"
)

var usuarios []models.User

// UserHandler maneja las solicitudes relacionadas con usuarios
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listarUsuarios(w)
	case "POST":
		crearUsuario(w, r)
	case "PUT":
		actualizarUsuario(w, r)
	case "DELETE":
		eliminarUsuario(w, r)
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

// actualizarUsuario actualiza un usuario existente a partir de los datos enviados en la solicitud
func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parsear los datos enviados en el cuerpo de la solicitud (en JSON)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"})
		return
	}

	// Buscar el usuario a actualizar
	for i, u := range usuarios {
		if u.ID == user.ID {
			// Actualizar los datos del usuario
			usuarios[i] = user
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// Si no se encuentra el usuario, devolver un error
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Usuario no encontrado"})
}

// eliminarUsuario elimina un usuario existente a partir del ID enviado como parámetro en la URL
func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obtener el ID del usuario desde los parámetros de la URL
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID del usuario no proporcionado"})
		return
	}

	// Convertir el ID de string a int
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
		return
	}

	// Buscar el usuario a eliminar
	for i, u := range usuarios {
		if u.ID == id {
			// Eliminar el usuario de la lista
			usuarios = append(usuarios[:i], usuarios[i+1:]...)

			// Devolver un mensaje de éxito en la respuesta
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Usuario eliminado"})
			return
		}
	}

	// Si no se encuentra el usuario, devolver un error
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Usuario no encontrado"})
}
