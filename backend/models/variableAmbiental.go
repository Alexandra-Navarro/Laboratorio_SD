package models

type VariableAmbiental struct {
	ID           int     `gorm:"primaryKey; autoIncrement" json:"id"`
	Nombre       string  `json:"nombre"`
	UnidadMedida string  `json:"unidad_medida"`
	UmbralBajo   float64 `json:"umbral_bajo"`
	UmbralAlto   float64 `json:"umbral_alto"`
}

func (VariableAmbiental) TableName() string {
	return "variable_ambiental"
}
