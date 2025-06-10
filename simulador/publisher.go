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
	return SensorData{
		SensorID:    sensorID,
		Timestamp:   time.Now().UTC(),
		Temperatura: 20 + rand.Float64()*15,    // 20-35°C
		Humedad:     30 + rand.Float64()*50,    // 30-80%
		CO2:         400 + rand.Float64()*1600, // 400-2000 ppm
		Ruido:       30 + rand.Float64()*70,    // 30-100 dB
		Luz:         100 + rand.Float64()*900,  // 100-1000 lux
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

	for i := 0; i <= maxMediciones; i++ {
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
		time.Sleep(15 * time.Second)
	}
	log.Println("Simulación finalizada: se enviaron 10 mediciones por sala.")
}
