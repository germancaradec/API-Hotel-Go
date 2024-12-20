
# Go API REST

Este proyecto es una API RESTful desarrollada en Go utilizando GORM para la interacción con una base de datos PostgreSQL. Su instalación se puede hacer clonando el repositorio o mediante una imagen de Docker. La API maneja cuatro modelos principales: Usuarios, Empleados, Reservas y Consultas. 

## Requisitos

Go 1.16 o superior + PostgreSQL

Docker (opcional)



## Instalación Estándar (Clonando el repositorio)

1 Clona el repositorio: git clone https://github.com/germancaradec/Go-API-REST-PostgresSQL.git

cd Go-API-REST-PostgresSQL

2 Configura tu base de datos PostgreSQL:

Crea una base de datos llamada gorm.

Actualiza la cadena de conexión en db/connection.go si es necesario.

3 Instala las dependencias:

go mod tidy

4 Ejecuta la aplicación:

go run main.go

La API estará disponible en http://localhost:8080.



## Instalación con Docker

### Descargar la Última Versión de la Imagen:

docker pull germancaradec/hotel-api:1.1

### Configurar y Ejecutar la Aplicación:

Crea un archivo docker-compose.yml en tu directorio de trabajo con el siguiente contenido:

services:
  db:
    image: postgres:latest
    environment:
      POSTGRES_DB: gorm
      POSTGRES_USER: tu_usuario
      POSTGRES_PASSWORD: tu_contraseña
    ports:
      - "5432:5432"

  api:
    image: germancaradec/hotel-api:1.1
    ports:
      - "8080:8080"
    depends_on:
      - db
    environment:
      DATABASE_URL: postgres://tu_usuario:tu_contraseña@db:5432/gorm?sslmode=disable

### Luego, ejecuta el siguiente comando para iniciar los contenedores:

docker-compose up --build


La API estará disponible en http://localhost:8080.


## Endpoints

### Usuarios

GET /users: Obtiene todos los usuarios.

GET /users/{id}: Obtiene un usuario específico por ID.

POST /users: Crea un nuevo usuario.

PUT /users/{id}: Actualiza un usuario existente por ID.

DELETE /users/{id}: Elimina un usuario específico por ID.

De la misma forma estan configurados los demás modelos.

### Empleados

GET /employees: Obtiene todos los empleados.

GET /employees/{id}: Obtiene un empleado específico por ID.

POST /employees: Crea un nuevo empleado.

PUT /employees/{id}: Actualiza un empleado existente por ID.

DELETE /employees/{id}: Elimina un empleado específico por ID.

## Consultas

### User

Ejemplo de Solicitud para Crear un Usuario

URL: http://localhost:8080/users

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

### Employee

Ejemplo de Solicitud para Crear un Empleado

URL: http://localhost:8080/employees

Método: POST

Cuerpo de la Solicitud:

{
  "user": {
    "first_name": "Carlos",
    "last_name": "Gómez",
    "email": "carlos.gomez@example.com",
    "reservations": [],
    "consultations": []
  },
  "position": "Gerente de Ventas",
  "salary": 75000.00,
  "department": "Ventas",
  "hire_date": "2024-11-01",
  "phone_number": "123456789"
}

Ejemplo de Respuesta al Crear un Empleado

{
    "user": {
        "ID": 8,
        "CreatedAt": "2024-11-05T10:06:32.8202087-03:00",
        "UpdatedAt": "2024-11-05T10:06:32.8202087-03:00",
        "DeletedAt": null,
        "first_name": "Carlos",
        "last_name": "Gómez",
        "email": "carlos.gomez@example.com",
        "reservations": [],
        "consultations": []
    },
    "position": "Gerente de Ventas",
    "salary": 75000,
    "department": "Ventas",
    "hire_date": "2024-11-01",
    "phone_number": "123456789"
}

### Consultation

Ejemplo de Solicitud para Crear una Consulta

URL: http://localhost:8080/consultations

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

### Reservation

Ejemplo de Solicitud para Crear una Reserva

URL: http://localhost:8080/reservations

Método: POST

Cuerpo de la Solicitud:

{
  "adults": 2,
  "check_in": "2024-11-10T14:00:00Z",
  "check_out": "2024-11-15T11:00:00Z",
  "children": 1,
  "email": "cliente@example.com",
  "number_of_rooms": 1,
  "room_type": "Suite",
  "user_id": 3
}


Ejemplo de Respuesta al Crear una Reserva

{
    "ID": 29,
    "CreatedAt": "2024-11-05T09:59:59.4954284-03:00",
    "UpdatedAt": "2024-11-05T09:59:59.4954284-03:00",
    "DeletedAt": null,
    "adults": 2,
    "check_in": "2024-11-10T14:00:00Z",
    "check_out": "2024-11-15T11:00:00Z",
    "children": 1,
    "email": "cliente@example.com",
    "number_of_rooms": 1,
    "room_type": "Suite",
    "user_id": 174
}

## Pruebas

Este proyecto incluye una serie de pruebas automatizadas para asegurar la funcionalidad de las rutas del API de usuarios y consultas. Las pruebas están implementadas utilizando el paquete de testing de Go y testify para realizar afirmaciones.

Para ejecutar las pruebas, utiliza el siguiente comando:

go test -v ./routes

### Tests Implementados

#### Tests de Usuarios

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

#### Tests de Consultas

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


### Garantías Ofrecidas por los Tests

Cobertura Completa: Los tests cubren operaciones CRUD completas para los modelos de usuarios y consultas, garantizando que cada operación (crear, leer, actualizar, eliminar) funcione correctamente.

Integridad de Datos: Aseguran que los datos sean gestionados adecuadamente, y los datos previos no interfieran con los tests a través de la limpieza y configuración de la base de datos.

Detección de Errores: Ayudan a identificar posibles errores en la implementación de los endpoints y en la interacción con la base de datos.

Todos los tests están diseñados para ejecutarse de manera independiente, asegurando que el estado de la base de datos se restablezca antes y después de cada prueba para evitar efectos colaterales.


## Conceptos y Funcionalidades Aplicadas

### Conceptos Implementados

- **API RESTful**: El proyecto implementa una API RESTful que permite la comunicación entre el cliente y el servidor a través de métodos HTTP (GET, POST, PUT, DELETE).
  
- **ORM (Object-Relational Mapping)**: Se utiliza GORM para facilitar la interacción con la base de datos PostgreSQL, permitiendo realizar operaciones CRUD sin escribir consultas SQL directamente.

- **Manejo de Errores**: Se proporciona un manejo adecuado de errores para informar al cliente sobre problemas en las solicitudes, garantizando una mejor experiencia de usuario.

- **Reutilización de estructuras**: Se reutiliza la estructura de Usuario en el modelo Empleado, mediante la composición a partir de un campo anónimo.

