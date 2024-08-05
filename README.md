
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




