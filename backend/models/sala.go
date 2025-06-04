package models

type Sala struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	EscuelaID int    `json:"escuela_id"`
}

func (Sala) TableName() string {
	return "sala"
}
