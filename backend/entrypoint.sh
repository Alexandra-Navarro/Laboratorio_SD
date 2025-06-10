#!/bin/sh

# Esperar a que PostgreSQL esté listo
echo "Esperando a que PostgreSQL esté listo..."
while ! nc -z postgres 5432; do
  sleep 1
done
echo "PostgreSQL está listo!"

# Iniciar la aplicación
exec ./main 