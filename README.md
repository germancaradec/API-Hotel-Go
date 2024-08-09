
API REST desarrollada en Go, utilizando GORM como orm, PostgreSQL como base de datos desde un contenedor de Docker. 
Creamos el servidor a partir del módulo gorilla mux, que trabaja por encima del módulo propio de Go, net.http 
En desarrollo utilizamos air como live reload (instalado de forma global)

go mod init + URL de github

go get -u github.com/gorilla/mux

go run .

go install github.com/air-verse/air@latest

air init

air

Instalar gorm y driver para postgres:
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

Crear un contenedor de postgres desde docker. 
Crear un usuario con contraseña, exponer postgresql en el puerto 5432, y usarlo en modo Detach 
para que se ejecute en segundo plano:
docker run --name some-postgres -e POSTGRES_USER=german -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

Para conectarnos:
docker exec -it some-postgres bash
psql -U german --password

Ver base de datos:
\l

Crear base de datos (necesitamos el nombre para colocarlo en gorm.Open):
CREATE DATABASE gorm;

Conectarnos con la base de datos:
\c gorm
password:...

Ver tablas:
\d

Ver estructura de tabla:
\d tasks



Ejemplo de solicitud para Update:
http://localhost:3000/users/2

{
  "first_name": "Gabriela",
  "last_name": "Gomez",
  "email": "ggomez@gmail.com"
}

El usuario de id 2 será actualizado y el servidor nos devolverá los datos:

{
  "ID": 2,
  "CreatedAt": "2024-08-03T10:45:33.932581-03:00",
  "UpdatedAt": "2024-08-06T11:30:10.9677082-03:00",
  "DeletedAt": null,
  "first_name": "Gabriela",
  "last_name": "Gomez",
  "email": "ggomez@gmail.com",
  "reservations": []
}





Tests: este proyecto incluye una serie de pruebas automatizadas para asegurar la funcionalidad de las rutas del API de usuarios. 
Los tests están implementados utilizando el paquete de testing de Go y testify para realizar afirmaciones.

Tests Implementados
TestGetUsersHandler

-Propósito: Verificar que la ruta GET /users devuelve la lista de usuarios.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene un usuario con el primer nombre "John" y el apellido "Doe".
TestGetUserHandler

-Propósito: Verificar que la ruta GET /users/{id} devuelve los detalles de un usuario específico.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene el usuario con el primer nombre "Jane" y el apellido "Doe".
TestPostUserHandler

-Propósito: Verificar que la ruta POST /users permite la creación de un nuevo usuario.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene el usuario con el primer nombre "Alice" y el apellido "Smith".
TestUpdateUserHandler

-Propósito: Verificar que la ruta PUT /users/{id} permite actualizar la información de un usuario.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El cuerpo de la respuesta contiene el usuario con el primer nombre "Robert" y el apellido "Johnson".
TestDeleteUserHandler

-Propósito: Verificar que la ruta DELETE /users/{id} elimina un usuario.
Validaciones:
La respuesta tiene un código de estado HTTP 200.
El usuario eliminado no se encuentra en la base de datos.

Garantías Ofrecidas por los Tests

Cobertura Completa: Los tests cubren operaciones CRUD completas para el modelo de usuario, garantizando que cada operación (crear, leer, actualizar, eliminar) funcione correctamente.

Integridad de Datos: Aseguran que los datos sean gestionados adecuadamente, y los datos previos no interfieran con los tests a través de la limpieza y configuración de la base de datos.

Detección de Errores: Ayudan a identificar posibles errores en la implementación de los endpoints y en la interacción con la base de datos.

Todos los tests están diseñados para ejecutarse de manera independiente, asegurando que el estado de la base de datos se restablezca antes y después de cada prueba para evitar efectos colaterales.


Agregar la Dependencia de testify:
Ejecuta el siguiente comando en tu directorio de proyecto:
go get github.com/stretchr/testify

Ejecutar tus pruebas con el siguiente comando:
go test -v ./routes

Si es necesario actualizar el Archivo go.sum:
Ejecutar el siguiente comando para actualizar el archivo go.sum con la entrada faltante:
go mod tidy






Ejemplos de json para cargar datos:
Usuarios
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com"
}
{
  "first_name": "Bob",
  "last_name": "Williams",
  "email": "bob.williams@example.com"
}


Consultas
{
  "phone": "+1234567890",
  "consultation": "Looking for advice on our upcoming project.",
  "more_info": true,
  "user_id": 119
}
{
  "phone": "+1-555-9876",
  "consultation": "Discussing new project ideas and potential collaboration opportunities.",
  "more_info": false,
  "user_id": 120
}


Reservas
{
  "adults": 2,
  "check_in": "2024-09-01T15:00:00Z",
  "check_out": "2024-09-05T11:00:00Z",
  "children": 1,
  "email": "john.doe@example.com",
  "number_of_rooms": 1,
  "room_type": "Double",
  "user_id": 119
}
{
  "adults": 2,
  "check_in": "2024-09-15T15:00:00Z",
  "check_out": "2024-09-20T11:00:00Z",
  "children": 1,
  "email": "bob.williams@example.com",
  "number_of_rooms": 1,
  "room_type": "Double",
  "user_id": 120
}

