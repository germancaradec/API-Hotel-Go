

# Usar la imagen base de Golang
FROM golang:1.23

# Crear y establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos desde el directorio actual al contenedor
COPY . /app

# Compilar la aplicación
RUN make build

# Exponer el puerto en el contenedor
EXPOSE 8080

# Descargar wait-for-it y hacerla ejecutable
RUN curl -o /wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && chmod +x /wait-for-it.sh


# Ejecutar la aplicación
CMD ["/wait-for-it.sh", "db:5432", "--", "./ms-api"]

