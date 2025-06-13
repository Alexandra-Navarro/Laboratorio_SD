package config

import (
	"os"
)

// Config contiene la configuración de la aplicación
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	MQTTBroker string
}

// GetConfig retorna la configuración de la aplicación
func GetConfig() *Config {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	mqttBroker := os.Getenv("MQTT_BROKER")

	return &Config{
		DBHost:     host,
		DBPort:     5432,
		DBUser:     user,
		DBPassword: password,
		DBName:     dbname,
		MQTTBroker: mqttBroker,
	}
}
