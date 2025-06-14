package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"fog-node/config"
	"fog-node/controllers"
	"fog-node/models"
	"fog-node/routes"
	"fog-node/services"
)

var db *sql.DB

func main() {
	// Cargar configuración
	cfg := config.GetConfig()

	// Inicializar base de datos
	initDB(cfg)

	// Inicializar servicios
	sensorService := services.NewSensorService(db)

	// Inicializar controladores
	sensorController := controllers.NewSensorController(sensorService)

	// Inicializar MQTT
	initMQTT(cfg, sensorService)

	// Inicializar API REST
	initAPI(sensorController)

	// Mantener el programa en ejecución
	select {}
}

func initDB(cfg *config.Config) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil && db.Ping() == nil {
			log.Println("Conexión a PostgreSQL establecida")
			return
		}
		log.Printf("Esperando conexión a PostgreSQL... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("No se pudo conectar a PostgreSQL: %v", err)
}

func initMQTT(cfg *config.Config, sensorService *services.SensorService) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.MQTTBroker)
	opts.SetClientID("fog_node_1")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		var data models.SensorData
		err := json.Unmarshal(msg.Payload(), &data)
		if err != nil {
			log.Printf("Error al decodificar JSON: %v", err)
			return
		}

		// Extraer nombre de sala del tópico MQTT
		topic := msg.Topic()
		parts := strings.Split(topic, "/")
		if len(parts) != 3 {
			log.Printf("Formato de tópico inesperado: %s", topic)
			return
		}
		nombreSala := parts[1]

		// Buscar el ID numérico de la sala en la base de datos
		var salaID int
		err = db.QueryRow("SELECT id FROM sala WHERE nombre = $1", nombreSala).Scan(&salaID)
		if err != nil {
			log.Printf("No se encontró la sala '%s' en la base de datos: %v", nombreSala, err)
			return
		}

		// Procesar datos del sensor
		err = sensorService.ProcessSensorData(data, salaID)
		if err != nil {
			log.Printf("Error al procesar datos del sensor: %v", err)
			return
		}

		log.Printf("Datos recibidos y almacenados para la sala %d (%s)", salaID, nombreSala)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Suscribirse a todos los tópicos de aulas
	if token := client.Subscribe("aula/+/datos", 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Println("Conexión MQTT establecida")
}

func initAPI(sensorController *controllers.SensorController) {
	r := gin.Default()

	// Configurar rutas
	routes.SetupRoutes(r, sensorController)

	// Iniciar servidor
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	}()
}
