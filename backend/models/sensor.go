package models

import "time"

type Sensor struct {
	ID               int       `gorm:"primaryKey; autoIncrement" json:"id"`
	Modelo           string    `json:"modelo"`
	Estado           string    `json:"estado"`
	FechaInstalacion time.Time `json:"fecha_instalacion"`
	SalaID           int       `json:"sala_id"`
	VariableID       int       `json:"variable_id"`
}

func (Sensor) TableName() string {
	return "sensor"
}
