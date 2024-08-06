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
