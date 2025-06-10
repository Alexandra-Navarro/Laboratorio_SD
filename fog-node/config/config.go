package config

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
	return &Config{
		DBHost:     "localhost",
		DBPort:     5432,
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "MonitoreoDB",
		MQTTBroker: "tcp://localhost:1883",
	}
}
