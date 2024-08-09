
Go API REST con PostgreSQL

Este proyecto es una API RESTful desarrollada en Go utilizando GORM para la interacción con una base de datos PostgreSQL. La API maneja tres modelos principales: Usuarios, Reservas y Consultas.

Requisitos
Go 1.16 o superior
PostgreSQL

Instalación
Clona el repositorio:
git clone https://github.com/germancaradec/Go-API-REST-PostgresSQL.git
cd Go-API-REST-PostgresSQL

Configura tu base de datos PostgreSQL:
Crea una base de datos llamada gorm.
Actualiza la cadena de conexión en db/connection.go si es necesario.

Instala las dependencias:
go mod tidy

Ejecuta la aplicación:
go run main.go
La API estará disponible en http://localhost:3000.

Endpoints

Usuarios
GET /users: Obtiene todos los usuarios.
GET /users/{id}: Obtiene un usuario específico por ID.
POST /users: Crea un nuevo usuario.
PUT /users/{id}: Actualiza un usuario existente por ID.
DELETE /users/{id}: Elimina un usuario específico por ID.

Reservas
GET /reservations: Obtiene todas las reservas.
GET /reservations/{id}: Obtiene una reserva específica por ID.
POST /reservations: Crea una nueva reserva.
PUT /reservations/{id}: Actualiza una reserva existente por ID.
DELETE /reservations/{id}: Elimina una reserva específica por ID.

Consultas
GET /consultations: Obtiene todas las consultas.
GET /consultations/{id}: Obtiene una consulta específica por ID.
POST /consultations: Crea una nueva consulta.
PUT /consultations/{id}: Actualiza una consulta existente por ID.
DELETE /consultations/{id}: Elimina una consulta específica por ID.

Modelos

User
Ejemplo de Solicitud para Crear un Usuario
URL: http://localhost:3000/users
Método: POST

Cuerpo de la Solicitud:

{
  "first_name": "Gabriela",
  "last_name": "Gomez",
  "email": "ggomez@gmail.com"
}

Ejemplo de Respuesta al Crear un Usuario

{
  "ID": 2,
  "CreatedAt": "2024-08-03T10:45:33.932581-03:00",
  "UpdatedAt": "2024-08-06T11:30:10.9677082-03:00",
  "DeletedAt": null,
  "first_name": "Gabriela",
  "last_name": "Gomez",
  "email": "ggomez@gmail.com",
  "reservations": [],
  "consultations": []
}

Consultation

Ejemplo de Solicitud para Crear una Consulta
URL: http://localhost:3000/consultations
Método: POST

Cuerpo de la Solicitud:

{
  "phone": "123456789",
  "consultation": "¿Información sobre la reserva?",
  "more_info": true,
  "user_id": 2
}

Ejemplo de Respuesta al Crear una Consulta

{
  "ID": 1,
  "CreatedAt": "2024-08-03T10:45:33.932581-03:00",
  "UpdatedAt": "2024-08-06T11:30:10.9677082-03:00",
  "DeletedAt": null,
  "phone": "123456789",
  "consultation": "¿Información sobre la reserva?",
  "more_info": true,
  "user_id": 2
}

Pruebas

Este proyecto incluye una serie de pruebas automatizadas para asegurar la funcionalidad de las rutas del API de usuarios y consultas. Las pruebas están implementadas utilizando el paquete de testing de Go y testify para realizar afirmaciones.

Para ejecutar las pruebas, utiliza el siguiente comando:

go test -v ./routes

Tests Implementados

Tests de Usuarios

TestGetUsersHandler:
Propósito: Verificar que la ruta GET /users devuelve la lista de usuarios.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene al menos un usuario con el primer nombre "John" y el apellido "Doe".

TestGetUserHandler:
Propósito: Verificar que la ruta GET /users/{id} devuelve los detalles de un usuario específico.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene el usuario con el primer nombre "Jane" y el apellido "Doe".

TestPostUserHandler:
Propósito: Verificar que la ruta POST /users permite la creación de un nuevo usuario.
Validaciones:
La respuesta tiene un código de estado HTTP 201.
El cuerpo de la respuesta contiene el usuario con el primer nombre "Alice" y el apellido "Smith".

TestUpdateUserHandler:
Propósito: Verificar que la ruta PUT /users/{id} permite actualizar la información de un usuario.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene el usuario con el primer nombre "Robert" y el apellido "Johnson".

TestDeleteUserHandler:
Propósito: Verificar que la ruta DELETE /users/{id} elimina un usuario.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El usuario eliminado no se encuentra en la base de datos.

Tests de Consultas

TestGetConsultationsHandler:
Propósito: Verificar que la ruta GET /consultations devuelve la lista de consultas.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene al menos una consulta con el teléfono "123456789" y la consulta "¿Información sobre la reserva?".

TestGetConsultationHandler:
Propósito: Verificar que la ruta GET /consultations/{id} devuelve los detalles de una consulta específica.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene la consulta con el teléfono "987654321".

TestPostConsultationHandler:
Propósito: Verificar que la ruta POST /consultations permite la creación de una nueva consulta.
Validaciones:
La respuesta tiene un código de estado HTTP 201.
El cuerpo de la respuesta contiene la consulta con el teléfono "555555555".

TestUpdateConsultationHandler:
Propósito: Verificar que la ruta PUT /consultations/{id} permite actualizar la información de una consulta.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene la consulta actualizada.

TestDeleteConsultationHandler:
Propósito: Verificar que la ruta DELETE /consultations/{id} elimina una consulta.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
La consulta eliminada no se encuentra en la base de datos.


Garantías Ofrecidas por los Tests

Cobertura Completa: Los tests cubren operaciones CRUD completas para los modelos de usuarios y consultas, garantizando que cada operación (crear, leer, actualizar, eliminar) funcione correctamente.

Integridad de Datos: Aseguran que los datos sean gestionados adecuadamente, y los datos previos no interfieran con los tests a través de la limpieza y configuración de la base de datos.

Detección de Errores: Ayudan a identificar posibles errores en la implementación de los endpoints y en la interacción con la base de datos.

Todos los tests están diseñados para ejecutarse de manera independiente, asegurando que el estado de la base de datos se restablezca antes y después de cada prueba para evitar efectos colaterales.

