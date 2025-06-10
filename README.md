# Sistema de Monitoreo Ambiental - Edge-Fog-Cloud

Este proyecto implementa un sistema distribuido de monitoreo ambiental para salas de clases, basado en una arquitectura Edge-Fog-Cloud.

## Estructura del Proyecto

```
.
├── esp32/              # Código para el ESP32
│   ├── main.go        # Programa principal del ESP32
│   └── go.mod         # Dependencias de Go
├── fog-node/          # Código del fog node
│   ├── main.go        # Programa principal del fog node
│   ├── go.mod         # Dependencias de Go
│   └── Dockerfile     # Configuración de Docker
└── README.md          # Este archivo
```

## Requisitos

### ESP32
- TinyGo instalado (versión 0.27.0 o superior)
- ESP32 con los siguientes sensores:
  - DHT22 (Temperatura y Humedad)
  - MH-Z19 (CO2)
  - KY-038 (Ruido)
  - LDR (Iluminación)
- Acceso a red WiFi
- Conexión al broker MQTT

### Fog Node
- Docker instalado
- PostgreSQL instalado
- Broker MQTT (por ejemplo, Mosquitto)

## Configuración

### ESP32
1. Edita el archivo `esp32/main.go` y configura:
   - `WIFI_SSID` y `WIFI_PASSWORD` con tus credenciales WiFi
   - `MQTT_BROKER` con la dirección de tu broker MQTT
   - `MQTT_USER` y `MQTT_PASSWORD` si tu broker requiere autenticación
   - `ROOM_ID` con el ID de la sala
   - `READ_INTERVAL` si deseas cambiar el intervalo de lectura (por defecto 15 segundos)

2. Verifica la configuración de pines:
   ```go
   const (
       DHT_PIN = machine.GPIO4    // DHT22
       CO2_RX_PIN = machine.GPIO16 // MH-Z19
       CO2_TX_PIN = machine.GPIO17 // MH-Z19
       NOISE_PIN = machine.ADC0    // KY-038
       LDR_PIN = machine.ADC1      // LDR
   )
   ```
   Ajusta estos pines según tu configuración física.

### Fog Node
1. Edita el archivo `fog-node/main.go` y configura:
   - Credenciales de PostgreSQL
   - Dirección del broker MQTT

## Compilación y Ejecución

### ESP32
1. Instalar TinyGo:
   ```bash
   # Windows (usando scoop)
   scoop install tinygo

   # Linux/macOS
   brew install tinygo
   ```

2. Compilar el código:
   ```bash
   cd esp32
   tinygo build -o esp32.uf2 -target esp32 main.go
   ```

3. Flashear el ESP32:
   ```bash
   tinygo flash -target esp32 main.go
   ```

4. Monitorear logs:
   ```bash
   # Encuentra el puerto serie
   tinygo monitor
   ```

### Fog Node
1. Construir la imagen Docker:
   ```bash
   cd fog-node
   docker build -t fog-node .
   ```

2. Ejecutar el contenedor:
   ```bash
   docker run -d \
     --name fog-node \
     -p 8080:8080 \
     -e DB_HOST=host.docker.internal \
     -e DB_PORT=5432 \
     -e DB_USER=postgres \
     -e DB_PASSWORD=postgres \
     -e DB_NAME=monitoreo_ambiental \
     fog-node
   ```

3. Verificar logs:
   ```bash
   docker logs fog-node
   ```

## Pruebas

### Prueba del ESP32
1. Verifica la conexión WiFi:
   - Los logs mostrarán "Conectado a WiFi" si la conexión es exitosa

2. Verifica la conexión MQTT:
   - Los logs mostrarán "Conexión MQTT establecida" si la conexión es exitosa

3. Verifica las lecturas de sensores:
   - Los logs mostrarán los valores leídos de cada sensor
   - Si hay errores, se mostrarán mensajes específicos

### Prueba del Fog Node
1. Verifica la conexión a PostgreSQL:
   - Los logs mostrarán "Conexión a PostgreSQL establecida"

2. Verifica la recepción de datos:
   - Los logs mostrarán "Datos recibidos y almacenados" cuando lleguen datos

3. Prueba la API REST:
   ```bash
   # Obtener mediciones recientes
   curl http://localhost:8080/api/mediciones/1

   # Obtener alertas
   curl http://localhost:8080/api/alertas/1
   ```

## Solución de Problemas

### ESP32
- **No se conecta al WiFi**:
  - Verifica las credenciales WiFi
  - Asegúrate de que la red esté en el rango del ESP32
  - Verifica que el adaptador WiFi esté correctamente conectado

- **No se conecta al MQTT**:
  - Verifica la dirección del broker
  - Verifica las credenciales MQTT si son necesarias
  - Asegúrate de que el broker esté accesible

- **Errores en sensores**:
  - Verifica las conexiones físicas
  - Verifica la configuración de pines
  - Asegúrate de que los sensores estén alimentados correctamente

### Fog Node
- **No se conecta a PostgreSQL**:
  - Verifica las credenciales
  - Asegúrate de que PostgreSQL esté accesible
  - Verifica que la base de datos exista

- **No recibe datos MQTT**:
  - Verifica la conexión al broker
  - Verifica los tópicos suscritos
  - Revisa los logs para errores específicos

## Contribuir

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request 