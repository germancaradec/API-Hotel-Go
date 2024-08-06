package main

import (
	"net/http"

	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/routes"
	"github.com/gorilla/mux"
)


func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Reservation{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	// s := r.PathPrefix("/api").Subrouter()

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/reservations", routes.GetReservationsHandler).Methods("GET")
	r.HandleFunc("/reservations/{id}", routes.GetReservationHandler).Methods("GET")
	r.HandleFunc("/reservations", routes.CreateReservationHandler).Methods("POST")
	r.HandleFunc("/reservations/{id}", routes.DeleteReservationHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}