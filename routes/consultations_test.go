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

// Configura el router para las pruebas de Consultations
func setupConsultationRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/consultations", GetConsultationsHandler).Methods("GET")
	r.HandleFunc("/consultations/{id}", GetConsultationHandler).Methods("GET")
	r.HandleFunc("/consultations", CreateConsultationHandler).Methods("POST")
	r.HandleFunc("/consultations/{id}", UpdateConsultationHandler).Methods("PUT")
	r.HandleFunc("/consultations/{id}", DeleteConsultationHandler).Methods("DELETE")
	return r
}

// Conecta a la base de datos y realiza las migraciones necesarias para Consultations
func setupConsultationDB() {
	db.DBConnection() // Conectar a la base de datos
	db.DB.AutoMigrate(&models.Consultation{}) // Migrar el modelo Consultation
}

// Limpia las tablas de la base de datos para evitar errores de clave duplicada y restricciones de clave externa
func cleanUpConsultationDB() {
	db.DB.Unscoped().Exec("DELETE FROM consultations")
}

func TestGetConsultationsHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	// Insertar solo un registro de prueba
	db.DB.Create(&models.Consultation{
		FirstName:   "Alice",
		LastName:    "Johnson",
		Email:       "alice.johnson@example.com",
		Phone:       "+1234567890",
		Consultation: "Necesito información sobre los servicios de consultoría disponibles.",
		MoreInfo:    true,
	})

	req, err := http.NewRequest("GET", "/consultations", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := setupConsultationRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var consultations []models.Consultation
	err = json.NewDecoder(rr.Body).Decode(&consultations)
	if err != nil {
		t.Fatal(err)
	}

	// Verificar que solo haya un registro y que los datos sean correctos
	assert.Len(t, consultations, 1)
	assert.Equal(t, "Alice", consultations[0].FirstName)
	assert.Equal(t, "Johnson", consultations[0].LastName)
}

func TestGetConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	consultation := models.Consultation{FirstName: "Bob", LastName: "Smith", Email: "bob.smith@example.com"}
	db.DB.Create(&consultation)

	consultationID := strconv.FormatUint(uint64(consultation.ID), 10) // Convertir ID a cadena

	req, err := http.NewRequest("GET", "/consultations/"+consultationID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := setupConsultationRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedConsultation models.Consultation
	err = json.NewDecoder(rr.Body).Decode(&returnedConsultation)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Bob", returnedConsultation.FirstName)
	assert.Equal(t, "Smith", returnedConsultation.LastName)
}

func TestCreateConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	consultation := models.Consultation{
		FirstName:    "Charlie",
		LastName:     "Brown",
		Email:        "charlie.brown@example.com",
		Phone:        "123-456-7890",
		Consultation: "I need some help with my project.",
		MoreInfo:     true,
	}
	consultationJson, err := json.Marshal(consultation)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/consultations", bytes.NewBuffer(consultationJson))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := setupConsultationRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var createdConsultation models.Consultation
	err = json.NewDecoder(rr.Body).Decode(&createdConsultation)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Charlie", createdConsultation.FirstName)
	assert.Equal(t, "Brown", createdConsultation.LastName)
}

func TestUpdateConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	consultation := models.Consultation{FirstName: "David", LastName: "Jones", Email: "david.jones@example.com"}
	db.DB.Create(&consultation)

	updatedConsultation := models.Consultation{
		FirstName:    "David",
		LastName:     "Jones",
		Email:        "david.jones@example.com",
		Phone:        "987-654-3210",
		Consultation: "Updated consultation text.",
		MoreInfo:     false,
	}
	consultationJson, err := json.Marshal(updatedConsultation)
	if err != nil {
		t.Fatal(err)
	}

	consultationID := strconv.FormatUint(uint64(consultation.ID), 10) // Convertir ID a cadena

	req, err := http.NewRequest("PUT", "/consultations/"+consultationID, bytes.NewBuffer(consultationJson))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := setupConsultationRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Consultation
	err = json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Updated consultation text.", result.Consultation)
	assert.False(t, result.MoreInfo)
}

func TestDeleteConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	consultation := models.Consultation{FirstName: "Emma", LastName: "Wilson", Email: "emma.wilson@example.com"}
	db.DB.Create(&consultation)

	consultationID := strconv.FormatUint(uint64(consultation.ID), 10) // Convertir ID a cadena

	req, err := http.NewRequest("DELETE", "/consultations/"+consultationID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := setupConsultationRouter()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar que el registro no existe
	var deletedConsultation models.Consultation
	err = db.DB.Unscoped().First(&deletedConsultation, consultation.ID).Error
	assert.Error(t, err)
}
