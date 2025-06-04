package models

import "time"

type Alerta struct {
	ID             int       `gorm:"primaryKey; autoIncrement" json:"id"`
	Tipo           string    `json:"tipo"`
	Descripcion    string    `json:"descripcion"`
	ValorDetectado float64   `json:"valor_detectado"`
	Umbral         string    `json:"umbral"`
	Fecha          time.Time `json:"fecha"`
	SalaID         int       `json:"sala_id"`
}

func (Alerta) TableName() string {
	return "alerta"
}
