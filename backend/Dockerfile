FROM golang:1.19-alpine

WORKDIR /app

# Instalar dependencias del sistema
RUN apk add --no-cache gcc musl-dev

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"] 