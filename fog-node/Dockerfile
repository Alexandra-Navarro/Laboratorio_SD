FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod ./
COPY go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o fog-node .

# Imagen final
FROM alpine:latest

WORKDIR /app

# Copiar el binario compilado
COPY --from=builder /app/fog-node .

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./fog-node"] 