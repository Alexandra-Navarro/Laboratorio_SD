package models

import "time"

type Medicion struct {
	ID         int       `gorm:"primaryKey; autoIncrement" json:"id"`
	Fecha      time.Time `json:"fecha"`
	Valor      float64   `json:"valor"`
	VariableID int       `json:"variable_id"`
	SalaID     int       `json:"sala_id"`
	SensorID   int       `json:"sensor_id"`
}

func (Medicion) TableName() string {
	return "medicion"
}
