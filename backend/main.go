package main

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	mqttClient mqtt.Client
)

func initDB() {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=5432 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func initMQTT() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + os.Getenv("MQTT_BROKER") + ":" + os.Getenv("MQTT_PORT"))
	opts.SetClientID("backend-service")

	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Failed to connect to MQTT broker:", token.Error())
	}
}

func main() {
}
