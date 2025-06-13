# Sistema de Monitoreo Ambiental en Aulas Académicas  
**Proyecto de Sistemas Distribuidos – Universidad de Santiago de Chile**

Este proyecto implementa un sistema distribuido de monitoreo ambiental en contextos educativos, basado en una arquitectura **Edge–Fog–Cloud**, con el objetivo de mejorar las condiciones físicas del entorno de aprendizaje. Desarrollado como parte del curso *Sistemas Distribuidos*, el sistema permite la simulación de sensores IoT, procesamiento local de datos y visualización en tiempo real a través de dashboards interactivos.

---

## Arquitectura del Sistema

El sistema está estructurado en componentes modulares y desplegado completamente mediante contenedores Docker:

- **Simulador (ESP32 simulado)**: Emula sensores ambientales (T°, CO₂, humedad, ruido, iluminación).
- **Fog Node (Raspberry Pi 4)**: Procesa datos, detecta alertas y sincroniza con la nube.
- **Backend (Go + Gin)**: Expone una API REST para consultas, alertas, y configuración.
- **Frontend (Vue.js)**: Permite visualizar alertas y gestionar umbrales de forma dinámica.
- **Grafana**: Visualización analítica en tiempo real conectada a PostgreSQL.
- **PostgreSQL**: Base de datos estructurada para almacenamiento local y en la nube.
- **Mosquitto**: Broker MQTT que intermedia la comunicación entre sensores y backend.

---

## Requisitos

- Docker y Docker Compose
- Git
- Navegador web moderno

---

## Variables de Entorno

Crear un archivo `.env` en la raíz del proyecto con la siguiente configuración base:

```env
# PostgreSQL
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=MonitoreoDB
POSTGRES_HOST=postgres
POSTGRES_PORT=5432

# MQTT
MQTT_USER=monitoring
MQTT_PASSWORD=monitoring123

# Grafana
GRAFANA_ADMIN_USER=admin
GRAFANA_ADMIN_PASSWORD=admin123
```

---

## Instalación y Ejecución

1. **Clonar el repositorio:**
```bash
git clone <url-del-repositorio>
cd <nombre-del-directorio>
```

2. **Levantar el sistema por primera vez:**
```bash
docker-compose up --build -d
```

3. **Acceder a las interfaces:**
- Frontend: http://localhost
- API REST (Backend): http://localhost:8080
- Dashboard (Grafana): http://localhost:3000

## Credenciales de Acceso

### Sistema Principal
- **Usuario**: usuario@demo.cl
- **Contraseña**: usuario123

### Grafana
- **Usuario**: admin
- **Contraseña**: admin123

---

## Simulación de Sensores

El simulador genera datos automáticamente hasta alcanzar un número predefinido de iteraciones (3 iteraciones por cada sala). Para generar nuevos datos se puede reiniciar la simulación:

```bash
docker-compose restart simulador
```

---

## Estructura del Proyecto

```
.
├── backend/             # API REST en Go (Gin)
├── frontend/            # Interfaz administrativa (Vue.js + Vuetify)
├── fog-node/            # Lógica de procesamiento local (Fog)
├── grafana/             # Dashboards y configuración de Grafana
├── mosquitto/           # Configuración del broker MQTT
├── simulador/           # Generador de datos simulados
├── CreateTables.sql     # Script de creación de la base de datos
├── InsertTable.sql      # Script de inserción de datos iniciales
├── DeleteTables.sql     # Script de borrado (para reinicio)
└── docker-compose.yml   # Orquestación de servicios
```

---

## Historias de Usuario Implementadas

Este proyecto incorpora tres historias de usuario clave, alineadas con la arquitectura distribuida y los patrones de comunicación requeridos:

1. **Captura y transmisión de datos ambientales (HU1)**  
   Como usuario responsable del monitoreo técnico, deseo que los sensores ESP32 recojan datos ambientales (temperatura, CO₂, humedad, etc.) y los envíen en tiempo real al servidor local mediante MQTT, para asegurar una supervisión oportuna y continua.

2. **Visualización y gestión de alertas (HU3)**  
   Como usuario del sistema de monitoreo, deseo ver las alertas generadas por los sensores (y su historial), para tomar decisiones rápidas ante anomalías.

3. **Configuración personalizada de umbrales por sala (HU adicional)**  
   Como usuario administrador del sistema, deseo poder modificar los umbrales de alerta para cada variable ambiental por sala, desde una interfaz accesible, para adaptar el sistema a las condiciones específicas de cada espacio educativo.

---

## Visualización y Análisis

- **Grafana** permite explorar datos históricos y en tiempo real por sala, variable y periodo.
- **Vue.js** permite configurar umbrales personalizados, ver estados de sensores y acceder a alertas activas e históricas.

---

## Desarrollo y Tecnología

- **Lenguajes**: Go, JavaScript, SQL
- **Frameworks**: Gin (Go), Vue.js, Vuetify
- **Comunicación**: MQTT, HTTPS
- **Persistencia**: PostgreSQL
- **Contenedores**: Docker + docker-compose