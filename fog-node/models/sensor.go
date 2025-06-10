package models

import "time"

// SensorData representa los datos recibidos de un sensor
type SensorData struct {
	SensorID    string    `json:"sensor_id"`
	Timestamp   time.Time `json:"timestamp"`
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	CO2         float32   `json:"co2"`
	Noise       float32   `json:"noise"`
	Light       float32   `json:"light"`
}

// Medicion representa una medición almacenada en la base de datos
type Medicion struct {
	Fecha        time.Time `json:"fecha"`
	Variable     string    `json:"variable"`
	Valor        float64   `json:"valor"`
	UnidadMedida string    `json:"unidad_medida"`
}

// Alerta representa una alerta generada por el sistema
type Alerta struct {
	Tipo           string    `json:"tipo"`
	Descripcion    string    `json:"descripcion"`
	ValorDetectado float64   `json:"valor_detectado"`
	Umbral         string    `json:"umbral"`
	Fecha          time.Time `json:"fecha"`
}
