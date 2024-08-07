package routes

import (
	"encoding/json"
	"net/http"

	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/gorilla/mux"
)

// GetUsersHandler obtiene todos los usuarios desde la base de datos en orden ascendente por ID y los devuelve en formato JSON
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	// Buscar todos los usuarios en la base de datos y ordenarlos por ID en orden ascendente
	if err := db.DB.Order("id asc").Find(&users).Error; err != nil {
		// Manejar el error si ocurre al buscar los usuarios
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	
	// Codificar los usuarios en formato JSON y enviarlos como respuesta
	if err := json.NewEncoder(w).Encode(&users); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}


// GetUserHandler obtiene un usuario específico por ID y lo devuelve en formato JSON
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var user models.User
	// Buscar un usuario específico por ID
	db.DB.First(&user, params["id"])

	// Verificar si el usuario fue encontrado
	if user.ID == 0 {
		// Si el usuario no existe, devolver un error 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// Cargar las reservas asociadas al usuario (opcional)
	if err := db.DB.Model(&user).Association("Reservations").Find(&user.Reservations); err != nil {
		http.Error(w, "Failed to load user reservations", http.StatusInternalServerError)
		return
	}

	// Codificar el usuario en formato JSON y enviarlo como respuesta
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// PostUserHandler crea un nuevo usuario en la base de datos
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Decodificar el cuerpo de la solicitud para obtener los datos del nuevo usuario
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// Manejar el error si la carga útil de la solicitud es inválida
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Crear el nuevo usuario en la base de datos
	createdUser := db.DB.Create(&user)
	if err := createdUser.Error; err != nil {
		// Manejar el error si ocurre al crear el usuario
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Codificar el usuario creado en formato JSON y enviarlo como respuesta
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// UpdateUserHandler actualiza un usuario existente por ID con los datos proporcionados
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var user models.User

	// Buscar el usuario existente por ID
	if err := db.DB.First(&user, params["id"]).Error; err != nil {
		if err.Error() == "record not found" {
			// Si el usuario no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
			return
		}
		// Manejar el error si ocurre al buscar el usuario
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	// Decodificar el cuerpo de la solicitud para obtener los datos actualizados
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		// Manejar el error si la carga útil de la solicitud es inválida
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Actualizar los campos del usuario existente con los datos proporcionados
	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email
	// Nota: No actualizamos el campo `Reservations` ya que es una relación y no suele actualizarse directamente en un PUT

	// Guardar los cambios en la base de datos
	if err := db.DB.Save(&user).Error; err != nil {
		// Manejar el error si ocurre al guardar el usuario actualizado
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Cargar las reservas asociadas para la respuesta (opcional)
	if err := db.DB.Model(&user).Association("Reservations").Find(&user.Reservations); err != nil {
		http.Error(w, "Failed to load user reservations", http.StatusInternalServerError)
		return
	}

	// Codificar el usuario actualizado en formato JSON y enviarlo como respuesta
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// DeleteUserHandler elimina un usuario específico por ID
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var user models.User
	// Buscar el usuario específico por ID
	db.DB.First(&user, params["id"])

	// Verificar si el usuario fue encontrado
	if user.ID == 0 {
		// Si el usuario no existe, devolver un error 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// Eliminar el usuario de la base de datos
	if err := db.DB.Unscoped().Delete(&user).Error; err != nil {
		// Manejar el error si ocurre al eliminar el usuario
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// Devolver un estado 200 OK si la eliminación fue exitosa
	w.WriteHeader(http.StatusOK)
}
