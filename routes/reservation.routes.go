package routes

import (
	"encoding/json"
	"net/http"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/gorilla/mux"
)

// GetReservationsHandler obtiene todas las reservas
func GetReservationsHandler(w http.ResponseWriter, r *http.Request) {
	var reservations []models.Reservation
	db.DB.Find(&reservations)
	
	// Encode the result into JSON and send it back to the client
	if err := json.NewEncoder(w).Encode(&reservations); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// GetReservationHandler obtiene una reserva específica por ID
func GetReservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reservation models.Reservation
	db.DB.First(&reservation, params["id"])

	if reservation.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Reservation not found"))
		return
	}

	// Encode the result into JSON and send it back to the client
	if err := json.NewEncoder(w).Encode(&reservation); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// CreateReservationHandler crea una nueva reserva
func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdReservation := db.DB.Create(&reservation)
	if err := createdReservation.Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Encode the result into JSON and send it back to the client
	if err := json.NewEncoder(w).Encode(&reservation); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// UpdateReservationHandler actualiza una reserva existente por ID
func UpdateReservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reservation models.Reservation

	// Buscar la reserva existente por ID
	if err := db.DB.First(&reservation, params["id"]).Error; err != nil {
		if err.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Reservation not found"))
			return
		}
		http.Error(w, "Failed to retrieve reservation", http.StatusInternalServerError)
		return
	}

	// Decodificar la solicitud para obtener los datos actualizados
	var updatedReservation models.Reservation
	if err := json.NewDecoder(r.Body).Decode(&updatedReservation); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Actualizar los campos de la reserva existente con los datos proporcionados
	reservation.Checkin = updatedReservation.Checkin
	reservation.Checkout = updatedReservation.Checkout
	reservation.Email = updatedReservation.Email
	reservation.UserID = updatedReservation.UserID

	// Guardar los cambios en la base de datos
	if err := db.DB.Save(&reservation).Error; err != nil {
		http.Error(w, "Failed to update reservation", http.StatusInternalServerError)
		return
	}

	// Enviar la reserva actualizada como respuesta
	if err := json.NewEncoder(w).Encode(&reservation); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// DeleteReservationHandler elimina una reserva específica por ID
func DeleteReservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reservation models.Reservation
	db.DB.First(&reservation, params["id"])

	if reservation.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Reservation not found"))
		return
	}

	if err := db.DB.Unscoped().Delete(&reservation).Error; err != nil {
		http.Error(w, "Failed to delete reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
