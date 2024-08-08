package routes

import (
	"encoding/json"
	"net/http"

	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/gorilla/mux"
)

// GetConsultationsHandler obtiene todas las consultas desde la base de datos en orden ascendente por ID y las devuelve en formato JSON
func GetConsultationsHandler(w http.ResponseWriter, r *http.Request) {
	var consultations []models.Consultation
	// Buscar todas las consultas en la base de datos y ordenarlas por ID en orden ascendente
	if err := db.DB.Order("id asc").Find(&consultations).Error; err != nil {
		// Manejar el error si ocurre al buscar las consultas
		http.Error(w, "Failed to retrieve consultations", http.StatusInternalServerError)
		return
	}
	
	// Codificar las consultas en formato JSON y enviarlas como respuesta
	if err := json.NewEncoder(w).Encode(&consultations); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// GetConsultationHandler obtiene una consulta específica por ID y la devuelve en formato JSON
func GetConsultationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var consultation models.Consultation
	// Buscar una consulta específica por ID
	if err := db.DB.First(&consultation, params["id"]).Error; err != nil {
		// Verificar si la consulta no fue encontrada
		if err.Error() == "record not found" {
			// Si la consulta no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Consultation not found"))
			return
		}
		// Manejar el error si ocurre al buscar la consulta
		http.Error(w, "Failed to retrieve consultation", http.StatusInternalServerError)
		return
	}

	// Codificar la consulta en formato JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(&consultation); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// CreateConsultationHandler crea una nueva consulta en la base de datos
func CreateConsultationHandler(w http.ResponseWriter, r *http.Request) {
	var consultation models.Consultation
	// Decodificar el cuerpo de la solicitud para obtener los datos de la nueva consulta
	if err := json.NewDecoder(r.Body).Decode(&consultation); err != nil {
		// Manejar el error si la carga útil de la solicitud es inválida
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Crear la nueva consulta en la base de datos
	if err := db.DB.Create(&consultation).Error; err != nil {
		// Manejar el error si ocurre al crear la consulta
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Codificar la consulta creada en formato JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(&consultation); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// UpdateConsultationHandler actualiza una consulta existente por ID con los datos proporcionados
func UpdateConsultationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var consultation models.Consultation

	// Buscar la consulta existente por ID
	if err := db.DB.First(&consultation, params["id"]).Error; err != nil {
		// Verificar si la consulta no fue encontrada
		if err.Error() == "record not found" {
			// Si la consulta no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Consultation not found"))
			return
		}
		// Manejar el error si ocurre al buscar la consulta
		http.Error(w, "Failed to retrieve consultation", http.StatusInternalServerError)
		return
	}

	// Decodificar el cuerpo de la solicitud para obtener los datos actualizados
	var updatedConsultation models.Consultation
	if err := json.NewDecoder(r.Body).Decode(&updatedConsultation); err != nil {
		// Manejar el error si la carga útil de la solicitud es inválida
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Actualizar los campos de la consulta existente con los datos proporcionados
	consultation.Phone = updatedConsultation.Phone
	consultation.Consultation = updatedConsultation.Consultation
	consultation.MoreInfo = updatedConsultation.MoreInfo
	consultation.UserID = updatedConsultation.UserID

	// Guardar los cambios en la base de datos
	if err := db.DB.Save(&consultation).Error; err != nil {
		// Manejar el error si ocurre al guardar la consulta actualizada
		http.Error(w, "Failed to update consultation", http.StatusInternalServerError)
		return
	}

	// Codificar la consulta actualizada en formato JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(&consultation); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// DeleteConsultationHandler elimina una consulta específica por ID
func DeleteConsultationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var consultation models.Consultation
	// Buscar la consulta específica por ID
	if err := db.DB.First(&consultation, params["id"]).Error; err != nil {
		// Verificar si la consulta no fue encontrada
		if err.Error() == "record not found" {
			// Si la consulta no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Consultation not found"))
			return
		}
		// Manejar el error si ocurre al buscar la consulta
		http.Error(w, "Failed to retrieve consultation", http.StatusInternalServerError)
		return
	}

	// Eliminar la consulta de la base de datos
	if err := db.DB.Unscoped().Delete(&consultation).Error; err != nil {
		// Manejar el error si ocurre al eliminar la consulta
		http.Error(w, "Failed to delete consultation", http.StatusInternalServerError)
		return
	}

	// Devolver un estado 200 OK si la eliminación fue exitosa
	w.WriteHeader(http.StatusOK)
}
