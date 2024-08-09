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
	db.DB.AutoMigrate(&models.User{}, &models.Consultation{}) // Migrar los modelos
}

// Limpia las tablas de la base de datos para evitar errores de clave duplicada y restricciones de clave externa
func cleanUpConsultationDB() {
	db.DB.Unscoped().Exec("DELETE FROM consultations")
	db.DB.Unscoped().Exec("DELETE FROM users")
}

func TestGetConsultationsHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	// Crear un usuario de prueba
	user := models.User{
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "alice.johnson@example.com",
	}
	db.DB.Create(&user)

	// Insertar un registro de prueba
	db.DB.Create(&models.Consultation{
		Phone:         "+1234567890",
		Consultation:  "Necesito información sobre los servicios de consultoría disponibles.",
		MoreInfo:      true,
		UserID:        user.ID, // Asociar la consulta con el usuario
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
	assert.Equal(t, "+1234567890", consultations[0].Phone)
}

func TestGetConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	// Crear un usuario de prueba
	user := models.User{
		FirstName: "Bob",
		LastName:  "Smith",
		Email:     "bob.smith@example.com",
	}
	db.DB.Create(&user)

	consultation := models.Consultation{
		Phone:         "+0987654321",
		Consultation:  "Consulta sobre servicios de consultoría.",
		MoreInfo:      false,
		UserID:        user.ID, // Asociar la consulta con el usuario
	}
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
	assert.Equal(t, "+0987654321", returnedConsultation.Phone)
	assert.Equal(t, "Consulta sobre servicios de consultoría.", returnedConsultation.Consultation)
}

func TestCreateConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	// Crear un usuario de prueba
	user := models.User{
		FirstName: "Carol",
		LastName:  "Doe",
		Email:     "carol.doe@example.com",
	}
	db.DB.Create(&user)

	consultation := models.Consultation{
		Phone:         "123-456-7890",
		Consultation:  "I need some help with my project.",
		MoreInfo:      true,
		UserID:        user.ID, // Asociar la consulta con el usuario
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
	assert.Equal(t, "123-456-7890", createdConsultation.Phone)
	assert.Equal(t, "I need some help with my project.", createdConsultation.Consultation)
}

func TestUpdateConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	// Crear un usuario de prueba
	user := models.User{
		FirstName: "Dan",
		LastName:  "Brown",
		Email:     "dan.brown@example.com",
	}
	db.DB.Create(&user)

	consultation := models.Consultation{
		Phone:         "321-654-0987",
		Consultation:  "Initial consultation text.",
		MoreInfo:      true,
		UserID:        user.ID, // Asociar la consulta con el usuario
	}
	db.DB.Create(&consultation)

	updatedConsultation := models.Consultation{
		Phone:         "987-654-3210",
		Consultation:  "Updated consultation text.",
		MoreInfo:      false,
		UserID:        user.ID, // Debe seguir asociado al mismo usuario
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
	assert.Equal(t, "987-654-3210", result.Phone)
	assert.Equal(t, "Updated consultation text.", result.Consultation)
	assert.False(t, result.MoreInfo)
}

func TestDeleteConsultationHandler(t *testing.T) {
	setupConsultationDB()
	defer cleanUpConsultationDB() // Limpiar después de la prueba

	// Crear un usuario de prueba
	user := models.User{
		FirstName: "Eva",
		LastName:  "Adams",
		Email:     "eva.adams@example.com",
	}
	db.DB.Create(&user)

	consultation := models.Consultation{
		Phone:         "654-321-0987",
		Consultation:  "Consulta para eliminar.",
		MoreInfo:      true,
		UserID:        user.ID, // Asociar la consulta con el usuario
	}
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
