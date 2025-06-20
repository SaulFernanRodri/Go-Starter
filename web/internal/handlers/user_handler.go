package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-starter/web/pkg/models"
	"go-starter/web/pkg/utils"
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
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}

// listarUsuarios devuelve los usuarios en formato JSON
func listarUsuarios(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if len(usuarios) == 0 {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "No hay usuarios registrados"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Devolver la lista de usuarios como JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// crearUsuario crea un nuevo usuario a partir de los datos enviados en la solicitud
func crearUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parsear los datos enviados en el cuerpo de la solicitud (en JSON)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "Datos invalidos"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Validar que los datos sean correctos
	if user.Nombre == "" || user.Correo == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Nombre o correo inválido"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Asignar un ID al nuevo usuario
	user.ID = len(usuarios) + 1

	user.Imagen, err = utils.GenerateMilsymbol(user.Milsymbol)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Error al generar el símbolo"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	usuarios = append(usuarios, user)

	// Devolver el usuario creado en la respuesta
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// actualizarUsuario actualiza un usuario existente a partir de los datos enviados en la solicitud
func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parsear los datos enviados en el cuerpo de la solicitud (en JSON)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"}); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
		return
	}

	// Buscar el usuario a actualizar
	for i, u := range usuarios {
		if u.ID == user.ID {
			// Actualizar los datos del usuario
			usuarios[i] = user
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
			}
			return
		}
	}

	// Si no se encuentra el usuario, devolver un error
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": "Usuario no encontrado"}); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// eliminarUsuario elimina un usuario existente a partir del ID enviado como parámetro en la URL
func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obtener el ID del usuario desde los parámetros de la URL
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "ID del usuario no proporcionado"}); err != nil {
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

	// Buscar el usuario a eliminar
	for i, u := range usuarios {
		if u.ID == id {
			// Eliminar el usuario de la lista
			usuarios = append(usuarios[:i], usuarios[i+1:]...)

			// Devolver un mensaje de éxito en la respuesta
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(map[string]string{"message": "Usuario eliminado"}); err != nil {
				http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
			}
			return
		}
	}

	// Si no se encuentra el usuario, devolver un error
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": "Usuario no encontrado"}); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

func ShowImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	// Obtener el ID del usuario desde los parámetros de la URL
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "ID del usuario no proporcionado"}); err != nil {
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

	// Buscar el usuario con el ID especificado
	for _, u := range usuarios {
		if u.ID == id {
			// Devolver la imagen del usuario
			if _, err := w.Write([]byte(u.Imagen)); err != nil {
				http.Error(w, "Error al escribir la imagen", http.StatusInternalServerError)
			}
			return
		}
	}

	// Si no se encuentra el usuario, devolver un error
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": "Usuario no encontrado"}); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}
