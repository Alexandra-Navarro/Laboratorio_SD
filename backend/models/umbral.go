package models

type Umbral struct {
	ID             int     `gorm:"primaryKey;autoIncrement" json:"id"`
	UmbralBajo     float64 `json:"umbral_bajo"`
	UmbralAlto     float64 `json:"umbral_alto"`
	UmbralBajoPrev float64 `json:"umbral_bajo_prev"`
	UmbralAltoPrev float64 `json:"umbral_alto_prev"`
	SalaID         int     `json:"sala_id"`
	VariableID     int     `json:"variable_id"`
}

func (Umbral) TableName() string {
	return "umbrales"
}
