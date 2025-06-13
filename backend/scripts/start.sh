#!/bin/sh

# Generar el archivo .env dinámicamente
cat <<EOF > .env
POSTGRES_HOST=${POSTGRES_HOST:-localhost}
POSTGRES_USER=${POSTGRES_USER:-postgres}
POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-postgres}
POSTGRES_DB=${POSTGRES_DB:-mydatabase}
POSTGRES_PORT=${POSTGRES_PORT:-5432}
POSTGRES_SSLMODE=${POSTGRES_SSLMODE:-disable}
POSTGRES_TIMEZONE=${POSTGRES_TIMEZONE:-UTC}
APP_PORT=${APP_PORT:-8080}
EOF

# Mostrar el contenido del .env para verificarlo (opcional)
echo "Archivo .env generado:"
cat .env

# Ejecutar la aplicación
./main