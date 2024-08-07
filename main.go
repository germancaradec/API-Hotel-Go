package main

import (
	"net/http"

	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	middleware "github.com/germancaradec/Go-API-REST-PostgresSQL.git/midle"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/routes"
	"github.com/gorilla/mux"
)


func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Reservation{})
	db.DB.AutoMigrate(models.Consultation{})

	r := mux.NewRouter()

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/reservations", routes.GetReservationsHandler).Methods("GET")
	r.HandleFunc("/reservations/{id}", routes.GetReservationHandler).Methods("GET")
	r.HandleFunc("/reservations", routes.CreateReservationHandler).Methods("POST")
	r.HandleFunc("/reservations/{id}", routes.UpdateReservationHandler).Methods("PUT")
	r.HandleFunc("/reservations/{id}", routes.DeleteReservationHandler).Methods("DELETE")
	
	r.HandleFunc("/consultations", routes.GetConsultationsHandler).Methods("GET")
	r.HandleFunc("/consultations/{id}", routes.GetConsultationHandler).Methods("GET")
	r.HandleFunc("/consultations", routes.CreateConsultationHandler).Methods("POST")
	r.HandleFunc("/consultations/{id}", routes.UpdateConsultationHandler).Methods("PUT")
	r.HandleFunc("/consultations/{id}", routes.DeleteConsultationHandler).Methods("DELETE")

	http.ListenAndServe(":3000", middleware.CORS(r))
}