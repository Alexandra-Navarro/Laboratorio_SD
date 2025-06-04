# Sistema de Monitoreo Ambiental - PoC

Este proyecto implementa una prueba de concepto para un sistema distribuido de monitoreo ambiental en aulas escolares, utilizando una arquitectura Edge-Fog-Cloud.

## 🏗️ Arquitectura

El sistema está compuesto por los siguientes componentes:

- **Simulador MQTT**: Genera datos simulados de sensores ambientales
- **Backend (Fog Node)**: Procesa y valida los datos, genera alertas
- **Frontend**: Interfaz web para visualización y configuración
- **Grafana**: Dashboard para visualización de métricas
- **PostgreSQL**: Base de datos para almacenamiento local
- **MQTT Broker**: Broker para la comunicación entre componentes

## 🚀 Requisitos Previos

- Docker y Docker Compose
- Git
- Go 1.19 o superior (para desarrollo local)
- Node.js 16 o superior (para desarrollo local)

## 🛠️ Instalación

1. Clonar el repositorio:
```bash
git clone <url-del-repositorio>
cd <nombre-del-directorio>
```

2. Configurar variables de entorno:
```bash
cp .env.example .env
# Editar .env con tus configuraciones
```

3. Iniciar los servicios:
```bash
docker-compose up -d
```

## 📊 Acceso a los Servicios

- Frontend: http://localhost
- Grafana: http://localhost:3000
- Backend API: http://localhost:8080
- MQTT Broker: localhost:1883
- PostgreSQL: localhost:5432

## 🔧 Configuración

### Credenciales por Defecto

- Grafana:
  - Usuario: admin
  - Contraseña: admin123

- PostgreSQL:
  - Usuario: monitoring
  - Contraseña: monitoring123
  - Base de datos: monitoring_db

## 📝 API Endpoints

### Backend

- `GET /api/sensors`: Lista todos los sensores
- `GET /api/alerts`: Obtiene alertas activas
- `POST /api/thresholds`: Configura umbrales por sala
- `GET /api/metrics`: Obtiene métricas históricas

## 🔍 Estructura del Proyecto

```
.
├── backend/          # Servidor Go con Gin
├── frontend/         # Aplicación Vue.js
├── simulador/        # Generador de datos MQTT
├── grafana/          # Configuración de Grafana
├── docker-compose.yml
└── README.md
```

## 🧪 Desarrollo

Para desarrollo local:

1. Backend:
```bash
cd backend
go mod download
go run main.go
```

2. Frontend:
```bash
cd frontend
npm install
npm run serve
```

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. 