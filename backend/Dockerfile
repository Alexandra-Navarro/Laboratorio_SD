FROM golang:1.23-alpine
WORKDIR /app

# Instala herramientas necesarias
RUN apk add --no-cache \
      netcat-openbsd \
      git \
      ca-certificates \
      gcc \
      musl-dev

# Copia todo el proyecto
COPY . .

# Limpia y actualiza go.mod/go.sum según el código
RUN go mod tidy

# Compila la aplicación
RUN go build -o main .

EXPOSE 8080
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]