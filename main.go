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

	// Migración de las tablas necesarias en la base de datos
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Reservation{})
	db.DB.AutoMigrate(&models.Consultation{})
	db.DB.AutoMigrate(&models.Employee{}) 

	// Creación del enrutador
	r := mux.NewRouter()

	// Rutas para User
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Rutas para Reservation
	r.HandleFunc("/reservations", routes.GetReservationsHandler).Methods("GET")
	r.HandleFunc("/reservations/{id}", routes.GetReservationHandler).Methods("GET")
	r.HandleFunc("/reservations", routes.CreateReservationHandler).Methods("POST")
	r.HandleFunc("/reservations/{id}", routes.UpdateReservationHandler).Methods("PUT")
	r.HandleFunc("/reservations/{id}", routes.DeleteReservationHandler).Methods("DELETE")

	// Rutas para Consultation
	r.HandleFunc("/consultations", routes.GetConsultationsHandler).Methods("GET")
	r.HandleFunc("/consultations/{id}", routes.GetConsultationHandler).Methods("GET")
	r.HandleFunc("/consultations", routes.CreateConsultationHandler).Methods("POST")
	r.HandleFunc("/consultations/{id}", routes.UpdateConsultationHandler).Methods("PUT")
	r.HandleFunc("/consultations/{id}", routes.DeleteConsultationHandler).Methods("DELETE")

	// Rutas para Employee
	r.HandleFunc("/employees", routes.GetEmployeesHandler).Methods("GET")
	r.HandleFunc("/employees/{id}", routes.GetEmployeeHandler).Methods("GET")
	r.HandleFunc("/employees", routes.PostEmployeeHandler).Methods("POST")
	r.HandleFunc("/employees/{id}", routes.UpdateEmployeeHandler).Methods("PUT")
	r.HandleFunc("/employees/{id}", routes.DeleteEmployeeHandler).Methods("DELETE")

	// Configuración del servidor HTTP con CORS habilitado
	http.ListenAndServe(":3000", middleware.CORS(r))
}
