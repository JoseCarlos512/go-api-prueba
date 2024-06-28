# Usar una imagen base de Go
FROM golang:1.18-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto en el que la app se ejecutará
EXPOSE 3100

# Ejecutar la aplicación
CMD ["./main"]
