package routes

import (
	"encoding/json"
	"net/http"

	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/gorilla/mux"
)

// GetReservationsHandler obtiene todas las reservas desde la base de datos en orden ascendente por ID y las devuelve en formato JSON
func GetReservationsHandler(w http.ResponseWriter, r *http.Request) {
	var reservations []models.Reservation
	// Buscar todas las reservas en la base de datos y ordenarlas por ID en orden ascendente
	if err := db.DB.Order("id asc").Find(&reservations).Error; err != nil {
		// Manejar el error si ocurre al buscar las reservas
		http.Error(w, "Failed to retrieve reservations", http.StatusInternalServerError)
		return
	}
	
	// Codificar las reservas en formato JSON y enviarlas como respuesta
	if err := json.NewEncoder(w).Encode(&reservations); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// GetReservationHandler obtiene una reserva específica por ID y la devuelve en formato JSON
func GetReservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var reservation models.Reservation
	// Buscar una reserva específica por ID
	if err := db.DB.First(&reservation, params["id"]).Error; err != nil {
		// Verificar si la reserva no fue encontrada
		if err.Error() == "record not found" {
			// Si la reserva no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Reservation not found"))
			return
		}
		// Manejar el error si ocurre al buscar la reserva
		http.Error(w, "Failed to retrieve reservation", http.StatusInternalServerError)
		return
	}

	// Codificar la reserva en formato JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(&reservation); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// CreateReservationHandler crea una nueva reserva en la base de datos
func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	// Decodificar el cuerpo de la solicitud para obtener los datos de la nueva reserva
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		// Manejar el error si la carga útil de la solicitud es inválida
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Crear la nueva reserva en la base de datos
	if err := db.DB.Create(&reservation).Error; err != nil {
		// Manejar el error si ocurre al crear la reserva
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Codificar la reserva creada en formato JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(&reservation); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// UpdateReservationHandler actualiza una reserva existente por ID con los datos proporcionados
func UpdateReservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var reservation models.Reservation

	// Buscar la reserva existente por ID
	if err := db.DB.First(&reservation, params["id"]).Error; err != nil {
		// Verificar si la reserva no fue encontrada
		if err.Error() == "record not found" {
			// Si la reserva no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Reservation not found"))
			return
		}
		// Manejar el error si ocurre al buscar la reserva
		http.Error(w, "Failed to retrieve reservation", http.StatusInternalServerError)
		return
	}

	// Decodificar el cuerpo de la solicitud para obtener los datos actualizados
	var updatedReservation models.Reservation
	if err := json.NewDecoder(r.Body).Decode(&updatedReservation); err != nil {
		// Manejar el error si la carga útil de la solicitud es inválida
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Actualizar los campos de la reserva existente con los datos proporcionados
	reservation.Adults = updatedReservation.Adults
	reservation.Checkin = updatedReservation.Checkin
	reservation.Checkout = updatedReservation.Checkout
	reservation.Children = updatedReservation.Children
	reservation.Email = updatedReservation.Email
	reservation.NumberOfRooms = updatedReservation.NumberOfRooms
	reservation.RoomType = updatedReservation.RoomType
	reservation.UserID = updatedReservation.UserID

	// Guardar los cambios en la base de datos
	if err := db.DB.Save(&reservation).Error; err != nil {
		// Manejar el error si ocurre al guardar la reserva actualizada
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	// Codificar la reserva actualizada en formato JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(&reservation); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// DeleteReservationHandler elimina una reserva específica por ID
func DeleteReservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var reservation models.Reservation
	// Buscar la reserva específica por ID
	if err := db.DB.First(&reservation, params["id"]).Error; err != nil {
		// Verificar si la reserva no fue encontrada
		if err.Error() == "record not found" {
			// Si la reserva no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Reservation not found"))
			return
		}
		// Manejar el error si ocurre al buscar la reserva
		http.Error(w, "Failed to retrieve reservation", http.StatusInternalServerError)
		return
	}

	// Eliminar la reserva de la base de datos
	if err := db.DB.Unscoped().Delete(&reservation).Error; err != nil {
		// Manejar el error si ocurre al eliminar la reserva
		http.Error(w, "Failed to delete reservation", http.StatusInternalServerError)
		return
	}

	// Devolver un estado 200 OK si la eliminación fue exitosa
	w.WriteHeader(http.StatusOK)
}
