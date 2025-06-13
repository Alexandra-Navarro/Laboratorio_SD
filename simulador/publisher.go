package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	SensorID    string    `json:"sensor_id"`
	Timestamp   time.Time `json:"timestamp"`
	Temperatura float64   `json:"temperatura"`
	Humedad     float64   `json:"humedad"`
	CO2         float64   `json:"co2"`
	Ruido       float64   `json:"ruido"`
	Luz         float64   `json:"luz"`
}

var (
	sensorID = "esp32_simulado_01"
	rooms    = []string{"A101", "A102", "A103", "B101", "B102"}
)

const maxMediciones = 3

func generateSensorData(roomID string) SensorData {
	// Generar valores base con una distribución más realista
	// Temperatura: 20-24°C normal, ocasionalmente fuera de rango
	temperatura := 20.0 + rand.Float64()*4.0
	if rand.Float64() < 0.2 { // 20% de probabilidad de temperatura fuera de rango
		if rand.Float64() < 0.5 {
			temperatura = 18.0 + rand.Float64()*2.0 // Temperatura baja
		} else {
			temperatura = 24.0 + rand.Float64()*4.0 // Temperatura alta
		}
	}

	// Humedad: 40-60% normal, ocasionalmente fuera de rango
	humedad := 40.0 + rand.Float64()*20.0
	if rand.Float64() < 0.15 { // 15% de probabilidad de humedad anormal
		if rand.Float64() < 0.5 {
			humedad = 25.0 + rand.Float64()*5.0 // Humedad baja
		} else {
			humedad = 60.0 + rand.Float64()*15.0 // Humedad alta
		}
	}

	// CO2: 400-1000 ppm normal, ocasionalmente alto
	co2 := 400.0 + rand.Float64()*600.0
	if rand.Float64() < 0.25 { // 25% de probabilidad de CO2 alto
		co2 = 1000.0 + rand.Float64()*1000.0
	}

	// Ruido: < 35 dB normal, ocasionalmente alto
	ruido := 20.0 + rand.Float64()*15.0
	if rand.Float64() < 0.1 { // 10% de probabilidad de ruido alto
		ruido = 35.0 + rand.Float64()*30.0
	}

	// Luz: 300-700 lux normal, ocasionalmente fuera de rango
	luz := 300.0 + rand.Float64()*400.0
	if rand.Float64() < 0.2 { // 20% de probabilidad de luz anormal
		luz = 100.0 + rand.Float64()*900.0
	}

	return SensorData{
		SensorID:    sensorID,
		Timestamp:   time.Now().UTC(),
		Temperatura: temperatura,
		Humedad:     humedad,
		CO2:         co2,
		Ruido:       ruido,
		Luz:         luz,
	}
}

func main() {
	// Configurar cliente MQTT
	opts := mqtt.NewClientOptions()
	broker := os.Getenv("MQTT_BROKER")
	port := os.Getenv("MQTT_PORT")
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	opts.SetClientID("sensor-simulator")
	opts.SetUsername(os.Getenv("MQTT_USER"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(1 * time.Minute)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Error connecting to MQTT broker:", token.Error())
	}
	log.Printf("Connected to MQTT broker at %s:%s", broker, port)

	log.Println("Esperando 1 minuto para asegurar que el subscriber esté conectado...")
	time.Sleep(1 * time.Minute)

	for i := 0; i < maxMediciones; i++ {
		for _, roomID := range rooms {
			data := generateSensorData(roomID)
			payload, err := json.Marshal(data)
			if err != nil {
				log.Printf("Error marshaling data: %v", err)
				continue
			}

			topic := fmt.Sprintf("aula/%s/datos", roomID)
			token := client.Publish(topic, 0, false, payload)
			token.Wait()

			if token.Error() != nil {
				log.Printf("Error publishing to %s: %v", topic, token.Error())
			} else {
				log.Printf("Published to %s: %s", topic, string(payload))
			}
		}
		time.Sleep(15 * time.Minute)
	}
	log.Printf("Simulación finalizada: se enviaron %d mediciones por sala.", maxMediciones)
}
