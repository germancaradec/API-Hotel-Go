package routes

import (
	"encoding/json"
	"net/http"

	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/db"
	"github.com/germancaradec/Go-API-REST-PostgresSQL.git/models"
	"github.com/gorilla/mux"
)

// GetEmployeesHandler obtiene todos los empleados desde la base de datos en orden ascendente por ID y los devuelve en formato JSON
func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee
	// Buscar todos los empleados en la base de datos y ordenarlos por ID en orden ascendente
	if err := db.DB.Order("id asc").Find(&employees).Error; err != nil {
		// Manejar el error si ocurre al buscar los empleados
		http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
		return
	}

	// Codificar los empleados en formato JSON y enviarlos como respuesta
	if err := json.NewEncoder(w).Encode(&employees); err != nil {
		// Manejar el error si ocurre al codificar el JSON
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var employee models.Employee

    // Buscar un empleado específico por ID e incluir reservas y consultas asociadas
    if err := db.DB.Preload("User.Reservations").Preload("User.Consultations").Preload("Reservations").Preload("Consultations").First(&employee, params["id"]).Error; err != nil {
        if err.Error() == "record not found" {
            w.WriteHeader(http.StatusNotFound)
            w.Write([]byte("Employee not found"))
            return
        }
        http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(&employee); err != nil {
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
    }
}



func PostEmployeeHandler(w http.ResponseWriter, r *http.Request) {
    var employee models.Employee
    if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    createdEmployee := db.DB.Create(&employee)
    if err := createdEmployee.Error; err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := json.NewEncoder(w).Encode(&employee); err != nil {
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
    }
}


func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var employee models.Employee

    if err := db.DB.Preload("User.Reservations").Preload("User.Consultations").First(&employee, params["id"]).Error; err != nil {
        if err.Error() == "record not found" {
            w.WriteHeader(http.StatusNotFound)
            w.Write([]byte("Employee not found"))
            return
        }
        http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
        return
    }

    var updatedEmployee struct {
        Position    string  `json:"position"`
        Salary      float64 `json:"salary"`
        Department  string  `json:"department"`
        HireDate    string  `json:"hire_date"`
        PhoneNumber string  `json:"phone_number"`
        User        struct {
            FirstName string `json:"first_name"`
            LastName  string `json:"last_name"`
            Email     string `json:"email"`
        } `json:"user"`
    }

    if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    employee.Position = updatedEmployee.Position
    employee.Salary = updatedEmployee.Salary
    employee.Department = updatedEmployee.Department
    employee.HireDate = updatedEmployee.HireDate
    employee.PhoneNumber = updatedEmployee.PhoneNumber
    employee.User.FirstName = updatedEmployee.User.FirstName
    employee.User.LastName = updatedEmployee.User.LastName
    employee.User.Email = updatedEmployee.User.Email

    if err := db.DB.Save(&employee).Error; err != nil {
        http.Error(w, "Failed to update employee", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(&employee); err != nil {
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
    }
}




// DeleteEmployeeHandler elimina un empleado específico por ID y sus reservas y consultas asociadas
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extraer parámetros de la URL
	var employee models.Employee

	// Buscar el empleado específico por ID
	if err := db.DB.First(&employee, params["id"]).Error; err != nil {
		// Verificar si el empleado no fue encontrado o si hubo un error
		if err.Error() == "record not found" {
			// Si el empleado no existe, devolver un error 404
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Employee not found"))
			return
		}
		// Manejar el error si ocurre al buscar el empleado
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	// Eliminar las reservas asociadas al empleado
	if err := db.DB.Where("employee_id = ?", employee.User.ID).Delete(&models.Reservation{}).Error; err != nil {
		// Manejar el error si ocurre al eliminar las reservas
		http.Error(w, "Failed to delete reservations", http.StatusInternalServerError)
		return
	}

	// Eliminar las consultas asociadas al empleado
	if err := db.DB.Where("employee_id = ?", employee.User.ID).Delete(&models.Consultation{}).Error; err != nil {
		// Manejar el error si ocurre al eliminar las consultas
		http.Error(w, "Failed to delete consultations", http.StatusInternalServerError)
		return
	}

	// Eliminar el empleado de la base de datos
	if err := db.DB.Unscoped().Delete(&employee).Error; err != nil {
		// Manejar el error si ocurre al eliminar el empleado
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	// Devolver un estado 200 OK si la eliminación fue exitosa
	w.WriteHeader(http.StatusOK)
}