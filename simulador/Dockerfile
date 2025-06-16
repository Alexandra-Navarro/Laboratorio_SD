# Usa Go 1.23 en Alpine para cumplir con go ≥ 1.22 del módulo simulador
FROM golang:1.23-alpine

WORKDIR /app

# Instala toolchain C, git y certificados TLS necesarios para go mod
RUN apk add --no-cache \
      gcc \
      musl-dev \
      git \
      ca-certificates

# Copia todo el proyecto
COPY . .

# Sincroniza go.mod/go.sum con las importaciones del código
RUN go mod tidy

# Compila la aplicación simulador en el binario 'publisher'
RUN go build -o publisher .

# Arranca el simulador
CMD ["./publisher"]
