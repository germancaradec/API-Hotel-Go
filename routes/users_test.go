package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Configura el router para las pruebas
func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	r.HandleFunc("/users", PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")
	return r
}

// Conecta a la base de datos y realiza las migraciones necesarias para los tests
func setupDB() {
	db.DBConnection() // Conectar a la base de datos
	db.DB.AutoMigrate(&models.User{}, &models.Reservation{}) // Migrar los modelos User y Reservation
}

// Limpia las tablas de la base de datos para evitar errores de clave duplicada y restricciones de clave externa
func cleanUpDB() {
	// Eliminar todos los registros de las tablas usando Unscoped para ignorar la eliminación lógica
	db.DB.Unscoped().Exec("DELETE FROM reservations")
	db.DB.Unscoped().Exec("DELETE FROM users")
}

func setupUserDB() {
	setupDB()
}

func cleanUpUserDB() {
	cleanUpDB()
}

func TestGetUsersHandler(t *testing.T) {
	setupUserDB()
	defer cleanUpUserDB()

	// Crear un usuario de prueba
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	db.DB.Create(&user)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := setupRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var users []models.User
	err = json.NewDecoder(rr.Body).Decode(&users)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, users, 1)
	assert.Equal(t, "John", users[0].FirstName)
	assert.Equal(t, "Doe", users[0].LastName)
}

func TestGetUserHandler(t *testing.T) {
	setupDB()
	defer cleanUpDB() // Limpiar después de la prueba

	user := models.User{FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"}
	db.DB.Create(&user)

	userID := strconv.FormatUint(uint64(user.ID), 10) // Convertir ID a cadena

	req, err := http.NewRequest("GET", "/users/"+userID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := setupRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedUser models.User
	err = json.NewDecoder(rr.Body).Decode(&returnedUser)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Jane", returnedUser.FirstName)
	assert.Equal(t, "Doe", returnedUser.LastName)
}

func TestPostUserHandler(t *testing.T) {
	setupDB()
	defer cleanUpDB() // Limpiar después de la prueba

	user := models.User{FirstName: "Alice", LastName: "Smith", Email: "alice.smith@example.com"}
	userJson, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJson))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := setupRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var createdUser models.User
	err = json.NewDecoder(rr.Body).Decode(&createdUser)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Alice", createdUser.FirstName)
	assert.Equal(t, "Smith", createdUser.LastName)
}

func TestUpdateUserHandler(t *testing.T) {
	setupDB()
	defer cleanUpDB() // Limpiar después de la prueba

	user := models.User{FirstName: "Bob", LastName: "Johnson", Email: "bob.johnson@example.com"}
	db.DB.Create(&user)

	updatedUser := models.User{FirstName: "Robert", LastName: "Johnson", Email: "robert.johnson@example.com"}
	userJson, err := json.Marshal(updatedUser)
	if err != nil {
		t.Fatal(err)
	}

	userID := strconv.FormatUint(uint64(user.ID), 10) // Convertir ID a cadena

	req, err := http.NewRequest("PUT", "/users/"+userID, bytes.NewBuffer(userJson))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := setupRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.User
	err = json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Robert", result.FirstName)
	assert.Equal(t, "Johnson", result.LastName)
}

func TestDeleteUserHandler(t *testing.T) {
	setupDB()
	defer cleanUpDB() // Limpiar después de la prueba

	user := models.User{FirstName: "Charlie", LastName: "Brown", Email: "charlie.brown@example.com"}
	db.DB.Create(&user)

	userID := strconv.FormatUint(uint64(user.ID), 10) // Convertir ID a cadena

	req, err := http.NewRequest("DELETE", "/users/"+userID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := setupRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var deletedUser models.User
	err = db.DB.Unscoped().First(&deletedUser, user.ID).Error
	assert.Error(t, err)
}
