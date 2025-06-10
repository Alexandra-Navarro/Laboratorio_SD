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
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
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
