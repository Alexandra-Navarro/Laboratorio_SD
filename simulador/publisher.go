package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	RoomID    string    `json:"room_id"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	sensorTypes = []string{"temperature", "humidity", "co2", "noise", "light"}
	rooms       = []string{"A101", "A102", "A103", "B101", "B102"}
)

func generateSensorData() SensorData {
	sensorType := sensorTypes[rand.Intn(len(sensorTypes))]
	roomID := rooms[rand.Intn(len(rooms))]
	var value float64

	switch sensorType {
	case "temperature":
		value = 20 + rand.Float64()*15 // 20-35°C
	case "humidity":
		value = 30 + rand.Float64()*50 // 30-80%
	case "co2":
		value = 400 + rand.Float64()*1600 // 400-2000 ppm
	case "noise":
		value = 30 + rand.Float64()*70 // 30-100 dB
	case "light":
		value = 100 + rand.Float64()*900 // 100-1000 lux
	}

	return SensorData{
		RoomID:    roomID,
		Type:      sensorType,
		Value:     value,
		Timestamp: time.Now(),
	}
}

func main() {
	// Configurar cliente MQTT
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + os.Getenv("MQTT_BROKER") + ":" + os.Getenv("MQTT_PORT"))
	opts.SetClientID("sensor-simulator")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Error connecting to MQTT broker:", token.Error())
	}

	// Generar y publicar datos cada 15 minutos
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data := generateSensorData()
			payload, err := json.Marshal(data)
			if err != nil {
				log.Printf("Error marshaling data: %v", err)
				continue
			}

			topic := "sensors/" + data.RoomID + "/" + data.Type
			token := client.Publish(topic, 0, false, payload)
			token.Wait()

			if token.Error() != nil {
				log.Printf("Error publishing to %s: %v", topic, token.Error())
			} else {
				log.Printf("Published to %s: %s", topic, string(payload))
			}
		}
	}
}
